
package controllers

import (
	"encoding/json"
	"strconv"
	"github.com/revel/revel"

	"matrix/core"
	"matrix/modules/itsm/models"
)

type ItsmEventLog struct {
	*revel.Controller
}

func (c ItsmEventLog) Index() revel.Result {
	return c.RenderTemplate("itsm/event_log/event_log_index.html")
}

func (c ItsmEventLog) ListData() revel.Result {
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

	event_log_list := make([]models.EventLog, 0, limit)
	err := data_query.Limit(limit, offset).Find(&event_log_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.EventLog))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  event_log_list,
		Total: count,
	})
}

type EventLogDetailForm struct {
	EventLog models.EventLog `json:"log"`
}

func (f *EventLogDetailForm) IsCreate() bool {
	return f.EventLog.Id == 0
}

func (f *EventLogDetailForm) Valid(validation *revel.Validation) bool { 
	validation.Required(f.EventLog.EventId).Message("事件不能为空！")

	validation.Required(f.EventLog.Reason).Message("事由不能为空！")


	return validation.HasErrors() == false
}

func (c ItsmEventLog) DetailData() revel.Result {
	db_session := c.DbSession

	var event_log_id int64
	c.Params.Bind(&event_log_id, "id")

	var event_log models.EventLog
	if event_log_id != 0 {
		has, err := db_session.Id(event_log_id).Get(&event_log)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的事件日志不存在！"})
		}
	}

	var detail_form EventLogDetailForm
	detail_form.EventLog = event_log

	return c.RenderJson(detail_form)
}

func (c ItsmEventLog) Save() revel.Result {
	db_session := c.DbSession

	var detail_form EventLogDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	event_log := detail_form.EventLog

	var affected int64
	if detail_form.IsCreate() { 
		affected, err = db_session.Insert(&event_log)
		core.HandleError(err)
	} else { 
		affected, err = db_session.Id(event_log.Id).Update(&event_log)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c ItsmEventLog) Delete() revel.Result {
	db_session := c.DbSession

	event_log_id_list := make([]int64, 0)
	c.Params.Bind(&event_log_id_list, "id_list")

	affected, err := db_session.In("id", event_log_id_list).Delete(new(models.EventLog))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}

