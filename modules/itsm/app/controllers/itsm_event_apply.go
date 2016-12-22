package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/revel/revel"

	"matrix/core"
	"matrix/modules/itsm/models"
	auth_models "matrix/modules/auth/models"
	"matrix/modules/itsm/service/entity_code_service"
	"fmt"
)

type ItsmEventApply struct {
	*revel.Controller
}

func (c ItsmEventApply) Index() revel.Result {
	return c.RenderTemplate("itsm/event_apply/event_apply_index.html")
}

type  EventView struct {
	models.EventInfo        `xorm:"extends" json:"evt"`
	models.EventTypeInfo    `xorm:"extends" json:"et"`
	models.EngineerInfo     `xorm:"extends" json:"egr"`
	auth_models.UserInfo    `xorm:"extends" json:"egru"`
	models.EventStatusInfo  `xorm:"extends" json:"es"`
}

func (c ItsmEventApply) ListData() revel.Result {
	db_session := c.DbSession

	filter, order, offset, limit := core.GetGridRequestParam(c.Request)
	/*
		SELECT evt.*, egr.*, egru.*, es.* FROM itsm_event evt
		  INNER JOIN itsm_event_type et on evt.type_id = et.id
		  INNER JOIN itsm_engineer egr on evt.engineer_id = egr.id
		  INNER JOIN auth_user egru on egr.user_id = egru.id
		  INNER JOIN itsm_event_status es on evt.status_id = es.id
	 */

	query := db_session.
	Select("evt.*, et.*, egr.*, egru.*, es.*").
		Table("itsm_event").Alias("evt").
		Join("inner", []string{"itsm_event_type", "et"}, "evt.type_id = et.id").
		Join("inner", []string{"itsm_engineer", "egr"}, "evt.engineer_id = egr.id").
		Join("inner", []string{"auth_user", "egru"}, "egr.user_id = egru.id").
		Join("inner", []string{"itsm_event_status", "es"}, "evt.status_id = es.id").
		Where(filter)

	//query extra filter here

	data_query := *query
	if order != "" {
		data_query = *data_query.OrderBy(order)
	} else {
		data_query = *data_query.Asc("evt.id")
	}

	event_list := make([]EventView, 0, limit)
	err := data_query.Limit(limit, offset).Find(&event_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(EventView))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  event_list,
		Total: count,
	})
}

type EngineerSelectItem struct {
	Id       int64 `xorm:"'id'" json:"id"`
	NickName string `xorm:"'nick_name'" json:"nick_name"`
}

type EventApplyDetailForm struct {
	Event               models.EventInfo                  `json:"event"`
	ApplyUser           auth_models.UserInfo              `json:"apply_user"`
	EventTypeList       []models.EventTypeInfo            `json:"event_type_list"`
	ServiceAreaList     []models.ServiceAreaInfo          `json:"service_area_list"`
	EngineerList        []EngineerSelectItem              `json:"engineer_list"`
	ApplyDepartmentList []models.EventApplyDepartmentInfo `json:"apply_department_list"`
}

func (f *EventApplyDetailForm) IsCreate() bool {
	return f.Event.Id == 0
}

func (f *EventApplyDetailForm) Valid(validation *revel.Validation) bool {
	validation.Required(f.Event.TypeId).Message("类型不能为空！")
	validation.Required(f.Event.ApplyDepartmentId).Message("提报部门不能为空！")
	validation.Required(f.Event.ApplyUserId).Message("提报用户不能为空！")
	validation.Required(f.Event.Contact).Message("联系电话不能为空！")
	validation.Required(f.Event.AreaId).Message("服务区域不能为空！")
	validation.Required(f.Event.Location).Message("地点不能为空！")
	validation.Required(f.Event.Description).Message("事件描述不能为空！")

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

	login_user := core.GetLoginUser(c.Session)
	login_user_id := login_user.UserId
	event.ApplyUserId = login_user_id

	var apply_user auth_models.UserInfo
	_, err = db_session.Id(login_user_id).Get(&apply_user)
	core.HandleError(err)

	var engineer_select_item_list = make([]EngineerSelectItem, 0)
	sql := `SELECT u.id, u.nick_name from auth_user u WHERE u.id IN (SELECT e.user_id from itsm_engineer e)`
	err = db_session.Sql(sql).Find(&engineer_select_item_list)
	core.HandleError(err)

	var apply_department_list = make([]models.EventApplyDepartmentInfo, 0)
	err = db_session.Find(&apply_department_list);
	core.HandleError(err)

	var detail_form EventApplyDetailForm
	detail_form.Event = event
	detail_form.EventTypeList = event_type_list
	detail_form.ServiceAreaList = service_area_list
	detail_form.ApplyUser = apply_user
	detail_form.EngineerList = engineer_select_item_list
	detail_form.ApplyDepartmentList = apply_department_list

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

	login_user := core.GetLoginUser(c.Session)
	login_user_id := login_user.UserId

	event_in_ui := detail_form.Event

	var affected int64
	if detail_form.IsCreate() {
		//按规则生成 code
		serial := entity_code_service.GetNextSerial(db_session, "event")
		event_in_ui.Code = fmt.Sprint(serial)
		event_in_ui.StatusId = models.Event_Status_TBZ_Id //提报中

		affected, err = db_session.Insert(&event_in_ui)
		core.HandleError(err)
	} else {
		var event_in_db models.EventInfo
		_, err := db_session.Id(event_in_ui.Id).Get(&event_in_db)
		core.HandleError(err)

		//更新数据库各字段
		event_in_db.Code = event_in_ui.Code
		event_in_db.ApplyDepartmentId = event_in_ui.ApplyDepartmentId
		event_in_db.ApplyUserId = login_user_id
		event_in_db.Contact = event_in_ui.Contact
		event_in_db.TypeId = event_in_ui.TypeId
		event_in_db.EngineerId = event_in_ui.EngineerId
		event_in_db.AreaId = event_in_ui.AreaId
		event_in_db.Location = event_in_ui.Location
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
