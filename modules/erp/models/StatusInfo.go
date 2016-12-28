package models

import (
	"matrix/core"
	"github.com/go-xorm/xorm"
)

const (

	//状态类型 begin

	// STATUS_TYPE_STOCK 仓库状态
	STATUS_TYPE_STOCK = 100

	// STATUS_TYPE_STOCK 仓库状态
	STATUS_TYPE_EXCHANGE_RATE = 200

	//状态类型 end

	//状态 id begin

	// 仓库状态 正常
	STOCK_STATUS_ENABLED = 100

	// 仓库状态 停用
	STOCK_STATUS_DISABLED = 200

	// 货币汇率状态 正常
	EXCHANGE_RATE_STATUS_ENABLED = 300

	// 货币汇率状态 停用
	EXCHANGE_RATE_STATUS_DISABLED = 400

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
	status.Desc =  core.NewNullString("仓库状态 正常", true)
	db_session.Insert(status)

	status = new(StatusInfo)
	status.Id = STOCK_STATUS_DISABLED
	status.Type = STATUS_TYPE_STOCK
	status.Name = "停用"
	status.Desc =  core.NewNullString("仓库状态 停用", true)
	db_session.Insert(status)


	//货币汇率状态

	status = new(StatusInfo)
	status.Id = EXCHANGE_RATE_STATUS_ENABLED
	status.Type = STATUS_TYPE_EXCHANGE_RATE
	status.Name = "正常"
	status.Desc =  core.NewNullString("货币汇率状态 正常", true)
	db_session.Insert(status)

	status = new(StatusInfo)
	status.Id = EXCHANGE_RATE_STATUS_DISABLED
	status.Type = STATUS_TYPE_EXCHANGE_RATE
	status.Name = "停用"
	status.Desc =  core.NewNullString("货币汇率状态 停用", true)
	db_session.Insert(status)
}
