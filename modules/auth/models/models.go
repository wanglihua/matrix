package models

import (
	"matrix/core"
)

var TablePrefix = "auth_"

//---------------------------------------------------------------------------------------------------------------

type UserInfo struct {
	Id         int64           `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	LoginName  string          `xorm:"nvarchar(255) notnull unique index 'login_name'" json:"login_name"`
	NickName   string          `xorm:"nvarchar(255) notnull 'nick_name'" json:"nick_name"`
	Password   string          `xorm:"nvarchar(255) notnull 'password'" json:"password"`
	Phone      core.NullString `xorm:"nvarchar(50) null 'phone'" json:"phone" smith:"verbose_name=联系电话,max=30"`
	WeixinId   core.NullString `xorm:"nvarchar(200) null 'weixin_id'" json:"weixin_id" smith:"verbose_name=企业号微信id,max=150"`
	Sex        core.NullString `xorm:"nvarchar(10) null 'sex'" json:"sex" smith:"verbose_name=企业号微信id,max=4"`
	CreateTime core.Time       `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time       `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int             `xorm:"version notnull 'version'" json:"version"`
}

func (e UserInfo) TableName() string {
	return TablePrefix + "user"
}

//---------------------------------------------------------------------------------------------------------------

type AdminInfo struct {
	Id      int64     `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	UserId  int64     `xorm:"bigint notnull unique index 'user_id'" json:"user_id"`
	AddTime core.Time `xorm:"created notnull 'add_time'" json:"add_time"`
}

func (e AdminInfo) TableName() string {
	return TablePrefix + "admin"
}

//---------------------------------------------------------------------------------------------------------------

type GroupInfo struct {
	Id         int64     `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	GroupName  string    `xorm:"nvarchar(255) notnull unique index 'group_name'" json:"group_name"`
	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e GroupInfo) TableName() string {
	return TablePrefix + "group"
}

//---------------------------------------------------------------------------------------------------------------

type GroupUserInfo struct {
	Id int64 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	GroupId int64 `xorm:"bigint notnull index 'group_id'" json:"group_id"`
	UserId  int64 `xorm:"bigint notnull index 'user_id'" json:"user_id"`

	AddTime core.Time `xorm:"created notnull 'add_time'" json:"add_time"`
}

func (e GroupUserInfo) TableName() string {
	return TablePrefix + "group_user"
}
