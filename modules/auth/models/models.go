package models

import (
    "time"
)

type User struct {
    Id              int64            `xorm: "bigint notnull pk autoincr"`
    LoginName       string           `xorm: "varchar(255) notnull unique index"`
    NickName        string           `xorm: "varchar(255) notnull"`
    Password        string           `xorm: "varchar(255) notnull"`
    CreateTime      *time.Time       `xorm: "created"`
    UpdateTime      *time.Time       `xorm: "updated"`
    Version         *int             `xorm: "version"`
}

type Group struct {
    Id              int64            `xorm: "bigint notnull pk autoincr"`
    GroupName       string           `xorm: "varchar(255) notnull unique index"`
    CreateTime      time.Time        `xorm: "created"`
    UpdateTime      time.Time        `xorm: "updated"`
    Version         int              `xorm: "version"`
}

type UserGroup struct {
    Id              int64            `xorm: "bigint notnull pk autoincr"`
    UserId          int64            `xorm: "bigint notnull index"`
    GroupId         int64            `xorm: "bigint notnull index"`
    CreateTime      time.Time        `xorm: "created"`
    UpdateTime      time.Time        `xorm: "updated"`
    Version         int              `xorm: "version"`
}
