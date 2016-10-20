package models

import (
	"matrix/core"
)

var TablePrefix = "hd_sys_"

//---------------------------------------------------------------------------------------------------------------

type Config struct {
	Id         int64     `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	SysName    string    `xorm:"nvarchar(255) notnull 'sys_name'" json:"sys_name"`
	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e Config) TableName() string {
	return TablePrefix + "config"
}
