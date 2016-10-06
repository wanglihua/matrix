package main

import (
	"fmt"
	"os"
	"github.com/revel/revel"
	"strings"
	"io"
	"compress/gzip"
	"archive/tar"
	"time"
	"path/filepath"
	"log"
	"matrix/hd/homedir"
)

var cmdPackage = &Command{
	UsageLine: "pack [web app name]",
	Short:     "package applicatoin",
	Long: `
    package application
    hd pack web_app_name(not be same with import path)
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

	home_dir, _ := homedir.Dir()
	publish_dir := home_dir + "/publish"

	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}

	cur_path_list := strings.Split(currentDir, string(os.PathSeparator))
	import_path := cur_path_list[len(cur_path_list) - 1]

	os.MkdirAll(publish_dir, 0777)
	destFile := publish_dir + "/" + web_app_name + "_" + time.Now().Format("200601021504") + ".tar.gz"

	archiveName := mustTarGzDir(destFile, currentDir, web_app_name, import_path)

	fmt.Println("Your archive is ready:", archiveName)

	return 0
}

func mustTarGzDir(destFilename, srcDir string, web_app_name string, import_path string) string {
	zipFile, err := os.Create(destFilename)
	panicOnError(err, "Failed to create archive")
	defer zipFile.Close()

	gzipWriter := gzip.NewWriter(zipFile)
	defer gzipWriter.Close()

	tarWriter := tar.NewWriter(gzipWriter)
	defer tarWriter.Close()

	revel.Walk(srcDir, func(srcPath string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if filterPathFiles(srcPath) {
			return nil
		}

		srcFile, err := os.Open(srcPath)
		panicOnError(err, "Failed to read source file")
		defer srcFile.Close()

		relative_path_file := strings.TrimLeft(srcPath[len(srcDir):], string(os.PathSeparator))

		file_header_name := web_app_name + "/" + import_path + "/" + relative_path_file

		path_file_list := strings.Split(srcPath, string(os.PathSeparator))
		path_file_list_len := len(path_file_list)

		//is the executable file
		if path_file_list_len >= 2 &&  path_file_list[path_file_list_len - 2] == path_file_list[path_file_list_len - 1] {
			// web_app_name as binary name
			file_header_name = web_app_name + "/" + web_app_name
		}

		err = tarWriter.WriteHeader(&tar.Header{
			Name:    file_header_name,
			Size:    info.Size(),
			Mode:    int64(info.Mode()),
			ModTime: info.ModTime(),
		})

		panicOnError(err, "Failed to write tar entry header")

		_, err = io.Copy(tarWriter, srcFile)
		panicOnError(err, "Failed to copy")

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

	if strings.Contains(srcPath, "/vendor/") {
		return true
	}

	if strings.Contains(srcPath, "/doc/") {
		return true
	}

	if strings.Contains(srcPath, "/smith/") {
		return true
	}

	if strings.HasSuffix(srcPath, "~") {
		return true
	}

	if strings.HasSuffix(srcPath, ".go") {
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

	if strings.HasSuffix(srcPath, "test") {
		return true
	}

	return false;
}
