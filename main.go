package main

import (
    _ "matrix/routers"
    "github.com/astaxie/beego"
)

func main() {
    test()

    beego.SetStaticPath("/static", "static")
    beego.Run()
}

func test() {
    //just simple test code
}
