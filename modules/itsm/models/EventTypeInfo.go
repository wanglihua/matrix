package models

import (
	"matrix/core"
)

type EventTypeInfo struct {
	Id         int64           `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	Name       string          `xorm:"nvarchar(50) notnull index 'type_name'" json:"type_name" smith:"verbose_name=名称,min=1,max=30"`
	Desc       core.NullString `xorm:"nvarchar(500) 'type_desc'" json:"type_desc" smith:"verbose_name=描述"`
	CreateTime core.Time       `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time       `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int             `xorm:"version notnull 'version'" json:"version"`
}

func (e EventTypeInfo) TableName() string {
	return TablePrefix + "event_type"
}

func (e EventTypeInfo) ModelDesc() string {
	return "verbose_name=事件类型,entity_json=type,route=event/type"
}
