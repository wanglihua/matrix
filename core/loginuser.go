package core

import (
    "github.com/revel/revel/cache"
    "github.com/revel/revel"
)

const LOGIN_USER_CACHE_KEY = "Login_User_b6db637be022550d"

//根据以后需要再在这里添加成员
type LoginUser struct {
    UserId    int64
    LoginName string
    NickName  string
}

func GetLoginUser(session revel.Session) *LoginUser {
    var loginUser = new(LoginUser)
    err := cache.Get(session.Id() + LOGIN_USER_CACHE_KEY, loginUser)
    if err != nil {
        revel.WARN.Println(err.Error())
    }

    if err != nil || loginUser.UserId == 0 {
        return nil
    } else {
        return loginUser
    }
}

func PutLoginUserToSession(session revel.Session, user LoginUser) {
    cache.Set(session.Id() + LOGIN_USER_CACHE_KEY, user, GetSessionExpire())
}

func RemoveLoginUserInSession(session revel.Session) {
    cache.Delete(session.Id() + LOGIN_USER_CACHE_KEY)
}
