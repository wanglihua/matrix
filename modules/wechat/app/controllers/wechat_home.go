package controllers

import (
	"github.com/revel/revel"
)

type WechatHome struct {
	*revel.Controller
}

func (c WechatHome) Index() revel.Result {
	return c.RenderTemplate("wechat/home/home_index.html")
}
