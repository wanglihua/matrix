package controllers

import (
    "github.com/revel/revel"
    "fmt"
    "matrix/db"
    "matrix/modules/auth/app"
)

type AuthUser struct {
    *revel.Controller
}

func (c AuthUser) Index() revel.Result {
    fmt.Println(c.Session.Id())
    db.Engine.Sync2(&app.User{})
    return c.RenderTemplate("user/user_index.html")
}
