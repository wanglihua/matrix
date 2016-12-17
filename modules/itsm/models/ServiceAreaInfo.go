package models

import (
	"matrix/core"
)

type ServiceAreaInfo struct {
	Id         int64           `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	Name       string          `xorm:"nvarchar(50) notnull index 'area_name'" json:"area_name" smith:"verbose_name=名称,min=1,max=30"`
	Desc       core.NullString `xorm:"nvarchar(500) 'area_desc'" json:"area_desc" smith:"verbose_name=描述"`
	CreateTime core.Time       `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time       `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int             `xorm:"version notnull 'version'" json:"version"`
}

func (e ServiceAreaInfo) TableName() string {
	return TablePrefix + "service_area"
}

func (e ServiceAreaInfo) ModelDesc() string {
	return "verbose_name=服务区域,entity_json=area,route=service/area"
}
