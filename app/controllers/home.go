package controllers

import (
    "github.com/revel/revel"
    "fmt"
)

type Home struct {
    *revel.Controller
}

func (c Home) Index() revel.Result {
    fmt.Println(c.Session.Id())
    return c.RenderTemplate("home/home_index.html")
}
