package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/revel/revel"

	"matrix/core"
	"matrix/modules/itsm/models"
)

type ItsmEventStatus struct {
	*revel.Controller
}

func (c ItsmEventStatus) Index() revel.Result {
	return c.RenderTemplate("itsm/event_status/event_status_index.html")
}

func (c ItsmEventStatus) ListData() revel.Result {
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

	event_status_list := make([]models.EventStatus, 0, limit)
	err := data_query.Limit(limit, offset).Find(&event_status_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.EventStatus))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  event_status_list,
		Total: count,
	})
}

type EventStatusDetailForm struct {
	EventStatus models.EventStatus `json:"status"`
}

func (f *EventStatusDetailForm) IsCreate() bool {
	return f.EventStatus.Id == 0
}

func (f *EventStatusDetailForm) Valid(validation *revel.Validation) bool {
	validation.Required(f.EventStatus.Name).Message("名称不能为空！")
	if f.EventStatus.Name != "" {
		validation.MinSize(f.EventStatus.Name, 1).Message("名称长度不能小于1！")
	}
	if f.EventStatus.Name != "" {
		validation.MaxSize(f.EventStatus.Name, 30).Message("名称长度不能大于30！")
	}
	if f.EventStatus.Desc.Valid == true && f.EventStatus.Desc.String != "" {
		validation.MaxSize(f.EventStatus.Desc.String, 300).Message("描述长度不能大于300！")
	}

	return validation.HasErrors() == false
}

func (c ItsmEventStatus) DetailData() revel.Result {
	db_session := c.DbSession

	var event_status_id int64
	c.Params.Bind(&event_status_id, "id")

	var event_status models.EventStatus
	if event_status_id != 0 {
		has, err := db_session.Id(event_status_id).Get(&event_status)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的事件状态不存在！"})
		}
	}

	var detail_form EventStatusDetailForm
	detail_form.EventStatus = event_status

	return c.RenderJson(detail_form)
}

func (c ItsmEventStatus) Save() revel.Result {
	db_session := c.DbSession

	var detail_form EventStatusDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	event_status := detail_form.EventStatus

	var affected int64
	if detail_form.IsCreate() {
		affected, err = db_session.Insert(&event_status)
		core.HandleError(err)
	} else {
		affected, err = db_session.Id(event_status.Id).Update(&event_status)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！"})
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c ItsmEventStatus) Delete() revel.Result {
	db_session := c.DbSession

	event_status_id_list := make([]int64, 0)
	c.Params.Bind(&event_status_id_list, "id_list")

	affected, err := db_session.In("id", event_status_id_list).Delete(new(models.EventStatus))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}
