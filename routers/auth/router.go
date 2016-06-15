package auth

import (
    "github.com/astaxie/beego"
    "matrix/controllers/auth"
)

func Namespace() *beego.Namespace {
    ns := beego.NewNamespace("/auth",
        beego.NSRouter("/user", &auth.User{}, "get:Index"),

        beego.NSRouter("/admin", &auth.Admin{}, "get:Index"),

        beego.NSRouter("/home", &auth.Home{}, "get:Index"),
    )

    return ns
}
