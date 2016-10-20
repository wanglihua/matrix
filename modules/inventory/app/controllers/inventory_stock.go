package controllers

import (
    "strconv"
    "github.com/revel/revel"

    "matrix/core"
    "matrix/modules/inventory/models"
)

type InventoryStock struct {
    *revel.Controller
}

func (c InventoryStock) Index() revel.Result {
    return c.RenderTemplate("inventory/stock/stock_index.html")
}

func (c InventoryStock) ListData() revel.Result {
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

    stockList := make([]models.Stock, 0, limit)
    err := dataQuery.Limit(limit, offset).Find(&stockList)
    core.HandleError(err)

    countQuery := *query
    count, err := countQuery.Count(new(models.Stock))
    core.HandleError(err)

    return c.RenderJson(core.GridResult{
        Data:  stockList,
        Total: count,
    })
}

type StockForm struct {
    Stock models.Stock
}

func (f *StockForm) IsCreate() bool {
    return f.Stock.Id == 0
}

func (f *StockForm) Valid(validation *revel.Validation) bool {
    validation.Required(f.Stock.Code).Message("编号不能为空！")
    if f.Stock.Code != "" {
        validation.MinSize(f.Stock.Code, 3).Message("编号长度不能小于3！")
    }
    if f.Stock.Code != "" {
        validation.MaxSize(f.Stock.Code, 10).Message("编号长度不能大于10！")
    }

    validation.Required(f.Stock.Name).Message("名称不能为空！")
    if f.Stock.Name != "" {
        validation.MinSize(f.Stock.Name, 3).Message("名称长度不能小于3！")
    }
    if f.Stock.Name != "" {
        validation.MaxSize(f.Stock.Name, 30).Message("名称长度不能大于30！")
    }

    if f.Stock.Address != "" {
        validation.MinSize(f.Stock.Address, 3).Message("仓库地址长度不能小于3！")
    }
    if f.Stock.Address != "" {
        validation.MaxSize(f.Stock.Address, 20).Message("仓库地址长度不能大于20！")
    }

    if f.Stock.Keeper != "" {
        validation.MaxSize(f.Stock.Keeper, 20).Message("负责人长度不能大于20！")
    }

    if f.Stock.Phone != "" {
        validation.MinSize(f.Stock.Phone, 2).Message("联系电话长度不能小于2！")
    }
    if f.Stock.Phone != "" {
        validation.MaxSize(f.Stock.Phone, 40).Message("联系电话长度不能大于40！")
    }

    if f.Stock.Remark != "" {
        validation.MinSize(f.Stock.Remark, 3).Message("备注长度不能小于3！")
    }
    if f.Stock.Remark != "" {
        validation.MaxSize(f.Stock.Remark, 250).Message("备注长度不能大于250！")
    }

    return validation.HasErrors() == false
}

func (c InventoryStock) Detail() revel.Result {
    session := c.DbSession

    var stockId int64
    c.Params.Bind(&stockId, "id")

    stock := new(models.Stock)
    if stockId != 0 {
        has, err := session.Id(stockId).Get(stock)
        core.HandleError(err)
        if has == false {
            panic("指定的仓库不存在！")
        }
    }

    form := new(StockForm)
    form.Stock = *stock

    c.UnbindToRenderArgs(form, "form")
    return c.RenderTemplate("inventory/stock/stock_detail.html")
}

func (c InventoryStock) Save() revel.Result {
    session := c.DbSession

    form := new(StockForm)
    c.Params.Bind(form, "form")

    if form.Valid(c.Validation) == false {
        return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage() })
    }

    stock := &form.Stock

    var affected int64
    if form.IsCreate() {
        codeCount, err := session.Where("stock_code = ?", stock.Code).Count(new(models.Stock))
        core.HandleError(err)
        if codeCount != 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，编号已存在！" })
        }

        nameCount, err := session.Where("stock_name = ?", stock.Name).Count(new(models.Stock))
        core.HandleError(err)
        if nameCount != 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，名称已存在！" })
        }

        affected, err = session.Insert(stock)
        core.HandleError(err)
    } else {
        codeCount, err := session.Where("id <> ? and stock_code = ?", stock.Id, stock.Code).Count(new(models.Stock))
        core.HandleError(err)
        if codeCount != 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，编号已存在！" })
        }

        nameCount, err := session.Where("id <> ? and stock_name = ?", stock.Id, stock.Name).Count(new(models.Stock))
        core.HandleError(err)
        if nameCount != 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，名称已存在！" })
        }

        affected, err = session.Id(stock.Id).Update(stock)
        core.HandleError(err)

        if affected == 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
        }
    }

    return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c InventoryStock) Delete() revel.Result {
    session := c.DbSession

    stockIdList := make([]int64, 0)
    c.Params.Bind(&stockIdList, "id_list")

    //删除库位
    _, err := session.In("stock_id", stockIdList).Delete(new(models.StorageLoc))
    core.HandleError(err)

    //删除仓库
    affected, err := session.In("id", stockIdList).Delete(new(models.Stock))
    core.HandleError(err)

    return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}

