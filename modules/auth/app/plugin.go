package app

import (
    "fmt"
    "github.com/revel/revel"
)

func init() {
    revel.OnAppStart(func() {
        fmt.Println("auth module start!")

        delimsConfig, _ := revel.Config.String("template.delimiters")
        fmt.Println("user app plugin.go " + delimsConfig)
        fmt.Println("user app plugin.go " + revel.TemplateDelims)
    })
}
