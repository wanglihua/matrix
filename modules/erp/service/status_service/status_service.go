package status_service

import (
	"matrix/modules/erp/models"
	"github.com/go-xorm/xorm"
	"matrix/core"
	"fmt"
)

func GetStatusList(db_session *xorm.Session, status_type int64) []models.StatusInfo {
	var status_list = make([]models.StatusInfo, 0)
	err := db_session.Where(fmt.Sprintf("%s = ?", models.StatusCols.Type), status_type).Find(&status_list)
	core.HandleError(err)

	return status_list
}
