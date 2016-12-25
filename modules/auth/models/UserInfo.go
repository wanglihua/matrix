package models

import (
	"matrix/core"
	"github.com/go-xorm/xorm"
)

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

func (e UserInfo) InitData(db_session *xorm.Session) {
	user := new(UserInfo)
	user.LoginName = "admin"
	user.NickName = "管理员"
	user.Password = core.EncryptPassword("111111")
	db_session.Insert(user)

	admin := new(AdminInfo)
	admin.UserId = user.Id
	db_session.Insert(admin)
}
