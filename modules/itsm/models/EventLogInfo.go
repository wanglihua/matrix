package models

import (
	"matrix/core"
)

type EventLogInfo struct {
	Id         int64           `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	EventId    int64           `xorm:"bigint notnull index 'event_id'" json:"event_id" smith:"verbose_name=事件"`
	Reason     string          `xorm:"nvarchar(200) notnull 'reason'" json:"reason" smith:"verbose_name=事由"`
	Remark     core.NullString `xorm:"nvarchar(300) 'remark'" json:"remark" smith:"verbose_name=备注"`
	CreateTime core.Time       `xorm:"created notnull 'create_time'" json:"create_time"`
}

func (e EventLogInfo) TableName() string {
	return TablePrefix + "event_log"
}

func (e EventLogInfo) ModelDesc() string {
	return "verbose_name=事件日志,entity_json=log,route=event/log"
}
