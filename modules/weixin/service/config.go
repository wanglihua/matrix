package service

import (
	"github.com/go-xorm/xorm"
	"matrix/core"
	"matrix/modules/weixin/models"
)

var token = ""
var encodingAesKey = ""
var appId = ""
var appSecret = ""

//从数据库配置中获取一次后，就不用获取第二次了

func GetToken(session *xorm.Session) string {
	if token == "" {
		config := new(models.ConfigInfo)
		_, err := session.Limit(1).Get(config)
		core.HandleError(err)

		token = config.Token
	}

	return token
}

func GetEncodingAesKey(session *xorm.Session) string {
	if encodingAesKey == "" {
		config := new(models.ConfigInfo)
		_, err := session.Limit(1).Get(config)
		core.HandleError(err)

		encodingAesKey = config.EncodingAesKey
	}

	return encodingAesKey
}

func GetAppId(session *xorm.Session) string {
	if appId == "" {
		config := new(models.ConfigInfo)
		_, err := session.Limit(1).Get(config)
		core.HandleError(err)

		appId = config.AppId
	}

	return appId
}

func GetAppSecret(session *xorm.Session) string {
	if appSecret == "" {
		config := new(models.ConfigInfo)
		_, err := session.Limit(1).Get(config)
		core.HandleError(err)

		appSecret = config.AppSecret
	}

	return appSecret
}
