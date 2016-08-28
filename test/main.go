package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "ReportCateIcon"
	charArray := strings.Split(str, "")
	newCharArray := make([]string, 0)
	for index, char_str := range charArray {
		if strings.ToUpper(char_str) == char_str {
			if index != 0 {
				newCharArray = append(newCharArray, "_")
			}
			newCharArray = append(newCharArray, strings.ToLower(char_str))
		} else {
			newCharArray = append(newCharArray, char_str)
		}
	}

	fmt.Println(strings.Join(newCharArray, ""))
}
