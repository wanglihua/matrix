package app

import (
    "github.com/revel/revel"
    "html/template"
    "matrix/app/layout"
    //"bytes"
    //"fmt"
    "matrix/service"
)

func registerTags() {
    revel.TemplateFuncs["test"] = testTemplateFunc
    revel.TemplateFuncs["header"] = headerTemplateFunc
    revel.TemplateFuncs["footer"] = footerTemplateFunc
}

func testTemplateFunc(renderArgs map[string]interface{}) template.HTML {
    session := renderArgs["session"].(revel.Session)
    return template.HTML("<span style='color:red;'>Session Id: " + session.Id() + "</span>")
    //return template.HTML("<span style='color:red;'>Hello World Just a test!</span>")
}

type headerTemplateData struct {
    Title string
}

//var headerTemplate *template.Template = nil

func headerTemplateFunc(title string, renderArgs map[string]interface{}) template.HTML {
    /*
    if headerTemplate == nil {
        headerTemplate, _ = template.New("site header").Parse(layout.HeaderTemplate)
    }

    headerData := headerTemplateData{
        Title: title,
    }

    var renderBuffer bytes.Buffer
    headerTemplate.Execute(&renderBuffer, headerData)
    return template.HTML(renderBuffer.String())
    */

    //session := renderArgs["session"].(revel.Session) //通过session得到当前登录人信息啥的

    session := renderArgs["session"].(revel.Session)
    loginUser := service.GetLoginUser(session)
    loginNickName := "未登录"
    if loginUser != nil {
        loginNickName = loginUser.NickName
    }
    return template.HTML(layout.GetHeader(title, loginNickName))
}

func footerTemplateFunc() template.HTML {
    //just plain string
    return template.HTML(layout.FooterTemplate)
}
