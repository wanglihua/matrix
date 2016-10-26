package main

const MAIN_TEMPLATE = `package main

import (
    "runtime/debug"
    "flag"
    "reflect"
    "github.com/revel/revel"{{range $k, $v := $.ImportPaths}}
    {{$v}} "{{$k}}"{{end}}
    "github.com/revel/revel/testing"
    "path/filepath"
    "os"
    "log"
    "{{.ImportPath}}/app"
	"{{.ImportPath}}/core/lic"
    "{{.ImportPath}}/db"
	"{{.ImportPath}}/core/winsvc"
    "fmt"
    "strconv"
	"golang.org/x/sys/windows/svc"
	"strings"
)

var (
	service_action *string = flag.String("service", "", "windows service action to perform: install remove start stop pause continue")
	runMode        *string = flag.String("runMode", "prod", "Run mode: dev prod")
	port           int = 0
	importPath     string = "matrix"
	srcPath        string = ""

    // So compiler won't complain if the generated code doesn't reference reflect package...
    _ = reflect.Invalid
)

func main() {
    defer func() {
        if err := recover(); err != nil {
            errorDesc := fmt.Sprint(err)
            revel.ERROR.Print(errorDesc, "\n", string(debug.Stack()))
        }
    }()

    flag.Parse()

	isIntSess, err := svc.IsAnInteractiveSession()
	if err != nil {
		log.Fatalf("failed to determine if we are running in an interactive session: %v", err)
	}
	if isIntSess == false {
		run_web()
	}

	if strings.TrimSpace(*service_action) == "" {
		run_web()
	} else if *service_action == "install" {
		winsvc.InstallService(app.AppName, app.AppName)
	} else if *service_action == "remove" {
		winsvc.RemoveService(app.AppName)
	} else if *service_action == "start" {
		winsvc.StartService(app.AppName)
	} else if *service_action == "stop" {
		winsvc.StopService(app.AppName)
	} else if *service_action == "pause" {
		winsvc.PauseService(app.AppName)
	} else if *service_action == "continue" {
		winsvc.ContinueService(app.AppName)
	}
}

func run_web() {
    currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        log.Fatal(err)
    }
    lic_file_path := currentDir + string(os.PathSeparator) + "app.lic"

    if *runMode == "prod" {
        importPath = "{{.ImportPath}}"
        srcPath = currentDir
        lic_file_path = currentDir + string(os.PathSeparator) + "{{.ImportPath}}" + string(os.PathSeparator) + "app.lic"
    }

    lic.ValidAppLic(app.AppName, lic_file_path)

    revel.Init(*runMode, importPath, srcPath)
    revel.INFO.Println("Running revel server")

    {{range $i, $c := .Controllers}}
    revel.RegisterController((*{{index $.ImportPaths .ImportPath}}.{{.StructName}})(nil),
        []*revel.MethodType{
            {{range .MethodSpecs}}&revel.MethodType{
                Name: "{{.Name}}",
                Args: []*revel.MethodArg{ {{range .Args}}
                    &revel.MethodArg{Name: "{{.Name}}", Type: reflect.TypeOf((*{{index $.ImportPaths .ImportPath | .TypeExpr.TypeName}})(nil)) },{{end}}
                },
                RenderArgNames: map[int][]string{ {{range .RenderCalls}}
                    {{.Line}}: []string{ {{range .Names}}
                        "{{.}}",{{end}}
                    },{{end}}
                },
            },
            {{end}}
        })
    {{end}}
    revel.DefaultValidationKeys = map[string]map[int]string{ {{range $path, $lines := .ValidationKeys}}
        "{{$path}}": { {{range $line, $key := $lines}}
            {{$line}}: "{{$key}}",{{end}}
        },{{end}}
    }
    testing.TestSuites = []interface{}{ {{range .TestSuites}}
        (*{{index $.ImportPaths .ImportPath}}.{{.StructName}})(nil),{{end}}
    }

	db.InitDatabase()

	if revel.DevMode == false {
		if revel.HttpPort == 80 {
			fmt.Println("请打开浏览器访问 http://localhost")
		} else {
			fmt.Println("请打开浏览器访问 http://localhost:" + strconv.Itoa(revel.HttpPort))
		}
	}

	revel.Run(port)
}
`
