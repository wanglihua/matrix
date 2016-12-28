package controllers

import (
	"encoding/json"
	"strconv"
	"github.com/revel/revel"

	"matrix/core"
	"matrix/modules/erp/models"
)

type ErpStorageLoc struct {
	*revel.Controller
}

func (c ErpStorageLoc) Index() revel.Result {
	return c.RenderTemplate("erp/storage_loc/storage_loc_index.html")
}

type ErpStorageLocViewItem struct {
	models.StorageLocInfo `xorm:"extends" json:"l"`
	models.StockInfo      `xorm:"extends" json:"s"`
}

func (c ErpStorageLoc) ListData() revel.Result {
	db_session := c.DbSession

	filter, order, offset, limit := core.GetGridRequestParam(c.Request)

	/*
		SELECT l.*, s.*
		FROM erp_storage_loc l
		    INNER JOIN erp_stock s ON l.stock_id = s.id
	*/

	query := db_session.
	Select("l.*, s.*").
		Table("erp_storage_loc").Alias("l").
		Join("INNER", []string{"erp_stock", "s"}, "l.stock_id = s.id").
		Where(filter)

	//query extra filter here

	data_query := *query
	if order != "" {
		data_query = *data_query.OrderBy(order)
	} else {
		data_query = *data_query.Asc("l.id")
	}

	storage_loc_list := make([]ErpStorageLocViewItem, 0, limit)
	err := data_query.Limit(limit, offset).Find(&storage_loc_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(ErpStorageLocViewItem))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  storage_loc_list,
		Total: count,
	})
}

type StorageLocDetailForm struct {
	StorageLoc models.StorageLocInfo `json:"storage_loc"`
	StockList  []models.StockInfo    `json:"stock_list"`
}

func (f *StorageLocDetailForm) IsCreate() bool {
	return f.StorageLoc.Id == 0
}

func (f *StorageLocDetailForm) Valid(validation *revel.Validation) bool {
	validation.Required(f.StorageLoc.StockId).Message("仓库不能为空！")

	validation.Required(f.StorageLoc.Name).Message("名称不能为空！")
	if f.StorageLoc.Name != "" {
		validation.MinSize(f.StorageLoc.Name, 2).Message("名称长度不能小于2！")
	}
	if f.StorageLoc.Name != "" {
		validation.MaxSize(f.StorageLoc.Name, 100).Message("名称长度不能大于100！")
	}

	if f.StorageLoc.Desc.Valid == true {
		validation.MaxSize(f.StorageLoc.Desc.String, 500).Message("描述长度不能大于500！")
	}

	validation.Required(f.StorageLoc.XAxes).Message("X轴坐标不能为空！")
	if f.StorageLoc.XAxes != "" {
		validation.MaxSize(f.StorageLoc.XAxes, 100).Message("X轴坐标长度不能大于100！")
	}

	validation.Required(f.StorageLoc.YAxes).Message("Y轴坐标不能为空！")
	if f.StorageLoc.YAxes != "" {
		validation.MaxSize(f.StorageLoc.YAxes, 100).Message("Y轴坐标长度不能大于100！")
	}

	return validation.HasErrors() == false
}

func (c ErpStorageLoc) DetailData() revel.Result {
	db_session := c.DbSession

	var storage_loc_id int64
	c.Params.Bind(&storage_loc_id, "id")

	var storage_loc models.StorageLocInfo
	if storage_loc_id != 0 {
		has, err := db_session.Id(storage_loc_id).Get(&storage_loc)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的仓位不存在！"})
		}
	}

	var stock_list = make([]models.StockInfo, 0)
	err := db_session.Find(&stock_list)
	core.HandleError(err)

	var detail_form StorageLocDetailForm
	detail_form.StorageLoc = storage_loc
	detail_form.StockList = stock_list

	return c.RenderJson(detail_form)
}

func (c ErpStorageLoc) Save() revel.Result {
	db_session := c.DbSession

	var detail_form StorageLocDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	storage_loc := detail_form.StorageLoc

	var affected int64
	if detail_form.IsCreate() {
		affected, err = db_session.Insert(&storage_loc)
		core.HandleError(err)
	} else {
		affected, err = db_session.Id(storage_loc.Id).Update(&storage_loc)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c ErpStorageLoc) Delete() revel.Result {
	db_session := c.DbSession

	storage_loc_id_list := make([]int64, 0)
	c.Params.Bind(&storage_loc_id_list, "id_list")

	affected, err := db_session.In("id", storage_loc_id_list).Delete(new(models.StorageLocInfo))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}
