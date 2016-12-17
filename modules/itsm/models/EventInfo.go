package models

import (
	"matrix/core"
)

type EventInfo struct {
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

func (e EventInfo) TableName() string {
	return TablePrefix + "event"
}

func (e EventInfo) ModelDesc() string {
	return "verbose_name=事件,entity_json=event,route=event"
}
