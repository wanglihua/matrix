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

	product_list := make([]models.Product, 0, limit)
	err := data_query.Limit(limit, offset).Find(&product_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.Product))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  product_list,
		Total: count,
	})
}

type ProductDetailForm struct {
	Product         models.Product `json:"product"`
	ProductCateList []models.ProductCate `json:"cate_list"`
}

func (f *ProductDetailForm) IsCreate() bool {
	return f.Product.Id == 0
}

func (f *ProductDetailForm) Valid(validation *revel.Validation) bool {
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

	var product_id int64
	c.Params.Bind(&product_id, "id")

	var detail_form ProductDetailForm

	var product models.Product
	if product_id != 0 {
		has, err := db_session.Id(product_id).Get(&product)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "货品信息不存在！"})
		}
	}

	product_cate_list := make([]models.ProductCate, 0)
	err := db_session.Find(&product_cate_list)
	core.HandleError(err)

	detail_form.Product = product
	detail_form.ProductCateList = product_cate_list

	return c.RenderJson(detail_form)
}

func (c InventoryProduct) Save() revel.Result {
	db_session := c.DbSession

	var detail_form ProductDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	product := detail_form.Product

	var affected int64
	if detail_form.IsCreate() {
		code_count, err := db_session.Where("code = ?", product.Code).Count(new(models.Product))
		core.HandleError(err)
		if code_count != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，编码已存在！"})
		}

		name_count, err := db_session.Where("name = ?", product.Name).Count(new(models.Product))
		core.HandleError(err)
		if name_count != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，品名已存在！"})
		}

		affected, err = db_session.Insert(&product)
		core.HandleError(err)
	} else {
		code_count, err := db_session.Where("id <> ? and code = ?", product.Id, product.Code).Count(new(models.Product))
		core.HandleError(err)
		if code_count != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，编码已存在！"})
		}

		name_count, err := db_session.Where("id <> ? and name = ?", product.Id, product.Name).Count(new(models.Product))
		core.HandleError(err)
		if name_count != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，品名已存在！"})
		}

		affected, err = db_session.Id(product.Id).Update(&product)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！"})
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c InventoryProduct) Delete() revel.Result {
	db_session := c.DbSession

	product_id_list := make([]int64, 0)
	c.Params.Bind(&product_id_list, "id_list")

	affected, err := db_session.In("id", product_id_list).Delete(new(models.Product))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}
