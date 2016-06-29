package controllers

import (
    "github.com/revel/revel"
    "matrix/core"
    "matrix/modules/auth/models"
    "matrix/db"
)

type Home struct {
    *revel.Controller
    core.BaseController
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

    session := c.DbSession
    user := new(models.User)
    user.LoginName = "admin"
    user.NickName = "管理员"
    user.Password = core.EncryptPassword("111111")
    session.Insert(user)

    admin := new(models.Admin)
    admin.UserId = user.Id
    session.Insert(admin)

    return c.RenderJson(core.JsonResult{Success:true, Message:"数据库同步成功!"})
}
