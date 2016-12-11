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
	Id         int64           `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	Name       string          `xorm:"nvarchar(50) notnull index 'type_name'" json:"type_name" smith:"verbose_name=名称,min=1,max=30"`
	Desc       core.NullString `xorm:"nvarchar(500) 'type_desc'" json:"type_desc" smith:"verbose_name=描述"`
	CreateTime core.Time       `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time       `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int             `xorm:"version notnull 'version'" json:"version"`
}

func (e EventType) TableName() string {
	return TablePrefix + "event_type"
}

func (e EventType) ModelDesc() string {
	return "verbose_name=事件类型,entity_json=type,route=event/type"
}

//---------------------------------------------------------------------------------------------------------------

type EventPriority struct {
	Id         int64           `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	Name       string          `xorm:"nvarchar(50) notnull index 'priority_name'" json:"priority_name" smith:"verbose_name=名称,min=1,max=30"`
	Desc       core.NullString `xorm:"nvarchar(500) 'priority_desc'" json:"priority_desc" smith:"verbose_name=描述"`
	CreateTime core.Time       `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time       `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int             `xorm:"version notnull 'version'" json:"version"`
}

func (e EventPriority) TableName() string {
	return TablePrefix + "event_priority"
}

func (e EventPriority) ModelDesc() string {
	return "verbose_name=事件优先级,entity_json=priority,route=event/priority"
}

//---------------------------------------------------------------------------------------------------------------

type Event struct {
	Id            int64           `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	Code          string          `xorm:"nvarchar(20) notnull index 'code'" json:"code" smith:"verbose_name=编号,min=1,max=10"`
	TypeId        int64           `xorm:"bigint notnull index 'type_id'" json:"type_id" smith:"verbose_name=类型"`
	PriorityId    int64           `xorm:"bigint notnull 'priority_id'" json:"priority_id" smith:"verbose_name=优先级"`
	RequestUserId int64           `xorm:"bigint notnull index 'request_user_id'" json:"request_user_id" smith:"verbose_name=请求用户"`
	EngineerId    core.NullInt    `xorm:"bigint index 'engineer_id'" json:"engineer_id" smith:"verbose_name=工程师"`
	AreaId        int64           `xorm:"bigint notnull index 'area_id'" json:"area_id" smith:"verbose_name=服务区域"`
	Location      string          `xorm:"nvarchar(500) notnull 'location'" json:"location" smith:"verbose_name=地点"`
	Description   core.NullString `xorm:"nvarchar(500) 'description'" json:"description" smith:"verbose_name=描述"`
	Solution      core.NullString `xorm:"nvarchar(500) 'solution'" json:"solution" smith:"verbose_name=解决方案"`
	SrcEventId    core.NullInt    `xorm:"bigint  index 'src_event_id'" json:"src_event_id" smith:"verbose_name=源事件"`
	StatusId      int64           `xorm:"bigint notnull index 'status_id'" json:"status_id" smith:"verbose_name=状态"`
	CreateTime    core.Time       `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime    core.Time       `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version       int             `xorm:"version notnull 'version'" json:"version"`
}

func (e Event) TableName() string {
	return TablePrefix + "event"
}

func (e Event) ModelDesc() string {
	return "verbose_name=事件,entity_json=event,route=event"
}

//---------------------------------------------------------------------------------------------------------------

type EventImage struct {
	Id         int64     `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	EventId    int64     `xorm:"bigint notnull index 'event_id'" json:"event_id" smith:"verbose_name=事件"`
	ImageUrl   string    `xorm:"nvarchar(300) notnull 'image_url'" json:"image_url" smith:"verbose_name=图片"`
	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e EventImage) TableName() string {
	return TablePrefix + "event_image"
}

func (e EventImage) ModelDesc() string {
	return "verbose_name=事件图片,entity_json=event_image,route=event/image"
}

//---------------------------------------------------------------------------------------------------------------

//提报中
var Event_Status_TBZ_Id int64 = 1000

//已派工
var Event_Status_YPG_Id int64 = 2000

//已接受
var Event_Status_YJS_Id int64 = 3000

//处理中
var Event_Status_CLZ_Id int64 = 4000

//已完成
var Event_Status_YWC_Id int64 = 5000

//已评价
var Event_Status_YPJ_Id int64 = 6000

//已取消
var Event_Status_YQC_Id int64 = 7000

type EventStatus struct {
	Id         int64           `xorm:"bigint notnull pk 'id'" json:"id"`
	Name       string          `xorm:"nvarchar(50) notnull index 'status_name'" json:"status_name" smith:"verbose_name=名称,min=1,max=30"`
	Desc       core.NullString `xorm:"nvarchar(500) 'status_desc'" json:"status_desc" smith:"verbose_name=描述,max=300"`
	CreateTime core.Time       `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time       `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int             `xorm:"version notnull 'version'" json:"version"`
}

func (e EventStatus) TableName() string {
	return TablePrefix + "event_status"
}

