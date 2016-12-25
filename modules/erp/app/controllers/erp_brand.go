
package controllers

import (
	"encoding/json"
	"strconv"
	"github.com/revel/revel"

	"matrix/core"
	"matrix/modules/erp/models"
)

type ErpBrand struct {
	*revel.Controller
}

func (c ErpBrand) Index() revel.Result {
	return c.RenderTemplate("erp/brand/brand_index.html")
}

func (c ErpBrand) ListData() revel.Result {
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

	brand_list := make([]models.BrandInfo, 0, limit)
	err := data_query.Limit(limit, offset).Find(&brand_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.BrandInfo))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  brand_list,
		Total: count,
	})
}

type BrandDetailForm struct {
	Brand models.BrandInfo `json:"brand"`
}

func (f *BrandDetailForm) IsCreate() bool {
	return f.Brand.Id == 0
}

func (f *BrandDetailForm) Valid(validation *revel.Validation) bool { 
	validation.Required(f.Brand.Name).Message("品牌名称不能为空！")
	if f.Brand.Name != "" {
		validation.MinSize(f.Brand.Name, 2).Message("品牌名称长度不能小于2！")
	}
	if f.Brand.Name != "" {
		validation.MaxSize(f.Brand.Name, 100).Message("品牌名称长度不能大于100！")
	}

	if f.Brand.Contact.Valid == true {
		validation.MinSize(f.Brand.Contact.String, 2).Message("联系人长度不能小于2！")
	}
	if f.Brand.Contact.Valid == true {
		validation.MaxSize(f.Brand.Contact.String, 100).Message("联系人长度不能大于100！")
	}

	if f.Brand.Address.Valid == true {
		validation.MinSize(f.Brand.Address.String, 2).Message("地址长度不能小于2！")
	}
	if f.Brand.Address.Valid == true {
		validation.MaxSize(f.Brand.Address.String, 200).Message("地址长度不能大于100！")
	}

	if f.Brand.Phone.Valid == true {
		validation.MinSize(f.Brand.Phone.String, 2).Message("电话长度不能小于2！")
	}
	if f.Brand.Phone.Valid == true {
		validation.MaxSize(f.Brand.Phone.String, 100).Message("电话长度不能大于100！")
	}

	return validation.HasErrors() == false
}

func (c ErpBrand) DetailData() revel.Result {
	db_session := c.DbSession

	var brand_id int64
	c.Params.Bind(&brand_id, "id")

	var brand models.BrandInfo
	if brand_id != 0 {
		has, err := db_session.Id(brand_id).Get(&brand)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的品牌不存在！"})
		}
	}

	var detail_form BrandDetailForm
	detail_form.Brand = brand

	return c.RenderJson(detail_form)
}

func (c ErpBrand) Save() revel.Result {
	db_session := c.DbSession

	var detail_form BrandDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	brand := detail_form.Brand

	var affected int64
	if detail_form.IsCreate() { 
		affected, err = db_session.Insert(&brand)
		core.HandleError(err)
	} else { 
		affected, err = db_session.Id(brand.Id).Update(&brand)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c ErpBrand) Delete() revel.Result {
	db_session := c.DbSession

	brand_id_list := make([]int64, 0)
	c.Params.Bind(&brand_id_list, "id_list")

	affected, err := db_session.In("id", brand_id_list).Delete(new(models.BrandInfo))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}

