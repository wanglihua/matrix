package models

import (
	"matrix/core"
)

type GroupUserInfo struct {
	Id int64 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	GroupId int64 `xorm:"bigint notnull index 'group_id'" json:"group_id"`
	UserId  int64 `xorm:"bigint notnull index 'user_id'" json:"user_id"`

	AddTime core.Time `xorm:"created notnull 'add_time'" json:"add_time"`
}

func (e GroupUserInfo) TableName() string {
	return TablePrefix + "group_user"
}
