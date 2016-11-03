package models

import (
	"matrix/core"
)

var ModuleTitleName = "Repair"
var ModuleLowerCase = "repair"
var ModuleChinese = "维修管理"

var TablePrefix = "hd_repair_"


//---------------------------------------------------------------------------------------------------------------

type KnowledgeCate struct {
	Id         int64            `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	
	Name       string           `xorm:"nvarchar(50) notnull index 'cate_name'" json:"cate_name" smith:"verbose_name=标题,min=2,max=20"`
	Desc       core.NullString  `xorm:"text 'cate_desc'" json:"cate_desc" smith:"verbose_name=描述"`
	ParentId   int64            `xorm:"bigint notnull index 'parent_id'" json:"parent_id" smith:"verbose_name=父类"`
	
	CreateTime core.Time        `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time        `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int              `xorm:"version notnull 'version'" json:"version"`
}

func (e KnowledgeCate) TableName() string {
	return TablePrefix + "knowledge_cate"
}

func (e KnowledgeCate) ModelDesc() string {
	return "verbose_name=知识库类别,entity_json=cate,route=knowledge/cate"
}

//---------------------------------------------------------------------------------------------------------------

type Knowledge struct {
	Id         int64     `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	
	CateId     int64     `xorm:"bigint notnull index 'cate_id'" json:"cate_id" smith:"verbose_name=类别"`
	Title      string    `xorm:"nvarchar(200) notnull index 'title'" json:"title" smith:"verbose_name=标题,min=2,max=20"`
	Content    string    `xorm:"text notnull 'content'" json:"content" smith:"verbose_name=内容"`
	
	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e Knowledge) TableName() string {
	return TablePrefix + "knowledge"
}

func (e Knowledge) ModelDesc() string {
	return "verbose_name=知识库,entity_json=knowledge,route=knowledge"
}

//---------------------------------------------------------------------------------------------------------------
