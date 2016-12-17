package models

import (
	"matrix/core"
)

type EquipmentTypeInfo struct {
	Id         int64     `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	Name       string    `xorm:"nvarchar(50) notnull index 'name'" json:"name" smith:"verbose_name=名称,min=2,max=20"`
	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e EquipmentTypeInfo) TableName() string {
	return TablePrefix + "equipment_type"
}

func (e EquipmentTypeInfo) ModelDesc() string {
	return "verbose_name=设备类型,entity_json=equipment_type,route=equipment/type"
}