func (e EventStatus) ModelDesc() string {
	return "verbose_name=事件状态,entity_json=status,route=event/status"
}

//---------------------------------------------------------------------------------------------------------------

type EventLog struct {
	Id         int64           `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	EventId    int64           `xorm:"bigint notnull index 'event_id'" json:"event_id" smith:"verbose_name=事件"`
	Reason     string          `xorm:"nvarchar(200) notnull 'reason'" json:"reason" smith:"verbose_name=事由"`
	Remark     core.NullString `xorm:"nvarchar(300) 'remark'" json:"remark" smith:"verbose_name=备注"`
	CreateTime core.Time       `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time       `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int             `xorm:"version notnull 'version'" json:"version"`
}

func (e EventLog) TableName() string {
	return TablePrefix + "event_log"
}

func (e EventLog) ModelDesc() string {
	return "verbose_name=事件日志,entity_json=log,route=event/log"
}

//---------------------------------------------------------------------------------------------------------------

type ServiceArea struct {
	Id         int64           `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	Name       string          `xorm:"nvarchar(50) notnull index 'area_name'" json:"area_name" smith:"verbose_name=名称,min=1,max=30"`
	Desc       core.NullString `xorm:"nvarchar(500) 'area_desc'" json:"area_desc" smith:"verbose_name=描述"`
	CreateTime core.Time       `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time       `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int             `xorm:"version notnull 'version'" json:"version"`
}

func (e ServiceArea) TableName() string {
	return TablePrefix + "service_area"
}

func (e ServiceArea) ModelDesc() string {
	return "verbose_name=服务区域,entity_json=area,route=service/area"
}

//---------------------------------------------------------------------------------------------------------------

type EngineerGroup struct {
	Id         int64           `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	Name       string          `xorm:"nvarchar(50) notnull index 'group_name'" json:"group_name" smith:"verbose_name=名称,min=1,max=30"`
	Desc       core.NullString `xorm:"nvarchar(500) 'group_desc'" json:"group_desc" smith:"verbose_name=描述"`
	CreateTime core.Time       `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time       `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int             `xorm:"version notnull 'version'" json:"version"`
}

func (e EngineerGroup) TableName() string {
	return TablePrefix + "engineer_group"
}

func (e EngineerGroup) ModelDesc() string {
	return "verbose_name=工程师分组,entity_json=group,route=engineer/group"
}

//---------------------------------------------------------------------------------------------------------------

//工程师特有的其他字段信息
type Engineer struct {
	Id         int64     `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	UserId     int64     `xorm:"bigint notnull index 'user_id'" json:"user_id" smith:"verbose_name=用户id"` //对应于auth用户表中的id
	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e Engineer) TableName() string {
	return TablePrefix + "engineer"
}

func (e Engineer) ModelDesc() string {
	return "verbose_name=工程师,entity_json=engineer,route=engineer"
}

//---------------------------------------------------------------------------------------------------------------

type EngineerGroupSetting struct {
	Id         int64     `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	GroupId    int64     `xorm:"bigint notnull index 'group_id'" json:"group_id" smith:"verbose_name=分组id"`
	EngineerId int64     `xorm:"bigint notnull index 'engineer_id'" json:"engineer_id" smith:"verbose_name=工程师id"`
	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e EngineerGroupSetting) TableName() string {
	return TablePrefix + "engineer_group_setting"
}

func (e EngineerGroupSetting) ModelDesc() string {
	return "verbose_name=工程师分组设定,entity_json=setting,route=engineer/group/setting"
}

//---------------------------------------------------------------------------------------------------------------

type EngineerServiceArea struct {
	Id            int64     `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	ServiceAreaId int64     `xorm:"bigint notnull index 'service_area_id'" json:"service_area_id" smith:"verbose_name=服务区域id"`
	EngineerId    int64     `xorm:"bigint notnull index 'engineer_id'" json:"engineer_id" smith:"verbose_name=工程师id"`
	CreateTime    core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime    core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version       int       `xorm:"version notnull 'version'" json:"version"`
}

func (e EngineerServiceArea) TableName() string {
	return TablePrefix + "engineer_service_area"
}

func (e EngineerServiceArea) ModelDesc() string {
	return "verbose_name=工程师服务区域设定,entity_json=area,route=engineer/service/area"
}

//---------------------------------------------------------------------------------------------------------------

type EngineerEventType struct {
	Id          int64     `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	EventTypeId int64     `xorm:"bigint notnull index 'type_id'" json:"type_id" smith:"verbose_name=事件类型id"`
	EngineerId  int64     `xorm:"bigint notnull index 'engineer_id'" json:"engineer_id" smith:"verbose_name=工程师id"`
	CreateTime  core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime  core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version     int       `xorm:"version notnull 'version'" json:"version"`
}

func (e EngineerEventType) TableName() string {
	return TablePrefix + "engineer_event_type"
}

func (e EngineerEventType) ModelDesc() string {
	return "verbose_name=工程师事件类型设定,entity_json=et,route=engineer/event/type"
}

//---------------------------------------------------------------------------------------------------------------

//别的可以考虑的字段，可以是 事件类型id，服务区域id等等
type EngineerManager struct {
	Id         int64     `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	EngineerId int64     `xorm:"bigint notnull index 'engineer_id'" json:"engineer_id" smith:"verbose_name=工程师id"`
	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e EngineerManager) TableName() string {
	return TablePrefix + "engineer_manager"
}

