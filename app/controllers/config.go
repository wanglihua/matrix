package controllers

import (
	"github.com/revel/revel"
	"matrix/app/layout"
	"matrix/app/models"
	"matrix/core"
)

type SysConfig struct {
	*revel.Controller
}

type ConfigForm struct {
	Config models.ConfigInfo
}

func (f ConfigForm) Valid(validation *revel.Validation) bool {
	validation.Required(f.Config.SysName).Message("系统名称不能为空！")
	validation.MinSize(f.Config.SysName, 3).Message("系统名称长度不能小于3！")
	validation.MaxSize(f.Config.SysName, 100).Message("系统名称长度不能大于100！")

	return validation.HasErrors() == false
}

func (c SysConfig) Index() revel.Result {
	session := c.DbSession

	config := new(models.ConfigInfo)
	_, err := session.Limit(1).Get(config)
	core.HandleError(err)

	form := new(ConfigForm)
	form.Config = *config

	c.UnbindToRenderArgs(form, "form")
	return c.RenderTemplate("config/config_index.html")
}

func (c SysConfig) Save() revel.Result {
	session := c.DbSession

	form := new(ConfigForm)
	c.Params.Bind(&form, "form")

	if form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	configInDb := new(models.ConfigInfo)
	has, err := session.Limit(1).Get(configInDb)
	core.HandleError(err)

	config := &form.Config

	if has == false {
		_, err = session.Insert(config)
		core.HandleError(err)
	} else {
		_, err = session.Id(configInDb.Id).Update(config)
	}

	layout.SetSysName(form.Config.SysName)

	return c.RenderJson(core.JsonResult{Success: true, Message: "保存成功!"})
}
