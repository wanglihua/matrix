package layout

import (
    "matrix/db"
    "matrix/app/models"
)

var sysName = ""

func GetSysName() string {
    if sysName == "" {
        session := db.Engine.NewSession()

        config := new(models.Config)
        has, _ := session.Limit(1).Get(config)
        if has {
            sysName = config.SysName
        }

        session.Close()
    }

    return sysName
}

func SetSysName(name string) {
    sysName = name
}