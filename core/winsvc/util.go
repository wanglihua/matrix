package winsvc

import (
	"os"
	"path/filepath"
	"fmt"
)

func get_exe_Path_file() (string, error) {
	prog := os.Args[0]
	p, err := filepath.Abs(prog)
	if err != nil {
		return "", err
	}
	fi, err := os.Stat(p)
	if err == nil {
		if !fi.Mode().IsDir() {
			return p, nil
		}
		err = fmt.Errorf("%s is directory", p)
	}
	if filepath.Ext(p) == "" {
		p += ".exe"
		fi, err := os.Stat(p)
		if err == nil {
			if !fi.Mode().IsDir() {
				return p, nil
			}
			err = fmt.Errorf("%s is directory", p)
		}
	}
	return "", err
}

func get_nssm_path_file(import_path string, run_mode string) string {
	exe_path_file, err := get_exe_Path_file()
	if err != nil {
		panic(err)
	}
	exe_path := filepath.Dir(exe_path_file)
	if run_mode == "dev" {
		return exe_path + "/tools/nssm.exe"
	} else {
		return exe_path + "/" + import_path + "/tools/nssm.exe"
	}
}
