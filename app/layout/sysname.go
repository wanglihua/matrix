package layout

import (
    "matrix/app/models"
    "github.com/go-xorm/xorm"
)

var sysName = ""

func GetSysName(dbsession *xorm.Session) string {
    if sysName == "" {
        config := new(models.Config)
        has, _ := dbsession.Limit(1).Get(config)
        if has {
            sysName = config.SysName
        }
    }

    return sysName
}

func SetSysName(name string) {
    sysName = name
}