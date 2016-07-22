package app

import (
    "github.com/revel/revel"
    "github.com/go-xorm/xorm"
    "html/template"
    "matrix/app/layout"
    //"bytes"
    //"fmt"
    "matrix/core"
    "strconv"
    //"time"
)

func registerTags() {
    revel.TemplateFuncs["test"] = testTemplateFunc
    revel.TemplateFuncs["header"] = headerTemplateFunc
    revel.TemplateFuncs["footer"] = footerTemplateFunc

    revel.TemplateFuncs["IntToStr"] = intToStrTemplateFunc
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
    //timeBegin := time.Now()

    dbsession := renderArgs["dbsession"].(*xorm.Session)

    sysName := layout.GetSysName(dbsession)

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

    //revel.TRACE.Println("headerTemplateFunc")

    session := renderArgs["session"].(revel.Session)
    loginUser := core.GetLoginUser(session)
    loginNickName := "未登录"
    if loginUser != nil {
        loginNickName = loginUser.NickName
    }

    //timeEnd := time.Now()
    //duration :=  int64(timeEnd.Sub(timeBegin)) / 1000000
    //revel.TRACE.Println("headerTemplateFunc duration: ", duration, "ms")

    return template.HTML(layout.GetHeader(title, sysName, loginNickName))
}

func footerTemplateFunc() template.HTML {
    //just plain string
    return template.HTML(layout.FooterTemplate)
}

func intToStrTemplateFunc(val int) string {
    return strconv.Itoa(val)
}
