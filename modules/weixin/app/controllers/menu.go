package controllers

import (
    "github.com/revel/revel"
    "matrix/core"
)

type WeixinMenu struct {
    *revel.Controller
    core.BaseController
}

func (c WeixinMenu) Index() revel.Result {
    return c.RenderTemplate("weixin/menu/menu_index.html")
}

func (c WeixinMenu) Save() revel.Result {
    return c.RenderTemplate("weixin/menu/menu_index.html")
}

func (c WeixinMenu) Download() revel.Result {
    return c.RenderTemplate("weixin/menu/menu_index.html")
}


func (c WeixinMenu) Upload() revel.Result {
    return c.RenderTemplate("weixin/menu/menu_index.html")
}
