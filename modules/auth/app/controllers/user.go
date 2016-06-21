package controllers

import (
    "github.com/revel/revel"
    "fmt"
    "matrix/db"
    "matrix/modules/auth/models"
    "matrix/service/gridrequest"
    "matrix/service"
)

type AuthUser struct {
    *revel.Controller
}

func (c AuthUser) Index() revel.Result {
    //session := db.Engine.NewSession()
    //defer session.Close()

    fmt.Println(c.Session.Id())

    db.Engine.Sync2(&models.User{}, &models.Group{}, &models.UserGroup{})

    return c.RenderTemplate("user/user_index.html")
}

func (c AuthUser) ListData() revel.Result {
    session := db.Engine.NewSession()
    defer session.Close()

    filter, orderList, offset, limit := GridRequest.GetParam(c.Request)

    query := session.Where(filter)
    dataQuery := *query
    countQuery := *query

    for _, orderParam := range orderList {
        if orderParam.OrderAsc {
            dataQuery = *dataQuery.Asc(orderParam.ColName)
        } else {
            dataQuery = *dataQuery.Desc(orderParam.ColName)
        }
    }

    users := make([]models.User, 0, limit)
    err := dataQuery.Limit(limit, offset).Find(&users)
    if err != nil {
        panic(err)
    }

    count, err := countQuery.Count(new(models.User))
    if err != nil {
        panic(err)
    }

    return c.RenderJson(service.GridResult{
        Data: users,
        Total: count,
    })
}
