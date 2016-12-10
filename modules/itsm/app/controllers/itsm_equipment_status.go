
package controllers

import (
	"encoding/json"
	"strconv"
	"github.com/revel/revel"

	"matrix/core"
	"matrix/modules/itsm/models"
)

type ItsmEquipmentStatus struct {
	*revel.Controller
}

func (c ItsmEquipmentStatus) Index() revel.Result {
	return c.RenderTemplate("itsm/equipment_status/equipment_status_index.html")
}

func (c ItsmEquipmentStatus) ListData() revel.Result {
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

	equipment_status_list := make([]models.EquipmentStatus, 0, limit)
	err := data_query.Limit(limit, offset).Find(&equipment_status_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.EquipmentStatus))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  equipment_status_list,
		Total: count,
	})
}

type EquipmentStatusDetailForm struct {
	EquipmentStatus models.EquipmentStatus `json:"equipment_status"`
}

func (f *EquipmentStatusDetailForm) IsCreate() bool {
	return f.EquipmentStatus.Id == 0
}

func (f *EquipmentStatusDetailForm) Valid(validation *revel.Validation) bool { 
	validation.Required(f.EquipmentStatus.Name).Message("名称不能为空！")
	if f.EquipmentStatus.Name != "" {
		validation.MinSize(f.EquipmentStatus.Name, 2).Message("名称长度不能小于2！")
	}
	if f.EquipmentStatus.Name != "" {
		validation.MaxSize(f.EquipmentStatus.Name, 20).Message("名称长度不能大于20！")
	}

	return validation.HasErrors() == false
}

func (c ItsmEquipmentStatus) DetailData() revel.Result {
	db_session := c.DbSession

	var equipment_status_id int64
	c.Params.Bind(&equipment_status_id, "id")

	var equipment_status models.EquipmentStatus
	if equipment_status_id != 0 {
		has, err := db_session.Id(equipment_status_id).Get(&equipment_status)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的设备状态不存在！"})
		}
	}

	var detail_form EquipmentStatusDetailForm
	detail_form.EquipmentStatus = equipment_status

	return c.RenderJson(detail_form)
}

func (c ItsmEquipmentStatus) Save() revel.Result {
	db_session := c.DbSession

	var detail_form EquipmentStatusDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	equipment_status := detail_form.EquipmentStatus

	var affected int64
	if detail_form.IsCreate() { 
		affected, err = db_session.Insert(&equipment_status)
		core.HandleError(err)
	} else { 
		affected, err = db_session.Id(equipment_status.Id).Update(&equipment_status)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c ItsmEquipmentStatus) Delete() revel.Result {
	db_session := c.DbSession

	equipment_status_id_list := make([]int64, 0)
	c.Params.Bind(&equipment_status_id_list, "id_list")

	affected, err := db_session.In("id", equipment_status_id_list).Delete(new(models.EquipmentStatus))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}

