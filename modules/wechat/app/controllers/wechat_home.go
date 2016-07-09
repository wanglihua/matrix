package controllers

import (
    "github.com/revel/revel"
    "matrix/core"
)

type WechatHome struct {
    *revel.Controller
    core.BaseController
}

func (c WechatHome) Index() revel.Result {
    return c.RenderTemplate("wechat/home/home_index.html")
}
