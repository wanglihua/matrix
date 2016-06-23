package controllers

import (
    "fmt"
    "matrix/modules/auth/models"
    "matrix/service"

    "strconv"

    "github.com/revel/revel"
)

type AuthUser struct {
    *revel.Controller
    service.BaseController
}

func (c AuthUser) Index() revel.Result {
    //session := c.DbSession

    fmt.Println(c.Session.Id())

    return c.RenderTemplate("user/user_index.html")
}

func (c AuthUser) ListData() revel.Result {
    session := c.DbSession

    filter, order, offset, limit := service.GetGridRequestParam(c.Request)
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
    session := c.DbSession

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
    fmt.Println(userId)
    fmt.Println(user.NickName)
    c.RenderArgs["user"] = user
    return c.RenderTemplate("user/user_detail.html")
}

func (c AuthUser) Save() revel.Result {
    session := c.DbSession

    user := new(models.User)
    c.Params.Bind(&user, "user")

    if user.Id == 0 {
        _, err := session.Insert(user)
        service.HandleError(err)
    } else {
        _, err := session.Id(user.Id).Update(user)
        //_, err := session.Table(new(User)).Id(user.Id).Update(map[string]interface{}{"password":"123456"})
        service.HandleError(err)
    }

    fmt.Println("user save")
    fmt.Println(user.Id)
    fmt.Println(user.NickName)

    return c.RenderJson(service.JsonResult{Success: true, Message: "保存成功!"})
}

func (c AuthUser) Delete() revel.Result {
    session := c.DbSession

    userIdList := make([]int, 0)
    c.Params.Bind(&userIdList, "id_list")

    user := new(models.User)
    _, err := session.In("id", userIdList).Delete(user)
    service.HandleError(err)

    return c.RenderJson(service.JsonResult{Success: true, Message: "删除成功!"})
}
