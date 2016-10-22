package core

import "os"

//判断一个目录是不是存在，常常在创建一个目录前作这个判断
func IsDirExists(path string) bool {
	fi, err := os.Stat(path)

	if err != nil {
		return os.IsExist(err)
	} else {
		return fi.IsDir()
	}
}
