package models

import (
	"matrix/core"
)

type KnowledgeCateInfo struct {
	Id         int64           `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	Name       string          `xorm:"nvarchar(50) notnull index 'cate_name'" json:"cate_name" smith:"verbose_name=名称,min=2,max=20"`
	Desc       core.NullString `xorm:"nvarchar(500) 'cate_desc'" json:"cate_desc" smith:"verbose_name=描述"`
	ParentId   int64           `xorm:"bigint notnull index 'parent_id'" json:"parent_id" smith:"verbose_name=父类"`
	CreateTime core.Time       `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time       `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int             `xorm:"version notnull 'version'" json:"version"`
}

func (e KnowledgeCateInfo) TableName() string {
	return TablePrefix + "knowledge_cate"
}

func (e KnowledgeCateInfo) ModelDesc() string {
	return "verbose_name=知识库类别,entity_json=cate,route=knowledge/cate"
}
