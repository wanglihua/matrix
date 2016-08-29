package controllers

import (
	"github.com/revel/revel"
	"matrix/service"
)

type WeixinHome struct {
	*revel.Controller
	service.BaseController
}

func (c WeixinHome) Index() revel.Result {
	return c.RenderTemplate("weixin/home/home_index.html")
}
