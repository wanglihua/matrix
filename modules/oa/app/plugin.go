package app

import (
    //"fmt"
    "github.com/revel/revel"
)

func init() {
    revel.OnAppStart(func() {
        //fmt.Println("oa module start!")
    })
}
