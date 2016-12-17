package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/revel/revel"

	"matrix/core"
	"matrix/modules/itsm/models"
)

type ItsmEventManage struct {
	*revel.Controller
}

func (c ItsmEventManage) Index() revel.Result {
	return c.RenderTemplate("itsm/event_manage/event_manage_index.html")
}

func (c ItsmEventManage) ListData() revel.Result {
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

	event_list := make([]models.EventInfo, 0, limit)
	err := data_query.Limit(limit, offset).Find(&event_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.EventInfo))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  event_list,
		Total: count,
	})
}

type EventManageDetailForm struct {
	Event models.EventInfo `json:"event"`
}

func (f *EventManageDetailForm) IsCreate() bool {
	return f.Event.Id == 0
}

func (f *EventManageDetailForm) Valid(validation *revel.Validation) bool {
	validation.Required(f.Event.Code).Message("编号不能为空！")
	if f.Event.Code != "" {
		validation.MinSize(f.Event.Code, 1).Message("编号长度不能小于1！")
	}
	if f.Event.Code != "" {
		validation.MaxSize(f.Event.Code, 10).Message("编号长度不能大于10！")
	}

	validation.Required(f.Event.TypeId).Message("类型不能为空！")

	validation.Required(f.Event.PriorityId).Message("优先级不能为空！")

	validation.Required(f.Event.RequestUserId).Message("请求用户不能为空！")

	validation.Required(f.Event.AreaId).Message("服务区域不能为空！")

	validation.Required(f.Event.Location).Message("地点不能为空！")

	validation.Required(f.Event.StatusId).Message("状态不能为空！")

	return validation.HasErrors() == false
}

func (c ItsmEventManage) DetailData() revel.Result {
	db_session := c.DbSession

	var event_id int64
	c.Params.Bind(&event_id, "id")

	var event models.EventInfo
	if event_id != 0 {
		has, err := db_session.Id(event_id).Get(&event)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的事件不存在！"})
		}
	}

	var detail_form EventManageDetailForm
	detail_form.Event = event

	return c.RenderJson(detail_form)
}

func (c ItsmEventManage) Save() revel.Result {
	db_session := c.DbSession

	var detail_form EventManageDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	event := detail_form.Event

	var affected int64
	if detail_form.IsCreate() {
		affected, err = db_session.Insert(&event)
		core.HandleError(err)
	} else {
		affected, err = db_session.Id(event.Id).Update(&event)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！"})
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c ItsmEventManage) Delete() revel.Result {
	db_session := c.DbSession

	event_id_list := make([]int64, 0)
	c.Params.Bind(&event_id_list, "id_list")

	affected, err := db_session.In("id", event_id_list).Delete(new(models.EventInfo))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}
