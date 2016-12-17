package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"strconv"

	"matrix/core"
	"matrix/modules/inventory/models"
)

type InventoryStockIoType struct {
	*revel.Controller
}

func (c InventoryStockIoType) Index() revel.Result {
	return c.RenderTemplate("inventory/stock_io_type/stock_io_type_index.html")
}

func (c InventoryStockIoType) ListData() revel.Result {
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

	stock_io_type_list := make([]models.StockIoTypeInfo, 0, limit)
	err := data_query.Limit(limit, offset).Find(&stock_io_type_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.StockIoTypeInfo))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  stock_io_type_list,
		Total: count,
	})
}

type StockIoTypeDetailForm struct {
	StockIoType models.StockIoTypeInfo `json:"iotype"`
}

func (f *StockIoTypeDetailForm) IsCreate() bool {
	return f.StockIoType.Id == 0
}

func (f *StockIoTypeDetailForm) Valid(validation *revel.Validation) bool {
	validation.Required(f.StockIoType.Cate).Message("类别不能为空！")
	if f.StockIoType.Cate != "" {
		validation.MinSize(f.StockIoType.Cate, 2).Message("类别长度不能小于2！")
	}
	if f.StockIoType.Cate != "" {
		validation.MaxSize(f.StockIoType.Cate, 20).Message("类别长度不能大于20！")
	}

	validation.Required(f.StockIoType.Name).Message("名称不能为空！")
	if f.StockIoType.Name != "" {
		validation.MinSize(f.StockIoType.Name, 2).Message("名称长度不能小于2！")
	}
	if f.StockIoType.Name != "" {
		validation.MaxSize(f.StockIoType.Name, 200).Message("名称长度不能大于200！")
	}

	validation.Required(f.StockIoType.Code).Message("代码不能为空！")
	if f.StockIoType.Code != "" {
		validation.MinSize(f.StockIoType.Code, 2).Message("代码长度不能小于2！")
	}
	if f.StockIoType.Code != "" {
		validation.MaxSize(f.StockIoType.Code, 20).Message("代码长度不能大于20！")
	}

	if f.StockIoType.Footer.Valid == true {
		validation.MinSize(f.StockIoType.Footer.String, 2).Message("单据页脚长度不能小于2！")
	}
	if f.StockIoType.Footer.Valid == true {
		validation.MaxSize(f.StockIoType.Footer.String, 200).Message("单据页脚长度不能大于200！")
	}

	validation.Required(f.StockIoType.State).Message("状态不能为空！")
	if f.StockIoType.State != "" {
		validation.MinSize(f.StockIoType.State, 1).Message("状态长度不能小于1！")
	}

	return validation.HasErrors() == false
}

func (c InventoryStockIoType) DetailData() revel.Result {
	db_session := c.DbSession

	var stock_io_type_id int64
	c.Params.Bind(&stock_io_type_id, "id")

	var stock_io_type models.StockIoTypeInfo
	if stock_io_type_id != 0 {
		has, err := db_session.Id(stock_io_type_id).Get(&stock_io_type)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的出入库类型不存在！"})
		}
	}

	var detail_form StockIoTypeDetailForm
	detail_form.StockIoType = stock_io_type

	return c.RenderJson(detail_form)
}

func (c InventoryStockIoType) Save() revel.Result {
	db_session := c.DbSession

	var detail_form StockIoTypeDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	stock_io_type := detail_form.StockIoType

	var affected int64
	if detail_form.IsCreate() {
		name_count, err := db_session.Where("type_name = ?", stock_io_type.Name).Count(&stock_io_type)
		core.HandleError(err)
		if name_count != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，名称已存在！"})
		}

		code_count, err := db_session.Where("type_code = ?", stock_io_type.Code).Count(&stock_io_type)
		core.HandleError(err)
		if code_count != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，代码已存在！"})
		}

		affected, err = db_session.Insert(&stock_io_type)
		core.HandleError(err)
	} else {
		name_count, err := db_session.Where("id <> ? and type_name = ?", stock_io_type.Id, stock_io_type.Name).Count(&stock_io_type)
		core.HandleError(err)
		if name_count != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，名称已存在！"})
		}

		code_count, err := db_session.Where("id <> ? and type_code = ?", stock_io_type.Id, stock_io_type.Code).Count(&stock_io_type)
		core.HandleError(err)
		if code_count != 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，代码已存在！"})
		}

		affected, err = db_session.Id(stock_io_type.Id).Update(&stock_io_type)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！"})
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c InventoryStockIoType) Delete() revel.Result {
	db_session := c.DbSession

	stock_io_type_id_list := make([]int64, 0)
	c.Params.Bind(&stock_io_type_id_list, "id_list")

	affected, err := db_session.In("id", stock_io_type_id_list).Delete(new(models.StockIoTypeInfo))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}
