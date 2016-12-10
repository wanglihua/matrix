
package controllers

import (
	"encoding/json"
	"strconv"
	"github.com/revel/revel"

	"matrix/core"
	"matrix/modules/itsm/models"
)

type ItsmApplicationType struct {
	*revel.Controller
}

func (c ItsmApplicationType) Index() revel.Result {
	return c.RenderTemplate("itsm/application_type/application_type_index.html")
}

func (c ItsmApplicationType) ListData() revel.Result {
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

	application_type_list := make([]models.ApplicationType, 0, limit)
	err := data_query.Limit(limit, offset).Find(&application_type_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.ApplicationType))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  application_type_list,
		Total: count,
	})
}

type ApplicationTypeDetailForm struct {
	ApplicationType models.ApplicationType `json:"application_type"`
}

func (f *ApplicationTypeDetailForm) IsCreate() bool {
	return f.ApplicationType.Id == 0
}

func (f *ApplicationTypeDetailForm) Valid(validation *revel.Validation) bool { 
	validation.Required(f.ApplicationType.Name).Message("名称不能为空！")
	if f.ApplicationType.Name != "" {
		validation.MinSize(f.ApplicationType.Name, 2).Message("名称长度不能小于2！")
	}
	if f.ApplicationType.Name != "" {
		validation.MaxSize(f.ApplicationType.Name, 20).Message("名称长度不能大于20！")
	}

	return validation.HasErrors() == false
}

func (c ItsmApplicationType) DetailData() revel.Result {
	db_session := c.DbSession

	var application_type_id int64
	c.Params.Bind(&application_type_id, "id")

	var application_type models.ApplicationType
	if application_type_id != 0 {
		has, err := db_session.Id(application_type_id).Get(&application_type)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的应用类型不存在！"})
		}
	}

	var detail_form ApplicationTypeDetailForm
	detail_form.ApplicationType = application_type

	return c.RenderJson(detail_form)
}

func (c ItsmApplicationType) Save() revel.Result {
	db_session := c.DbSession

	var detail_form ApplicationTypeDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	application_type := detail_form.ApplicationType

	var affected int64
	if detail_form.IsCreate() { 
		affected, err = db_session.Insert(&application_type)
		core.HandleError(err)
	} else { 
		affected, err = db_session.Id(application_type.Id).Update(&application_type)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c ItsmApplicationType) Delete() revel.Result {
	db_session := c.DbSession

	application_type_id_list := make([]int64, 0)
	c.Params.Bind(&application_type_id_list, "id_list")

	affected, err := db_session.In("id", application_type_id_list).Delete(new(models.ApplicationType))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}

