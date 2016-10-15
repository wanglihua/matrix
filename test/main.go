package main

import (
	"log"
	"matrix/core/lic"
	"os"
	"path/filepath"
)

func main() {
	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	lic_file_path := currentDir + string(os.PathSeparator) + "app.lic"

	lic.ValidAppLic("salestarget", lic_file_path)
}
