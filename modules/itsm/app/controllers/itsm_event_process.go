package controllers

import (
	"encoding/json"
	"strconv"
	"github.com/revel/revel"
	"matrix/core"
	"matrix/modules/itsm/models"
	"matrix/modules/itsm/service/engineer_service"
	auth_models "matrix/modules/auth/models"
)

type ItsmEventProcess struct {
	*revel.Controller
}

func (c ItsmEventProcess) Index() revel.Result {
	db_session := c.DbSession
	engineer_id := engineer_service.GetLoginedEngineerId(c.Session, db_session)
	if engineer_id == 0 {
		return c.NotFound("您不是工程师，不能处理事件！")
	}

	return c.RenderTemplate("itsm/event_process/event_process_index.html")
}

type EventProcessView struct {
	models.EventInfo       `xorm:"extends" json:"evt"`
	models.EventTypeInfo   `xorm:"extends" json:"et"`
	auth_models.UserInfo   `xorm:"extends" json:"eu"`
	models.EventStatusInfo `xorm:"extends" json:"es"`
}

func (c ItsmEventProcess) ListData() revel.Result {
	db_session := c.DbSession

	engineer_id := engineer_service.GetLoginedEngineerId(c.Session, db_session)
	if engineer_id == 0 {
		return c.RenderJson(core.JsonResult{Success: false, Message: "您不是工程师，不能处理事件！"})
	}

	filter, order, offset, limit := core.GetGridRequestParam(c.Request)
	/*
		SELECT evt.*, et.*, eu.*, es.* FROM itsm_event evt
		  INNER JOIN itsm_event_type et on evt.type_id = et.id
		  INNER JOIN auth_user eu on evt.apply_user_id = eu.id
		  INNER JOIN itsm_event_status es on evt.status_id = es.id
	*/

	query := db_session.
	Select("evt.*, et.*, eu.*, es.*").
		Table("itsm_event").Alias("evt").
		Join("inner", []string{"itsm_event_type", "et"}, "evt.type_id = et.id").
		Join("inner", []string{"auth_user", "eu"}, "evt.apply_user_id = eu.id").
		Join("inner", []string{"itsm_event_status", "es"}, "evt.status_id = es.id").
		Where(filter)

	//只取当前工程师的事件
	query = query.Where("evt.engineer_id = ?", engineer_id)

	//query extra filter here

	data_query := *query
	if order != "" {
		data_query = *data_query.OrderBy(order)
	} else {
		data_query = *data_query.Asc("evt.id")
	}

	event_list := make([]EventProcessView, 0, limit)
	err := data_query.Limit(limit, offset).Find(&event_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(EventProcessView))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  event_list,
		Total: count,
	})
}

type EventProcessDetailForm struct {
	Event           models.EventInfo                `json:"event"`
	ApplyUser       auth_models.UserInfo            `json:"apply_user"`
	EventType       models.EventTypeInfo            `json:"event_type"`
	ServiceArea     models.ServiceAreaInfo          `json:"service_area"`
	ApplyDepartment models.EventApplyDepartmentInfo `json:"apply_department"`
}

func (c ItsmEventProcess) DetailData() revel.Result {
	db_session := c.DbSession

	engineer_id := engineer_service.GetLoginedEngineerId(c.Session, db_session)
	if engineer_id == 0 {
		return c.RenderJson(core.JsonResult{Success: false, Message: "您不是工程师，不能处理事件！"})
	}

	var event_id int64
	c.Params.Bind(&event_id, "id")

	var event models.EventInfo
	has, err := db_session.Id(event_id).Get(&event)
	core.HandleError(err)
	if has == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: "指定的事件不存在！"})
	}

	var event_type models.EventTypeInfo
	_, err = db_session.Id(event.TypeId).Get(&event_type)
	core.HandleError(err)

	var service_area models.ServiceAreaInfo
	_, err = db_session.Id(event.AreaId).Get(&service_area)
	core.HandleError(err)

	var apply_user auth_models.UserInfo
	_, err = db_session.Id(event.ApplyUserId).Get(&apply_user)
	core.HandleError(err)

	var apply_department models.EventApplyDepartmentInfo
	_, err = db_session.Id(event.ApplyDepartmentId).Get(&apply_department)
	core.HandleError(err)

	var detail_form EventGrabDetailForm
	detail_form.Event = event
	detail_form.EventType = event_type
	detail_form.ServiceArea = service_area
	detail_form.ApplyUser = apply_user
	detail_form.ApplyDepartment = apply_department

	return c.RenderJson(detail_form)
}

func (c ItsmEventProcess) Save() revel.Result {
	db_session := c.DbSession

	engineer_id := engineer_service.GetLoginedEngineerId(c.Session, db_session)
	if engineer_id == 0 {
		return c.RenderJson(core.JsonResult{Success: false, Message: "您不是工程师，不能处理事件！"})
	}

	var detail_form EventProcessDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	event_in_ui := detail_form.Event

	//只是保存，除了 Solution 字段外，不用更新其它字段

	event_cols := models.EventCols
	affected, err := db_session.Id(event_in_ui.Id).Cols(event_cols.Solution).Update(&event_in_ui)
	core.HandleError(err)

	if affected == 0 {
		return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！"})
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}
