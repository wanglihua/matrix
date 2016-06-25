package controllers

import (
    "github.com/revel/revel"

    "matrix/modules/auth/models"
    "matrix/service"
    "strconv"
    "time"
    "matrix/core"
    "fmt"
)

type AuthUser struct {
    *revel.Controller
    service.BaseController
}

func (c AuthUser) Index() revel.Result {
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
        dataQuery = *dataQuery.Asc("id")
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

    userId := service.GetInt64FromRequest(c.Request, "id")


    user := new(models.User)
    if userId != 0 {
        has, err := session.Id(userId).Get(user)
        service.HandleError(err)
        if has == false {
            panic("指定的用户不存在！")
        }
    }


    fmt.Println(time.Now())
    fmt.Println(time.Time(user.AddTime).Format("2006-01-02 15:04"))
    fmt.Println(time.Time(user.AddTime))

    c.RenderArgs["user"] = user
    return c.RenderTemplate("user/user_detail.html")
}

func (c AuthUser) Save() revel.Result {
    session := c.DbSession

    user := new(models.User)
    c.Params.Bind(&user, "user")

    user.AddTime = core.Time(time.Now())

    var affected int64
    var err error
    if user.Id == 0 {
        affected, err = session.Insert(user)
        service.HandleError(err)
    } else {
        affected, err = session.Id(user.Id).Update(user)
        //affected, err := session.Id(user.Id).Cols("nick_name").Update(user)
        //affected, err := session.Table(new(User)).Id(user.Id).Update(map[string]interface{}{"password":"123456"})
        service.HandleError(err)
    }

    return c.RenderJson(service.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c AuthUser) Delete() revel.Result {
    session := c.DbSession

    userIdList := make([]int, 0)
    c.Params.Bind(&userIdList, "id_list")

    user := new(models.User)
    affected, err := session.In("id", userIdList).Delete(user)
    service.HandleError(err)

    return c.RenderJson(service.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}
