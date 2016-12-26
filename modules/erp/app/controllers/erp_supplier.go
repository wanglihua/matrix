
package controllers

import (
	"encoding/json"
	"strconv"
	"github.com/revel/revel"

	"matrix/core"
	"matrix/modules/erp/models"
)

type ErpSupplier struct {
	*revel.Controller
}

func (c ErpSupplier) Index() revel.Result {
	return c.RenderTemplate("erp/supplier/supplier_index.html")
}

func (c ErpSupplier) ListData() revel.Result {
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

	supplier_list := make([]models.SupplierInfo, 0, limit)
	err := data_query.Limit(limit, offset).Find(&supplier_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.SupplierInfo))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  supplier_list,
		Total: count,
	})
}

type SupplierDetailForm struct {
	Supplier models.SupplierInfo `json:"supplier"`
}

func (f *SupplierDetailForm) IsCreate() bool {
	return f.Supplier.Id == 0
}

func (f *SupplierDetailForm) Valid(validation *revel.Validation) bool { 
	validation.Required(f.Supplier.Name).Message("名称不能为空！")
	if f.Supplier.Name != "" {
		validation.MinSize(f.Supplier.Name, 2).Message("名称长度不能小于2！")
	}
	if f.Supplier.Name != "" {
		validation.MaxSize(f.Supplier.Name, 100).Message("名称长度不能大于100！")
	}

	if f.Supplier.Contact.Valid == true {
		validation.MinSize(f.Supplier.Contact.String, 2).Message("联系人长度不能小于2！")
	}
	if f.Supplier.Contact.Valid == true {
		validation.MaxSize(f.Supplier.Contact.String, 100).Message("联系人长度不能大于100！")
	}

	if f.Supplier.ContactInfo.Valid == true {
		validation.MaxSize(f.Supplier.ContactInfo.String, 500).Message("联系方式长度不能大于500！")
	}

	if f.Supplier.Address.Valid == true {
		validation.MinSize(f.Supplier.Address.String, 2).Message("地址长度不能小于2！")
	}
	if f.Supplier.Address.Valid == true {
		validation.MaxSize(f.Supplier.Address.String, 200).Message("地址长度不能大于200！")
	}

	if f.Supplier.Phone.Valid == true {
		validation.MinSize(f.Supplier.Phone.String, 2).Message("电话长度不能小于2！")
	}
	if f.Supplier.Phone.Valid == true {
		validation.MaxSize(f.Supplier.Phone.String, 100).Message("电话长度不能大于100！")
	}

	if f.Supplier.Desc.Valid == true {
		validation.MaxSize(f.Supplier.Desc.String, 500).Message("描述长度不能大于500！")
	}

	return validation.HasErrors() == false
}

func (c ErpSupplier) DetailData() revel.Result {
	db_session := c.DbSession

	var supplier_id int64
	c.Params.Bind(&supplier_id, "id")

	var supplier models.SupplierInfo
	if supplier_id != 0 {
		has, err := db_session.Id(supplier_id).Get(&supplier)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的供应商不存在！"})
		}
	}

	var detail_form SupplierDetailForm
	detail_form.Supplier = supplier

	return c.RenderJson(detail_form)
}

func (c ErpSupplier) Save() revel.Result {
	db_session := c.DbSession

	var detail_form SupplierDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	supplier := detail_form.Supplier

	var affected int64
	if detail_form.IsCreate() { 
		affected, err = db_session.Insert(&supplier)
		core.HandleError(err)
	} else { 
		affected, err = db_session.Id(supplier.Id).Update(&supplier)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c ErpSupplier) Delete() revel.Result {
	db_session := c.DbSession

	supplier_id_list := make([]int64, 0)
	c.Params.Bind(&supplier_id_list, "id_list")

	affected, err := db_session.In("id", supplier_id_list).Delete(new(models.SupplierInfo))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}

