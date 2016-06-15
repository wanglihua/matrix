package routers

import (
    "github.com/astaxie/beego"
    "matrix/controllers/auth"
)

func AuthNamespace() *beego.Namespace {
    ns := beego.NewNamespace("/auth",
        beego.NSRouter("/user", &auth.AuthUser{}, "get:Index"),

        beego.NSRouter("/admin", &auth.AuthAdmin{}, "get:Index"),

        //beego.NSRouter("/home", &auth.AuthHome{}, "get:Index"),
    )

    return ns
}
