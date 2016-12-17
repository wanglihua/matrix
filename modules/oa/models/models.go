package models

import (
	"matrix/core"
)

var TablePrefix = "oa_"

//---------------------------------------------------------------------------------------------------------------

type Worklog struct {
	Id int64 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	UserId    int64         `xorm:"bigint notnull index 'user_id'" json:"user_id"`
	Title     string        `xorm:"nvarchar(300) notnull 'title'" json:"title"`
	Content   string        `xorm:"nvarchar(2000) notnull 'content'" json:"content"`
	WorkDate  core.NullDate `xorm:"datetime null 'work_date'" json:"work_date"`
	BeginTime string        `xorm:"nvarchar(20) null 'begin_time'" json:"begin_time"`
	EndTime   string        `xorm:"nvarchar(20) null 'end_time'" json:"end_time"`
	Remark    string        `xorm:"nvarchar(2000) null 'remark'" json:"remark"`

	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e Worklog) TableName() string {
	return TablePrefix + "worklog"
}
