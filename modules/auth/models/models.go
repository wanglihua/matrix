package models

import (
    "time"
)

var tablePrefix = "hd_auth_"

//---------------------------------------------------------------------------------------------------------------

type User struct {
    Id         int64            `xorm:"bigint notnull pk autoincr 'id'" json:"Id"`
    LoginName  string           `xorm:"nvarchar(255) notnull unique index 'login_name'" json:"login_name"`
    NickName   string           `xorm:"nvarchar(255) notnull 'nick_name'" json:"nick_name"`
    Password   string           `xorm:"nvarchar(255) notnull 'password'" json:"password"`
    CreateTime *time.Time       `xorm:"created 'create_time'" json:"create_time"`
    UpdateTime *time.Time       `xorm:"updated 'update_time'" json:"update_time"`
    Version    int              `xorm:"version 'version'" json:"version"`
}

func (e User) TableName() string {
    return tablePrefix + "user"
}

//---------------------------------------------------------------------------------------------------------------

type Group struct {
    Id         int64            `xorm:"bigint notnull pk autoincr 'Id'"`
    GroupName  string           `xorm:"nvarchar(255) notnull unique index 'GroupName'"`
    CreateTime *time.Time       `xorm:"created 'create_time'" json:"create_time"`
    UpdateTime *time.Time       `xorm:"updated 'update_time'" json:"update_time"`
    Version    int              `xorm:"version 'version'" json:"version"`
}

func (e Group) TableName() string {
    return tablePrefix + "group"
}

//---------------------------------------------------------------------------------------------------------------

type UserGroup struct {
    Id         int64            `xorm:"bigint notnull pk autoincr 'Id'"`
    UserId     int64            `xorm:"bigint notnull index 'UserId'"`
    GroupId    int64            `xorm:"bigint notnull index 'GroupId'"`
    CreateTime *time.Time       `xorm:"created 'create_time'" json:"create_time"`
    UpdateTime *time.Time       `xorm:"updated 'update_time'" json:"update_time"`
    Version    int              `xorm:"version 'version'" json:"version"`
}

func (e UserGroup) TableName() string {
    return tablePrefix + "user_group"
}

//---------------------------------------------------------------------------------------------------------------