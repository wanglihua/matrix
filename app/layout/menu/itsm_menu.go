package menu

import (
	"bytes"
	"github.com/go-xorm/xorm"
	"github.com/revel/revel"
)

func GetItsmMenu(db_session *xorm.Session, web_session revel.Session) string {
	if _, found := revel.ModuleByName("itsm"); found == false {
		return ""
	}

	var buffer bytes.Buffer

	buffer.WriteString(`
	
	`)

	return buffer.String()
}
