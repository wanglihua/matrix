package models

import (
	"matrix/core"
)

type EventApplyDepartmentInfo struct {
	Id          int64             `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	Name        string            `xorm:"nvarchar(50) notnull index 'name'" json:"name" smith:"verbose_name=名称,min=2,max=20"`
	Description core.NullString   `xorm:"nvarchar(300) null 'description'" json:"description" smith:"verbose_name=描述,max=30"`
	CreateTime  core.Time         `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime  core.Time         `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version     int               `xorm:"version notnull 'version'" json:"version"`
}

func (e EventApplyDepartmentInfo) TableName() string {
	return TablePrefix + "event_apply_department"
}

func (e EventApplyDepartmentInfo) ModelDesc() string {
	return "verbose_name=事件提报部门,entity_json=apply_department,route=event/apply/department"
}
