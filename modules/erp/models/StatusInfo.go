package models

import (
	"matrix/core"
	"github.com/go-xorm/xorm"
)

const (
	//状态类型 begin

	STATUS_TYPE_STOCK         = 100 //仓库状态
	STATUS_TYPE_EXCHANGE_RATE = 200 //仓库状态
	STATUS_TYPE_PRODUCT_CATE  = 300 //商品目录状态
	STATUS_TYPE_PRODUCT       = 400 //商品状态

	//状态类型 end

	//状态 id begin

	// 仓库状态 正常 停用
	STOCK_STATUS_ENABLED  = 100
	STOCK_STATUS_DISABLED = 200

	// 货币汇率状态 正常 停用
	EXCHANGE_RATE_STATUS_ENABLED  = 300
	EXCHANGE_RATE_STATUS_DISABLED = 400

	// 商品目录状态 正常 停用
	PRODUCT_CATE_STATUS_ENABLED  = 500
	PRODUCT_CATE_STATUS_DISABLED = 600

	//商品状态 重点维护 待开发 正常 清仓 停产 自动创建 暂停销售
	PRODUCT_STATUS_ZDWH = 700
	PRODUCT_STATUS_DKF  = 800
	PRODUCT_STATUS_ZC   = 900
	PRODUCT_STATUS_QC   = 1000
	PRODUCT_STATUS_TC   = 1100
	PRODUCT_STATUS_ZDCJ = 1200
	PRODUCT_STATUS_ZTXC = 1300

	//状态 id end
)

type StatusInfo struct {
	Id         int64                 `xorm:"bigint notnull pk 'id'" json:"id"`
	Type       int64                 `xorm:"bigint notnull index 'type'" json:"type" smith:"verbose_name=类型"`
	Name       string                `xorm:"nvarchar(200) notnull index 'name'" json:"name" smith:"verbose_name=名称,min=2,max=100"`
	Desc       core.NullString       `xorm:"nvarchar(500) 'desc'" json:"desc" smith:"verbose_name=描述,min=2,max=400"`
	CreateTime core.Time             `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time             `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int                   `xorm:"version notnull 'version'" json:"version"`
}

func (e StatusInfo) TableName() string {
	return TablePrefix + "status"
}

func (e StatusInfo) ModelDesc() string {
	return "verbose_name=状态,entity_json=status,route=status"
}

func (e StatusInfo) InitData(db_session *xorm.Session) {
	//仓库状态

	status := new(StatusInfo)
	status.Id = STOCK_STATUS_ENABLED
	status.Type = STATUS_TYPE_STOCK
	status.Name = "正常"
	status.Desc = core.NewNullString("仓库状态 正常", true)
	db_session.Insert(status)

	status = new(StatusInfo)
	status.Id = STOCK_STATUS_DISABLED
	status.Type = STATUS_TYPE_STOCK
	status.Name = "停用"
	status.Desc = core.NewNullString("仓库状态 停用", true)
	db_session.Insert(status)

	//货币汇率状态

	status = new(StatusInfo)
	status.Id = EXCHANGE_RATE_STATUS_ENABLED
	status.Type = STATUS_TYPE_EXCHANGE_RATE
	status.Name = "正常"
	status.Desc = core.NewNullString("货币汇率状态 正常", true)
	db_session.Insert(status)

	status = new(StatusInfo)
	status.Id = EXCHANGE_RATE_STATUS_DISABLED
	status.Type = STATUS_TYPE_EXCHANGE_RATE
	status.Name = "停用"
	status.Desc = core.NewNullString("货币汇率状态 停用", true)
	db_session.Insert(status)

	//商品目录状态

	status = new(StatusInfo)
	status.Id = PRODUCT_CATE_STATUS_ENABLED
	status.Type = STATUS_TYPE_PRODUCT_CATE
	status.Name = "正常"
	status.Desc = core.NewNullString("商品目录状态 正常", true)
	db_session.Insert(status)

	status = new(StatusInfo)
	status.Id = PRODUCT_CATE_STATUS_DISABLED
	status.Type = STATUS_TYPE_PRODUCT_CATE
	status.Name = "停用"
	status.Desc = core.NewNullString("商品目录状态 停用", true)
	db_session.Insert(status)

	//商品状态

	status = new(StatusInfo)
	status.Id = PRODUCT_STATUS_ZDWH
	status.Type = STATUS_TYPE_PRODUCT
	status.Name = "重点维护"
	status.Desc = core.NewNullString("商品状态 重点维护", true)
	db_session.Insert(status)

	status = new(StatusInfo)
	status.Id = PRODUCT_STATUS_DKF
	status.Type = STATUS_TYPE_PRODUCT
	status.Name = "待开发"
	status.Desc = core.NewNullString("商品状态 待开发", true)
	db_session.Insert(status)

	status = new(StatusInfo)
	status.Id = PRODUCT_STATUS_ZC
	status.Type = STATUS_TYPE_PRODUCT
	status.Name = "正常"
	status.Desc = core.NewNullString("商品状态 正常", true)
	db_session.Insert(status)

	status = new(StatusInfo)
	status.Id = PRODUCT_STATUS_QC
	status.Type = STATUS_TYPE_PRODUCT
	status.Name = "清仓"
	status.Desc = core.NewNullString("商品状态 清仓", true)
	db_session.Insert(status)

	status = new(StatusInfo)
	status.Id = PRODUCT_STATUS_TC
	status.Type = STATUS_TYPE_PRODUCT
	status.Name = "停产"
	status.Desc = core.NewNullString("商品状态 停产", true)
	db_session.Insert(status)

	status = new(StatusInfo)
	status.Id = PRODUCT_STATUS_ZDCJ
	status.Type = STATUS_TYPE_PRODUCT
	status.Name = "自动创建"
	status.Desc = core.NewNullString("商品状态 自动创建", true)
	db_session.Insert(status)

	status = new(StatusInfo)
	status.Id = PRODUCT_STATUS_ZTXC
	status.Type = STATUS_TYPE_PRODUCT
	status.Name = "暂停销售"
	status.Desc = core.NewNullString("商品状态 暂停销售", true)
	db_session.Insert(status)
}
