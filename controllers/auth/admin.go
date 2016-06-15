package auth

import (
    "github.com/astaxie/beego"
)

type AuthAdmin struct {
    beego.Controller
}

func (c *AuthAdmin) Index() {
    //beego.URLFor()
    c.Layout = "layout/layout.tpl"
    c.TplName = "auth/user/user_index.tpl"
}

func (c *AuthAdmin) List() {

}

func (c *AuthAdmin) Detail() {

}

func (c *AuthAdmin) Save() {

}

func (c *AuthAdmin) Delete() {

}

