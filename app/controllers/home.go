package controllers

import (
    "github.com/revel/revel"
    "fmt"
    "matrix/app/routes"
)

type Home struct {
    *revel.Controller
}

func (c Home) Index() revel.Result {
    fmt.Println(c.Session.Id())
    fmt.Println(routes.User.Index())
    return c.RenderTemplate("home/home_index.html")
}
