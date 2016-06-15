package routers

import (
    "matrix/controllers"
    "matrix/routers/auth"
    "github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.Home{}, "get:Index")

    beego.AddNamespace(auth.Namespace())
}
