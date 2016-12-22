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

	model_list = append(model_list, &models.EntityCodeInfo{})

	model_list = append(model_list, &auth_models.UserInfo{})
	model_list = append(model_list, &auth_models.AdminInfo{})
	model_list = append(model_list, &auth_models.GroupInfo{})
	model_list = append(model_list, &auth_models.GroupUserInfo{})

	model_list = append(model_list, &weixin_models.ConfigInfo{})
	model_list = append(model_list, &weixin_models.MenuInfo{})
	model_list = append(model_list, &models.ConfigInfo{})
	model_list = append(model_list, &inventory_models.SupplierInfo{})
	model_list = append(model_list, &inventory_models.ClientInfo{})
	model_list = append(model_list, &inventory_models.ProductCateInfo{})

	model_list = append(model_list, &inventory_models.ProductInfo{})
	model_list = append(model_list, &inventory_models.StockInfo{})
	model_list = append(model_list, &inventory_models.StorageLocInfo{})
	model_list = append(model_list, &inventory_models.CapitalAccountInfo{})
	model_list = append(model_list, &inventory_models.StockIoTypeInfo{})

	model_list = append(model_list, &inventory_models.PayTypeInfo{})
	model_list = append(model_list, &inventory_models.HandlerInfo{})
	model_list = append(model_list, &inventory_models.StockBillInfo{})
	model_list = append(model_list, &inventory_models.StockBillDetailInfo{})
	model_list = append(model_list, &inventory_models.ConfigInfo{})
	model_list = append(model_list, &oa_models.WorklogInfo{})

	model_list = append(model_list, &itsm_models.ServiceAreaInfo{})
	model_list = append(model_list, &itsm_models.EngineerGroupInfo{})
	model_list = append(model_list, &itsm_models.EngineerInfo{})
	model_list = append(model_list, &itsm_models.EngineerGroupSettingInfo{})
	model_list = append(model_list, &itsm_models.EngineerServiceAreaInfo{})
	model_list = append(model_list, &itsm_models.EngineerEventTypeInfo{})
	model_list = append(model_list, &itsm_models.EngineerManagerInfo{})
	model_list = append(model_list, &itsm_models.EventInfo{})
	model_list = append(model_list, &itsm_models.EventTypeInfo{})
	model_list = append(model_list, &itsm_models.EventStatusInfo{})
	model_list = append(model_list, &itsm_models.EventPriorityInfo{})
	model_list = append(model_list, &itsm_models.EventApplyDepartmentInfo{})
	model_list = append(model_list, &itsm_models.EventLogInfo{})
	model_list = append(model_list, &itsm_models.EquipmentStatusInfo{})
	model_list = append(model_list, &itsm_models.EquipmentTypeInfo{})
	model_list = append(model_list, &itsm_models.EquipmentInfo{})
	model_list = append(model_list, &itsm_models.ApplicationStatusInfo{})
	model_list = append(model_list, &itsm_models.ApplicationTypeInfo{})
	model_list = append(model_list, &itsm_models.ApplicationInfo{})

	err := db.Engine.Sync2(model_list...)
	core.HandleError(err)

	session := c.DbSession
	user := new(auth_models.UserInfo)
	user.LoginName = "admin"
	user.NickName = "管理员"
	user.Password = core.EncryptPassword("111111")
	session.Insert(user)

	admin := new(auth_models.AdminInfo)
	admin.UserId = user.Id
	session.Insert(admin)

	//event status insert begin

	event_status := new(itsm_models.EventStatusInfo)
	event_status.Id = itsm_models.Event_Status_TBZ_Id
	event_status.Name = "提报中"
	event_status.Desc = core.NewNullString("提报中", true)
	session.Insert(event_status)

	event_status = new(itsm_models.EventStatusInfo)
	event_status.Id = itsm_models.Event_Status_YPG_Id
	event_status.Name = "已派工"
	event_status.Desc = core.NewNullString("已派工", true)
	session.Insert(event_status)

	event_status = new(itsm_models.EventStatusInfo)
	event_status.Id = itsm_models.Event_Status_YJS_Id
	event_status.Name = "已接受"
	event_status.Desc = core.NewNullString("已接受", true)
	session.Insert(event_status)

	event_status = new(itsm_models.EventStatusInfo)
	event_status.Id = itsm_models.Event_Status_CLZ_Id
	event_status.Name = "处理中"
	event_status.Desc = core.NewNullString("处理中", true)
	session.Insert(event_status)

	event_status = new(itsm_models.EventStatusInfo)
	event_status.Id = itsm_models.Event_Status_YWC_Id
	event_status.Name = "已完成"
	event_status.Desc = core.NewNullString("已完成", true)
	session.Insert(event_status)

	event_status = new(itsm_models.EventStatusInfo)
	event_status.Id = itsm_models.Event_Status_YPJ_Id
	event_status.Name = "已评价"
	event_status.Desc = core.NewNullString("已评价", true)
	session.Insert(event_status)

	event_status = new(itsm_models.EventStatusInfo)
	event_status.Id = itsm_models.Event_Status_YQC_Id
	event_status.Name = "已取消"
	event_status.Desc = core.NewNullString("已取消", true)
	session.Insert(event_status)

	//event status insert end

	return c.RenderJson(core.JsonResult{Success: true, Message: "数据库同步成功!"})
}
