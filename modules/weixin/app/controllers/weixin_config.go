package controllers

import (
    "github.com/revel/revel"
    "matrix/core"
    "matrix/modules/weixin/models"
)

type WeixinConfig struct {
    *revel.Controller
}

type ConfigForm struct {
    Config models.Config
}

func (f *ConfigForm) Valid(validation *revel.Validation) bool {
    validation.Required(f.Config.Token).Message("Token不能为空！")
    validation.MinSize(f.Config.Token, 3).Message("Token长度不能小于3！")
    validation.MaxSize(f.Config.Token, 32).Message("Token长度不能大于32！")

    if f.Config.EncodingAesKey != "" {
        validation.MinSize(f.Config.EncodingAesKey, 43).Message("EncodingAESKey长度应为43！")
        validation.MaxSize(f.Config.EncodingAesKey, 43).Message("EncodingAESKey长度应为43！")
    }

    if f.Config.AppId != "" {
        validation.MaxSize(f.Config.AppId, 100).Message("AppId长度不能大于100！")
    }

    if f.Config.AppSecret != "" {
        validation.MaxSize(f.Config.AppSecret, 100).Message("AppSecret长度不能大于100！")
    }

    return validation.HasErrors() == false
}

func (c WeixinConfig) Index() revel.Result {
    session := c.DbSession

    config := new(models.Config)
    _, err := session.Limit(1).Get(config)
    core.HandleError(err)

    form := new(ConfigForm)
    form.Config = *config

    c.UnbindToRenderArgs(form, "form")
    return c.RenderTemplate("weixin/config/config_index.html")
}

func (c WeixinConfig) Save() revel.Result {
    session := c.DbSession

    form := new(ConfigForm)
    c.Params.Bind(&form, "form")

    if form.Valid(c.Validation) == false {
        return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage() })
    }

    configInDb := new(models.Config)
    has, err := session.Limit(1).Get(configInDb)
    core.HandleError(err)

    config := &form.Config

    if has == false {
        _, err = session.Insert(config)
        core.HandleError(err)
    } else {
        _, err = session.Id(configInDb.Id).Update(config)
    }

    return c.RenderJson(core.JsonResult{Success: true, Message: "保存成功!"})
}
