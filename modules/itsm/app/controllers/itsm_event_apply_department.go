
package controllers

import (
	"encoding/json"
	"strconv"
	"github.com/revel/revel"

	"matrix/core"
	"matrix/modules/itsm/models"
)

type ItsmEventApplyDepartment struct {
	*revel.Controller
}

func (c ItsmEventApplyDepartment) Index() revel.Result {
	return c.RenderTemplate("itsm/event_apply_department/event_apply_department_index.html")
}

func (c ItsmEventApplyDepartment) ListData() revel.Result {
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

	event_apply_department_list := make([]models.EventApplyDepartmentInfo, 0, limit)
	err := data_query.Limit(limit, offset).Find(&event_apply_department_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.EventApplyDepartmentInfo))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  event_apply_department_list,
		Total: count,
	})
}

type EventApplyDepartmentDetailForm struct {
	EventApplyDepartment models.EventApplyDepartmentInfo `json:"apply_department"`
}

func (f *EventApplyDepartmentDetailForm) IsCreate() bool {
	return f.EventApplyDepartment.Id == 0
}

func (f *EventApplyDepartmentDetailForm) Valid(validation *revel.Validation) bool { 
	validation.Required(f.EventApplyDepartment.Name).Message("名称不能为空！")
	if f.EventApplyDepartment.Name != "" {
		validation.MinSize(f.EventApplyDepartment.Name, 2).Message("名称长度不能小于2！")
	}
	if f.EventApplyDepartment.Name != "" {
		validation.MaxSize(f.EventApplyDepartment.Name, 20).Message("名称长度不能大于20！")
	}

	if f.EventApplyDepartment.Description.Valid == true {
		validation.MaxSize(f.EventApplyDepartment.Description.String, 30).Message("描述长度不能大于30！")
	}

	return validation.HasErrors() == false
}

func (c ItsmEventApplyDepartment) DetailData() revel.Result {
	db_session := c.DbSession

	var event_apply_department_id int64
	c.Params.Bind(&event_apply_department_id, "id")

	var event_apply_department models.EventApplyDepartmentInfo
	if event_apply_department_id != 0 {
		has, err := db_session.Id(event_apply_department_id).Get(&event_apply_department)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的事件提报部门不存在！"})
		}
	}

	var detail_form EventApplyDepartmentDetailForm
	detail_form.EventApplyDepartment = event_apply_department

	return c.RenderJson(detail_form)
}

func (c ItsmEventApplyDepartment) Save() revel.Result {
	db_session := c.DbSession

	var detail_form EventApplyDepartmentDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	event_apply_department := detail_form.EventApplyDepartment

	var affected int64
	if detail_form.IsCreate() { 
		affected, err = db_session.Insert(&event_apply_department)
		core.HandleError(err)
	} else { 
		affected, err = db_session.Id(event_apply_department.Id).Update(&event_apply_department)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c ItsmEventApplyDepartment) Delete() revel.Result {
	db_session := c.DbSession

	event_apply_department_id_list := make([]int64, 0)
	c.Params.Bind(&event_apply_department_id_list, "id_list")

	affected, err := db_session.In("id", event_apply_department_id_list).Delete(new(models.EventApplyDepartmentInfo))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}

