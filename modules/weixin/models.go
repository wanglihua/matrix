package weixin

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
