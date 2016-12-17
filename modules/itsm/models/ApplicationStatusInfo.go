package models

import (
	"matrix/core"
)

type ApplicationStatusInfo struct {
	Id          int64           `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	Name        string          `xorm:"nvarchar(50) notnull index 'name'" json:"name" smith:"verbose_name=名称,min=2,max=20"`
	Description core.NullString `xorm:"nvarchar(500) 'description'" json:"description" smith:"verbose_name=描述, max=300"`
	CreateTime  core.Time       `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime  core.Time       `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version     int             `xorm:"version notnull 'version'" json:"version"`
}

func (e ApplicationStatusInfo) TableName() string {
	return TablePrefix + "application_status"
}

func (e ApplicationStatusInfo) ModelDesc() string {
	return "verbose_name=应用状态,entity_json=application_status,route=application/status"
}
