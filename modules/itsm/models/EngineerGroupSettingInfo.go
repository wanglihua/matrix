package models

import (
	"matrix/core"
)

type EngineerGroupSettingInfo struct {
	Id         int64     `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	GroupId    int64     `xorm:"bigint notnull index 'group_id'" json:"group_id" smith:"verbose_name=分组id"`
	EngineerId int64     `xorm:"bigint notnull index 'engineer_id'" json:"engineer_id" smith:"verbose_name=工程师id"`
	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e EngineerGroupSettingInfo) TableName() string {
	return TablePrefix + "engineer_group_setting"
}

func (e EngineerGroupSettingInfo) ModelDesc() string {
	return "verbose_name=工程师分组设定,entity_json=setting,route=engineer/group/setting"
}
