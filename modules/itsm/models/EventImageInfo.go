package models

import (
	"matrix/core"
)

type EventImageInfo struct {
	Id         int64     `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	EventId    int64     `xorm:"bigint notnull index 'event_id'" json:"event_id" smith:"verbose_name=事件"`
	ImageUrl   string    `xorm:"nvarchar(300) notnull 'image_url'" json:"image_url" smith:"verbose_name=图片"`
	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e EventImageInfo) TableName() string {
	return TablePrefix + "event_image"
}

func (e EventImageInfo) ModelDesc() string {
	return "verbose_name=事件图片,entity_json=event_image,route=event/image"
}
