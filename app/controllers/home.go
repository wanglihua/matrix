package controllers

import (
    "github.com/revel/revel"
    "matrix/core"
    "matrix/db"
    models "matrix/app/models"
    userModels "matrix/modules/auth/models"
    weixinModels "matrix/modules/weixin/models"
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
        &userModels.User{},
        &userModels.Admin{},
        &userModels.Group{},
        &userModels.UserGroup{},
        &weixinModels.Config{},
        &models.Config{},
    )

    session := c.DbSession
    user := new(userModels.User)
    user.LoginName = "admin"
    user.NickName = "管理员"
    user.Password = core.EncryptPassword("111111")
    session.Insert(user)

    admin := new(userModels.Admin)
    admin.UserId = user.Id
    session.Insert(admin)

    return c.RenderJson(core.JsonResult{Success:true, Message:"数据库同步成功!"})
}
