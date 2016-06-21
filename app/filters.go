package app

import (
    "runtime/debug"
    "github.com/revel/revel"
    "matrix/service"
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
    //if siteError != nil && revel.DevMode {
    if siteError != nil {
        //在生产环境下也应该把错误栈写入日志
        revel.ERROR.Print(siteError, "\n", string(debug.Stack()))
        //revel.ERROR.Print(siteError, "\n", siteError.Stack)

        if c.Request.Header["X-Requested-With"][0] == "XMLHttpRequest" {
            c.Result = c.RenderJson(service.JsonResult{Success:false, Message:"操作失败，程序错误！"})
        } else {
            c.Response.Out.WriteHeader(500)
            c.Response.Out.Write(debug.Stack())
            c.Result = c.RenderError(siteError)
        }
    }
}

var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
    // Add some common security headers
    c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
    c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
    c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

    //add copyright
    c.Response.Out.Header().Add("copyright", "Jiaxing Huadu Information Technolony Co.,Ltd")
    //add email and mobile

    fc[0](c, fc[1:]) // Execute the next filter stage.
}