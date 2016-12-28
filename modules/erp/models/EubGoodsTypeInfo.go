package models

import (
	"matrix/core"
)

type EubGoodsTypeInfo struct {
	Id         int64                 `xorm:"bigint notnull pk 'id'" json:"id"`
	Name       string                `xorm:"nvarchar(200) notnull index 'name'" json:"name" smith:"verbose_name=名称,min=2,max=100"`
	Desc       core.NullString       `xorm:"nvarchar(500) 'desc'" json:"desc" smith:"verbose_name=描述,min=2,max=400"`
	CreateTime core.Time             `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time             `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int                   `xorm:"version notnull 'version'" json:"version"`
}

func (e EubGoodsTypeInfo) TableName() string {
	return TablePrefix + "eub_goods_type"
}

func (e EubGoodsTypeInfo) ModelDesc() string {
	return "verbose_name=EUB物品类型,entity_json=eub_goods_type,route=eub/goods/type"
}
