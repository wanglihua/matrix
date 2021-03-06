package app

import (
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
			c.Result = c.RenderError(fmt.Errorf("%s \r\n %s", errorDesc, string(debug.Stack())))
		} else {
			c.Result = c.RenderError(fmt.Errorf("%s \r\n %s", errorDesc, string(debug.Stack())))
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
