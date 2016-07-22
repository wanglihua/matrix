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
        //revel.TRACE.Println("session id: " + c.Session.Id())

        //在这之前的数据库访问请自行开启 db session
        dbSession := db.Engine.NewSession()
        c.DbSession = dbSession.NoAutoCondition(true)
        //c.DbSession = db.Engine.NoAutoCondition() 这个不行，xorm会有问题
        c.RenderArgs["dbsession"] = dbSession

        if revel.DevMode == false {
            err := c.DbSession.Begin()
            HandleError(err)
        }

        authResult := userAuth(c)
        if authResult != nil {
            return authResult
        }

        c.User = GetLoginUser(c.Session) //将LoginUser从Session Cache中取出放入Controller上下文中，方便接下来的访问
    }

    return nil
}

func (c *BaseController) Finally() revel.Result {
    if (isStaticRequest(c) == false) {
        if revel.DevMode == false {
            err := c.DbSession.Commit()
            HandleError(err)
        }

        c.DbSession.Close() //模板tags中的数据库访问请自行开启 db session

        //revel.TRACE.Println("session id: " + c.Session.Id())
        revel.TRACE.Println("end --------------------------------------------------------------------")
    }

    return nil
}

func (c *BaseController) Panic() revel.Result {
    if (isStaticRequest(c) == false) {
        revel.TRACE.Println("BaseController Panic")

        if revel.DevMode == false {
            err := c.DbSession.Rollback()
            HandleError(err)
        }
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

func userAuth(c *BaseController) revel.Result {
    /*
    Static.Serve
    Static.ServeModule
    TestRunner.Index
    TestRunner.Run
     */
    if (c.Name == "Static" || c.Name == "TestRunner") {
        return nil
    }

    revel.TRACE.Println("action: " + c.Action)

    if c.Action == "Login.Index" || c.Action == "Login.Login" {
        return nil
    }

    if c.Action == "Home.SyncDb" || c.Action == "Home.SyncDbPost" {
        return nil
    }

    if strings.HasPrefix(c.Name, "Wechat") {
        return nil //暂时先 return true
    }

    loginuser := GetLoginUser(c.Session)
    //revel.TRACE.Println("userAuth session id: " + c.Session.Id())

    if loginuser == nil {
        revel.TRACE.Println("loginuser == nil")

        if IsAjaxRequest(c.Request) {
            //c.Result = c.RenderJson(JsonResult{Success: false, Message: "操作失败，未登陆或没相应权限！"})
            return c.RenderJson(JsonResult{Success: false, Message: "操作失败，未登陆或没相应权限！"})
        } else {
            //c.Result = c.Redirect(routes.Login.Index())
            return c.Redirect(routes.Login.Index())
        }
    }

    return nil
}

func (c BaseController) UnbindToRenderArgs(val interface{}, name string) {
    datas := make(map[string]string)
    revel.Unbind(datas, name, val)
    for key, value := range datas {
        c.RenderArgs[strings.Replace(key, ".", "_", -1)] = value
    }
}
