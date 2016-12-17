package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"strconv"

	"matrix/core"
	"matrix/modules/itsm/models"
)

type ItsmEquipment struct {
	*revel.Controller
}

func (c ItsmEquipment) Index() revel.Result {
	return c.RenderTemplate("itsm/equipment/equipment_index.html")
}

func (c ItsmEquipment) ListData() revel.Result {
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

	equipment_list := make([]models.EquipmentInfo, 0, limit)
	err := data_query.Limit(limit, offset).Find(&equipment_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.EquipmentInfo))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  equipment_list,
		Total: count,
	})
}

type EquipmentDetailForm struct {
	Equipment models.EquipmentInfo `json:"equipment"`
}

func (f *EquipmentDetailForm) IsCreate() bool {
	return f.Equipment.Id == 0
}

func (f *EquipmentDetailForm) Valid(validation *revel.Validation) bool {
	validation.Required(f.Equipment.Code).Message("编码不能为空！")
	if f.Equipment.Code != "" {
		validation.MinSize(f.Equipment.Code, 2).Message("编码长度不能小于2！")
	}
	if f.Equipment.Code != "" {
		validation.MaxSize(f.Equipment.Code, 20).Message("编码长度不能大于20！")
	}

	validation.Required(f.Equipment.Name).Message("名称不能为空！")
	if f.Equipment.Name != "" {
		validation.MinSize(f.Equipment.Name, 2).Message("名称长度不能小于2！")
	}
	if f.Equipment.Name != "" {
		validation.MaxSize(f.Equipment.Name, 20).Message("名称长度不能大于20！")
	}

	validation.Required(f.Equipment.StatusId).Message("状态不能为空！")

	return validation.HasErrors() == false
}

func (c ItsmEquipment) DetailData() revel.Result {
	db_session := c.DbSession

	var equipment_id int64
	c.Params.Bind(&equipment_id, "id")

	var equipment models.EquipmentInfo
	if equipment_id != 0 {
		has, err := db_session.Id(equipment_id).Get(&equipment)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的设备不存在！"})
		}
	}

	var detail_form EquipmentDetailForm
	detail_form.Equipment = equipment

	return c.RenderJson(detail_form)
}

func (c ItsmEquipment) Save() revel.Result {
	db_session := c.DbSession

	var detail_form EquipmentDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	equipment := detail_form.Equipment

	var affected int64
	if detail_form.IsCreate() {
		code_count, err := db_session.Where("code = ?", equipment.Code).Count(&equipment)
		core.HandleError(err)
		if code_count != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，编码已存在！"})
		}

		affected, err = db_session.Insert(&equipment)
		core.HandleError(err)
	} else {
		code_count, err := db_session.Where("id <> ? and code = ?", equipment.Id, equipment.Code).Count(&equipment)
		core.HandleError(err)
		if code_count != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，编码已存在！"})
		}

		affected, err = db_session.Id(equipment.Id).Update(&equipment)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！"})
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c ItsmEquipment) Delete() revel.Result {
	db_session := c.DbSession

	equipment_id_list := make([]int64, 0)
	c.Params.Bind(&equipment_id_list, "id_list")

	affected, err := db_session.In("id", equipment_id_list).Delete(new(models.EquipmentInfo))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}
