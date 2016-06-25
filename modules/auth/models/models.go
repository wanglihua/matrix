package models

import (
    "matrix/core"
)

var tablePrefix = "hd_auth_"

//---------------------------------------------------------------------------------------------------------------

type User struct {
    Id         int64            `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
    LoginName  string           `xorm:"nvarchar(255) notnull unique index 'login_name'" json:"login_name"`
    NickName   string           `xorm:"nvarchar(255) notnull 'nick_name'" json:"nick_name"`
    Password   string           `xorm:"nvarchar(255) notnull 'password'" json:"password"`
    AddTime    core.Time        `xorm:"datetime notnull 'add_time'" json:"add_time"`
    CreateTime core.Time        `xorm:"created notnull 'create_time'" json:"create_time"`
    UpdateTime core.Time        `xorm:"updated notnull 'update_time'" json:"update_time"`
    Version    int              `xorm:"version notnull 'version'" json:"version"`
}

func (e User) TableName() string {
    return tablePrefix + "user"
}

//---------------------------------------------------------------------------------------------------------------

type Group struct {
    Id         int64            `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
    GroupName  string           `xorm:"nvarchar(255) notnull unique index 'group_name'" json:"group_name"`
    CreateTime core.Time        `xorm:"created notnull 'create_time'" json:"create_time"`
    UpdateTime core.Time        `xorm:"updated notnull 'update_time'" json:"update_time"`
    Version    int              `xorm:"version notnull 'version'" json:"version"`
}

func (e Group) TableName() string {
    return tablePrefix + "group"
}

//---------------------------------------------------------------------------------------------------------------

type UserGroup struct {
    Id         int64            `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
    UserId     int64            `xorm:"bigint notnull index 'user_id'" json:"user_id"`
    GroupId    int64            `xorm:"bigint notnull index 'group_id'" json:"group_id"`
    CreateTime core.Time        `xorm:"created notnull 'create_time'" json:"create_time"`
    UpdateTime core.Time        `xorm:"updated notnull 'update_time'" json:"update_time"`
    Version    int              `xorm:"version notnull 'version'" json:"version"`
}

func (e UserGroup) TableName() string {
    return tablePrefix + "user_group"
}

//---------------------------------------------------------------------------------------------------------------