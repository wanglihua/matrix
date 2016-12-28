package controllers

import (
	"encoding/json"
	"strconv"
	"github.com/revel/revel"

	"matrix/core"
	"matrix/modules/erp/models"
	"matrix/modules/erp/service/status_service"
)

type ErpStock struct {
	*revel.Controller
}

func (c ErpStock) Index() revel.Result {
	return c.RenderTemplate("erp/stock/stock_index.html")
}

type ErpStockViewItem struct {
	models.StockInfo  `xorm:"extends" json:"stock"`
	models.StatusInfo `xorm:"extends" json:"status"`
}

func (c ErpStock) ListData() revel.Result {
	db_session := c.DbSession

	filter, order, offset, limit := core.GetGridRequestParam(c.Request)

	/*
		SELECT stock.*, status.*
		FROM erp_stock stock
		  INNER JOIN erp_status status ON stock.status_id = status.id
	*/
	query := db_session.
	Select("stock.*, status.*").
		Table("erp_stock").Alias("stock").
		Join("INNER", []string{"erp_status", "status"}, "stock.status_id = status.id").
		Where(filter)

	//query extra filter here

	data_query := *query
	if order != "" {
		data_query = *data_query.OrderBy(order)
	} else {
		data_query = *data_query.Asc("stock.id")
	}

	stock_list := make([]ErpStockViewItem, 0, limit)
	err := data_query.Limit(limit, offset).Find(&stock_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(ErpStockViewItem))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  stock_list,
		Total: count,
	})
}

type StockDetailForm struct {
	Stock      models.StockInfo    `json:"stock"`
	StatusList []models.StatusInfo `json:"status_list"`
}

func (f *StockDetailForm) IsCreate() bool {
	return f.Stock.Id == 0
}

func (f *StockDetailForm) Valid(validation *revel.Validation) bool {
	validation.Required(f.Stock.Name).Message("名称不能为空！")
	if f.Stock.Name != "" {
		validation.MinSize(f.Stock.Name, 2).Message("名称长度不能小于2！")
	}
	if f.Stock.Name != "" {
		validation.MaxSize(f.Stock.Name, 100).Message("名称长度不能大于100！")
	}

	if f.Stock.ShortName.Valid == true {
		validation.MinSize(f.Stock.ShortName.String, 2).Message("简称长度不能小于2！")
	}
	if f.Stock.ShortName.Valid == true {
		validation.MaxSize(f.Stock.ShortName.String, 100).Message("简称长度不能大于100！")
	}

	if f.Stock.Contact.Valid == true {
		validation.MinSize(f.Stock.Contact.String, 2).Message("联系人长度不能小于2！")
	}
	if f.Stock.Contact.Valid == true {
		validation.MaxSize(f.Stock.Contact.String, 100).Message("联系人长度不能大于100！")
	}

	if f.Stock.Address.Valid == true {
		validation.MinSize(f.Stock.Address.String, 2).Message("地址长度不能小于2！")
	}
	if f.Stock.Address.Valid == true {
		validation.MaxSize(f.Stock.Address.String, 200).Message("地址长度不能大于200！")
	}

	if f.Stock.Phone.Valid == true {
		validation.MinSize(f.Stock.Phone.String, 2).Message("电话长度不能小于2！")
	}
	if f.Stock.Phone.Valid == true {
		validation.MaxSize(f.Stock.Phone.String, 100).Message("电话长度不能大于100！")
	}

	validation.Required(f.Stock.StatusId).Message("状态不能为空！")

	return validation.HasErrors() == false
}

func (c ErpStock) DetailData() revel.Result {
	db_session := c.DbSession

	var stock_id int64
	c.Params.Bind(&stock_id, "id")

	var stock models.StockInfo
	if stock_id != 0 {
		has, err := db_session.Id(stock_id).Get(&stock)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的仓库不存在！"})
		}
	}

	status_list := status_service.GetStatusList(db_session, models.STATUS_TYPE_STOCK)

	var detail_form StockDetailForm
	detail_form.Stock = stock
	detail_form.StatusList = status_list

	return c.RenderJson(detail_form)
}

func (c ErpStock) Save() revel.Result {
	db_session := c.DbSession

	var detail_form StockDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	stock := detail_form.Stock

	var affected int64
	if detail_form.IsCreate() {
		affected, err = db_session.Insert(&stock)
		core.HandleError(err)
	} else {
		affected, err = db_session.Id(stock.Id).Update(&stock)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c ErpStock) Delete() revel.Result {
	db_session := c.DbSession

	stock_id_list := make([]int64, 0)
	c.Params.Bind(&stock_id_list, "id_list")

	affected, err := db_session.In("id", stock_id_list).Delete(new(models.StockInfo))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}
