package controllers

import (
    "strconv"
    "github.com/revel/revel"

    "matrix/core"
	"matrix/service"
    "matrix/modules/inventory/models"
)

type InventoryStorageLoc struct {
    *revel.Controller
    service.BaseController
}

func (c InventoryStorageLoc) Index() revel.Result {
    var stockId int64
    c.Params.Bind(&stockId, "stockId")

    c.RenderArgs["stockId"] = stockId
    return c.RenderTemplate("inventory/storageloc/storageloc_index.html")
}

func (c InventoryStorageLoc) ListData() revel.Result {
    session := c.DbSession

    var stockId int64
    c.Params.Bind(&stockId, "stockId")

    filter, order, offset, limit := core.GetGridRequestParam(c.Request)
    query := session.Where(filter)

    //query extra filter here
    query = query.Where("stock_id = ?", stockId);

    dataQuery := *query
    if order != "" {
        dataQuery = *dataQuery.OrderBy(order)
    } else {
        dataQuery = *dataQuery.Asc("id")
    }

    storageLocList := make([]models.StorageLoc, 0, limit)
    err := dataQuery.Limit(limit, offset).Find(&storageLocList)
    core.HandleError(err)

    countQuery := *query
    count, err := countQuery.Count(new(models.StorageLoc))
    core.HandleError(err)

    return c.RenderJson(core.GridResult{
        Data:  storageLocList,
        Total: count,
    })
}

type StorageLocForm struct {
    StorageLoc models.StorageLoc
}

func (f *StorageLocForm) IsCreate() bool {
    return f.StorageLoc.Id == 0
}

func (f *StorageLocForm) Valid(validation *revel.Validation) bool {
    validation.Required(f.StorageLoc.Code).Message("编号不能为空！")
    if f.StorageLoc.Code != "" {
        validation.MinSize(f.StorageLoc.Code, 3).Message("编号长度不能小于3！")
    }
    if f.StorageLoc.Code != "" {
        validation.MaxSize(f.StorageLoc.Code, 10).Message("编号长度不能大于10！")
    }

    validation.Required(f.StorageLoc.Name).Message("名称不能为空！")
    if f.StorageLoc.Name != "" {
        validation.MinSize(f.StorageLoc.Name, 3).Message("名称长度不能小于3！")
    }
    if f.StorageLoc.Name != "" {
        validation.MaxSize(f.StorageLoc.Name, 30).Message("名称长度不能大于30！")
    }

    validation.Required(f.StorageLoc.StockId).Message("仓库不能为空！")

    if f.StorageLoc.Remark != "" {
        validation.MinSize(f.StorageLoc.Remark, 3).Message("备注长度不能小于3！")
    }
    if f.StorageLoc.Remark != "" {
        validation.MaxSize(f.StorageLoc.Remark, 250).Message("备注长度不能大于250！")
    }

    return validation.HasErrors() == false
}

func (c InventoryStorageLoc) Detail() revel.Result {
    session := c.DbSession

    var storageLocId int64
    c.Params.Bind(&storageLocId, "id")

    storageLoc := new(models.StorageLoc)
    if storageLocId != 0 {
        has, err := session.Id(storageLocId).Get(storageLoc)
        core.HandleError(err)
        if has == false {
            panic("指定的库位不存在！")
        }
    } else {
        var stockId int64
        c.Params.Bind(&stockId, "stockId")
        storageLoc.StockId = stockId
    }

    form := new(StorageLocForm)
    form.StorageLoc = *storageLoc

    c.UnbindToRenderArgs(form, "form")
    return c.RenderTemplate("inventory/storageloc/storageloc_detail.html")
}

func (c InventoryStorageLoc) Save() revel.Result {
    session := c.DbSession

    form := new(StorageLocForm)
    c.Params.Bind(form, "form")

    if form.Valid(c.Validation) == false {
        return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage() })
    }

    storageLoc := &form.StorageLoc

    var affected int64
    if form.IsCreate() {
        codeCount, err := session.Where("stock_id = ? and stock_code = ?", storageLoc.StockId, storageLoc.Code).Count(new(models.StorageLoc))
        core.HandleError(err)
        if codeCount != 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，编号已存在！" })
        }

        nameCount, err := session.Where("stock_id = ? and stock_name = ?", storageLoc.StockId, storageLoc.Name).Count(new(models.StorageLoc))
        core.HandleError(err)
        if nameCount != 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，名称已存在！" })
        }

        affected, err = session.Insert(storageLoc)
        core.HandleError(err)
    } else {
        codeCount, err := session.Where("id <> ? and stock_id = ? and stock_code = ?", storageLoc.Id, storageLoc.StockId, storageLoc.Code).Count(new(models.StorageLoc))
        core.HandleError(err)
        if codeCount != 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，编号已存在！" })
        }

        nameCount, err := session.Where("id <> ? and stock_id = ? and stock_name = ?", storageLoc.Id, storageLoc.StockId, storageLoc.Name).Count(new(models.StorageLoc))
        core.HandleError(err)
        if nameCount != 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，名称已存在！" })
        }

        affected, err = session.Id(storageLoc.Id).Update(storageLoc)
        core.HandleError(err)

        if affected == 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
        }
    }

    return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c InventoryStorageLoc) Delete() revel.Result {
    session := c.DbSession

    storageLocIdList := make([]int64, 0)
    c.Params.Bind(&storageLocIdList, "id_list")

    storageLoc := new(models.StorageLoc)
    affected, err := session.In("id", storageLocIdList).Delete(storageLoc)
    core.HandleError(err)

    return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}