func (e EngineerManager) ModelDesc() string {
	return "verbose_name=工程师经理,entity_json=manager,route=engineer/manager"
}

//---------------------------------------------------------------------------------------------------------------

type Equipment struct {
	Id         int64     `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	Code       string    `xorm:"nvarchar(50) notnull index unique 'code'" json:"code" smith:"verbose_name=编码,min=2,max=20"`
	Name       string    `xorm:"nvarchar(50) notnull index 'name'" json:"name" smith:"verbose_name=名称,min=2,max=20"`
	StatusId   int64     `xorm:"bigint notnull index 'status_id'" json:"status_id" smith:"verbose_name=状态"`
	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e Equipment) TableName() string {
	return TablePrefix + "equipment"
}

func (e Equipment) ModelDesc() string {
	return "verbose_name=设备,entity_json=equipment,route=equipment"
}

//---------------------------------------------------------------------------------------------------------------

type EquipmentStatus struct {
	Id          int64           `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	Name        string          `xorm:"nvarchar(50) notnull index 'name'" json:"name" smith:"verbose_name=名称,min=2,max=20"`
	Description core.NullString `xorm:"nvarchar(500) 'description'" json:"description" smith:"verbose_name=描述,max=300"`
	CreateTime  core.Time       `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime  core.Time       `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version     int             `xorm:"version notnull 'version'" json:"version"`
}

func (e EquipmentStatus) TableName() string {
	return TablePrefix + "equipment_status"
}

func (e EquipmentStatus) ModelDesc() string {
	return "verbose_name=设备状态,entity_json=equipment_status,route=equipment/status"
}

//---------------------------------------------------------------------------------------------------------------

type EquipmentType struct {
	Id         int64     `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	Name       string    `xorm:"nvarchar(50) notnull index 'name'" json:"name" smith:"verbose_name=名称,min=2,max=20"`
	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e EquipmentType) TableName() string {
	return TablePrefix + "equipment_type"
}

func (e EquipmentType) ModelDesc() string {
	return "verbose_name=设备类型,entity_json=equipment_type,route=equipment/type"
}

//---------------------------------------------------------------------------------------------------------------

type Application struct {
	Id         int64     `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	Code       string    `xorm:"nvarchar(50) notnull index unique 'code'" json:"code" smith:"verbose_name=编码,min=2,max=20"`
	Name       string    `xorm:"nvarchar(50) notnull index 'name'" json:"name" smith:"verbose_name=名称,min=2,max=20"`
	StatusId   int64     `xorm:"bigint notnull index 'status_id'" json:"status_id" smith:"verbose_name=状态"`
	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e Application) TableName() string {
	return TablePrefix + "application"
}

func (e Application) ModelDesc() string {
	return "verbose_name=应用,entity_json=application,route=application"
}

//---------------------------------------------------------------------------------------------------------------

type ApplicationStatus struct {
	Id          int64           `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	Name        string          `xorm:"nvarchar(50) notnull index 'name'" json:"name" smith:"verbose_name=名称,min=2,max=20"`
	Description core.NullString `xorm:"nvarchar(500) 'description'" json:"description" smith:"verbose_name=描述, max=300"`
	CreateTime  core.Time       `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime  core.Time       `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version     int             `xorm:"version notnull 'version'" json:"version"`
}

func (e ApplicationStatus) TableName() string {
	return TablePrefix + "application_status"
}

func (e ApplicationStatus) ModelDesc() string {
	return "verbose_name=应用状态,entity_json=application_status,route=application/status"
}

//---------------------------------------------------------------------------------------------------------------

type ApplicationType struct {
	Id         int64     `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	Name       string    `xorm:"nvarchar(50) notnull index 'name'" json:"name" smith:"verbose_name=名称,min=2,max=20"`
	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e ApplicationType) TableName() string {
	return TablePrefix + "application_type"
}

func (e ApplicationType) ModelDesc() string {
	return "verbose_name=应用类型,entity_json=application_type,route=application/type"
}

//---------------------------------------------------------------------------------------------------------------

type KnowledgeCate struct {
	Id         int64           `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	Name       string          `xorm:"nvarchar(50) notnull index 'cate_name'" json:"cate_name" smith:"verbose_name=名称,min=2,max=20"`
	Desc       core.NullString `xorm:"nvarchar(500) 'cate_desc'" json:"cate_desc" smith:"verbose_name=描述"`
	ParentId   int64           `xorm:"bigint notnull index 'parent_id'" json:"parent_id" smith:"verbose_name=父类"`
	CreateTime core.Time       `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time       `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int             `xorm:"version notnull 'version'" json:"version"`
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
