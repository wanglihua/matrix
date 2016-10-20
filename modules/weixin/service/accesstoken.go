package service

import (
	"errors"
	"github.com/go-xorm/xorm"
	"io/ioutil"
	"matrix/core"
	"matrix/modules/weixin/service/json"
	"net/http"
	"strings"
	"time"
)

var accessToken = ""
var accessTokenExpireTime = time.Now().Add(-10000)

func GetAccessToken(session *xorm.Session) string {
	if accessToken == "" || time.Now().After(accessTokenExpireTime) {
		//如果没获取过，或者已经过期
		//{"access_token":"ACCESS_TOKEN","expires_in":7200}
		//{"errcode":40013,"errmsg":"invalid appid"}
		url := "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=" + GetAppId(session) + "&secret=" + GetAppSecret(session)
		httpResp, err := http.Get(url)
		core.HandleError(err)
		defer httpResp.Body.Close()

		responeBytes, err := ioutil.ReadAll(httpResp.Body)
		core.HandleError(err)
		var responeStr = string(responeBytes)

		if strings.Contains(responeStr, "errcode") {
			//var weixinError WeixinError
			//err = json.NewDecoder(httpResp.Body).Decode(&weixinError)
			//core.HandleError(err)
			core.HandleError(errors.New(responeStr))
			return ""
		}

		var accessTokenGot AccessToken
		//err = json.NewDecoder(httpResp.Body).Decode(&accessTokenGot) 会有EOF错误
		err = json.NewDecoder(strings.NewReader(responeStr)).Decode(&accessTokenGot)
		core.HandleError(err)

		accessToken = accessTokenGot.Token

		now := time.Now()
		accessTokenExpireTime = time.Date(now.Year(), now.Month(), now.Day(),
			now.Hour(), now.Minute(), now.Second()+accessTokenGot.ExpiresIn-120, //减掉120秒是为了保证在 access token 失效前重新获取
			now.Nanosecond(), time.Local)

	}

	return accessToken
}
