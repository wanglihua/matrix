package main

import (
	"fmt"
	"io/ioutil"
	smith_core "matrix/smith/core"
	"os"
	"path/filepath"
	"strings"
)

//功能没完成
func main() {
	root_dir := "/home/wanglihua/gopath/src/matrix/modules/itsm"
	filepath.Walk(root_dir, func(path string, info os.FileInfo, err error) error {
		if info == nil || info.IsDir() {
			return nil
		}
		if strings.ToUpper(filepath.Ext(path)) == strings.ToUpper(".sql") {
			if smith_core.IsFileExist(path+".go") == false {
				generate_data_access(path)
			}
		}
		return nil
	})
}

func generate_data_access(path_file string) {
	file_bytes, _ := ioutil.ReadFile(path_file)
	sql_code_lines := strings.Split(string(file_bytes), "\n")
	for _, code_line := range sql_code_lines {
		fmt.Println(code_line)
	}
}
