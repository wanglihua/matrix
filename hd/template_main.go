package main

const MAIN_TEMPLATE = `package main

import (
    "flag"
    "reflect"
    "github.com/revel/revel"{{range $k, $v := $.ImportPaths}}
    {{$v}} "{{$k}}"{{end}}
    "github.com/revel/revel/testing"
    "path/filepath"
    "os"
    "log"
    "{{.ImportPath}}/db"
    //"github.com/go-xorm/xorm"
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
    flag.Parse()
    if *runMode == "prod" {
        currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
        if err != nil {
            log.Fatal(err)
        }
        *srcPath = currentDir + "\\..\\";
    }

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

    revel.Run(*port)
}
`
