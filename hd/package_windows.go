package main

import (
	"archive/zip"
	"fmt"
	"github.com/revel/revel"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var cmdPackage = &Command{
	UsageLine: "pack [web app name]",
	Short:     "package applicatoin",
	Long: `
    package application
    hd pack web_app_name
`,
}

func init() {
	cmdPackage.Run = packageApp
}

func packageApp(cmd *Command, args []string) int {
	if len(args) == 0 {
		fmt.Fprint(os.Stderr, cmdPackage.Long)
		return 1
	}

	web_app_name := args[0]

	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	cur_path_list := strings.Split(currentDir, string(os.PathSeparator))
	import_path := cur_path_list[len(cur_path_list)-1]

	os.MkdirAll("d:\\publish", 0777)
	destFile := "d:\\publish\\" + web_app_name + "_" + time.Now().Format("200601021504") + ".zip"

	archiveName := mustZipDir(destFile, currentDir, web_app_name, import_path)

	fmt.Println("Your archive is ready:", archiveName)

	return 0
}

func mustZipDir(destFilename, srcDir string, web_app_name string, import_path string) string {
	zipFile, err := os.Create(destFilename)
	panicOnError(err, "Failed to create archive")
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	revel.Walk(srcDir, func(srcPath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if filterPathFiles(srcPath) {
			return nil
		}

		relative_path_file := strings.TrimLeft(srcPath[len(srcDir):], string(os.PathSeparator))

		file_header_name := web_app_name + "\\" + import_path + "\\" + relative_path_file

		path_file_list := strings.Split(srcPath, string(os.PathSeparator))
		path_file_list_len := len(path_file_list)

		//is the executable file
		if path_file_list_len >= 2 && path_file_list[path_file_list_len-2]+".exe" == path_file_list[path_file_list_len-1] {
			// web_app_name as binary name
			file_header_name = web_app_name + "\\" + web_app_name + ".exe"
		}

		fileWriter, err := zipWriter.Create(file_header_name)
		panicOnError(err, "Failed to create file writer")

		fileContent, err := ioutil.ReadFile(srcPath)
		panicOnError(err, "Failed to read source file")

		_, err = fileWriter.Write(fileContent)
		panicOnError(err, "Failed to write file content")

		return nil
	})

	return zipFile.Name()
}

func filterPathFiles(srcPath string) bool {

	if strings.Contains(srcPath, ".svn") {
		return true
	}

	if strings.Contains(srcPath, ".git") {
		return true
	}

	if strings.Contains(srcPath, ".idea") {
		return true
	}

	if strings.Contains(srcPath, ".vscode") {
		return true
	}

	if strings.Contains(srcPath, "\\vendor\\") {
		return true
	}

	if strings.Contains(srcPath, "\\doc\\") {
		return true
	}

	if strings.Contains(srcPath, "\\smith\\") {
		return true
	}

	if strings.HasSuffix(srcPath, ".exe~") {
		return true
	}

	if strings.HasSuffix(srcPath, ".go") {
		return true
	}

	if strings.HasSuffix(srcPath, ".syso") {
		return true
	}

	if strings.HasSuffix(srcPath, ".less") {
		return true
	}

	if strings.HasSuffix(srcPath, ".py") {
		return true
	}

	if strings.HasSuffix(srcPath, ".log") {
		return true
	}

	if strings.HasSuffix(srcPath, "debug") {
		return true
	}

	if strings.HasSuffix(srcPath, "test.exe") {
		return true
	}

	return false
}
