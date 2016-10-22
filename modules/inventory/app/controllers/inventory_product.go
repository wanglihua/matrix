package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"matrix/core"
	"matrix/modules/inventory/models"
	"strconv"
)

type InventoryProduct struct {
	*revel.Controller
}

func (c InventoryProduct) Index() revel.Result {
	return c.RenderTemplate("inventory/product/product_index.html")
}

func (c InventoryProduct) ListData() revel.Result {
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

	productList := make([]models.Product, 0, limit)
	err := dataQuery.Limit(limit, offset).Find(&productList)
	core.HandleError(err)

	countQuery := *query
	count, err := countQuery.Count(new(models.Product))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  productList,
		Total: count,
	})
}

type ProductForm struct {
	Product models.Product `json:"product"`
}

func (f *ProductForm) IsCreate() bool {
	return f.Product.Id == 0
}

func (f *ProductForm) Valid(validation *revel.Validation) bool {
	validation.Required(f.Product.Code).Message("编码不能为空！")
	if f.Product.Code != "" {
		validation.MinSize(f.Product.Code, 2).Message("编码长度不能小于2！")
	}
	if f.Product.Code != "" {
		validation.MaxSize(f.Product.Code, 20).Message("编码长度不能大于20！")
	}

	validation.Required(f.Product.Name).Message("品名不能为空！")
	if f.Product.Name != "" {
		validation.MinSize(f.Product.Name, 2).Message("品名长度不能小于2！")
	}
	if f.Product.Name != "" {
		validation.MaxSize(f.Product.Name, 100).Message("品名长度不能大于100！")
	}

	if f.Product.Model.Valid == true {
		validation.MinSize(f.Product.Model.String, 2).Message("规格型号长度不能小于2！")
		validation.MaxSize(f.Product.Model.String, 200).Message("规格型号长度不能大于200！")
	}

	validation.Required(f.Product.Unit).Message("单位不能为空！")
	if f.Product.Unit != "" {
		validation.MinSize(f.Product.Unit, 1).Message("单位长度不能小于1！")
	}
	if f.Product.Unit != "" {
		validation.MaxSize(f.Product.Unit, 10).Message("单位长度不能大于10！")
	}

	validation.Required(f.Product.State).Message("状态不能为空！")
	if f.Product.State != "" {
		validation.MinSize(f.Product.State, 1).Message("状态长度不能小于1！")
	}

	return validation.HasErrors() == false
}

func (c InventoryProduct) DetailData() revel.Result {
	db_session := c.DbSession

	var id int64
	c.Params.Bind(&id, "id")

	var product models.Product
	if id != 0 {
		has, err := db_session.Id(id).Get(&product)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "货品信息不存在！"})
		}
	}

	var detail_form ProductForm
	detail_form.Product = product

	return c.RenderJson(detail_form)
}

func (c InventoryProduct) Save() revel.Result {
	session := c.DbSession

	var detail_form ProductForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	product := detail_form.Product

	var affected int64
	if detail_form.IsCreate() {
		codeCount, err := session.Where("code = ?", product.Code).Count(new(models.Product))
		core.HandleError(err)
		if codeCount != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，编码已存在！"})
		}

		nameCount, err := session.Where("name = ?", product.Name).Count(new(models.Product))
		core.HandleError(err)
		if nameCount != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，品名已存在！"})
		}

		affected, err = session.Insert(&product)
		core.HandleError(err)
	} else {
		codeCount, err := session.Where("id <> ? and code = ?", product.Id, product.Code).Count(new(models.Product))
		core.HandleError(err)
		if codeCount != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，编码已存在！"})
		}

		nameCount, err := session.Where("id <> ? and name = ?", product.Id, product.Name).Count(new(models.Product))
		core.HandleError(err)
		if nameCount != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，品名已存在！"})
		}

		affected, err = session.Id(product.Id).Update(&product)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！"})
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c InventoryProduct) Delete() revel.Result {
	session := c.DbSession

	productIdList := make([]int64, 0)
	c.Params.Bind(&productIdList, "id_list")

	affected, err := session.In("id", productIdList).Delete(new(models.Product))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}
