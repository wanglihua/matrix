// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package winsvc

import (
	"os/exec"
)

func StartService(name string, import_path string) error {
	exe_path_file, err := get_exe_Path_file()
	if err != nil {
		return err
	}
	cmd := exec.Command(get_nssm_path_file(import_path), "start", name, exe_path_file)
	return cmd.Run()
}

func StopService(name string, import_path string) error {
	exe_path_file, err := get_exe_Path_file()
	if err != nil {
		return err
	}
	cmd := exec.Command(get_nssm_path_file(import_path), "stop", name, exe_path_file)
	return cmd.Run()
}

func RestartService(name string, import_path string) error {
	exe_path_file, err := get_exe_Path_file()
	if err != nil {
		return err
	}
	cmd := exec.Command(get_nssm_path_file(import_path), name, exe_path_file)
	return cmd.Run()
}
