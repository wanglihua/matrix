package app

import (
	//"fmt"
	"github.com/revel/revel"
)

func init() {
	revel.OnAppStart(func() {
		//fmt.Println("weixin module start!")
	})
}
