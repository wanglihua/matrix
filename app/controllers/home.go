package controllers

import (
    "github.com/revel/revel"
    "matrix/service"
    "matrix/modules/auth/models"
    "matrix/db"
)

type Home struct {
    *revel.Controller
    service.BaseController
}

func (c Home) Index() revel.Result {
    return c.RenderTemplate("home/home_index.html")
}

func (c Home) SyncDb() revel.Result {
    return c.RenderTemplate("home/home_syncdb.html")
}

func (c Home) SyncDbPost() revel.Result {
    db.Engine.Sync2(
        &models.User{},
        &models.Admin{},
        &models.Group{},
        &models.UserGroup{},
    )
    return c.RenderJson(service.JsonResult{Success:true, Message:"数据库同步成功!"})
}
