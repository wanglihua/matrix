package app

import (
	"github.com/go-xorm/xorm"
	"github.com/revel/revel"
	"html/template"
	"matrix/app/layout"
	"strconv"
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
}

type headerTemplateData struct {
	Title string
}

func headerTemplateFunc(title string, renderArgs map[string]interface{}) template.HTML {
	db_session := renderArgs["dbsession"].(*xorm.Session)
	web_session := renderArgs["session"].(revel.Session)
	return template.HTML(layout.GetHeader(title, db_session, web_session))
}

func footerTemplateFunc() template.HTML {
	//just plain string
	return template.HTML(layout.GetFooter())
}

func intToStrTemplateFunc(val interface{}) string {
	return fmt.Sprintf("%d", val)
}

