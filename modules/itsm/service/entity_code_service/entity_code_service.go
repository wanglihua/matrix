package entity_code_service

import (
	"github.com/go-xorm/xorm"
	"matrix/app/models"
	"matrix/core"
)

func GetNextSerial(db_session *xorm.Session, entity string) int64 {
	var entity_code models.EntityCodeInfo
	has, err := db_session.Id(entity).Get(&entity_code)
	core.HandleError(err)

	var serial int64 = 1
	if has == false {
		_, err = db_session.Exec("INSERT INTO sys_entity_code (entity, serial) VALUES (?, 1)", entity)
		core.HandleError(err)
	} else {
		_, err = db_session.Exec("UPDATE sys_entity_code SET serial = serial + 1 WHERE entity = ?", entity)
		core.HandleError(err)
		serial = entity_code.Serial + 1
	}
	return serial
}
