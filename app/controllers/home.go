package controllers

import (
	"github.com/revel/revel"
	"matrix/app/models"
	"matrix/core"
	"matrix/db"
	authModels "matrix/modules/auth/models"
	inventoryModels "matrix/modules/inventory/models"
	oaModels "matrix/modules/oa/models"
	weixinModels "matrix/modules/weixin/models"
)

type Home struct {
	*revel.Controller
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
	modelList = append(modelList, &inventoryModels.Client{})
	modelList = append(modelList, &inventoryModels.ProductCate{})
	modelList = append(modelList, &inventoryModels.Product{})
	modelList = append(modelList, &inventoryModels.Stock{})
	modelList = append(modelList, &inventoryModels.StorageLoc{})
	modelList = append(modelList, &inventoryModels.CapitalAccount{})
	modelList = append(modelList, &inventoryModels.StockIoType{})
	modelList = append(modelList, &inventoryModels.PayType{})
	modelList = append(modelList, &inventoryModels.Handler{})
	modelList = append(modelList, &inventoryModels.StockBill{})
	modelList = append(modelList, &inventoryModels.StockBillDetail{})
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

	return c.RenderJson(core.JsonResult{Success: true, Message: "数据库同步成功!"})
}
