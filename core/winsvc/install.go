package winsvc

import (
	"fmt"
	"os/exec"
)

func InstallService(name, desc string, import_path string, run_mode string) error {
	exe_path_file, err := get_exe_Path_file()
	if err != nil {
		return err
	}
	cmd := exec.Command(get_nssm_path_file(import_path, run_mode), "install", name, exe_path_file)
	err = cmd.Run()
	if err == nil {
		fmt.Println("Windows服务安装成功！")
		return nil
	} else {
		return err
	}
}

func RemoveService(name string, import_path string, run_mode string) error {
	cmd := exec.Command(get_nssm_path_file(import_path, run_mode), "remove", name, "confirm")
	err := cmd.Run()
	if err == nil {
		fmt.Println("Windows服务移除成功！")
		return nil
	} else {
		return err
	}
}
