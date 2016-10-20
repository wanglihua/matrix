package controllers

import (
	"github.com/revel/revel"
	"matrix/core"
)

type Help struct {
	*revel.Controller
}

func (c Help) Index() revel.Result {
	return c.RenderTemplate("help/help_index.html")
}

func (c Help) Faq() revel.Result {
	return c.RenderTemplate("help/help_faq.html")
}

func (c Help) Bug() revel.Result {
	return c.RenderTemplate("help/help_bug.html")
}

func (c Help) BugPost() revel.Result {
	return c.RenderJson(core.JsonResult{Success: true, Message: "保存成功!"})
}

func (c Help) Requirement() revel.Result {
	return c.RenderTemplate("help/help_requirement.html")
}

func (c Help) RequirementPost() revel.Result {
	return c.RenderJson(core.JsonResult{Success: true, Message: "保存成功!"})
}

func (c Help) About() revel.Result {
	return c.RenderTemplate("help/help_about.html")
}

func (c Help) License() revel.Result {
	return c.RenderTemplate("help/help_license.html")
}
