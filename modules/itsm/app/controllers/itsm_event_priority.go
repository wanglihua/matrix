package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"strconv"

	"matrix/core"
	"matrix/modules/itsm/models"
)

type ItsmEventPriority struct {
	*revel.Controller
}

func (c ItsmEventPriority) Index() revel.Result {
	return c.RenderTemplate("itsm/event_priority/event_priority_index.html")
}

func (c ItsmEventPriority) ListData() revel.Result {
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

	event_priority_list := make([]models.EventPriorityInfo, 0, limit)
	err := data_query.Limit(limit, offset).Find(&event_priority_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.EventPriorityInfo))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  event_priority_list,
		Total: count,
	})
}

type EventPriorityDetailForm struct {
	EventPriority models.EventPriorityInfo `json:"priority"`
}

func (f *EventPriorityDetailForm) IsCreate() bool {
	return f.EventPriority.Id == 0
}

func (f *EventPriorityDetailForm) Valid(validation *revel.Validation) bool {
	validation.Required(f.EventPriority.Name).Message("名称不能为空！")
	if f.EventPriority.Name != "" {
		validation.MinSize(f.EventPriority.Name, 1).Message("名称长度不能小于1！")
	}
	if f.EventPriority.Name != "" {
		validation.MaxSize(f.EventPriority.Name, 30).Message("名称长度不能大于30！")
	}

	return validation.HasErrors() == false
}

func (c ItsmEventPriority) DetailData() revel.Result {
	db_session := c.DbSession

	var event_priority_id int64
	c.Params.Bind(&event_priority_id, "id")

	var event_priority models.EventPriorityInfo
	if event_priority_id != 0 {
		has, err := db_session.Id(event_priority_id).Get(&event_priority)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的事件优先级不存在！"})
		}
	}

	var detail_form EventPriorityDetailForm
	detail_form.EventPriority = event_priority

	return c.RenderJson(detail_form)
}

func (c ItsmEventPriority) Save() revel.Result {
	db_session := c.DbSession

	var detail_form EventPriorityDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	event_priority := detail_form.EventPriority

	var affected int64
	if detail_form.IsCreate() {
		affected, err = db_session.Insert(&event_priority)
		core.HandleError(err)
	} else {
		affected, err = db_session.Id(event_priority.Id).Update(&event_priority)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！"})
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c ItsmEventPriority) Delete() revel.Result {
	db_session := c.DbSession

	event_priority_id_list := make([]int64, 0)
	c.Params.Bind(&event_priority_id_list, "id_list")

	affected, err := db_session.In("id", event_priority_id_list).Delete(new(models.EventPriorityInfo))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}
