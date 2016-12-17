package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"strconv"

	"matrix/core"
	"matrix/modules/inventory/models"
)

type InventoryStockBill struct {
	*revel.Controller
}

func (c InventoryStockBill) Index() revel.Result {
	return c.RenderTemplate("inventory/stock_bill/stock_bill_index.html")
}

func (c InventoryStockBill) ListData() revel.Result {
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

	stock_bill_list := make([]models.StockBillInfo, 0, limit)
	err := data_query.Limit(limit, offset).Find(&stock_bill_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.StockBillInfo))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  stock_bill_list,
		Total: count,
	})
}

type StockBillDetailForm struct {
	StockBill models.StockBillInfo `json:"bill"`
}

func (f *StockBillDetailForm) IsCreate() bool {
	return f.StockBill.Id == 0
}

func (f *StockBillDetailForm) Valid(validation *revel.Validation) bool {
	validation.Required(f.StockBill.Code).Message("出入库单号不能为空！")
	if f.StockBill.Code != "" {
		validation.MinSize(f.StockBill.Code, 2).Message("出入库单号长度不能小于2！")
	}
	if f.StockBill.Code != "" {
		validation.MaxSize(f.StockBill.Code, 100).Message("出入库单号长度不能大于100！")
	}

	validation.Required(f.StockBill.Date).Message("出入库日期不能为空！")

	validation.Required(f.StockBill.TypeId).Message("出入库类型不能为空！")

	validation.Required(f.StockBill.StockId).Message("仓库不能为空！")

	validation.Required(f.StockBill.HandlerId).Message("经手人不能为空！")

	validation.Required(f.StockBill.ObjectId).Message("往来对象不能为空！")

	if f.StockBill.Remark.Valid == true {
		validation.MaxSize(f.StockBill.Remark.String, 300).Message("备注长度不能大于300！")
	}

	validation.Required(f.StockBill.State).Message("状态不能为空！")
	if f.StockBill.State != "" {
		validation.MinSize(f.StockBill.State, 1).Message("状态长度不能小于1！")
	}

	return validation.HasErrors() == false
}

func (c InventoryStockBill) DetailData() revel.Result {
	db_session := c.DbSession

	var stock_bill_id int64
	c.Params.Bind(&stock_bill_id, "id")

	var stock_bill models.StockBillInfo
	if stock_bill_id != 0 {
		has, err := db_session.Id(stock_bill_id).Get(&stock_bill)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的出入库单不存在！"})
		}
	}

	var detail_form StockBillDetailForm
	detail_form.StockBill = stock_bill

	return c.RenderJson(detail_form)
}

func (c InventoryStockBill) Save() revel.Result {
	db_session := c.DbSession

	var detail_form StockBillDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	stock_bill := detail_form.StockBill

	var affected int64
	if detail_form.IsCreate() {
		code_count, err := db_session.Where("bill_code = ?", stock_bill.Code).Count(&stock_bill)
		core.HandleError(err)
		if code_count != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，出入库单号已存在！"})
		}

		affected, err = db_session.Insert(&stock_bill)
		core.HandleError(err)
	} else {
		code_count, err := db_session.Where("id <> ? and bill_code = ?", stock_bill.Id, stock_bill.Code).Count(&stock_bill)
		core.HandleError(err)
		if code_count != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，出入库单号已存在！"})
		}

		affected, err = db_session.Id(stock_bill.Id).Update(&stock_bill)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！"})
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c InventoryStockBill) Delete() revel.Result {
	db_session := c.DbSession

	stock_bill_id_list := make([]int64, 0)
	c.Params.Bind(&stock_bill_id_list, "id_list")

	affected, err := db_session.In("id", stock_bill_id_list).Delete(new(models.StockBillInfo))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}
