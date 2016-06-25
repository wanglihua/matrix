package controllers

import (
    "github.com/revel/revel"

    "matrix/service"
    "matrix/modules/auth/models"
)

type AuthAdmin struct {
    *revel.Controller
    service.BaseController
}

func (c AuthAdmin) Index() revel.Result {
    return c.RenderTemplate("admin/admin_index.html")
}

type AdminView struct {
    models.User     `xorm:"extends" json:"u"`
    models.Admin    `xorm:"extends" json:"a"`
}

func (c AuthAdmin) ListData() revel.Result {
    session := c.DbSession

    filter, order, offset, limit := service.GetGridRequestParam(c.Request)
    query := session.
    Select("u.*, a.*").
    Table(models.TablePrefix + "user").Alias("u").
    Join("inner", []string{models.TablePrefix + "admin", "a"}, "u.id = a.user_id").
    Where(filter)

    //query extra filter here

    dataQuery := *query
    if order != "" {
        dataQuery = *dataQuery.OrderBy(order)
    } else {
        dataQuery = *dataQuery.Asc("admin.id")
    }

    adminList := make([]AdminView, 0, limit)
    err := dataQuery.Limit(limit, offset).Find(&adminList)
    service.HandleError(err)

    countQuery := *query
    count, err := countQuery.Count(new(AdminView))
    service.HandleError(err)

    return c.RenderJson(service.GridResult{
        Data:  adminList,
        Total: count,
    })
}

func (c AuthAdmin) Add() revel.Result {
    //session := c.DbSession
    return c.RenderTemplate("admin/admin_add.html")
}

func (c AuthAdmin) AddSave() revel.Result {
    //session := c.DbSession
    return c.RenderJson(service.JsonResult{Success: true, Message: "添加成功!"})
}

func (c AuthAdmin) Remove() revel.Result {
    //session := c.DbSession
    return c.RenderJson(service.JsonResult{Success: true, Message: "移除成功!"})
}
