package models

import (
	"matrix/core"
	"github.com/go-xorm/xorm"
)

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

type EventStatusInfo struct {
	Id         int64           `xorm:"bigint notnull pk 'id'" json:"id"`
	Name       string          `xorm:"nvarchar(50) notnull index 'status_name'" json:"status_name" smith:"verbose_name=名称,min=1,max=30"`
	Desc       core.NullString `xorm:"nvarchar(500) 'status_desc'" json:"status_desc" smith:"verbose_name=描述,max=300"`
	CreateTime core.Time       `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time       `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int             `xorm:"version notnull 'version'" json:"version"`
}

func (e EventStatusInfo) TableName() string {
	return TablePrefix + "event_status"
}

func (e EventStatusInfo) ModelDesc() string {
	return "verbose_name=事件状态,entity_json=status,route=event/status"
}

func (e EventStatusInfo) InitData(db_session *xorm.Session) {
	event_status := new(EventStatusInfo)
	event_status.Id = Event_Status_TBZ_Id
	event_status.Name = "提报中"
	event_status.Desc = core.NewNullString("提报中", true)
	db_session.Insert(event_status)

	event_status = new(EventStatusInfo)
	event_status.Id = Event_Status_YPG_Id
	event_status.Name = "已派工"
	event_status.Desc = core.NewNullString("已派工", true)
	db_session.Insert(event_status)

	event_status = new(EventStatusInfo)
	event_status.Id = Event_Status_YJS_Id
	event_status.Name = "已接受"
	event_status.Desc = core.NewNullString("已接受", true)
	db_session.Insert(event_status)

	event_status = new(EventStatusInfo)
	event_status.Id = Event_Status_CLZ_Id
	event_status.Name = "处理中"
	event_status.Desc = core.NewNullString("处理中", true)
	db_session.Insert(event_status)

	event_status = new(EventStatusInfo)
	event_status.Id = Event_Status_YWC_Id
	event_status.Name = "已完成"
	event_status.Desc = core.NewNullString("已完成", true)
	db_session.Insert(event_status)

	event_status = new(EventStatusInfo)
	event_status.Id = Event_Status_YPJ_Id
	event_status.Name = "已评价"
	event_status.Desc = core.NewNullString("已评价", true)
	db_session.Insert(event_status)

	event_status = new(EventStatusInfo)
	event_status.Id = Event_Status_YQC_Id
	event_status.Name = "已取消"
	event_status.Desc = core.NewNullString("已取消", true)
	db_session.Insert(event_status)
}