package controllers

import (
    "github.com/revel/revel"
    "matrix/core"
)

type WeixinConfig struct {
    *revel.Controller
    core.BaseController
}

type ConfigForm struct {
    Token          string
    EncodingAesKey string
    AppId          string
    AppSecret      string
}

func (f ConfigForm) Valid(validation *revel.Validation) bool {
    validation.Required(f.Token).Message("Token不能为空！")
    validation.MinSize(f.Token, 3).Message("Token长度不能小于3！")
    validation.MaxSize(f.Token, 32).Message("Token长度不能大于32！")

    if f.EncodingAesKey != "" {
        validation.MinSize(f.EncodingAesKey, 43).Message("EncodingAESKey长度应为43！")
        validation.MaxSize(f.EncodingAesKey, 43).Message("EncodingAESKey长度应为43！")
    }

    if f.AppId != "" {
        validation.MaxSize(f.AppId, 100).Message("AppId长度不能大于100！")
    }

    if f.AppSecret != "" {
        validation.MaxSize(f.AppSecret, 100).Message("AppSecret长度不能大于100！")
    }

    return validation.HasErrors() == false
}

func (c WeixinConfig) Index() revel.Result {
    return c.RenderTemplate("weixin/config/config_index.html")
}

func (c WeixinConfig) Save() revel.Result {
    form := new(ConfigForm)
    c.Params.Bind(&form, "form")

    if form.Valid(c.Validation) == false {
        return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage() })
    }

    return c.RenderJson(core.JsonResult{Success: true, Message: "保存成功!"})
}
