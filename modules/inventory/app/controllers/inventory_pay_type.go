
package controllers

import (
	"encoding/json"
	"strconv"
	"github.com/revel/revel"

	"matrix/core"
	"matrix/modules/inventory/models"
)

type InventoryPayType struct {
	*revel.Controller
}

func (c InventoryPayType) Index() revel.Result {
	return c.RenderTemplate("inventory/pay_type/pay_type_index.html")
}

func (c InventoryPayType) ListData() revel.Result {
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

	pay_type_list := make([]models.PayType, 0, limit)
	err := data_query.Limit(limit, offset).Find(&pay_type_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.PayType))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  pay_type_list,
		Total: count,
	})
}

type PayTypeDetailForm struct {
	PayType models.PayType `json:"paytype"`
}

func (f *PayTypeDetailForm) IsCreate() bool {
	return f.PayType.Id == 0
}

func (f *PayTypeDetailForm) Valid(validation *revel.Validation) bool { 
	validation.Required(f.PayType.Cate).Message("类别不能为空！")
	if f.PayType.Cate != "" {
		validation.MinSize(f.PayType.Cate, 2).Message("类别长度不能小于2！")
	}
	if f.PayType.Cate != "" {
		validation.MaxSize(f.PayType.Cate, 20).Message("类别长度不能大于20！")
	}

	validation.Required(f.PayType.Name).Message("名称不能为空！")
	if f.PayType.Name != "" {
		validation.MinSize(f.PayType.Name, 2).Message("名称长度不能小于2！")
	}
	if f.PayType.Name != "" {
		validation.MaxSize(f.PayType.Name, 200).Message("名称长度不能大于200！")
	}

	validation.Required(f.PayType.Prefix).Message("单据前缀不能为空！")
	if f.PayType.Prefix != "" {
		validation.MinSize(f.PayType.Prefix, 2).Message("单据前缀长度不能小于2！")
	}
	if f.PayType.Prefix != "" {
		validation.MaxSize(f.PayType.Prefix, 10).Message("单据前缀长度不能大于10！")
	}

	if f.PayType.Remark.Valid == true {
		validation.MinSize(f.PayType.Remark.String, 2).Message("描述长度不能小于2！")
	}
	if f.PayType.Remark.Valid == true {
		validation.MaxSize(f.PayType.Remark.String, 300).Message("描述长度不能大于300！")
	}

	validation.Required(f.PayType.State).Message("状态不能为空！")
	if f.PayType.State != "" {
		validation.MinSize(f.PayType.State, 1).Message("状态长度不能小于1！")
	}

	return validation.HasErrors() == false
}

func (c InventoryPayType) DetailData() revel.Result {
	db_session := c.DbSession

	var pay_type_id int64
	c.Params.Bind(&pay_type_id, "id")

	var pay_type models.PayType
	if pay_type_id != 0 {
		has, err := db_session.Id(pay_type_id).Get(&pay_type)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的收付款类型不存在！"})
		}
	}

	var detail_form PayTypeDetailForm
	detail_form.PayType = pay_type

	return c.RenderJson(detail_form)
}

func (c InventoryPayType) Save() revel.Result {
	db_session := c.DbSession

	var detail_form PayTypeDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	pay_type := detail_form.PayType

	var affected int64
	if detail_form.IsCreate() {
		name_count, err := db_session.Where("type_name = ?", pay_type.Name).Count(&pay_type)
		core.HandleError(err)
		if name_count != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，名称已存在！" })
		}

		prefix_count, err := db_session.Where("prefix = ?", pay_type.Prefix).Count(&pay_type)
		core.HandleError(err)
		if prefix_count != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，单据前缀已存在！" })
		}

		affected, err = db_session.Insert(&pay_type)
		core.HandleError(err)
	} else { 
		name_count, err := db_session.Where("id <> ? and type_name = ?", pay_type.Id, pay_type.Name).Count(&pay_type)
		core.HandleError(err)
		if name_count != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，名称已存在！" })
		}

		prefix_count, err := db_session.Where("id <> ? and prefix = ?", pay_type.Id, pay_type.Prefix).Count(&pay_type)
		core.HandleError(err)
		if prefix_count != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，单据前缀已存在！" })
		}

		affected, err = db_session.Id(pay_type.Id).Update(&pay_type)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c InventoryPayType) Delete() revel.Result {
	db_session := c.DbSession

	pay_type_id_list := make([]int64, 0)
	c.Params.Bind(&pay_type_id_list, "id_list")

	affected, err := db_session.In("id", pay_type_id_list).Delete(new(models.PayType))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}

