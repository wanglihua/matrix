package main

import (
	"fmt"
	"os"
	"github.com/revel/revel"
	"strings"
	"io"
	"compress/gzip"
	"archive/tar"
	"archive/zip"
	"time"
	"io/ioutil"
	"path/filepath"
	"log"
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

	os.MkdirAll("~/publish", 0777)
	//destFile := "d:\\publish\\" + filepath.Base(revel.BasePath) + "_" + time.Now().Format("200601021504") + ".tar.gz"
	destFile := "~/publish/" + web_app_name + "_" + time.Now().Format("200601021504") + ".zip"
	// Create the zip file.
	//archiveName := mustTarGzDir(destFile, revel.BasePath)
	currentDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	archiveName := mustZipDir(destFile, currentDir, web_app_name)

	fmt.Println("Your archive is ready:", archiveName)

	return 0
}

func mustZipDir(destFilename, srcDir string, web_app_name string) string {
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

		file_header_name := web_app_name + "/matrix/" + strings.TrimLeft(srcPath[len(srcDir):], string(os.PathSeparator))
		if strings.HasSuffix(strings.ToLower(srcPath), "matrix.exe") {
			file_header_name = web_app_name + "/" + strings.TrimLeft(srcPath[len(srcDir):], string(os.PathSeparator))
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

func mustTarGzDir(destFilename, srcDir string, web_app_name string) string {
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

		file_header_name := web_app_name + "/matrix/" + strings.TrimLeft(srcPath[len(srcDir):], string(os.PathSeparator))
		if strings.HasSuffix(strings.ToLower(srcPath), "matrix.exe") {
			file_header_name = web_app_name + "/" + strings.TrimLeft(srcPath[len(srcDir):], string(os.PathSeparator))
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

	if strings.HasSuffix(srcPath, ".exe~") {
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

	if strings.HasSuffix(srcPath, "test.exe") {
		return true
	}

	return false;
}
