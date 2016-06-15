package controllers

import (
    "github.com/astaxie/beego"
)

type Home struct {
    beego.Controller
}

func (c *Home) Index() {
    c.Data["Website"] = "beego.me"
    c.Data["Email"] = "astaxie@gmail.com"

    c.Layout = "layout/layout.tpl"
    c.TplName = "home/index.tpl"
}
