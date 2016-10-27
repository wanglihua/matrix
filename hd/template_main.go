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
	service_action string = ""
	runMode        string = "prod"
	port           int = 0
	importPath     string = "{{.ImportPath}}"
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

	arg_len := len(os.Args)
	if arg_len == 1 {
		service_action = ""
	} else if arg_len != 3 {
		show_cmd_usage()
		os.Exit(0)
	} else  { // arg_len == 3
		if strings.ToLower(os.Args[1]) != "service" {
			show_cmd_usage()
			os.Exit(0)
		}

		service_action = strings.ToLower(os.Args[2])
		if service_action == "debug" {
			service_action = ""
			runMode = "dev"
		}
	}

	isIntSess, err := svc.IsAnInteractiveSession()
	if err != nil {
		log.Fatalf("failed to determine if we are running in an interactive session: %v", err)
	}
	if isIntSess == false {
		run_web()
	}

	if strings.TrimSpace(service_action) == "" {
		run_web()
	} else if service_action == "install" {
		winsvc.InstallService(app.AppName, app.AppName)
	} else if service_action == "remove" {
		winsvc.RemoveService(app.AppName)
	} else if service_action == "start" {
		winsvc.StartService(app.AppName)
	} else if service_action == "stop" {
		winsvc.StopService(app.AppName)
	} else if service_action == "pause" {
		winsvc.PauseService(app.AppName)
	} else if service_action == "continue" {
		winsvc.ContinueService(app.AppName)
	} else {
		show_cmd_usage()
		os.Exit(0)
	}
}

func show_cmd_usage() {
	_, file_name :=  filepath.Split(os.Args[0])
	file_name = strings.TrimSuffix(file_name, filepath.Ext(file_name))

	usage := ` + "`" + `命令行用法：
1 %[1]s
  直接运行应用，不安装成Windows服务
2 %[1]s service install
  将应用安装成Windows服务
3 %[1]s service remove
  移除应用已安装的Windows服务
4 %[1]s service start
  启动应用Windows服务
5 %[1]s service stop
  停止应用Windows服务
6 %[1]s service pause
  暂停应用Windows服务
7 %[1]s service continue
  继续应用Windows服务
` + "`" + `
	fmt.Printf(usage, file_name)
}

func run_web() {
    currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        log.Fatal(err)
    }
    lic_file_path := currentDir + string(os.PathSeparator) + "app.lic"

    if runMode == "prod" {
        importPath = "{{.ImportPath}}"
        srcPath = currentDir
        lic_file_path = currentDir + string(os.PathSeparator) + "{{.ImportPath}}" + string(os.PathSeparator) + "app.lic"
    }

    lic.ValidAppLic(app.AppName, lic_file_path)

    revel.Init(runMode, importPath, srcPath)
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
