package app

import (
    "github.com/revel/revel"
    "html/template"
    "matrix/app/layout"
    //"bytes"
    //"fmt"
    "matrix/core"
    "matrix/app/models"
    "matrix/db"
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
    if core.SysName == "" {
        session := db.Engine.NewSession()

        config := new(models.Config)
        has, err := session.Limit(1).Get(config)
        core.HandleError(err)
        if has {
            core.SysName = config.SysName
        }

        session.Close()
    }

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

    revel.TRACE.Println("headerTemplateFunc")

    session := renderArgs["session"].(revel.Session)
    loginUser := core.GetLoginUser(session)
    loginNickName := "未登录"
    if loginUser != nil {
        loginNickName = loginUser.NickName
    }
    return template.HTML(layout.GetHeader(title, core.SysName, loginNickName))
}

func footerTemplateFunc() template.HTML {
    //just plain string
    return template.HTML(layout.FooterTemplate)
}
