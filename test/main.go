package main

import (
    "fmt"
)

func main() {
    var f []func()

    nums := []int{1, 2, 3, 4, 5}

    for _, n := range(nums) {
        f = append(f, func() {
            fmt.Println(n)
        })
    }

    for _, fn := range f {
        fn()
    }
}
