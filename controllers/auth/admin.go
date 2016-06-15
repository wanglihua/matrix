package auth

import (
    "github.com/astaxie/beego"
)

type Admin struct {
    beego.Controller
}

func (c *Admin) Index() {
    //beego.URLFor()
    c.Layout = "layout/layout.tpl"
    c.TplName = "auth/user/user_index.tpl"
}

func (c *Admin) List() {

}

func (c *Admin) Detail() {

}

func (c *Admin) Save() {

}

func (c *Admin) Delete() {

}

