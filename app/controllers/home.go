package controllers

import (
    "github.com/revel/revel"
    "matrix/core"
	"matrix/service"
    "matrix/db"
    models "matrix/app/models"
    authModels "matrix/modules/auth/models"
    weixinModels "matrix/modules/weixin/models"
    inventoryModels "matrix/modules/inventory/models"
    oaModels "matrix/modules/oa/models"
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

    modelList := make([]interface{}, 0)
    modelList = append(modelList, &authModels.User{})
    modelList = append(modelList, &authModels.Admin{})
    modelList = append(modelList, &authModels.Group{})
    modelList = append(modelList, &authModels.GroupUser{})
    modelList = append(modelList, &weixinModels.Config{})
    modelList = append(modelList, &weixinModels.Menu{})
    modelList = append(modelList, &models.Config{})
    modelList = append(modelList, &inventoryModels.Supplier{})
    modelList = append(modelList, &inventoryModels.Stock{})
    modelList = append(modelList, &inventoryModels.StorageLoc{})
    modelList = append(modelList, &oaModels.Worklog{})

    db.Engine.Sync2(modelList...)

    session := c.DbSession
    user := new(authModels.User)
    user.LoginName = "admin"
    user.NickName = "管理员"
    user.Password = core.EncryptPassword("111111")
    session.Insert(user)

    admin := new(authModels.Admin)
    admin.UserId = user.Id
    session.Insert(admin)

    return c.RenderJson(core.JsonResult{Success:true, Message:"数据库同步成功!"})
}
