package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"strconv"

	"matrix/core"
	"matrix/modules/itsm/models"
)

type ItsmApplication struct {
	*revel.Controller
}

func (c ItsmApplication) Index() revel.Result {
	return c.RenderTemplate("itsm/application/application_index.html")
}

func (c ItsmApplication) ListData() revel.Result {
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

	application_list := make([]models.ApplicationInfo, 0, limit)
	err := data_query.Limit(limit, offset).Find(&application_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.ApplicationInfo))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  application_list,
		Total: count,
	})
}

type ApplicationDetailForm struct {
	Application models.ApplicationInfo `json:"application"`
}

func (f *ApplicationDetailForm) IsCreate() bool {
	return f.Application.Id == 0
}

func (f *ApplicationDetailForm) Valid(validation *revel.Validation) bool {
	validation.Required(f.Application.Code).Message("编码不能为空！")
	if f.Application.Code != "" {
		validation.MinSize(f.Application.Code, 2).Message("编码长度不能小于2！")
	}
	if f.Application.Code != "" {
		validation.MaxSize(f.Application.Code, 20).Message("编码长度不能大于20！")
	}

	validation.Required(f.Application.Name).Message("名称不能为空！")
	if f.Application.Name != "" {
		validation.MinSize(f.Application.Name, 2).Message("名称长度不能小于2！")
	}
	if f.Application.Name != "" {
		validation.MaxSize(f.Application.Name, 20).Message("名称长度不能大于20！")
	}

	validation.Required(f.Application.StatusId).Message("状态不能为空！")

	return validation.HasErrors() == false
}

func (c ItsmApplication) DetailData() revel.Result {
	db_session := c.DbSession

	var application_id int64
	c.Params.Bind(&application_id, "id")

	var application models.ApplicationInfo
	if application_id != 0 {
		has, err := db_session.Id(application_id).Get(&application)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的应用不存在！"})
		}
	}

	var detail_form ApplicationDetailForm
	detail_form.Application = application

	return c.RenderJson(detail_form)
}

func (c ItsmApplication) Save() revel.Result {
	db_session := c.DbSession

	var detail_form ApplicationDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	application := detail_form.Application

	var affected int64
	if detail_form.IsCreate() {
		code_count, err := db_session.Where("code = ?", application.Code).Count(&application)
		core.HandleError(err)
		if code_count != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，编码已存在！"})
		}

		affected, err = db_session.Insert(&application)
		core.HandleError(err)
	} else {
		code_count, err := db_session.Where("id <> ? and code = ?", application.Id, application.Code).Count(&application)
		core.HandleError(err)
		if code_count != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，编码已存在！"})
		}

		affected, err = db_session.Id(application.Id).Update(&application)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！"})
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c ItsmApplication) Delete() revel.Result {
	db_session := c.DbSession

	application_id_list := make([]int64, 0)
	c.Params.Bind(&application_id_list, "id_list")

	affected, err := db_session.In("id", application_id_list).Delete(new(models.ApplicationInfo))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}
