
package controllers

import (
	"encoding/json"
	"strconv"
	"github.com/revel/revel"

	"matrix/core"
	"matrix/modules/inventory/models"
)

type InventoryHandler struct {
	*revel.Controller
}

func (c InventoryHandler) Index() revel.Result {
	return c.RenderTemplate("inventory/handler/handler_index.html")
}

func (c InventoryHandler) ListData() revel.Result {
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

	handler_list := make([]models.Handler, 0, limit)
	err := data_query.Limit(limit, offset).Find(&handler_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.Handler))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  handler_list,
		Total: count,
	})
}

type HandlerDetailForm struct {
	Handler models.Handler `json:"handler"`
}

func (f *HandlerDetailForm) IsCreate() bool {
	return f.Handler.Id == 0
}

func (f *HandlerDetailForm) Valid(validation *revel.Validation) bool { 
	validation.Required(f.Handler.Name).Message("名称不能为空！")
	if f.Handler.Name != "" {
		validation.MinSize(f.Handler.Name, 2).Message("名称长度不能小于2！")
	}
	if f.Handler.Name != "" {
		validation.MaxSize(f.Handler.Name, 100).Message("名称长度不能大于100！")
	}

	if f.Handler.Desc.Valid == true {
		validation.MinSize(f.Handler.Desc.String, 2).Message("描述长度不能小于2！")
	}
	if f.Handler.Desc.Valid == true {
		validation.MaxSize(f.Handler.Desc.String, 300).Message("描述长度不能大于300！")
	}

	validation.Required(f.Handler.State).Message("状态不能为空！")
	if f.Handler.State != "" {
		validation.MinSize(f.Handler.State, 1).Message("状态长度不能小于1！")
	}

	return validation.HasErrors() == false
}

func (c InventoryHandler) DetailData() revel.Result {
	db_session := c.DbSession

	var handler_id int64
	c.Params.Bind(&handler_id, "id")

	var handler models.Handler
	if handler_id != 0 {
		has, err := db_session.Id(handler_id).Get(&handler)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的经手人不存在！"})
		}
	}

	var detail_form HandlerDetailForm
	detail_form.Handler = handler

	return c.RenderJson(detail_form)
}

func (c InventoryHandler) Save() revel.Result {
	db_session := c.DbSession

	var detail_form HandlerDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	handler := detail_form.Handler

	var affected int64
	if detail_form.IsCreate() {
		name_count, err := db_session.Where("handler_name = ?", handler.Name).Count(&handler)
		core.HandleError(err)
		if name_count != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，名称已存在！" })
		}

		affected, err = db_session.Insert(&handler)
		core.HandleError(err)
	} else { 
		name_count, err := db_session.Where("id <> ? and handler_name = ?", handler.Id, handler.Name).Count(&handler)
		core.HandleError(err)
		if name_count != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，名称已存在！" })
		}

		affected, err = db_session.Id(handler.Id).Update(&handler)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c InventoryHandler) Delete() revel.Result {
	db_session := c.DbSession

	handler_id_list := make([]int64, 0)
	c.Params.Bind(&handler_id_list, "id_list")

	affected, err := db_session.In("id", handler_id_list).Delete(new(models.Handler))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}

