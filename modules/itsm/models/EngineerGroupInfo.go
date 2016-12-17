package models

import (
	"matrix/core"
)

type EngineerGroupInfo struct {
	Id         int64           `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	Name       string          `xorm:"nvarchar(50) notnull index 'group_name'" json:"group_name" smith:"verbose_name=名称,min=1,max=30"`
	Desc       core.NullString `xorm:"nvarchar(500) 'group_desc'" json:"group_desc" smith:"verbose_name=描述"`
	CreateTime core.Time       `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time       `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int             `xorm:"version notnull 'version'" json:"version"`
}

func (e EngineerGroupInfo) TableName() string {
	return TablePrefix + "engineer_group"
}

func (e EngineerGroupInfo) ModelDesc() string {
	return "verbose_name=工程师分组,entity_json=group,route=engineer/group"
}
