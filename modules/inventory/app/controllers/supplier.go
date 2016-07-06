
package controllers

import (
    "strconv"
    "github.com/revel/revel"

    "matrix/core"
    "matrix/modules/inventory/models"
)

type InventorySupplier struct {
    *revel.Controller
    core.BaseController
}

func (c InventorySupplier) Index() revel.Result {
    return c.RenderTemplate("inventory/supplier/supplier_index.html")
}

func (c InventorySupplier) ListData() revel.Result {
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

    supplierList := make([]models.Supplier, 0, limit)
    err := dataQuery.Limit(limit, offset).Find(&supplierList)
    core.HandleError(err)

    countQuery := *query
    count, err := countQuery.Count(new(models.Supplier))
    core.HandleError(err)

    return c.RenderJson(core.GridResult{
        Data:  supplierList,
        Total: count,
    })
}

type SupplierForm struct {
    Supplier models.Supplier
}

func (f SupplierForm) IsCreate() bool {
    return f.Supplier.Id == 0
}

func (f SupplierForm) Valid(validation *revel.Validation) bool { 
    validation.Required(f.Supplier.Name).Message("供应商名称不能为空！")
    if f.Supplier.Name != "" {
        validation.MinSize(f.Supplier.Name, 3).Message("供应商名称长度不能小于3！")
    }
    if f.Supplier.Name != "" {
        validation.MaxSize(f.Supplier.Name, 30).Message("供应商名称长度不能大于30！")
    }

    validation.Required(f.Supplier.Code).Message("供应商编号不能为空！")
    if f.Supplier.Code != "" {
        validation.MinSize(f.Supplier.Code, 3).Message("供应商编号长度不能小于3！")
    }
    if f.Supplier.Code != "" {
        validation.MaxSize(f.Supplier.Code, 10).Message("供应商编号长度不能大于10！")
    }

    if f.Supplier.Contact != "" {
        validation.MaxSize(f.Supplier.Contact, 20).Message("联系人长度不能大于20！")
    }

    if f.Supplier.Phone != "" {
        validation.MinSize(f.Supplier.Phone, 2).Message("联系电话长度不能小于2！")
    }
    if f.Supplier.Phone != "" {
        validation.MaxSize(f.Supplier.Phone, 40).Message("联系电话长度不能大于40！")
    }

    if f.Supplier.Address != "" {
        validation.MinSize(f.Supplier.Address, 3).Message("联系地址长度不能小于3！")
    }
    if f.Supplier.Address != "" {
        validation.MaxSize(f.Supplier.Address, 20).Message("联系地址长度不能大于20！")
    }

    if f.Supplier.Bank != "" {
        validation.MaxSize(f.Supplier.Bank, 20).Message("开户银行长度不能大于20！")
    }

    if f.Supplier.BankAccount != "" {
        validation.MinSize(f.Supplier.BankAccount, 4).Message("银行账号长度不能小于4！")
    }
    if f.Supplier.BankAccount != "" {
        validation.MaxSize(f.Supplier.BankAccount, 50).Message("银行账号长度不能大于50！")
    }

    if f.Supplier.Remark != "" {
        validation.MinSize(f.Supplier.Remark, 3).Message("备注长度不能小于3！")
    }
    if f.Supplier.Remark != "" {
        validation.MaxSize(f.Supplier.Remark, 250).Message("备注长度不能大于250！")
    }

    return validation.HasErrors() == false
}

func (c InventorySupplier) Detail() revel.Result {
    session := c.DbSession

    supplierId := core.GetInt64FromRequest(c.Request, "id")

    supplier := new(models.Supplier)
    if supplierId != 0 {
        has, err := session.Id(supplierId).Get(supplier)
        core.HandleError(err)
        if has == false {
            panic("指定的供应商不存在！")
        }
    }

    form := new(SupplierForm)
    form.Supplier = *supplier

    c.RenderArgs["form"] = form
    return c.RenderTemplate("inventory/supplier/supplier_detail.html")
}

func (c InventorySupplier) Save() revel.Result {
    session := c.DbSession

    form := new(SupplierForm)
    c.Params.Bind(form, "form")

    if form.Valid(c.Validation) == false {
        return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage() })
    }

    supplier := &form.Supplier

    var affected int64
    if form.IsCreate() { 
        nameCount, err := session.Where("sup_name = ?", supplier.Name).Count(new(models.Supplier))
            core.HandleError(err)
            if nameCount != 0 {
                return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，供应商名称已存在！" })
            }
        
        codeCount, err := session.Where("sup_code = ?", supplier.Code).Count(new(models.Supplier))
            core.HandleError(err)
            if codeCount != 0 {
                return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，供应商编号已存在！" })
            }
        
        affected, err = session.Insert(supplier)
        core.HandleError(err)
    } else { 
        nameCount, err := session.Where("id <> ? and sup_name = ?", supplier.Id, supplier.Name).Count(new(models.Supplier))
            core.HandleError(err)
            if nameCount != 0 {
                return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，供应商名称已存在！" })
            }
        
        codeCount, err := session.Where("id <> ? and sup_code = ?", supplier.Id, supplier.Code).Count(new(models.Supplier))
            core.HandleError(err)
            if codeCount != 0 {
                return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，供应商编号已存在！" })
            }
        
        affected, err = session.Id(supplier.Id).Update(supplier)
        core.HandleError(err)

        if affected == 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
        }
    }

    return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c InventorySupplier) Delete() revel.Result {
    session := c.DbSession

    supplierIdList := make([]int64, 0)
    c.Params.Bind(&supplierIdList, "id_list")

    supplier := new(models.Supplier)
    affected, err := session.In("id", supplierIdList).Delete(supplier)
    core.HandleError(err)

    return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}

