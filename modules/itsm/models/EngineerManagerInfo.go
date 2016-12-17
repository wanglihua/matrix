package models

import (
	"matrix/core"
)

//别的可以考虑的字段，可以是 事件类型id，服务区域id等等
type EngineerManagerInfo struct {
	Id         int64     `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	EngineerId int64     `xorm:"bigint notnull index 'engineer_id'" json:"engineer_id" smith:"verbose_name=工程师id"`
	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e EngineerManagerInfo) TableName() string {
	return TablePrefix + "engineer_manager"
}

func (e EngineerManagerInfo) ModelDesc() string {
	return "verbose_name=工程师经理,entity_json=manager,route=engineer/manager"
}
