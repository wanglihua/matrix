
package controllers

import (
    "strconv"
    "github.com/revel/revel"

    "matrix/core"
    "matrix/modules/inventory/models"
)

type InventoryClient struct {
    *revel.Controller
}

func (c InventoryClient) Index() revel.Result {
    return c.RenderTemplate("inventory/client/client_index.html")
}

func (c InventoryClient) ListData() revel.Result {
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

    clientList := make([]models.Client, 0, limit)
    err := dataQuery.Limit(limit, offset).Find(&clientList)
    core.HandleError(err)

    countQuery := *query
    count, err := countQuery.Count(new(models.Client))
    core.HandleError(err)

    return c.RenderJson(core.GridResult{
        Data:  clientList,
        Total: count,
    })
}

type ClientForm struct {
    Client models.Client
}

func (f *ClientForm) IsCreate() bool {
    return f.Client.Id == 0
}

func (f *ClientForm) Valid(validation *revel.Validation) bool { 
    validation.Required(f.Client.Name).Message("客户名称不能为空！")
    if f.Client.Name != "" {
        validation.MinSize(f.Client.Name, 3).Message("客户名称长度不能小于3！")
    }
    if f.Client.Name != "" {
        validation.MaxSize(f.Client.Name, 30).Message("客户名称长度不能大于30！")
    }

    validation.Required(f.Client.Code).Message("客户编号不能为空！")
    if f.Client.Code != "" {
        validation.MinSize(f.Client.Code, 3).Message("客户编号长度不能小于3！")
    }
    if f.Client.Code != "" {
        validation.MaxSize(f.Client.Code, 10).Message("客户编号长度不能大于10！")
    }

    if f.Client.Contact != "" {
        validation.MaxSize(f.Client.Contact, 20).Message("联系人长度不能大于20！")
    }

    if f.Client.Phone != "" {
        validation.MinSize(f.Client.Phone, 2).Message("联系电话长度不能小于2！")
    }
    if f.Client.Phone != "" {
        validation.MaxSize(f.Client.Phone, 40).Message("联系电话长度不能大于40！")
    }

    if f.Client.Address != "" {
        validation.MinSize(f.Client.Address, 3).Message("联系地址长度不能小于3！")
    }
    if f.Client.Address != "" {
        validation.MaxSize(f.Client.Address, 20).Message("联系地址长度不能大于20！")
    }

    if f.Client.Bank != "" {
        validation.MaxSize(f.Client.Bank, 20).Message("开户银行长度不能大于20！")
    }

    if f.Client.BankAccount != "" {
        validation.MinSize(f.Client.BankAccount, 4).Message("银行账号长度不能小于4！")
    }
    if f.Client.BankAccount != "" {
        validation.MaxSize(f.Client.BankAccount, 50).Message("银行账号长度不能大于50！")
    }

    if f.Client.Remark != "" {
        validation.MinSize(f.Client.Remark, 3).Message("备注长度不能小于3！")
    }
    if f.Client.Remark != "" {
        validation.MaxSize(f.Client.Remark, 250).Message("备注长度不能大于250！")
    }

    return validation.HasErrors() == false
}

func (c InventoryClient) Detail() revel.Result {
    session := c.DbSession

    var clientId int64
    c.Params.Bind(&clientId, "id")

    client := new(models.Client)
    if clientId != 0 {
        has, err := session.Id(clientId).Get(client)
        core.HandleError(err)
        if has == false {
            panic("指定的客户不存在！")
        }
    }

    form := new(ClientForm)
    form.Client = *client

    c.UnbindToRenderArgs(form, "form")
    return c.RenderTemplate("inventory/client/client_detail.html")
}

func (c InventoryClient) Save() revel.Result {
    session := c.DbSession

    form := new(ClientForm)
    c.Params.Bind(form, "form")

    if form.Valid(c.Validation) == false {
        return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage() })
    }

    client := &form.Client

    var affected int64
    if form.IsCreate() { 
        nameCount, err := session.Where("client_name = ?", client.Name).Count(new(models.Client))
            core.HandleError(err)
            if nameCount != 0 {
                return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，客户名称已存在！" })
            }
        
        codeCount, err := session.Where("client_code = ?", client.Code).Count(new(models.Client))
            core.HandleError(err)
            if codeCount != 0 {
                return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，客户编号已存在！" })
            }
        
        affected, err = session.Insert(client)
        core.HandleError(err)
    } else { 
        nameCount, err := session.Where("id <> ? and client_name = ?", client.Id, client.Name).Count(new(models.Client))
            core.HandleError(err)
            if nameCount != 0 {
                return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，客户名称已存在！" })
            }
        
        codeCount, err := session.Where("id <> ? and client_code = ?", client.Id, client.Code).Count(new(models.Client))
            core.HandleError(err)
            if codeCount != 0 {
                return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，客户编号已存在！" })
            }
        
        affected, err = session.Id(client.Id).Update(client)
        core.HandleError(err)

        if affected == 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
        }
    }

    return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c InventoryClient) Delete() revel.Result {
    session := c.DbSession

    clientIdList := make([]int64, 0)
    c.Params.Bind(&clientIdList, "id_list")

    client := new(models.Client)
    affected, err := session.In("id", clientIdList).Delete(client)
    core.HandleError(err)

    return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}

