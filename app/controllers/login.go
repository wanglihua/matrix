package controllers

import (
    "github.com/revel/revel"
)

type Login struct {
    *revel.Controller
}

func (c Login) Index() revel.Result {
    return c.RenderTemplate("login/login_index.html")
}

func (c Login) Login() revel.Result {
    return c.RenderJson(JsonResult{Success:true, Message:"success!"})
}

func (c Login) Logout() revel.Result {
    return c.RenderJson(JsonResult{Success:true, Message:"success!"})
}