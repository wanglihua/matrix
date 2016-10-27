// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build windows

package winsvc

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
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

func InstallService(name, desc string) error {
	exe_path_file, err := get_exe_Path_file()
	if err != nil {
		return err
	}
	exe_path := filepath.Dir(exe_path_file)
	cmd := exec.Command(exe_path + "/tools/nssm.exe", "install", name, exe_path_file)
	return cmd.Run()
}

func RemoveService(name string) error {
	exe_path_file, err := get_exe_Path_file()
	if err != nil {
		return err
	}
	exe_path := filepath.Dir(exe_path_file)
	cmd := exec.Command(exe_path + "/tools/nssm.exe", "remove", name, "confirm")
	return cmd.Run()
}
