package controllers

import (
    "fmt"
    "matrix/db"
    "matrix/modules/auth/models"
    "matrix/service"
    "matrix/service/gridrequest"

    "github.com/revel/revel"
    "strconv"
)

type AuthUser struct {
    *revel.Controller
}

func (c AuthUser) Index() revel.Result {
    //session := db.Engine.NewSession()
    //defer session.Close()

    fmt.Println(c.Session.Id())

    return c.RenderTemplate("user/user_index.html")
}

func (c AuthUser) ListData() revel.Result {
    session := db.Engine.NewSession()
    defer session.Close()

    filter, order, offset, limit := GridRequest.GetParam(c.Request)
    query := session.Where(filter)

    //query extra filter here

    dataQuery := *query
    if order != "" {
        dataQuery = *dataQuery.OrderBy(order)
    } else {
        dataQuery = *dataQuery.Asc("Id")
    }

    users := make([]models.User, 0, limit)
    err := dataQuery.Limit(limit, offset).Find(&users)
    service.HandleError(err)

    countQuery := *query
    count, err := countQuery.Count(new(models.User))
    service.HandleError(err)

    return c.RenderJson(service.GridResult{
        Data:  users,
        Total: count,
    })
}

func (c AuthUser) Detail() revel.Result {
    session := db.Engine.NewSession()
    defer session.Close()

    userId, err := strconv.ParseInt(c.Request.Form["id"][0], 10, 64) //可以做成一个通用函数 service package
    service.HandleError(err)

    user := new(models.User)
    if userId != 0 {
        has, err := session.Id(userId).Get(user)
        service.HandleError(err)
        if has == false {
            panic("指定的用户不存在！")
        }
    }

    c.Render(user)
    return c.RenderTemplate("user/user_detail.html")
}

func (c AuthUser) Save() revel.Result {

    return c.RenderJson(service.JsonResult{Success:true, Message:"保存成功!"})
}

func (c AuthUser) Delete() revel.Result {

    return c.RenderJson(service.JsonResult{Success:true, Message:"删除成功!"})
}