
package controllers

import (
	"encoding/json"
	"strconv"
	"github.com/revel/revel"

	"matrix/core"
	"matrix/modules/inventory/models"
)

type InventoryConfig struct {
	*revel.Controller
}

func (c InventoryConfig) Index() revel.Result {
	return c.RenderTemplate("inventory/config/config_index.html")
}

func (c InventoryConfig) ListData() revel.Result {
	db_session := c.DbSession

	filter, order, offset, limit := core.GetGridRequestParam(c.Request)
	query := db_session.Where(filter)

	//query extra filter here

	data_query := *query
	if order != "" {
		data_query = *data_query.OrderBy(order)
	} else {
		data_query = *data_query.Asc("id")
	}

	config_list := make([]models.Config, 0, limit)
	err := data_query.Limit(limit, offset).Find(&config_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.Config))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  config_list,
		Total: count,
	})
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

	var config_id int64
	c.Params.Bind(&config_id, "id")

	var config models.Config
	if config_id != 0 {
		has, err := db_session.Id(config_id).Get(&config)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的系统配置不存在！"})
		}
	}

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
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c InventoryConfig) Delete() revel.Result {
	db_session := c.DbSession

	config_id_list := make([]int64, 0)
	c.Params.Bind(&config_id_list, "id_list")

	affected, err := db_session.In("id", config_id_list).Delete(new(models.Config))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}

