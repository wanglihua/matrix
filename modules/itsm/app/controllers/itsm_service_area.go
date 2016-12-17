package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"strconv"

	"matrix/core"
	"matrix/modules/itsm/models"
)

type ItsmServiceArea struct {
	*revel.Controller
}

func (c ItsmServiceArea) Index() revel.Result {
	return c.RenderTemplate("itsm/service_area/service_area_index.html")
}

func (c ItsmServiceArea) ListData() revel.Result {
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

	service_area_list := make([]models.ServiceAreaInfo, 0, limit)
	err := data_query.Limit(limit, offset).Find(&service_area_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.ServiceAreaInfo))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  service_area_list,
		Total: count,
	})
}

type ServiceAreaDetailForm struct {
	ServiceArea models.ServiceAreaInfo `json:"area"`
}

func (f *ServiceAreaDetailForm) IsCreate() bool {
	return f.ServiceArea.Id == 0
}

func (f *ServiceAreaDetailForm) Valid(validation *revel.Validation) bool {
	validation.Required(f.ServiceArea.Name).Message("名称不能为空！")
	if f.ServiceArea.Name != "" {
		validation.MinSize(f.ServiceArea.Name, 1).Message("名称长度不能小于1！")
	}
	if f.ServiceArea.Name != "" {
		validation.MaxSize(f.ServiceArea.Name, 30).Message("名称长度不能大于30！")
	}

	return validation.HasErrors() == false
}

func (c ItsmServiceArea) DetailData() revel.Result {
	db_session := c.DbSession

	var service_area_id int64
	c.Params.Bind(&service_area_id, "id")

	var service_area models.ServiceAreaInfo
	if service_area_id != 0 {
		has, err := db_session.Id(service_area_id).Get(&service_area)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的服务区域不存在！"})
		}
	}

	var detail_form ServiceAreaDetailForm
	detail_form.ServiceArea = service_area

	return c.RenderJson(detail_form)
}

func (c ItsmServiceArea) Save() revel.Result {
	db_session := c.DbSession

	var detail_form ServiceAreaDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	service_area := detail_form.ServiceArea

	var affected int64
	if detail_form.IsCreate() {
		affected, err = db_session.Insert(&service_area)
		core.HandleError(err)
	} else {
		affected, err = db_session.Id(service_area.Id).Update(&service_area)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！"})
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c ItsmServiceArea) Delete() revel.Result {
	db_session := c.DbSession

	service_area_id_list := make([]int64, 0)
	c.Params.Bind(&service_area_id_list, "id_list")

	affected, err := db_session.In("id", service_area_id_list).Delete(new(models.ServiceAreaInfo))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}
