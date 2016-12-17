package models

import (
	"matrix/core"
)

type EngineerEventTypeInfo struct {
	Id          int64     `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	EventTypeId int64     `xorm:"bigint notnull index 'type_id'" json:"type_id" smith:"verbose_name=事件类型id"`
	EngineerId  int64     `xorm:"bigint notnull index 'engineer_id'" json:"engineer_id" smith:"verbose_name=工程师id"`
	CreateTime  core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime  core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version     int       `xorm:"version notnull 'version'" json:"version"`
}

func (e EngineerEventTypeInfo) TableName() string {
	return TablePrefix + "engineer_event_type"
}

func (e EngineerEventTypeInfo) ModelDesc() string {
	return "verbose_name=工程师事件类型设定,entity_json=et,route=engineer/event/type"
}
