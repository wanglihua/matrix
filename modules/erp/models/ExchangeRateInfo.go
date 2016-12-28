package models

import (
	"matrix/core"
)

type ExchangeRateInfo struct {
	Id         int64                 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	Code       string                `xorm:"nvarchar(20) notnull index 'code'" json:"code" smith:"verbose_name=符号,min=1,max=10"`
	Name       string                `xorm:"nvarchar(50) notnull index 'name'" json:"name" smith:"verbose_name=名称,min=1,max=20"`
	Rate       float64               `xorm:"float notnull 'rate'" json:"rate" smith:"verbose_name=汇率,min=0"`
	StatusId   int64                 `xorm:"bigint notnull index 'status_id'" json:"status_id" smith:"verbose_name=状态"`
	CreateTime core.Time             `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time             `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int                   `xorm:"version notnull 'version'" json:"version"`
}

func (e ExchangeRateInfo) TableName() string {
	return TablePrefix + "exchange_rate"
}

func (e ExchangeRateInfo) ModelDesc() string {
	return "verbose_name=货币汇率,entity_json=exchange_rate,route=exchange/rate"
}
