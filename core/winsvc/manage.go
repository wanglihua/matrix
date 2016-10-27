// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package winsvc

import (
	"path/filepath"
	"os/exec"
)

func StartService(name string) error {
	exe_path_file, err := get_exe_Path_file()
	if err != nil {
		return err
	}
	exe_path := filepath.Dir(exe_path_file)
	cmd := exec.Command(exe_path + "/tools/nssm.exe", "start", name, exe_path_file)
	return cmd.Run()
}

func StopService(name string) error {
	exe_path_file, err := get_exe_Path_file()
	if err != nil {
		return err
	}
	exe_path := filepath.Dir(exe_path_file)
	cmd := exec.Command(exe_path + "/tools/nssm.exe", "stop", name, exe_path_file)
	return cmd.Run()
}

func RestartService(name string) error {
	exe_path_file, err := get_exe_Path_file()
	if err != nil {
		return err
	}
	exe_path := filepath.Dir(exe_path_file)
	cmd := exec.Command(exe_path + "/tools/nssm.exe", "restart", name, exe_path_file)
	return cmd.Run()
}
