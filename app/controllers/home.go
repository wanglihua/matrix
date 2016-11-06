package controllers

import (
	"github.com/revel/revel"
	"matrix/app/models"
	"matrix/core"
	"matrix/db"
	auth_models "matrix/modules/auth/models"
	inventory_models "matrix/modules/inventory/models"
	itsm_models "matrix/modules/itsm/models"
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

	model_list := make([]interface{}, 0)
	model_list = append(model_list, &auth_models.User{})
	model_list = append(model_list, &auth_models.Admin{})
	model_list = append(model_list, &auth_models.Group{})
	model_list = append(model_list, &auth_models.GroupUser{})
	model_list = append(model_list, &weixin_models.Config{})
	model_list = append(model_list, &weixin_models.Menu{})
	model_list = append(model_list, &models.Config{})
	model_list = append(model_list, &inventory_models.Supplier{})
	model_list = append(model_list, &inventory_models.Client{})
	model_list = append(model_list, &inventory_models.ProductCate{})

	model_list = append(model_list, &inventory_models.Product{})
	model_list = append(model_list, &inventory_models.Stock{})
	model_list = append(model_list, &inventory_models.StorageLoc{})
	model_list = append(model_list, &inventory_models.CapitalAccount{})
	model_list = append(model_list, &inventory_models.StockIoType{})

	model_list = append(model_list, &inventory_models.PayType{})
	model_list = append(model_list, &inventory_models.Handler{})
	model_list = append(model_list, &inventory_models.StockBill{})
	model_list = append(model_list, &inventory_models.StockBillDetail{})
	model_list = append(model_list, &inventory_models.Config{})
	model_list = append(model_list, &oa_models.Worklog{})

	model_list = append(model_list, &itsm_models.ServiceArea{})
	model_list = append(model_list, &itsm_models.EventType{})
	model_list = append(model_list, &itsm_models.Engineer{})

	err := db.Engine.Sync2(model_list...)
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
