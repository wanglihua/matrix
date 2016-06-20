package controllers

import (
    "github.com/revel/revel"
    "fmt"
    "matrix/db"
    "matrix/modules/auth/models"
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

    users := make([]models.User, 0)
    session.Limit(10, 0).Find(&users)
    user := new(models.User)
    count, _ := session.Count(user)

    gridResult := map[string]interface{}{
        "aaData": users,
        "iTotalDisplayRecords":count,
        "iTotalRecords":count,
    }

    return c.RenderJson(gridResult)
}
