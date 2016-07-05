package main

import (
    "strings"
    "fmt"
)

func main() {
    str := "GroupName"
    str = strings.ToLower(str[0:1]) + str[1:]
    fmt.Println(str)
}
