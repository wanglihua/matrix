package controllers

import (
    "github.com/revel/revel"
    "fmt"
    "matrix/db"
    "matrix/modules/auth/models"
    "matrix/serice/gridrequest"
)

type AuthUser struct {
    *revel.Controller
}

func (c AuthUser) Index() revel.Result {
    //session := db.Engine.NewSession()
    //defer session.Close()

    fmt.Println(c.Session.Id())

    /*
    db.Engine.Sync2(&models.User{}, &models.Group{}, &models.UserGroup{})
    db.Engine.CreateIndexes(&models.User{})
    db.Engine.CreateUniques(&models.User{})
    */

    return c.RenderTemplate("user/user_index.html")
}

//type GridResponse map[string]interface{}


func (c AuthUser) ListData() revel.Result {
    session := db.Engine.NewSession()
    defer session.Close()

    filter, orderList, offset, limit := GridRequest.GetParam(c.Request)

    query := session.Where(filter)
    dataQuery := *query
    countQuery := *query

    fmt.Println(orderList)
    /*
    for _, orderParam := range orderList {
        if orderParam.OrderAsc {
            dataQuery = *dataQuery.Asc(orderParam.OrderColumn)
        } else {
            dataQuery = *dataQuery.Desc(orderParam.OrderColumn)
        }
    }
    */

    users := make([]models.User, 0)
    dataQuery.Limit(limit, offset).Find(&users)
    count, _ := countQuery.Count(new(models.User))

    gridResult := map[string]interface{}{
        "aaData": users,
        "iTotalDisplayRecords":count,
        "iTotalRecords":count,
    }

    return c.RenderJson(gridResult)
}
