
package controllers

import (
	"encoding/json"
	"strconv"
	"github.com/revel/revel"

	"matrix/core"
	"matrix/modules/erp/models"
)

type ErpEubInsuranceType struct {
	*revel.Controller
}

func (c ErpEubInsuranceType) Index() revel.Result {
	return c.RenderTemplate("erp/eub_insurance_type/eub_insurance_type_index.html")
}

func (c ErpEubInsuranceType) ListData() revel.Result {
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

	eub_insurance_type_list := make([]models.EubInsuranceTypeInfo, 0, limit)
	err := data_query.Limit(limit, offset).Find(&eub_insurance_type_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.EubInsuranceTypeInfo))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  eub_insurance_type_list,
		Total: count,
	})
}

type EubInsuranceTypeDetailForm struct {
	EubInsuranceType models.EubInsuranceTypeInfo `json:"eub_insurance_type"`
}

func (f *EubInsuranceTypeDetailForm) IsCreate() bool {
	return f.EubInsuranceType.Id == 0
}

func (f *EubInsuranceTypeDetailForm) Valid(validation *revel.Validation) bool { 
	validation.Required(f.EubInsuranceType.Name).Message("名称不能为空！")
	if f.EubInsuranceType.Name != "" {
		validation.MinSize(f.EubInsuranceType.Name, 2).Message("名称长度不能小于2！")
	}
	if f.EubInsuranceType.Name != "" {
		validation.MaxSize(f.EubInsuranceType.Name, 100).Message("名称长度不能大于100！")
	}

	if f.EubInsuranceType.Desc.Valid == true {
		validation.MinSize(f.EubInsuranceType.Desc.String, 2).Message("描述长度不能小于2！")
	}
	if f.EubInsuranceType.Desc.Valid == true {
		validation.MaxSize(f.EubInsuranceType.Desc.String, 400).Message("描述长度不能大于400！")
	}

	return validation.HasErrors() == false
}

func (c ErpEubInsuranceType) DetailData() revel.Result {
	db_session := c.DbSession

	var eub_insurance_type_id int64
	c.Params.Bind(&eub_insurance_type_id, "id")

	var eub_insurance_type models.EubInsuranceTypeInfo
	if eub_insurance_type_id != 0 {
		has, err := db_session.Id(eub_insurance_type_id).Get(&eub_insurance_type)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的EUB保险类型不存在！"})
		}
	}

	var detail_form EubInsuranceTypeDetailForm
	detail_form.EubInsuranceType = eub_insurance_type

	return c.RenderJson(detail_form)
}

func (c ErpEubInsuranceType) Save() revel.Result {
	db_session := c.DbSession

	var detail_form EubInsuranceTypeDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	eub_insurance_type := detail_form.EubInsuranceType

	var affected int64
	if detail_form.IsCreate() { 
		affected, err = db_session.Insert(&eub_insurance_type)
		core.HandleError(err)
	} else { 
		affected, err = db_session.Id(eub_insurance_type.Id).Update(&eub_insurance_type)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c ErpEubInsuranceType) Delete() revel.Result {
	db_session := c.DbSession

	eub_insurance_type_id_list := make([]int64, 0)
	c.Params.Bind(&eub_insurance_type_id_list, "id_list")

	affected, err := db_session.In("id", eub_insurance_type_id_list).Delete(new(models.EubInsuranceTypeInfo))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}

