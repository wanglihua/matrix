package models

import (
	"matrix/core"
)

type EngineerServiceAreaInfo struct {
	Id            int64     `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	ServiceAreaId int64     `xorm:"bigint notnull index 'service_area_id'" json:"service_area_id" smith:"verbose_name=服务区域id"`
	EngineerId    int64     `xorm:"bigint notnull index 'engineer_id'" json:"engineer_id" smith:"verbose_name=工程师id"`
	CreateTime    core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime    core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version       int       `xorm:"version notnull 'version'" json:"version"`
}

func (e EngineerServiceAreaInfo) TableName() string {
	return TablePrefix + "engineer_service_area"
}

func (e EngineerServiceAreaInfo) ModelDesc() string {
	return "verbose_name=工程师服务区域设定,entity_json=area,route=engineer/service/area"
}
