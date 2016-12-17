package models

import (
	"matrix/core"
)

type EventPriorityInfo struct {
	Id         int64           `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	Name       string          `xorm:"nvarchar(50) notnull index 'priority_name'" json:"priority_name" smith:"verbose_name=名称,min=1,max=30"`
	Desc       core.NullString `xorm:"nvarchar(500) 'priority_desc'" json:"priority_desc" smith:"verbose_name=描述"`
	CreateTime core.Time       `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time       `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int             `xorm:"version notnull 'version'" json:"version"`
}

func (e EventPriorityInfo) TableName() string {
	return TablePrefix + "event_priority"
}

func (e EventPriorityInfo) ModelDesc() string {
	return "verbose_name=事件优先级,entity_json=priority,route=event/priority"
}
