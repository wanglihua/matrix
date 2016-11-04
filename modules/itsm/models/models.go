package models

import (
	"matrix/core"
)

var ModuleTitleName = "Itsm"
var ModuleLowerCase = "itsm"
var ModuleChinese = "ITSM"

var TablePrefix = "hd_itsm_"

//---------------------------------------------------------------------------------------------------------------

type EventType struct {
	Id         int64            `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	Name       string           `xorm:"nvarchar(50) notnull index 'type_name'" json:"type_name" smith:"verbose_name=名称,min=1,max=30"`
	Desc       core.NullString  `xorm:"nvarchar(500) 'type_desc'" json:"type_desc" smith:"verbose_name=描述"`

	CreateTime core.Time        `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time        `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int              `xorm:"version notnull 'version'" json:"version"`
}

func (e EventType) TableName() string {
	return TablePrefix + "event_type"
}

func (e EventType) ModelDesc() string {
	return "verbose_name=事件类型,entity_json=type,route=event/type"
}

//---------------------------------------------------------------------------------------------------------------

type ServiceArea struct {
	Id         int64            `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	Name       string           `xorm:"nvarchar(50) notnull index 'area_name'" json:"area_name" smith:"verbose_name=名称,min=1,max=30"`
	Desc       core.NullString  `xorm:"nvarchar(500) 'area_desc'" json:"area_desc" smith:"verbose_name=描述"`

	CreateTime core.Time        `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time        `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int              `xorm:"version notnull 'version'" json:"version"`
}

func (e ServiceArea) TableName() string {
	return TablePrefix + "service_area"
}

func (e ServiceArea) ModelDesc() string {
	return "verbose_name=服务区域,entity_json=area,route=service/area"
}

//---------------------------------------------------------------------------------------------------------------

type EngineerGroup struct {
	Id         int64            `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	Name       string           `xorm:"nvarchar(50) notnull index 'group_name'" json:"group_name" smith:"verbose_name=名称,min=1,max=30"`
	Desc       core.NullString  `xorm:"nvarchar(500) 'group_desc'" json:"group_desc" smith:"verbose_name=描述"`

	CreateTime core.Time        `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time        `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int              `xorm:"version notnull 'version'" json:"version"`
}

func (e EngineerGroup) TableName() string {
	return TablePrefix + "engineer_group"
}

func (e EngineerGroup) ModelDesc() string {
	return "verbose_name=工程师分组,entity_json=group,route=engineer/group"
}

//---------------------------------------------------------------------------------------------------------------
/*
ITSM_Engineer

	[EgrName] [varchar](20) NOT NULL, --user
	[EgrWxId] [varchar](20) NOT NULL, --user
	[Phone] [nvarchar](50) NULL, --user
	[sex] --user
	[IsMgr] [int] NULL,-- manager

*/

type Engineer struct {
	Id         int64        `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	Name       string       `xorm:"nvarchar(20) notnull index 'egr_name'" json:"egr_name" smith:"verbose_name=姓名,min=1,max=10"`
	Phone      string       `xorm:"nvarchar(50) null 'phone'" json:"phone" smith:"verbose_name=联系电话,min=1,max=30"`

	CreateTime core.Time    `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time    `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int          `xorm:"version notnull 'version'" json:"version"`
}

func (e Engineer) TableName() string {
	return TablePrefix + "engineer"
}

func (e Engineer) ModelDesc() string {
	return "verbose_name=工程师,entity_json=engineer,route=engineer"
}

//---------------------------------------------------------------------------------------------------------------

type EngineerGroupSetting struct {
	Id         int64        `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	GroupId    int64        `xorm:"bigint notnull index 'group_id'" json:"group_id" smith:"verbose_name=分组id"`
	EngineerId int64        `xorm:"bigint notnull index 'engineer_id'" json:"engineer_id" smith:"verbose_name=工程师id"`

	CreateTime core.Time    `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time    `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int          `xorm:"version notnull 'version'" json:"version"`
}

func (e EngineerGroupSetting) TableName() string {
	return TablePrefix + "engineer_group_setting"
}

func (e EngineerGroupSetting) ModelDesc() string {
	return "verbose_name=工程师分组设定,entity_json=setting,route=engineer/group/setting"
}

//---------------------------------------------------------------------------------------------------------------

type EngineerServiceArea struct {
	Id         int64        `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	ServiceId  int64        `xorm:"bigint notnull index 'service_id'" json:"service_id" smith:"verbose_name=服务区域id"`
	EngineerId int64        `xorm:"bigint notnull index 'engineer_id'" json:"engineer_id" smith:"verbose_name=工程师id"`

	CreateTime core.Time    `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time    `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int          `xorm:"version notnull 'version'" json:"version"`
}

func (e EngineerServiceArea) TableName() string {
	return TablePrefix + "engineer_service_area"
}

func (e EngineerServiceArea) ModelDesc() string {
	return "verbose_name=工程师服务区域设定,entity_json=area,route=engineer/service/area"
}

//---------------------------------------------------------------------------------------------------------------

type EngineerEventType struct {
	Id          int64        `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	EventTypeId int64        `xorm:"bigint notnull index 'type_id'" json:"type_id" smith:"verbose_name=事件类型id"`
	EngineerId  int64        `xorm:"bigint notnull index 'engineer_id'" json:"engineer_id" smith:"verbose_name=工程师id"`

	CreateTime  core.Time    `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime  core.Time    `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version     int          `xorm:"version notnull 'version'" json:"version"`
}

func (e EngineerEventType) TableName() string {
	return TablePrefix + "engineer_event_type"
}

func (e EngineerEventType) ModelDesc() string {
	return "verbose_name=工程师事件类型设定,entity_json=et,route=engineer/event/type"
}

//---------------------------------------------------------------------------------------------------------------

type KnowledgeCate struct {
	Id         int64            `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	Name       string           `xorm:"nvarchar(50) notnull index 'cate_name'" json:"cate_name" smith:"verbose_name=标题,min=2,max=20"`
	Desc       core.NullString  `xorm:"nvarchar(500) 'cate_desc'" json:"cate_desc" smith:"verbose_name=描述"`
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
	Id         int64        `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	CateId     int64        `xorm:"bigint notnull index 'cate_id'" json:"cate_id" smith:"verbose_name=类别"`
	Title      string       `xorm:"nvarchar(200) notnull index 'title'" json:"title" smith:"verbose_name=标题,min=2,max=20"`
	Content    string       `xorm:"text notnull 'content'" json:"content" smith:"verbose_name=内容"`

	CreateTime core.Time    `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time    `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int          `xorm:"version notnull 'version'" json:"version"`
}

func (e Knowledge) TableName() string {
	return TablePrefix + "knowledge"
}

func (e Knowledge) ModelDesc() string {
	return "verbose_name=知识库,entity_json=knowledge,route=knowledge"
}

//---------------------------------------------------------------------------------------------------------------
