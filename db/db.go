package db

import (
    "github.com/go-xorm/xorm"
    _ "github.com/mattn/go-sqlite3"
    "github.com/revel/revel"
)

var Engine *xorm.Engine = nil

func init() {
    if Engine == nil {
        var err error
        Engine, err = xorm.NewEngine("sqlite3", "./matrix.db")
        revel.ERROR.Println(err)
    }
}
