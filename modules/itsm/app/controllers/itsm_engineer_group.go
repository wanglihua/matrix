package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"strconv"

	"matrix/core"
	"matrix/modules/itsm/models"
)

type ItsmEngineerGroup struct {
	*revel.Controller
}

func (c ItsmEngineerGroup) Index() revel.Result {
	return c.RenderTemplate("itsm/engineer_group/engineer_group_index.html")
}

func (c ItsmEngineerGroup) ListData() revel.Result {
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

	engineer_group_list := make([]models.EngineerGroupInfo, 0, limit)
	err := data_query.Limit(limit, offset).Find(&engineer_group_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.EngineerGroupInfo))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  engineer_group_list,
		Total: count,
	})
}

type EngineerGroupDetailForm struct {
	EngineerGroup models.EngineerGroupInfo `json:"group"`
}

func (f *EngineerGroupDetailForm) IsCreate() bool {
	return f.EngineerGroup.Id == 0
}

func (f *EngineerGroupDetailForm) Valid(validation *revel.Validation) bool {
	validation.Required(f.EngineerGroup.Name).Message("名称不能为空！")
	if f.EngineerGroup.Name != "" {
		validation.MinSize(f.EngineerGroup.Name, 1).Message("名称长度不能小于1！")
	}
	if f.EngineerGroup.Name != "" {
		validation.MaxSize(f.EngineerGroup.Name, 30).Message("名称长度不能大于30！")
	}

	return validation.HasErrors() == false
}

func (c ItsmEngineerGroup) DetailData() revel.Result {
	db_session := c.DbSession

	var engineer_group_id int64
	c.Params.Bind(&engineer_group_id, "id")

	var engineer_group models.EngineerGroupInfo
	if engineer_group_id != 0 {
		has, err := db_session.Id(engineer_group_id).Get(&engineer_group)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的工程师分组不存在！"})
		}
	}

	var detail_form EngineerGroupDetailForm
	detail_form.EngineerGroup = engineer_group

	return c.RenderJson(detail_form)
}

func (c ItsmEngineerGroup) Save() revel.Result {
	db_session := c.DbSession

	var detail_form EngineerGroupDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	engineer_group := detail_form.EngineerGroup

	var affected int64
	if detail_form.IsCreate() {
		affected, err = db_session.Insert(&engineer_group)
		core.HandleError(err)
	} else {
		affected, err = db_session.Id(engineer_group.Id).Update(&engineer_group)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！"})
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c ItsmEngineerGroup) Delete() revel.Result {
	db_session := c.DbSession

	engineer_group_id_list := make([]int64, 0)
	c.Params.Bind(&engineer_group_id_list, "id_list")

	affected, err := db_session.In("id", engineer_group_id_list).Delete(new(models.EngineerGroupInfo))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}
