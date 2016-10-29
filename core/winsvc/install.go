// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package winsvc

import (
	"os/exec"
)

func InstallService(name, desc string, import_path string, run_mode string) error {
	exe_path_file, err := get_exe_Path_file()
	if err != nil {
		return err
	}
	cmd := exec.Command(get_nssm_path_file(import_path, run_mode), "install", name, exe_path_file)
	return cmd.Run()
}

func RemoveService(name string, import_path string, run_mode string) error {
	cmd := exec.Command(get_nssm_path_file(import_path, run_mode), "remove", name, "confirm")
	return cmd.Run()
}
