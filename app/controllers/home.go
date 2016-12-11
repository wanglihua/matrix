package controllers

import (
	"matrix/app/models"
	"matrix/core"
	"matrix/db"
	auth_models "matrix/modules/auth/models"
	inventory_models "matrix/modules/inventory/models"
	itsm_models "matrix/modules/itsm/models"
	oa_models "matrix/modules/oa/models"
	weixin_models "matrix/modules/weixin/models"

	"github.com/revel/revel"
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

	session := c.DbSession
	user := new(auth_models.User)
	user.LoginName = "admin"
	user.NickName = "管理员"
	user.Password = core.EncryptPassword("111111")
	session.Insert(user)

	admin := new(auth_models.Admin)
	admin.UserId = user.Id
	session.Insert(admin)

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
	model_list = append(model_list, &itsm_models.EngineerGroup{})
	model_list = append(model_list, &itsm_models.Engineer{})
	model_list = append(model_list, &itsm_models.EngineerGroupSetting{})
	model_list = append(model_list, &itsm_models.EngineerServiceArea{})
	model_list = append(model_list, &itsm_models.EngineerEventType{})
	model_list = append(model_list, &itsm_models.EngineerManager{})
	model_list = append(model_list, &itsm_models.Event{})
	model_list = append(model_list, &itsm_models.EventType{})
	model_list = append(model_list, &itsm_models.EventStatus{})
	model_list = append(model_list, &itsm_models.EventPriority{})
	model_list = append(model_list, &itsm_models.EventLog{})
	model_list = append(model_list, &itsm_models.EquipmentStatus{})
	model_list = append(model_list, &itsm_models.EquipmentType{})
	model_list = append(model_list, &itsm_models.Equipment{})
	model_list = append(model_list, &itsm_models.ApplicationStatus{})
	model_list = append(model_list, &itsm_models.ApplicationType{})
	model_list = append(model_list, &itsm_models.Application{})

	err := db.Engine.Sync2(model_list...)
	core.HandleError(err)

	//event status insert begin

	event_status := new(itsm_models.EventStatus)
	event_status.Id = itsm_models.Event_Status_TBZ_Id
	event_status.Name = "提报中"
	event_status.Desc = core.NewNullString("提报中", true)
	session.Insert(event_status)

	event_status = new(itsm_models.EventStatus)
	event_status.Id = itsm_models.Event_Status_YPG_Id
	event_status.Name = "已派工"
	event_status.Desc = core.NewNullString("已派工", true)
	session.Insert(event_status)

	event_status = new(itsm_models.EventStatus)
	event_status.Id = itsm_models.Event_Status_YJS_Id
	event_status.Name = "已接受"
	event_status.Desc = core.NewNullString("已接受", true)
	session.Insert(event_status)

	event_status = new(itsm_models.EventStatus)
	event_status.Id = itsm_models.Event_Status_CLZ_Id
	event_status.Name = "处理中"
	event_status.Desc = core.NewNullString("处理中", true)
	session.Insert(event_status)

	event_status = new(itsm_models.EventStatus)
	event_status.Id = itsm_models.Event_Status_YWC_Id
	event_status.Name = "已完成"
	event_status.Desc = core.NewNullString("已完成", true)
	session.Insert(event_status)

	event_status = new(itsm_models.EventStatus)
	event_status.Id = itsm_models.Event_Status_YPJ_Id
	event_status.Name = "已评价"
	event_status.Desc = core.NewNullString("已评价", true)
	session.Insert(event_status)

	event_status = new(itsm_models.EventStatus)
	event_status.Id = itsm_models.Event_Status_YQC_Id
	event_status.Name = "已取消"
	event_status.Desc = core.NewNullString("已取消", true)
	session.Insert(event_status)

	//event status insert end

	return c.RenderJson(core.JsonResult{Success: true, Message: "数据库同步成功!"})
}
