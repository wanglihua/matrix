package app

import (
	"errors"
	"fmt"
	"github.com/revel/revel"
	"matrix/core"
	"runtime/debug"
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
	//这个函数里面不能出错（panic），这里出错，程序就挂了
	defer func() {
		recover() //不做任何事，保证程序不退出就行
	}()

	errorDesc := fmt.Sprint(err)
	revel.ERROR.Print(errorDesc, "\n", string(debug.Stack()))

	if core.IsAjaxRequest(c.Request) {
		c.Result = c.RenderJson(core.JsonResult{Success: false, Message: "操作失败！详细:" + errorDesc})
	} else {
		if revel.DevMode == true {
			siteError := revel.NewErrorFromPanic(err) //这个方法有点问题，在Windows下会区分不了大小写盘符，所以不在生产环境中使用
			c.Result = c.RenderError(siteError)
		} else {
			//c.Response.Out.WriteHeader(500) 会报错？？ http: multiple response.WriteHeader calls
			c.Response.Out.Write(debug.Stack())
			c.Result = c.RenderError(errors.New(errorDesc))
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
