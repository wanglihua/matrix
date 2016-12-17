package models

import (
	"matrix/core"
)

type ApplicationInfo struct {
	Id         int64     `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	Code       string    `xorm:"nvarchar(50) notnull index unique 'code'" json:"code" smith:"verbose_name=编码,min=2,max=20"`
	Name       string    `xorm:"nvarchar(50) notnull index 'name'" json:"name" smith:"verbose_name=名称,min=2,max=20"`
	StatusId   int64     `xorm:"bigint notnull index 'status_id'" json:"status_id" smith:"verbose_name=状态"`
	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e ApplicationInfo) TableName() string {
	return TablePrefix + "application"
}

func (e ApplicationInfo) ModelDesc() string {
	return "verbose_name=应用,entity_json=application,route=application"
}
