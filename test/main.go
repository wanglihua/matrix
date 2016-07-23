package main

import (
    "fmt"
    "strings"
)

func main() {
    str := "ReportCateIcon"
    charArray := strings.Split(str, "")
    newCharArray := make([]string, 0)
    for index, char := range charArray {
        if strings.ToUpper(char) == char {
            if index != 0 {
                newCharArray = append(newCharArray, "_")
            }
            newCharArray = append(newCharArray, strings.ToLower(char))
        } else {
            newCharArray = append(newCharArray, char)
        }
    }
    fmt.Println(strings.Join(newCharArray, ""))
}
