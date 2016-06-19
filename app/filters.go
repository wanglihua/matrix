package app

import (
    "runtime/debug"
    "github.com/revel/revel"
    "fmt"
)
//捕获和处理整站异常
func PanicFilter(c *revel.Controller, fc []revel.Filter) {
    defer func() {
        if err := recover(); err != nil {
            handleInvocationPanic(c, err)
        }
    }()
    fc[0](c, fc[1:])
}

func handleInvocationPanic(c *revel.Controller, err interface{}) {
    siteError := revel.NewErrorFromPanic(err)
    //if siteError == nil && revel.DevMode {
    if siteError == nil { //在生产环境下也应该把错误栈写入日志
        revel.ERROR.Print(err, "\n", string(debug.Stack()))
        c.Response.Out.WriteHeader(500)
        c.Response.Out.Write(debug.Stack())
        return
    }

    //revel.ERROR.Print(err, "\n", error.Stack)
    c.Result = c.RenderError(siteError)


    //c.Result = c.RenderText("ddddddddddd")

    fmt.Println(c.Request.Header["X-Requested-With"]) //判定是不是 ajax 请求，XMLHttpRequest

    //c.Result = c.RenderJson(controllers.JsonResult{Success:true, Message:"Hello"})
}

var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
    // Add some common security headers
    c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
    c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
    c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

    fc[0](c, fc[1:]) // Execute the next filter stage.
}