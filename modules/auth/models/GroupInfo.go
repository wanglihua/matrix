package models

import (
	"matrix/core"
)

type GroupInfo struct {
	Id         int64     `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	GroupName  string    `xorm:"nvarchar(255) notnull unique index 'group_name'" json:"group_name"`
	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e GroupInfo) TableName() string {
	return TablePrefix + "group"
}