package app

import (
	"matrix/core"
    "matrix/db"
    "matrix/app/routes"
    "github.com/revel/revel"
    "strings"
)

func BeforeInterceptor(c *revel.Controller) revel.Result {
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
            core.HandleError(err)
        }

        authResult := userAuth(c)
        if authResult != nil {
            return authResult
        }
    }

    return nil
}

func FinallyInterceptor(c *revel.Controller) revel.Result {
    if (isStaticRequest(c) == false) {
        if revel.DevMode == false {
            err := c.DbSession.Commit()
            core.HandleError(err)
        }

        c.DbSession.Close() //模板tags中的数据库访问请自行开启 db session

        //revel.TRACE.Println("session id: " + c.Session.Id())
        revel.TRACE.Println("end --------------------------------------------------------------------")
    }

    return nil
}

func PanicInterceptor(c *revel.Controller) revel.Result {
    if (isStaticRequest(c) == false) {
        if revel.DevMode == false {
            err := c.DbSession.Rollback()
            core.HandleError(err)
        }
    }

    return nil
}

func isStaticRequest(c *revel.Controller) bool {
    if (c.Name != "Static") {
        return false
    } else {
        return true
    }
}

func userAuth(c *revel.Controller) revel.Result {
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

    loginuser := core.GetLoginUser(c.Session)
    //revel.TRACE.Println("userAuth session id: " + c.Session.Id())

    if loginuser == nil {
        revel.TRACE.Println("loginuser == nil")

        if core.IsAjaxRequest(c.Request) {
            //c.Result = c.RenderJson(JsonResult{Success: false, Message: "操作失败，未登陆或没相应权限！"})
            return c.RenderJson(core.JsonResult{Success: false, Message: "操作失败，未登陆或没相应权限！"})
        } else {
            //c.Result = c.Redirect(routes.Login.Index())
            return c.Redirect(routes.Login.Index())
        }
    }

    return nil
}

