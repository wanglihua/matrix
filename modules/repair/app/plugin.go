package app

import (
	"github.com/revel/revel"
)

func init() {
	revel.OnAppStart(func() {
		//revel.TRACE.Println("repair moudle start!")
	})
}
