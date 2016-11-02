package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"strconv"

	"matrix/core"
	"matrix/modules/inventory/models"
)

type InventoryConfig struct {
	*revel.Controller
}

func (c InventoryConfig) Index() revel.Result {
	return c.RenderTemplate("inventory/config/config_index.html")
}

type ConfigDetailForm struct {
	Config models.Config `json:"config"`
}

func (f *ConfigDetailForm) IsCreate() bool {
	return f.Config.Id == 0
}

func (f *ConfigDetailForm) Valid(validation *revel.Validation) bool {
	return validation.HasErrors() == false
}

func (c InventoryConfig) DetailData() revel.Result {
	db_session := c.DbSession

	var config models.Config
	_, err := db_session.Limit(1, 0).Get(&config)
	core.HandleError(err)

	var detail_form ConfigDetailForm
	detail_form.Config = config

	return c.RenderJson(detail_form)
}

func (c InventoryConfig) Save() revel.Result {
	db_session := c.DbSession

	var detail_form ConfigDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	config := detail_form.Config

	var affected int64
	if detail_form.IsCreate() {
		affected, err = db_session.Insert(&config)
		core.HandleError(err)
	} else {
		affected, err = db_session.Id(config.Id).Update(&config)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！"})
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}
