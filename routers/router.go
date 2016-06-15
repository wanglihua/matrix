package routers

import (
    "matrix/controllers"
    "github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.Home{}, "get:Index")

    beego.AddNamespace(AuthNamespace())
}
