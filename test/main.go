package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "hello world!"

	split_str_list := strings.Split(str, "he3llo")
	fmt.Println(len(split_str_list))

	split_str_list_len := len(split_str_list)
	for i := 0; i < split_str_list_len; i++ {
		fmt.Println(split_str_list[i])
	}
}

