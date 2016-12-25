package core

import (
	"github.com/go-xorm/xorm"
)

type TableDataIniter interface {
	InitData(db_session *xorm.Session)
}
