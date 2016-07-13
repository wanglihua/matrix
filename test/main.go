package main

import (
    "fmt"
)

func main() {
    fmt.Println(fmt.Sprintf("/xml/entity[%d]/fieldlist/field[%d]", 1, 2))
}
