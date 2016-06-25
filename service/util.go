package service

import (
    "strconv"
    "github.com/revel/revel"
    "strings"
)

func HandleError(err error) {
    if err != nil {
        panic(err)
    }
}

func GetInt64FromRequest(request *revel.Request, key string) int64 {
    value, err := strconv.ParseInt(request.Form[key][0], 10, 64)
    HandleError(err)
    return value
}

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
