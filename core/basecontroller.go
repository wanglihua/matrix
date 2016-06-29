package core

import (
    "fmt"
    "matrix/db"
    "github.com/revel/revel"
    "github.com/go-xorm/xorm"
)

type BaseController struct {
    *revel.Controller

    DbSession *xorm.Session
    User      *LoginUser
}

func (c *BaseController) Before() revel.Result {
    if (isStaticRequest(c) == false) {
        fmt.Println("BaseController Before")

        c.User = GetLoginUser(c.Session) //将LoginUser从Session Cache中取出放入Controller上下文中，方便接下来的访问

        c.DbSession = db.Engine.NewSession()

        //先不启用
        //err := c.DbSession.Begin()
        //HandleError(err)
    }

    return nil
}

func (c *BaseController) After() revel.Result {
    if (isStaticRequest(c) == false) {
        fmt.Println("BaseController After")

        //先不启用
        //err := c.DbSession.Commit()
        //HandleError(err)

        c.DbSession.Close()
    }

    return nil
}

func (c *BaseController) Panic() revel.Result {
    if (isStaticRequest(c) == false) {
        fmt.Println("BaseController Panic")

        //先不启用
        //err := c.DbSession.Rollback()
        //HandleError(err)
    }

    return nil
}

func (c *BaseController) GetValidationErrorMessage() string {
    errorMessage := ""
    for _, validationError := range c.Validation.Errors {
        errorMessage += "<p>" + validationError.Message + "</p>"
    }
    return errorMessage
}

func isStaticRequest(c *BaseController) bool {
    if (c.Name != "Static") {
        return false
    } else {
        return true
    }
}
