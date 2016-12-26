package models

import (
	"matrix/core"
)

type StockInfo struct {
	Id         int64                 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	Name       string                `xorm:"nvarchar(200) notnull index 'name'" json:"name" smith:"verbose_name=名称,min=2,max=100"`
	ShortName  core.NullString       `xorm:"nvarchar(200) 'short_name'" json:"short_name" smith:"verbose_name=简称,min=2,max=100"`
	Contact    core.NullString       `xorm:"nvarchar(200) 'contact'" json:"contact" smith:"verbose_name=联系人,min=2,max=100"`
	Address    core.NullString       `xorm:"nvarchar(200) 'address'" json:"address" smith:"verbose_name=地址,min=2,max=200"`
	Phone      core.NullString       `xorm:"nvarchar(200) 'phone'" json:"phone" smith:"verbose_name=电话,min=2,max=100"`
	StatusId   int64                 `xorm:"bigint notnull index 'status_id'" json:"status_id" smith:"verbose_name=状态"`
	CreateTime core.Time             `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time             `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int                   `xorm:"version notnull 'version'" json:"version"`
}

func (e StockInfo) TableName() string {
	return TablePrefix + "stock"
}

func (e StockInfo) ModelDesc() string {
	return "verbose_name=仓库,entity_json=stock,route=stock"
}
