package auth

import (
    "github.com/astaxie/beego"
)

type AuthHome struct {
    beego.Controller
}

func (c *AuthHome) Index() {
    c.Data["Website"] = "beego.me"
    c.Data["Email"] = "astaxie@gmail.com"

    c.Layout = "layout/layout.tpl"
    c.TplName = "home/index.tpl"
}
