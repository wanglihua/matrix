package models

import (
	"matrix/core"
)

//工程师特有的其他字段信息
type EngineerInfo struct {
	Id         int64     `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	UserId     int64     `xorm:"bigint notnull index 'user_id'" json:"user_id" smith:"verbose_name=用户id"` //对应于auth用户表中的id
	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e EngineerInfo) TableName() string {
	return TablePrefix + "engineer"
}

func (e EngineerInfo) ModelDesc() string {
	return "verbose_name=工程师,entity_json=engineer,route=engineer"
}
