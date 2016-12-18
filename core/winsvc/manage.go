package winsvc

import (
	"fmt"
	"os/exec"
)

func StartService(name string, import_path string, run_mode string) error {
	exe_path_file, err := get_exe_Path_file()
	if err != nil {
		return err
	}
	cmd := exec.Command(get_nssm_path_file(import_path, run_mode), "start", name, exe_path_file)
	err = cmd.Run()
	if err == nil {
		fmt.Println("Windows服务启动成功！")
		return nil
	} else {
		return err
	}
}

func StopService(name string, import_path string, run_mode string) error {
	exe_path_file, err := get_exe_Path_file()
	if err != nil {
		return err
	}
	cmd := exec.Command(get_nssm_path_file(import_path, run_mode), "stop", name, exe_path_file)
	err = cmd.Run()
	if err == nil {
		fmt.Println("Windows服务停止成功！")
		return nil
	} else {
		return err
	}
}

func RestartService(name string, import_path string, run_mode string) error {
	exe_path_file, err := get_exe_Path_file()
	if err != nil {
		return err
	}
	cmd := exec.Command(get_nssm_path_file(import_path, run_mode), name, exe_path_file)
	err = cmd.Run()
	if err == nil {
		fmt.Println("Windows服务重启成功！")
		return nil
	} else {
		return err
	}
}
