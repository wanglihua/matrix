package layout

import (
	"github.com/go-xorm/xorm"
	"matrix/app/models"
)

var sysName = ""

func GetSysName(dbsession *xorm.Session) string {
	if sysName == "" {
		config := new(models.ConfigInfo)
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
