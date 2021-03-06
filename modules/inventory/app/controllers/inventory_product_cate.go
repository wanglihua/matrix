package controllers

import (
	"github.com/revel/revel"
	"strconv"

	"encoding/json"
	"matrix/core"
	"matrix/modules/inventory/models"
)

type InventoryProductCate struct {
	*revel.Controller
}

func (c InventoryProductCate) Index() revel.Result {
	return c.RenderTemplate("inventory/product_cate/product_cate_index.html")
}

func (c InventoryProductCate) ListData() revel.Result {
	session := c.DbSession

	filter, order, offset, limit := core.GetGridRequestParam(c.Request)
	query := session.Where(filter)

	//query extra filter here

	dataQuery := *query
	if order != "" {
		dataQuery = *dataQuery.OrderBy(order)
	} else {
		dataQuery = *dataQuery.Asc("id")
	}

	productCateList := make([]models.ProductCateInfo, 0, limit)
	err := dataQuery.Limit(limit, offset).Find(&productCateList)
	core.HandleError(err)

	countQuery := *query
	count, err := countQuery.Count(new(models.ProductCateInfo))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  productCateList,
		Total: count,
	})
}

type ProductCateForm struct {
	ProductCate models.ProductCateInfo `json:"cate"`
}

func (f *ProductCateForm) IsCreate() bool {
	return f.ProductCate.Id == 0
}

func (f *ProductCateForm) Valid(validation *revel.Validation) bool {
	validation.Required(f.ProductCate.Name).Message("名称不能为空！")
	if f.ProductCate.Name != "" {
		validation.MinSize(f.ProductCate.Name, 2).Message("名称长度不能小于2！")
	}
	if f.ProductCate.Name != "" {
		validation.MaxSize(f.ProductCate.Name, 100).Message("名称长度不能大于100！")
	}

	if f.ProductCate.Desc.Valid {
		validation.MinSize(f.ProductCate.Desc.String, 2).Message("描述长度不能小于2！")
	}
	if f.ProductCate.Desc.Valid {
		validation.MaxSize(f.ProductCate.Desc.String, 300).Message("描述长度不能大于300！")
	}

	return validation.HasErrors() == false
}

func (c InventoryProductCate) DetailData() revel.Result {
	db_session := c.DbSession

	var id int64
	c.Params.Bind(&id, "id")

	var product_cate models.ProductCateInfo
	if id != 0 {
		has, err := db_session.Id(id).Get(&product_cate)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "货品分类不存在！"})
		}
	}

	var detail_form ProductCateForm
	detail_form.ProductCate = product_cate

	return c.RenderJson(detail_form)
}

func (c InventoryProductCate) Save() revel.Result {
	session := c.DbSession

	var detail_form ProductCateForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	productCate := detail_form.ProductCate

	var affected int64
	if detail_form.IsCreate() {
		nameCount, err := session.Where("cate_name = ?", productCate.Name).Count(new(models.ProductCateInfo))
		core.HandleError(err)
		if nameCount != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，名称已存在！"})
		}

		affected, err = session.Insert(&productCate)
		core.HandleError(err)
	} else {
		nameCount, err := session.Where("id <> ? and cate_name = ?", productCate.Id, productCate.Name).Count(new(models.ProductCateInfo))
		core.HandleError(err)
		if nameCount != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，名称已存在！"})
		}

		affected, err = session.Id(productCate.Id).Update(&productCate)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！"})
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c InventoryProductCate) Delete() revel.Result {
	session := c.DbSession

	productCateIdList := make([]int64, 0)
	c.Params.Bind(&productCateIdList, "id_list")

	affected, err := session.In("id", productCateIdList).Delete(new(models.ProductCateInfo))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}
