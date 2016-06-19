package app

import (
    "github.com/revel/revel"
    "html/template"
)

func testTemplateFunc(renderArgs map[string]interface{}) template.HTML {
    session := renderArgs["session"].(revel.Session)
    return template.HTML("<span style='color:red;'>Hello World Just a test " + session.Id() + "!</span>")
    //return template.HTML("<span style='color:red;'>Hello World Just a test!</span>")
}
