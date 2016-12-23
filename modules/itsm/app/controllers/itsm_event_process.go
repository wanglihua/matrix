package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/revel/revel"

	"matrix/core"
	"matrix/modules/itsm/models"
	"matrix/modules/itsm/service/engineer_service"
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

func (c ItsmEventProcess) ListData() revel.Result {
	db_session := c.DbSession

	filter, order, offset, limit := core.GetGridRequestParam(c.Request)
	query := db_session.Where(filter)

	engineer_id := engineer_service.GetLoginedEngineerId(c.Session, db_session)
	if engineer_id == 0 {
		return c.RenderJson(core.JsonResult{Success: false, Message: "您不是工程师，不能处理事件！"})
	}

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

type EventProcessDetailForm struct {
	Event models.EventInfo `json:"event"`
}

func (f *EventProcessDetailForm) IsCreate() bool {
	return f.Event.Id == 0
}

func (f *EventProcessDetailForm) Valid(validation *revel.Validation) bool {
	validation.Required(f.Event.Code).Message("编号不能为空！")
	if f.Event.Code != "" {
		validation.MinSize(f.Event.Code, 1).Message("编号长度不能小于1！")
	}
	if f.Event.Code != "" {
		validation.MaxSize(f.Event.Code, 10).Message("编号长度不能大于10！")
	}

	validation.Required(f.Event.TypeId).Message("类型不能为空！")

	validation.Required(f.Event.PriorityId).Message("优先级不能为空！")

	validation.Required(f.Event.ApplyUserId).Message("提报用户不能为空！")

	validation.Required(f.Event.AreaId).Message("服务区域不能为空！")

	validation.Required(f.Event.Location).Message("地点不能为空！")

	validation.Required(f.Event.StatusId).Message("状态不能为空！")

	return validation.HasErrors() == false
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
	if event_id != 0 {
		has, err := db_session.Id(event_id).Get(&event)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的事件不存在！"})
		}
	}

	var detail_form EventProcessDetailForm
	detail_form.Event = event

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
