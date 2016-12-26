package controllers

import (
	"matrix/app/models"
	"matrix/core"
	"matrix/db"
	erp_models "matrix/modules/erp/models"
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

	model_list = append(model_list, &erp_models.StatusInfo{})
	model_list = append(model_list, &erp_models.BrandInfo{})
	model_list = append(model_list, &erp_models.StockInfo{})
	model_list = append(model_list, &erp_models.StorageLocInfo{})

	err := db.Engine.Sync2(model_list...)
	core.HandleError(err)

	//执行数据初始化
	db_session := c.DbSession
	for _, model := range model_list {
		if table_initer, ok := model.(core.TableDataIniter); ok {
			table_initer.InitData(db_session)
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: "数据库同步成功!"})
}
