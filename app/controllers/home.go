package controllers

import (
	"github.com/revel/revel"
	"matrix/app/models"
	"matrix/core"
	"matrix/db"
	auth_models "matrix/modules/auth/models"
	inventory_models "matrix/modules/inventory/models"
	oa_models "matrix/modules/oa/models"
	weixin_models "matrix/modules/weixin/models"
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
	modelList = append(modelList, &auth_models.User{})
	modelList = append(modelList, &auth_models.Admin{})
	modelList = append(modelList, &auth_models.Group{})
	modelList = append(modelList, &auth_models.GroupUser{})
	modelList = append(modelList, &weixin_models.Config{})
	modelList = append(modelList, &weixin_models.Menu{})
	modelList = append(modelList, &models.Config{})
	modelList = append(modelList, &inventory_models.Supplier{})
	modelList = append(modelList, &inventory_models.Client{})
	modelList = append(modelList, &inventory_models.ProductCate{})

	modelList = append(modelList, &inventory_models.Product{})
	modelList = append(modelList, &inventory_models.Stock{})
	modelList = append(modelList, &inventory_models.StorageLoc{})
	modelList = append(modelList, &inventory_models.CapitalAccount{})
	modelList = append(modelList, &inventory_models.StockIoType{})

	modelList = append(modelList, &inventory_models.PayType{})
	modelList = append(modelList, &inventory_models.Handler{})
	modelList = append(modelList, &inventory_models.StockBill{})
	modelList = append(modelList, &inventory_models.StockBillDetail{})
	modelList = append(modelList, &oa_models.Worklog{})

	err := db.Engine.Sync2(modelList...)
	core.HandleError(err)

	session := c.DbSession
	user := new(auth_models.User)
	user.LoginName = "admin"
	user.NickName = "管理员"
	user.Password = core.EncryptPassword("111111")
	session.Insert(user)

	admin := new(auth_models.Admin)
	admin.UserId = user.Id
	session.Insert(admin)

	return c.RenderJson(core.JsonResult{Success: true, Message: "数据库同步成功!"})
}
