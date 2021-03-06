package db

import (
	"github.com/go-xorm/xorm"
	//_ "github.com/mattn/go-sqlite3"
	//_ "github.com/denisenkom/go-mssqldb"
	_ "github.com/go-sql-driver/mysql"
	//_ "github.com/lib/pq"
	"github.com/revel/revel"
	//"github.com/go-xorm/core"
	"time"
)

var Engine *xorm.Engine = nil

func InitDatabase() {
	if Engine == nil {
		db_driver, _ := revel.Config.String("db.driver")
		db_url, _ := revel.Config.String("db.url")

		var err error
		Engine, err = xorm.NewEngine(db_driver, db_url)
		//Engine, err = xorm.NewEngine("sqlite3", "./matrix.db")
		//Engine, err = xorm.NewEngine("mssql", "server=localhost;database=Matrix;user id=sa;password=sa;")
		if err != nil {
			panic(err)
		}

		//Engine.TZLocation, err = time.LoadLocation("Asia/Shanghai") //设置时区
		//Engine.TZLocation, err = time.LoadLocation("Local") //设置时区
		Engine.TZLocation = time.Local //设置时区
		if err != nil {
			panic(err)
		}

		//Engine.TZLocation = time.UTC //设置时区

		if revel.DevMode == false {
			//连接池
			Engine.SetMaxIdleConns(3)   //最小
			Engine.SetMaxOpenConns(100) //最大
		}

		if revel.DevMode == true {
			Engine.ShowSQL()
			Engine.ShowExecTime()
		}

		//实体定义时明确指明了表名和列名，下面这些定义就不起作用了
		//tableMapper := core.NewPrefixMapper(core.SameMapper{}, "hd_")
		//Engine.SetTableMapper(tableMapper)
		//Engine.SetColumnMapper(core.SameMapper{})
	}
}
