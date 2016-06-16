package main

import (
    "flag"
    "reflect"
    "github.com/revel/revel"
    staticControllers "github.com/revel/modules/static/app/controllers"
    _ "github.com/revel/modules/testrunner/app"
    testRunnerControllers "github.com/revel/modules/testrunner/app/controllers"
    _ "matrix/app"
    controllers "matrix/app/controllers"
    tests "matrix/tests"
    "github.com/revel/revel/testing"
    "path/filepath"
    "os"
    "log"
)

var (
    runMode    *string = flag.String("runMode", "dev", "Run mode.")
    port       *int = flag.Int("port", 0, "By default, read from app.conf")
    importPath *string = flag.String("importPath", "matrix", "Go Import Path for the app.")
    srcPath    *string = flag.String("srcPath", "", "Path to the source root.")

    _ = reflect.Invalid
)

func main() {

    if *runMode != "dev" {
        //部署的时候把当前目录下的src目录作为 srcPath
        currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
        if err != nil {
            log.Fatal(err)
        }
        *srcPath = currentDir + "\\src";
    }

    flag.Parse()

    revel.Init(*runMode, *importPath, *srcPath)
    revel.INFO.Println("Running revel server")

    revel.RegisterController((*controllers.Home)(nil),
        []*revel.MethodType{
            &revel.MethodType{
                Name: "Index",
                Args: []*revel.MethodArg{
                },
                RenderArgNames: map[int][]string{
                    10: []string{
                    },
                },
            },
        })

    revel.RegisterController((*testRunnerControllers.TestRunner)(nil),
        []*revel.MethodType{
            &revel.MethodType{
                Name: "Index",
                Args: []*revel.MethodArg{
                },
                RenderArgNames: map[int][]string{
                    72: []string{
                        "testSuites",
                    },
                },
            },
            &revel.MethodType{
                Name: "Suite",
                Args: []*revel.MethodArg{
                    &revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
                },
                RenderArgNames: map[int][]string{
                },
            },
            &revel.MethodType{
                Name: "Run",
                Args: []*revel.MethodArg{
                    &revel.MethodArg{Name: "suite", Type: reflect.TypeOf((*string)(nil)) },
                    &revel.MethodArg{Name: "test", Type: reflect.TypeOf((*string)(nil)) },
                },
                RenderArgNames: map[int][]string{
                    125: []string{
                    },
                },
            },
            &revel.MethodType{
                Name: "List",
                Args: []*revel.MethodArg{
                },
                RenderArgNames: map[int][]string{
                },
            },

        })

    revel.RegisterController((*staticControllers.Static)(nil),
        []*revel.MethodType{
            &revel.MethodType{
                Name: "Serve",
                Args: []*revel.MethodArg{
                    &revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
                    &revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
                },
                RenderArgNames: map[int][]string{
                },
            },
            &revel.MethodType{
                Name: "ServeModule",
                Args: []*revel.MethodArg{
                    &revel.MethodArg{Name: "moduleName", Type: reflect.TypeOf((*string)(nil)) },
                    &revel.MethodArg{Name: "prefix", Type: reflect.TypeOf((*string)(nil)) },
                    &revel.MethodArg{Name: "filepath", Type: reflect.TypeOf((*string)(nil)) },
                },
                RenderArgNames: map[int][]string{
                },
            },

        })

    revel.RegisterController((*controllers.Login)(nil),
        []*revel.MethodType{
            &revel.MethodType{
                Name: "Index",
                Args: []*revel.MethodArg{
                },
                RenderArgNames: map[int][]string{
                },
            },
            &revel.MethodType{
                Name: "Login",
                Args: []*revel.MethodArg{
                },
                RenderArgNames: map[int][]string{
                },
            },

        })

    revel.DefaultValidationKeys = map[string]map[int]string{
    }
    testing.TestSuites = []interface{}{
        (*tests.AppTest)(nil),
    }

    revel.Run(*port)
}
