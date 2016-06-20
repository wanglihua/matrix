package db

import (
    "github.com/go-xorm/xorm"
    //_ "github.com/mattn/go-sqlite3"
    _ "github.com/denisenkom/go-mssqldb"
    "github.com/revel/revel"
    "github.com/go-xorm/core"
)

var Engine *xorm.Engine = nil

func init() {
    if Engine == nil {
        var err error
        //Engine, err = xorm.NewEngine("sqlite3", "./matrix.db")
        Engine, err = xorm.NewEngine("mssql", "server=localhost;database=Matrix;user id=sa;password=sa;")
        if err != nil {
            revel.ERROR.Print(err)
        }

        Engine.ShowExecTime()
        Engine.ShowSQL()

        tableMapper := core.NewPrefixMapper(core.SameMapper{}, "hd_")
        Engine.SetTableMapper(tableMapper)
        Engine.SetColumnMapper(core.SameMapper{})

        if err != nil {
            revel.ERROR.Println(err)
        }
    }
}
