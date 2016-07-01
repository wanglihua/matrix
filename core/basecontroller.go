package core

import (
    "fmt"
    "matrix/db"
    "matrix/app/routes"
    "github.com/revel/revel"
    "github.com/go-xorm/xorm"
    "strings"
)

type BaseController struct {
    *revel.Controller

    DbSession *xorm.Session
    User      *LoginUser
}

func (c *BaseController) Before() revel.Result {
    if (isStaticRequest(c) == false) {
        fmt.Println("BaseController Before")

        userAuth(c)

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


func userAuth(c * BaseController) revel.Result {
    /*
    Static.Serve
    Static.ServeModule
    TestRunner.Index
    TestRunner.Run
     */
    if (c.Name == "Static" || c.Name == "TestRunner") {
        return nil
    }

    fmt.Println(c.Action)

    if c.Action == "Login.Index" || c.Action == "Login.Login" {
        return nil
    }

    if c.Action == "Home.SyncDb" || c.Action == "Home.SyncDbPost" {
        return nil
    }

    if strings.HasPrefix(c.Name, "Wechat") {
        return nil //暂时先 return nil
    }

    loginuser := GetLoginUser(c.Session)

    if loginuser == nil {
        if IsAjaxRequest(c.Request) {
            c.Result = c.RenderJson(JsonResult{Success:false, Message:"操作失败，未登陆或没相应权限！"})
        } else {
            return c.Redirect(routes.Login.Index())
        }
    }

    return nil
}
