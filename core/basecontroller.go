package core

import (
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
        revel.TRACE.Println("begin ------------------------------------------------------------------")

        if userAuth(c) == false {
            return  nil
        }

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
        //先不启用
        //err := c.DbSession.Commit()
        //HandleError(err)

        c.DbSession.Close()

        revel.TRACE.Println("end --------------------------------------------------------------------")
    }

    return nil
}

func (c *BaseController) Panic() revel.Result {
    if (isStaticRequest(c) == false) {
        revel.TRACE.Println("BaseController Panic")

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

func userAuth(c *BaseController) bool {
    /*
    Static.Serve
    Static.ServeModule
    TestRunner.Index
    TestRunner.Run
     */
    if (c.Name == "Static" || c.Name == "TestRunner") {
        return true
    }

    revel.TRACE.Println("action: " + c.Action)

    if c.Action == "Login.Index" || c.Action == "Login.Login" {
        return true
    }

    if c.Action == "Home.SyncDb" || c.Action == "Home.SyncDbPost" {
        return true
    }

    if strings.HasPrefix(c.Name, "Wechat") {
        return true //暂时先 return true
    }

    loginuser := GetLoginUser(c.Session)
    revel.TRACE.Println("userAuth session id: " + c.Session.Id())

    if loginuser == nil {
        revel.TRACE.Println("loginuser == nil")

        if IsAjaxRequest(c.Request) {
            c.Result = c.RenderJson(JsonResult{Success: false, Message: "操作失败，未登陆或没相应权限！"})
        } else {
            c.Result = c.Redirect(routes.Login.Index())
        }

        return  false
    }

    return true
}
