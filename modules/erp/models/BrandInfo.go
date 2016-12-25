package models

import (
	"matrix/core"
)

type BrandInfo struct {
	Id         int64                 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	Name       string                `xorm:"nvarchar(200) notnull index 'name'" json:"name" smith:"verbose_name=名称,min=2,max=100"`
	Contact    core.NullString       `xorm:"nvarchar(200) 'contact'" json:"contact" smith:"verbose_name=联系人,min=2,max=100"`
	Address    core.NullString       `xorm:"nvarchar(200) 'address'" json:"address" smith:"verbose_name=地址,min=2,max=100"`
	Phone      core.NullString       `xorm:"nvarchar(200) 'phone'" json:"phone" smith:"verbose_name=电话,min=2,max=100"`
	CreateTime core.Time             `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time             `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int                   `xorm:"version notnull 'version'" json:"version"`
}

func (e BrandInfo) TableName() string {
	return TablePrefix + "brand"
}

func (e BrandInfo) ModelDesc() string {
	return "verbose_name=品牌,entity_json=brand,route=brand"
}
