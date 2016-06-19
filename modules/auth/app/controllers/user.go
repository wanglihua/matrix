package controllers

import (
    "github.com/revel/revel"
    "fmt"
)

type AuthUser struct {
    *revel.Controller
}

func (c AuthUser) Index() revel.Result {
    fmt.Println(c.Session.Id())
    return c.RenderTemplate("user/user_index.html")
}
