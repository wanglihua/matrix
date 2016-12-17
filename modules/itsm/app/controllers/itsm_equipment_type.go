package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"strconv"

	"matrix/core"
	"matrix/modules/itsm/models"
)

type ItsmEquipmentType struct {
	*revel.Controller
}

func (c ItsmEquipmentType) Index() revel.Result {
	return c.RenderTemplate("itsm/equipment_type/equipment_type_index.html")
}

func (c ItsmEquipmentType) ListData() revel.Result {
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

	equipment_type_list := make([]models.EquipmentTypeInfo, 0, limit)
	err := data_query.Limit(limit, offset).Find(&equipment_type_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.EquipmentTypeInfo))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  equipment_type_list,
		Total: count,
	})
}

type EquipmentTypeDetailForm struct {
	EquipmentType models.EquipmentTypeInfo `json:"equipment_type"`
}

func (f *EquipmentTypeDetailForm) IsCreate() bool {
	return f.EquipmentType.Id == 0
}

func (f *EquipmentTypeDetailForm) Valid(validation *revel.Validation) bool {
	validation.Required(f.EquipmentType.Name).Message("名称不能为空！")
	if f.EquipmentType.Name != "" {
		validation.MinSize(f.EquipmentType.Name, 2).Message("名称长度不能小于2！")
	}
	if f.EquipmentType.Name != "" {
		validation.MaxSize(f.EquipmentType.Name, 20).Message("名称长度不能大于20！")
	}

	return validation.HasErrors() == false
}

func (c ItsmEquipmentType) DetailData() revel.Result {
	db_session := c.DbSession

	var equipment_type_id int64
	c.Params.Bind(&equipment_type_id, "id")

	var equipment_type models.EquipmentTypeInfo
	if equipment_type_id != 0 {
		has, err := db_session.Id(equipment_type_id).Get(&equipment_type)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的设备类型不存在！"})
		}
	}

	var detail_form EquipmentTypeDetailForm
	detail_form.EquipmentType = equipment_type

	return c.RenderJson(detail_form)
}

func (c ItsmEquipmentType) Save() revel.Result {
	db_session := c.DbSession

	var detail_form EquipmentTypeDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	equipment_type := detail_form.EquipmentType

	var affected int64
	if detail_form.IsCreate() {
		affected, err = db_session.Insert(&equipment_type)
		core.HandleError(err)
	} else {
		affected, err = db_session.Id(equipment_type.Id).Update(&equipment_type)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！"})
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c ItsmEquipmentType) Delete() revel.Result {
	db_session := c.DbSession

	equipment_type_id_list := make([]int64, 0)
	c.Params.Bind(&equipment_type_id_list, "id_list")

	affected, err := db_session.In("id", equipment_type_id_list).Delete(new(models.EquipmentTypeInfo))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}
