package models

import (
	"matrix/core"
)

type ApplicationTypeInfo struct {
	Id         int64     `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	Name       string    `xorm:"nvarchar(50) notnull index 'name'" json:"name" smith:"verbose_name=名称,min=2,max=20"`
	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e ApplicationTypeInfo) TableName() string {
	return TablePrefix + "application_type"
}

func (e ApplicationTypeInfo) ModelDesc() string {
	return "verbose_name=应用类型,entity_json=application_type,route=application/type"
}
