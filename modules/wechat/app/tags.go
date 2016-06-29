package app

import (
    "github.com/revel/revel"
    "html/template"
)

func registerTags() {
    revel.TemplateFuncs["wechat_header"] = headerTemplateFunc
    revel.TemplateFuncs["wechat_footer"] = footerTemplateFunc
}


func headerTemplateFunc(title string, renderArgs map[string]interface{}) template.HTML {
    return template.HTML("wechat header")
}

func footerTemplateFunc() template.HTML {
    return template.HTML("wechat footer")
}
