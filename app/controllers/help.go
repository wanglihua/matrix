package controllers

import (
    "github.com/revel/revel"
    "matrix/core"
)

type Help struct {
    *revel.Controller
    core.BaseController
}

func (c Help) Index() revel.Result {
    return c.RenderTemplate("help/help_index.html")
}

func (c Help) Faq() revel.Result {
    return c.RenderTemplate("help/help_faq.html")
}

func (c Help) About() revel.Result {
    return c.RenderTemplate("help/help_about.html")
}

func (c Help) License() revel.Result {
    return c.RenderTemplate("help/help_license.html")
}
