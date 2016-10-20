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
    "fmt"
    "strconv"
)

var (
    runMode    *string = flag.String("runMode", "prod", "Run mode.")
    port       *int    = flag.Int("port", 0, "By default, read from app.conf")
    importPath *string = flag.String("importPath", "{{.ImportPath}}", "Go Import Path for the app.")
    srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

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

    currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        log.Fatal(err)
    }
    lic_file_path := currentDir + string(os.PathSeparator) + "app.lic"

    if *runMode == "prod" {
        *importPath = "{{.ImportPath}}"
        *srcPath = currentDir
        lic_file_path = currentDir + string(os.PathSeparator) + "{{.ImportPath}}" + string(os.PathSeparator) + "app.lic"
    }

    lic.ValidAppLic(app.AppName, lic_file_path)

    revel.Init(*runMode, *importPath, *srcPath)
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

    /*
    revel.HttpAddr = "127.0.0.1"
    var err error
    db.Engine, err = xorm.NewEngine("sqlite3", "./app.db")
    if err != nil  {
        panic(err)
    }
    */

    if revel.DevMode == false {
        if revel.HttpPort == 80 {
            fmt.Println("请打开浏览器访问 http://localhost")
        } else {
            fmt.Println("请打开浏览器访问 http://localhost:" + strconv.Itoa(revel.HttpPort))
        }
    }

    revel.Run(*port)
}
`
