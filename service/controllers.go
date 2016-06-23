package service

import (
    "fmt"
    "matrix/db"
    "github.com/revel/revel"
    "github.com/go-xorm/xorm"
)

type BaseController struct {
    *revel.Controller
    DbSession *xorm.Session
}

/*
    DbSession := db.Engine.NewSession()
    defer DbSession.Close()
*/

func (c *BaseController) Before() revel.Result {
    if (isStaticRequest(c) == false) {
        fmt.Println("BaseController Before")

        c.DbSession = db.Engine.NewSession() // begin transaction
        err := c.DbSession.Begin()
        HandleError(err)
    }

    return nil
}

func (c *BaseController) After() revel.Result {
    if (isStaticRequest(c) == false) {
        fmt.Println("BaseController After")

        err := c.DbSession.Commit()
        HandleError(err)

        c.DbSession.Close()
    }

    return nil
}

func (c *BaseController) Panic() revel.Result {
    if (isStaticRequest(c) == false) {
        fmt.Println("BaseController Panic")

        err := c.DbSession.Rollback()
        HandleError(err)
    }

    return nil
}

func isStaticRequest(c *BaseController) bool {
    if (c.Name != "Static") {
        return false
    } else {
        return true
    }
}
