package main

import (
	"fmt"
	"os"
	"path"
	"text/template"

	"github.com/revel/revel"
	"log"
	"path/filepath"
	"strings"
)

var cmdRoute = &Command{
	UsageLine: "route [import path]",
	Short:     "generate routes.go main.go",
	Long: `
    generate routes.go main.go
    hd route

`,
}

func init() {
	cmdRoute.Run = generateRoute
}

func generateRoute(cmd *Command, args []string) int {
	/*
	    if len(args) < 1 {
	        fmt.Fprintf(os.Stderr, "%s\n%s", cmdRoute.UsageLine, cmdRoute.Long)
	        return 1
	    }

		importPath := args[0]

	    runMode := "dev"
	    if len(args) >= 2 {
	        runMode = args[1]
	    }
	*/

	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	cur_path_list := strings.Split(currentDir, string(os.PathSeparator))
	import_path := cur_path_list[len(cur_path_list)-1]

	if !revel.Initialized {
		revel.Init("dev", import_path, "")
	}

	sourceInfo, compileError := ProcessSource(revel.CodePaths)
	if compileError != nil {
		//panicOnError(compileError, "Failed to build")
		revel.ERROR.Fatalf("Failed to build: %v", compileError)
		os.Exit(2)
	}

	// Add the db.import to the import paths.
	if dbImportPath, found := revel.Config.String("db.import"); found {
		sourceInfo.InitImportPaths = append(sourceInfo.InitImportPaths, dbImportPath)
	}

	// Generate two source files.
	templateArgs := map[string]interface{}{
		"Controllers":    sourceInfo.ControllerSpecs(),
		"ValidationKeys": sourceInfo.ValidationKeys,
		"ImportPaths":    CalcImportAliases(sourceInfo),
		"TestSuites":     sourceInfo.TestSuites(),
		"ImportPath":     revel.ImportPath,
		//"RunMode":        runMode,
		"RunMode": "dev",
	}

	//harness.GenSource("tmp", "main.go", harness.MAIN, templateArgs)

	mainCode := revel.ExecuteTemplate(
		template.Must(template.New("").Parse(MAIN_TEMPLATE)),
		templateArgs)

	// Create the main file
	mainFile, mainError := os.Create(path.Join(revel.BasePath, "main.go"))
	defer mainFile.Close()
	if mainError != nil {
		revel.ERROR.Fatalf("Failed to create file: %v", mainError)
	}
	_, mainError = mainFile.WriteString(mainCode)
	if mainError != nil {
		revel.ERROR.Fatalf("Failed to write to file: %v", mainError)
	}

	//harness.GenSource("routes", "routes.go", harness.ROUTES, templateArgs)

	var routesDir = "routes"
	sourceCode := revel.ExecuteTemplate(
		template.Must(template.New("").Parse(ROUTES_TEMPLATE)),
		templateArgs)

	// Create a fresh dir.
	cleanDir(routesDir)
	//routesPath := path.Join(revel.BasePath, routesDir)
	routesPath := path.Join(revel.AppPath, routesDir)
	makeDirError := os.Mkdir(routesPath, 0777)
	if makeDirError != nil && !os.IsExist(makeDirError) {
		revel.ERROR.Fatalf("Failed to make '%v' directory: %v", routesDir, makeDirError)
	}

	// Create the file
	routesFile, routesError := os.Create(path.Join(routesPath, "routes.go"))
	defer routesFile.Close()
	if routesError != nil {
		revel.ERROR.Fatalf("Failed to create file: %v", routesError)
	}
	_, routesError = routesFile.WriteString(sourceCode)
	if routesError != nil {
		revel.ERROR.Fatalf("Failed to write to file: %v", routesError)
	}

	fmt.Println("Generate main.go routes.go success!")

	return 0
}

func cleanDir(dir string) {
	//tmpPath := path.Join(revel.BasePath, dir)
	tmpPath := path.Join(revel.AppPath, dir)
	f, err := os.Open(tmpPath)
	if err != nil {
		if !os.IsNotExist(err) {
			revel.ERROR.Println("Failed to clean dir:", err)
		}
	} else {
		defer f.Close()
		infos, err := f.Readdir(0)
		if err != nil {
			if !os.IsNotExist(err) {
				revel.ERROR.Println("Failed to clean dir:", err)
			}
		} else {
			for _, info := range infos {
				path := path.Join(tmpPath, info.Name())
				if info.IsDir() {
					err := os.RemoveAll(path)
					if err != nil {
						revel.ERROR.Println("Failed to remove dir:", err)
					}
				} else {
					err := os.Remove(path)
					if err != nil {
						revel.ERROR.Println("Failed to remove file:", err)
					}
				}
			}
		}
	}
}

// Looks through all the method args and returns a set of unique import paths
// that cover all the method arg types.
// Additionally, assign package aliases when necessary to resolve ambiguity.
func CalcImportAliases(src *SourceInfo) map[string]string {
	aliases := make(map[string]string)
	typeArrays := [][]*TypeInfo{src.ControllerSpecs(), src.TestSuites()}
	for _, specs := range typeArrays {
		for _, spec := range specs {
			addAlias(aliases, spec.ImportPath, spec.PackageName)

			for _, methSpec := range spec.MethodSpecs {
				for _, methArg := range methSpec.Args {
					if methArg.ImportPath == "" {
						continue
					}

					addAlias(aliases, methArg.ImportPath, methArg.TypeExpr.PkgName)
				}
			}
		}
	}

	// Add the "InitImportPaths", with alias "_"
	for _, importPath := range src.InitImportPaths {
		if _, ok := aliases[importPath]; !ok {
			aliases[importPath] = "_"
		}
	}

	return aliases
}

func addAlias(aliases map[string]string, importPath, pkgName string) {
	alias, ok := aliases[importPath]
	if ok {
		return
	}
	alias = makePackageAlias(aliases, pkgName)
	aliases[importPath] = alias
}

func makePackageAlias(aliases map[string]string, pkgName string) string {
	i := 0
	alias := pkgName
	for containsValue(aliases, alias) {
		alias = fmt.Sprintf("%s%d", pkgName, i)
		i++
	}
	return alias
}

func containsValue(m map[string]string, val string) bool {
	for _, v := range m {
		if v == val {
			return true
		}
	}
	return false
}
