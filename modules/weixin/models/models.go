package models

import (
    "matrix/core"
)

var TablePrefix = "hd_weixin_"

//---------------------------------------------------------------------------------------------------------------

type Config struct {
    Id             int64            `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
    Token          string           `xorm:"nvarchar(255) notnull 'token'" json:"token"`
    EncodingAesKey string           `xorm:"nvarchar(255) 'encoding_aes_key'" json:"encoding_aes_key"`
    AppId          string           `xorm:"nvarchar(255) 'app_id'" json:"app_id"`
    AppSecret      string           `xorm:"nvarchar(255) 'app_secret'" json:"app_secret"`
    CreateTime     core.Time        `xorm:"created notnull 'create_time'" json:"create_time"`
    UpdateTime     core.Time        `xorm:"updated notnull 'update_time'" json:"update_time"`
    Version        int              `xorm:"version notnull 'version'" json:"version"`
}

func (e Config) TableName() string {
    return TablePrefix + "config"
}

//---------------------------------------------------------------------------------------------------------------

type Menu struct {
    Id         int64                `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
    Name       string               `xorm:"nvarchar(30) notnull 'menu_name'" json:"menu_name"`
    Type       string               `xorm:"nvarchar(30) null 'menu_type'" json:"menu_type"`
    Data       string               `xorm:"nvarchar(300) null 'menu_data'" json:"menu_data"` //菜单数据url, key, media等
    Level      int                  `xorm:"int notnull 'menu_level'" json:"menu_level"`
    Order      int                  `xorm:"int notnull 'menu_order'" json:"menu_order"`
    CreateTime core.Time            `xorm:"created notnull 'create_time'" json:"create_time"`
    UpdateTime core.Time            `xorm:"updated notnull 'update_time'" json:"update_time"`
    Version    int                  `xorm:"version notnull 'version'" json:"version"`
}

func (e Menu) TableName() string {
    return TablePrefix + "menu"
}

type MenuList []Menu

func (l MenuList) Len() int {
    return len(l)
}
func (l MenuList) Swap(i, j int) {
    l[i], l[j] = l[j], l[i]
}
func (l MenuList) Less(i, j int) bool {
    return l[i].Order < l[j].Order
}


//---------------------------------------------------------------------------------------------------------------

