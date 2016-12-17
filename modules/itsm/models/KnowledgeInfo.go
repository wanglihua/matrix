package models

import (
	"matrix/core"
)

type KnowledgeInfo struct {
	Id         int64     `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	CateId     int64     `xorm:"bigint notnull index 'cate_id'" json:"cate_id" smith:"verbose_name=类别"`
	Title      string    `xorm:"nvarchar(200) notnull index 'title'" json:"title" smith:"verbose_name=标题,min=2,max=20"`
	Content    string    `xorm:"text notnull 'content'" json:"content" smith:"verbose_name=内容"`
	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e KnowledgeInfo) TableName() string {
	return TablePrefix + "knowledge"
}

func (e KnowledgeInfo) ModelDesc() string {
	return "verbose_name=知识库,entity_json=knowledge,route=knowledge"
}
