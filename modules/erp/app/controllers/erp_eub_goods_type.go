
package controllers

import (
	"encoding/json"
	"strconv"
	"github.com/revel/revel"

	"matrix/core"
	"matrix/modules/erp/models"
)

type ErpEubGoodsType struct {
	*revel.Controller
}

func (c ErpEubGoodsType) Index() revel.Result {
	return c.RenderTemplate("erp/eub_goods_type/eub_goods_type_index.html")
}

func (c ErpEubGoodsType) ListData() revel.Result {
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

	eub_goods_type_list := make([]models.EubGoodsTypeInfo, 0, limit)
	err := data_query.Limit(limit, offset).Find(&eub_goods_type_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.EubGoodsTypeInfo))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  eub_goods_type_list,
		Total: count,
	})
}

type EubGoodsTypeDetailForm struct {
	EubGoodsType models.EubGoodsTypeInfo `json:"eub_goods_type"`
}

func (f *EubGoodsTypeDetailForm) IsCreate() bool {
	return f.EubGoodsType.Id == 0
}

func (f *EubGoodsTypeDetailForm) Valid(validation *revel.Validation) bool { 
	validation.Required(f.EubGoodsType.Name).Message("名称不能为空！")
	if f.EubGoodsType.Name != "" {
		validation.MinSize(f.EubGoodsType.Name, 2).Message("名称长度不能小于2！")
	}
	if f.EubGoodsType.Name != "" {
		validation.MaxSize(f.EubGoodsType.Name, 100).Message("名称长度不能大于100！")
	}

	if f.EubGoodsType.Desc.Valid == true {
		validation.MinSize(f.EubGoodsType.Desc.String, 2).Message("描述长度不能小于2！")
	}
	if f.EubGoodsType.Desc.Valid == true {
		validation.MaxSize(f.EubGoodsType.Desc.String, 400).Message("描述长度不能大于400！")
	}

	return validation.HasErrors() == false
}

func (c ErpEubGoodsType) DetailData() revel.Result {
	db_session := c.DbSession

	var eub_goods_type_id int64
	c.Params.Bind(&eub_goods_type_id, "id")

	var eub_goods_type models.EubGoodsTypeInfo
	if eub_goods_type_id != 0 {
		has, err := db_session.Id(eub_goods_type_id).Get(&eub_goods_type)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的EUB物品类型不存在！"})
		}
	}

	var detail_form EubGoodsTypeDetailForm
	detail_form.EubGoodsType = eub_goods_type

	return c.RenderJson(detail_form)
}

func (c ErpEubGoodsType) Save() revel.Result {
	db_session := c.DbSession

	var detail_form EubGoodsTypeDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	eub_goods_type := detail_form.EubGoodsType

	var affected int64
	if detail_form.IsCreate() { 
		affected, err = db_session.Insert(&eub_goods_type)
		core.HandleError(err)
	} else { 
		affected, err = db_session.Id(eub_goods_type.Id).Update(&eub_goods_type)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c ErpEubGoodsType) Delete() revel.Result {
	db_session := c.DbSession

	eub_goods_type_id_list := make([]int64, 0)
	c.Params.Bind(&eub_goods_type_id_list, "id_list")

	affected, err := db_session.In("id", eub_goods_type_id_list).Delete(new(models.EubGoodsTypeInfo))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}

