package controllers

import (
    "github.com/revel/revel"
    "matrix/service"
)

type WechatHome struct {
    *revel.Controller
    service.BaseController
}

func (c WechatHome) Index() revel.Result {
    return c.RenderTemplate("wechat/home/home_index.html")
}
