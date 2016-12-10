
package controllers

import (
	"encoding/json"
	"strconv"
	"github.com/revel/revel"

	"matrix/core"
	"matrix/modules/itsm/models"
)

type ItsmApplicationStatus struct {
	*revel.Controller
}

func (c ItsmApplicationStatus) Index() revel.Result {
	return c.RenderTemplate("itsm/application_status/application_status_index.html")
}

func (c ItsmApplicationStatus) ListData() revel.Result {
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

	application_status_list := make([]models.ApplicationStatus, 0, limit)
	err := data_query.Limit(limit, offset).Find(&application_status_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.ApplicationStatus))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  application_status_list,
		Total: count,
	})
}

type ApplicationStatusDetailForm struct {
	ApplicationStatus models.ApplicationStatus `json:"application_status"`
}

func (f *ApplicationStatusDetailForm) IsCreate() bool {
	return f.ApplicationStatus.Id == 0
}

func (f *ApplicationStatusDetailForm) Valid(validation *revel.Validation) bool { 
	validation.Required(f.ApplicationStatus.Name).Message("名称不能为空！")
	if f.ApplicationStatus.Name != "" {
		validation.MinSize(f.ApplicationStatus.Name, 2).Message("名称长度不能小于2！")
	}
	if f.ApplicationStatus.Name != "" {
		validation.MaxSize(f.ApplicationStatus.Name, 20).Message("名称长度不能大于20！")
	}

	return validation.HasErrors() == false
}

func (c ItsmApplicationStatus) DetailData() revel.Result {
	db_session := c.DbSession

	var application_status_id int64
	c.Params.Bind(&application_status_id, "id")

	var application_status models.ApplicationStatus
	if application_status_id != 0 {
		has, err := db_session.Id(application_status_id).Get(&application_status)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的应用状态不存在！"})
		}
	}

	var detail_form ApplicationStatusDetailForm
	detail_form.ApplicationStatus = application_status

	return c.RenderJson(detail_form)
}

func (c ItsmApplicationStatus) Save() revel.Result {
	db_session := c.DbSession

	var detail_form ApplicationStatusDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	application_status := detail_form.ApplicationStatus

	var affected int64
	if detail_form.IsCreate() { 
		affected, err = db_session.Insert(&application_status)
		core.HandleError(err)
	} else { 
		affected, err = db_session.Id(application_status.Id).Update(&application_status)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c ItsmApplicationStatus) Delete() revel.Result {
	db_session := c.DbSession

	application_status_id_list := make([]int64, 0)
	c.Params.Bind(&application_status_id_list, "id_list")

	affected, err := db_session.In("id", application_status_id_list).Delete(new(models.ApplicationStatus))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}

