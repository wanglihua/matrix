package main

import (
	"fmt"
)

func main() {
	var h *Hello

	var hh Hello1

	hh = h
	hh.Hello1()

	if hh == nil {
		fmt.Println("nil")
	} else {
		fmt.Println("not nil")
	}
}

type Hello1 interface {
	Hello1()
}

type Hello struct {
}

func (h *Hello) Hello1() {
	fmt.Println("Hello 1")
}
