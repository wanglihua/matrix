
package controllers

import (
	"encoding/json"
	"strconv"
	"github.com/revel/revel"

	"matrix/core"
	"matrix/modules/inventory/models"
)

type InventoryStockBillDetail struct {
	*revel.Controller
}

func (c InventoryStockBillDetail) Index() revel.Result {
	return c.RenderTemplate("inventory/stock_bill_detail/stock_bill_detail_index.html")
}

func (c InventoryStockBillDetail) ListData() revel.Result {
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

	stock_bill_detail_list := make([]models.StockBillDetail, 0, limit)
	err := data_query.Limit(limit, offset).Find(&stock_bill_detail_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.StockBillDetail))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  stock_bill_detail_list,
		Total: count,
	})
}

type StockBillDetailDetailForm struct {
	StockBillDetail models.StockBillDetail `json:"billdetail"`
}

func (f *StockBillDetailDetailForm) IsCreate() bool {
	return f.StockBillDetail.Id == 0
}

func (f *StockBillDetailDetailForm) Valid(validation *revel.Validation) bool { 

	validation.Required(f.StockBillDetail.Code).Message("编码不能为空！")
	if f.StockBillDetail.Code != "" {
		validation.MinSize(f.StockBillDetail.Code, 2).Message("编码长度不能小于2！")
	}
	if f.StockBillDetail.Code != "" {
		validation.MaxSize(f.StockBillDetail.Code, 20).Message("编码长度不能大于20！")
	}

	validation.Required(f.StockBillDetail.Name).Message("品名不能为空！")
	if f.StockBillDetail.Name != "" {
		validation.MinSize(f.StockBillDetail.Name, 2).Message("品名长度不能小于2！")
	}
	if f.StockBillDetail.Name != "" {
		validation.MaxSize(f.StockBillDetail.Name, 100).Message("品名长度不能大于100！")
	}

	if f.StockBillDetail.Model.Valid == true {
		validation.MaxSize(f.StockBillDetail.Model.String, 200).Message("规格型号长度不能大于200！")
	}

	validation.Required(f.StockBillDetail.Unit).Message("单位不能为空！")
	if f.StockBillDetail.Unit != "" {
		validation.MinSize(f.StockBillDetail.Unit, 1).Message("单位长度不能小于1！")
	}
	if f.StockBillDetail.Unit != "" {
		validation.MaxSize(f.StockBillDetail.Unit, 10).Message("单位长度不能大于10！")
	}

	validation.Required(f.StockBillDetail.Price).Message("价格不能为空！")

	validation.Required(f.StockBillDetail.Quantity).Message("数量不能为空！")

	validation.Required(f.StockBillDetail.Amount).Message("金额不能为空！")

	if f.StockBillDetail.Batch.Valid == true {
		validation.MaxSize(f.StockBillDetail.Batch.String, 50).Message("批号长度不能大于50！")
	}

	if f.StockBillDetail.Remark.Valid == true {
		validation.MaxSize(f.StockBillDetail.Remark.String, 300).Message("备注长度不能大于300！")
	}

	return validation.HasErrors() == false
}

func (c InventoryStockBillDetail) DetailData() revel.Result {
	db_session := c.DbSession

	var stock_bill_detail_id int64
	c.Params.Bind(&stock_bill_detail_id, "id")

	var stock_bill_detail models.StockBillDetail
	if stock_bill_detail_id != 0 {
		has, err := db_session.Id(stock_bill_detail_id).Get(&stock_bill_detail)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的出入库单详细不存在！"})
		}
	}

	var detail_form StockBillDetailDetailForm
	detail_form.StockBillDetail = stock_bill_detail

	return c.RenderJson(detail_form)
}

func (c InventoryStockBillDetail) Save() revel.Result {
	db_session := c.DbSession

	var detail_form StockBillDetailDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	stock_bill_detail := detail_form.StockBillDetail

	var affected int64
	if detail_form.IsCreate() {
		code_count, err := db_session.Where("code = ?", stock_bill_detail.Code).Count(&stock_bill_detail)
		core.HandleError(err)
		if code_count != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，编码已存在！" })
		}

		name_count, err := db_session.Where("name = ?", stock_bill_detail.Name).Count(&stock_bill_detail)
		core.HandleError(err)
		if name_count != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，品名已存在！" })
		}

		affected, err = db_session.Insert(&stock_bill_detail)
		core.HandleError(err)
	} else { 
		code_count, err := db_session.Where("id <> ? and code = ?", stock_bill_detail.Id, stock_bill_detail.Code).Count(&stock_bill_detail)
		core.HandleError(err)
		if code_count != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，编码已存在！" })
		}

		name_count, err := db_session.Where("id <> ? and name = ?", stock_bill_detail.Id, stock_bill_detail.Name).Count(&stock_bill_detail)
		core.HandleError(err)
		if name_count != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，品名已存在！" })
		}

		affected, err = db_session.Id(stock_bill_detail.Id).Update(&stock_bill_detail)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c InventoryStockBillDetail) Delete() revel.Result {
	db_session := c.DbSession

	stock_bill_detail_id_list := make([]int64, 0)
	c.Params.Bind(&stock_bill_detail_id_list, "id_list")

	affected, err := db_session.In("id", stock_bill_detail_id_list).Delete(new(models.StockBillDetail))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}

