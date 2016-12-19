package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/revel/revel"

	"matrix/core"
	"matrix/modules/itsm/models"
)

type ItsmEventApply struct {
	*revel.Controller
}

func (c ItsmEventApply) Index() revel.Result {
	return c.RenderTemplate("itsm/event_apply/event_apply_index.html")
}

func (c ItsmEventApply) ListData() revel.Result {
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

type EventApplyDetailForm struct {
	Event           models.EventInfo         `json:"event"`
	EventTypeList   []models.EventTypeInfo   `json:"event_type_list"`
	ServiceAreaList []models.ServiceAreaInfo `json:"service_area_list"`
}

func (f *EventApplyDetailForm) IsCreate() bool {
	return f.Event.Id == 0
}

func (f *EventApplyDetailForm) Valid(validation *revel.Validation) bool {
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

func (c ItsmEventApply) DetailData() revel.Result {
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

	var event_type_list = make([]models.EventTypeInfo, 0)
	err := db_session.Find(&event_type_list)
	core.HandleError(err)

	var service_area_list = make([]models.ServiceAreaInfo, 0)
	err = db_session.Find(&service_area_list)
	core.HandleError(err)

	var detail_form EventApplyDetailForm
	detail_form.Event = event
	detail_form.EventTypeList = event_type_list
	detail_form.ServiceAreaList = service_area_list

	return c.RenderJson(detail_form)
}

func (c ItsmEventApply) Save() revel.Result {
	db_session := c.DbSession

	var detail_form EventApplyDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	event_in_ui := detail_form.Event

	var affected int64
	if detail_form.IsCreate() {
		//todo: 按规则生成 code
		affected, err = db_session.Insert(&event_in_ui)
		core.HandleError(err)
	} else {
		var event_in_db models.EventInfo
		_, err := db_session.Id(event_in_ui.Id).Get(&event_in_db)
		core.HandleError(err)

		//todo: 更新数据库各字段
		event_in_db.Description = event_in_ui.Description

		affected, err = db_session.Id(event_in_ui.Id).Update(&event_in_db)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！"})
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c ItsmEventApply) Delete() revel.Result {
	db_session := c.DbSession

	event_id_list := make([]int64, 0)
	c.Params.Bind(&event_id_list, "id_list")

	affected, err := db_session.In("id", event_id_list).Delete(new(models.EventInfo))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}
