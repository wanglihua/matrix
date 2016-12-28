package controllers

import (
	"encoding/json"

	"github.com/revel/revel"

	"fmt"
	"matrix/core"
	auth_models "matrix/modules/auth/models"
	"matrix/modules/itsm/models"
	"matrix/modules/itsm/service/engineer_service"
)

type ItsmEventGrab struct {
	*revel.Controller
}

func (c ItsmEventGrab) Index() revel.Result {
	db_session := c.DbSession
	engineer_id := engineer_service.GetLoginedEngineerId(c.Session, db_session)
	if engineer_id == 0 {
		return c.NotFound("您不是工程师，不能抢单！")
	}

	return c.RenderTemplate("itsm/event_grab/event_grab_index.html")
}

type EventGrabViewItem struct {
	models.EventInfo       `xorm:"extends" json:"evt"`
	models.EventTypeInfo   `xorm:"extends" json:"et"`
	models.EventStatusInfo `xorm:"extends" json:"es"`
}

func (c ItsmEventGrab) ListData() revel.Result {
	db_session := c.DbSession

	engineer_id := engineer_service.GetLoginedEngineerId(c.Session, db_session)
	if engineer_id == 0 {
		return c.RenderJson(core.JsonResult{Success: false, Message: "您不是工程师，不能抢单！"})
	}

	filter, order, offset, limit := core.GetGridRequestParam(c.Request)
	/*
		SELECT evt.*, et.*, es.* FROM itsm_event evt
		  INNER JOIN itsm_event_type et on evt.type_id = et.id
		  INNER JOIN itsm_event_status es on evt.status_id = es.id
	*/

	query := db_session.
	Select("evt.*, et.*, es.*").
		Table("itsm_event").Alias("evt").
		Join("inner", []string{"itsm_event_type", "et"}, "evt.type_id = et.id").
		Join("inner", []string{"itsm_event_status", "es"}, "evt.status_id = es.id").
		Where(filter)

	//query extra filter here

	event_cols := models.EventCols
	query = query.Where(fmt.Sprintf("%s = ?", event_cols.StatusId), models.Event_Status_TBZ_Id)

	data_query := *query
	if order != "" {
		data_query = *data_query.OrderBy(order)
	} else {
		data_query = *data_query.Asc("evt.id")
	}

	event_list := make([]EventGrabViewItem, 0, limit)
	err := data_query.Limit(limit, offset).Find(&event_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(EventGrabViewItem))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  event_list,
		Total: count,
	})
}

type EventGrabDetailForm struct {
	Event           models.EventInfo                `json:"event"`
	ApplyUser       auth_models.UserInfo            `json:"apply_user"`
	EventType       models.EventTypeInfo            `json:"event_type"`
	ServiceArea     models.ServiceAreaInfo          `json:"service_area"`
	ApplyDepartment models.EventApplyDepartmentInfo `json:"apply_department"`
}

func (c ItsmEventGrab) DetailData() revel.Result {
	db_session := c.DbSession

	engineer_id := engineer_service.GetLoginedEngineerId(c.Session, db_session)
	if engineer_id == 0 {
		return c.RenderJson(core.JsonResult{Success: false, Message: "您不是工程师，不能抢单！"})
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

func (c ItsmEventGrab) Save() revel.Result {
	db_session := c.DbSession

	var detail_form EventGrabDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	engineer_id := engineer_service.GetLoginedEngineerId(c.Session, db_session)
	if engineer_id == 0 {
		return c.RenderJson(core.JsonResult{Success: false, Message: "您不是工程师，抢单失败！"})
	}

	event_in_ui := detail_form.Event

	var event_in_db models.EventInfo
	_, err = db_session.Id(event_in_ui.Id).Get(&event_in_db)
	core.HandleError(err)

	if event_in_db.StatusId != models.Event_Status_TBZ_Id {
		return c.RenderJson(core.JsonResult{Success: false, Message: "当前事件不是提报状态，不能抢单！"})
	}

	//更新工程师字段和事件状态字段
	event_in_db.EngineerId = core.NewNullInt(engineer_id, true)
	event_in_db.StatusId = models.Event_Status_YPG_Id // 已派工

	var affected int64
	event_cols := models.EventCols
	affected, err = db_session.Id(event_in_ui.Id).Cols(event_cols.EngineerId, event_cols.StatusId).Update(&event_in_db)
	core.HandleError(err)

	if affected == 0 {
		return c.RenderJson(core.JsonResult{Success: false, Message: "抢单失败，请重试！"})
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: "抢单成功!"})
}
