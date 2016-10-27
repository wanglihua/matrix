package main

import (
	"fmt"
	"path/filepath"
	"strings"
)

func main() {
	dir, file :=  filepath.Split("c:\\xx\\test.exe")
	fmt.Println(dir, file, strings.TrimSuffix(file, filepath.Ext(file)))
}

