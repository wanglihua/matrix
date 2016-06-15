package auth

import (
    "github.com/astaxie/beego"
)

type AuthUser struct {
    beego.Controller
}

func (c *AuthUser) Index() {
    //beego.URLFor()
    c.Layout = "layout/layout.tpl"
    c.TplName = "auth/user/user_index.tpl"
}

func (c *AuthUser) List() {

}

func (c *AuthUser) Detail() {

}

func (c *AuthUser) Save() {

}

func (c *AuthUser) Delete() {

}
