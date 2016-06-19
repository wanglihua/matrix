package controllers

import (
    "github.com/revel/revel"
    "fmt"
)

type User struct {
    *revel.Controller
}

func (c User) Index() revel.Result {
    fmt.Println(c.Session.Id())
    return c.RenderTemplate("user/user_index.html")
}
