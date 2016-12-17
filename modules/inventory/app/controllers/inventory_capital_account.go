package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"strconv"

	"matrix/core"
	"matrix/modules/inventory/models"
)

type InventoryCapitalAccount struct {
	*revel.Controller
}

func (c InventoryCapitalAccount) Index() revel.Result {
	return c.RenderTemplate("inventory/capital_account/capital_account_index.html")
}

func (c InventoryCapitalAccount) ListData() revel.Result {
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

	capital_account_list := make([]models.CapitalAccountInfo, 0, limit)
	err := data_query.Limit(limit, offset).Find(&capital_account_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.CapitalAccountInfo))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  capital_account_list,
		Total: count,
	})
}

type CapitalAccountDetailForm struct {
	CapitalAccount models.CapitalAccountInfo `json:"ca"`
}

func (f *CapitalAccountDetailForm) IsCreate() bool {
	return f.CapitalAccount.Id == 0
}

func (f *CapitalAccountDetailForm) Valid(validation *revel.Validation) bool {
	validation.Required(f.CapitalAccount.Name).Message("名称不能为空！")
	if f.CapitalAccount.Name != "" {
		validation.MinSize(f.CapitalAccount.Name, 2).Message("名称长度不能小于2！")
	}
	if f.CapitalAccount.Name != "" {
		validation.MaxSize(f.CapitalAccount.Name, 200).Message("名称长度不能大于200！")
	}

	if f.CapitalAccount.Remark.Valid == true {
		validation.MinSize(f.CapitalAccount.Remark.String, 2).Message("描述长度不能小于2！")
	}
	if f.CapitalAccount.Remark.Valid == true {
		validation.MaxSize(f.CapitalAccount.Remark.String, 300).Message("描述长度不能大于300！")
	}

	return validation.HasErrors() == false
}

func (c InventoryCapitalAccount) DetailData() revel.Result {
	db_session := c.DbSession

	var capital_account_id int64
	c.Params.Bind(&capital_account_id, "id")

	var capital_account models.CapitalAccountInfo
	if capital_account_id != 0 {
		has, err := db_session.Id(capital_account_id).Get(&capital_account)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的资金账户不存在！"})
		}
	}

	var detail_form CapitalAccountDetailForm
	detail_form.CapitalAccount = capital_account

	return c.RenderJson(detail_form)
}

func (c InventoryCapitalAccount) Save() revel.Result {
	db_session := c.DbSession

	var detail_form CapitalAccountDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	capital_account := detail_form.CapitalAccount

	var affected int64
	if detail_form.IsCreate() {
		name_count, err := db_session.Where("account_name = ?", capital_account.Name).Count(&capital_account)
		core.HandleError(err)
		if name_count != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，名称已存在！"})
		}

		affected, err = db_session.Insert(&capital_account)
		core.HandleError(err)
	} else {
		name_count, err := db_session.Where("id <> ? and account_name = ?", capital_account.Id, capital_account.Name).Count(&capital_account)
		core.HandleError(err)
		if name_count != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，名称已存在！"})
		}

		affected, err = db_session.Id(capital_account.Id).Update(&capital_account)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！"})
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c InventoryCapitalAccount) Delete() revel.Result {
	db_session := c.DbSession

	capital_account_id_list := make([]int64, 0)
	c.Params.Bind(&capital_account_id_list, "id_list")

	affected, err := db_session.In("id", capital_account_id_list).Delete(new(models.CapitalAccountInfo))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}
