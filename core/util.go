package core

import (
    "strconv"
    "github.com/revel/revel"
    "strings"
    "time"
    "fmt"
)

func HandleError(err error) {
    if err != nil {
        panic(err)
    }
}

/*
func GetInt64FromRequest(request *revel.Request, key string) int64 {
    value, err := strconv.ParseInt(request.Form[key][0], 10, 64)
    HandleError(err)
    return value
}
*/

func PreventSQLInjection(sqlStr string) string {
    sqlStr = strings.Replace(sqlStr, "--", "", -1)
    sqlStr = strings.Replace(sqlStr, "#", "", -1)
    sqlStr = strings.Replace(sqlStr, ";", "", -1)
    sqlStr = strings.Replace(sqlStr, "\n", "", -1)
    sqlStr = strings.Replace(sqlStr, "\r", "", -1)
    sqlStr = strings.Replace(sqlStr, "''", "", -1)
    sqlStr = strings.Replace(sqlStr, "\"", "", -1)

    return sqlStr
}

var expireAfterDuration time.Duration = -1

func GetSessionExpire() time.Duration {
    if expireAfterDuration == -1 {
        var err error
        if expiresString, ok := revel.Config.String("session.expires"); !ok {
            expireAfterDuration = 30 * 24 * time.Hour
        } else if expiresString == "session" {
            expireAfterDuration = 0
        } else if expireAfterDuration, err = time.ParseDuration(expiresString); err != nil {
            panic(fmt.Errorf("session.expires invalid: %s", err))
        }
    }

    return expireAfterDuration
}

func IsAjaxRequest(request *revel.Request) bool {
    if len(request.Header["X-Requested-With"]) != 0 &&  request.Header["X-Requested-With"][0] == "XMLHttpRequest" {
        return true
    } else {
        return false
    }
}