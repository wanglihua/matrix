package controllers

import (
	"github.com/revel/revel"
)

type WeixinHome struct {
	*revel.Controller
}

func (c WeixinHome) Index() revel.Result {
	return c.RenderTemplate("weixin/home/home_index.html")
}
