package auth

import (
    "github.com/astaxie/beego"
)

type User struct {
    beego.Controller
}

func (c *User) Index() {
    //beego.URLFor()
    c.Layout = "layout/layout.tpl"
    c.TplName = "auth/user/user_index.tpl"
}

func (c *User) List() {

}

func (c *User) Detail() {

}

func (c *User) Save() {

}

func (c *User) Delete() {

}
