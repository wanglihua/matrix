package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"strconv"

	"matrix/core"
	"matrix/modules/itsm/models"
)

type ItsmEventType struct {
	*revel.Controller
}

func (c ItsmEventType) Index() revel.Result {
	return c.RenderTemplate("itsm/event_type/event_type_index.html")
}

func (c ItsmEventType) ListData() revel.Result {
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

	event_type_list := make([]models.EventTypeInfo, 0, limit)
	err := data_query.Limit(limit, offset).Find(&event_type_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.EventTypeInfo))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  event_type_list,
		Total: count,
	})
}

type EventTypeDetailForm struct {
	EventType models.EventTypeInfo `json:"type"`
}

func (f *EventTypeDetailForm) IsCreate() bool {
	return f.EventType.Id == 0
}

func (f *EventTypeDetailForm) Valid(validation *revel.Validation) bool {
	validation.Required(f.EventType.Name).Message("名称不能为空！")
	if f.EventType.Name != "" {
		validation.MinSize(f.EventType.Name, 1).Message("名称长度不能小于1！")
	}
	if f.EventType.Name != "" {
		validation.MaxSize(f.EventType.Name, 30).Message("名称长度不能大于30！")
	}

	return validation.HasErrors() == false
}

func (c ItsmEventType) DetailData() revel.Result {
	db_session := c.DbSession

	var event_type_id int64
	c.Params.Bind(&event_type_id, "id")

	var event_type models.EventTypeInfo
	if event_type_id != 0 {
		has, err := db_session.Id(event_type_id).Get(&event_type)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的事件类型不存在！"})
		}
	}

	var detail_form EventTypeDetailForm
	detail_form.EventType = event_type

	return c.RenderJson(detail_form)
}

func (c ItsmEventType) Save() revel.Result {
	db_session := c.DbSession

	var detail_form EventTypeDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	event_type := detail_form.EventType

	var affected int64
	if detail_form.IsCreate() {
		affected, err = db_session.Insert(&event_type)
		core.HandleError(err)
	} else {
		affected, err = db_session.Id(event_type.Id).Update(&event_type)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！"})
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c ItsmEventType) Delete() revel.Result {
	db_session := c.DbSession

	event_type_id_list := make([]int64, 0)
	c.Params.Bind(&event_type_id_list, "id_list")

	affected, err := db_session.In("id", event_type_id_list).Delete(new(models.EventTypeInfo))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}
