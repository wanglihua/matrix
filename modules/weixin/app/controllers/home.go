package controllers

import (
    "github.com/revel/revel"
    "matrix/core"
)

type WeixinHome struct {
    *revel.Controller
    core.BaseController
}

func (c WeixinHome) Index() revel.Result {
    return c.RenderTemplate("weixin/home/home_index.html")
}
