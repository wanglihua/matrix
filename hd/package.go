package main

import (
    "fmt"
    "os"
    "path/filepath"

    "github.com/revel/revel"
    "strings"
    "io"
    "compress/gzip"
    "archive/tar"
    "archive/zip"
    "time"
    "io/ioutil"
)

var cmdPackage = &Command{
    UsageLine: "pack [import path] [run mode]",
    Short:     "package applicatoin",
    Long: `
    package application
    hd pack matrix
    hd pack matrix prod
    hd pack matrix dev
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

    // Determine the run mode.
    mode := "prod"
    if len(args) >= 2 {
        mode = args[1]
    }

    appImportPath := args[0]
    revel.Init(mode, appImportPath, "")

    os.MkdirAll("d:\\publish", 0777)
    //destFile := "d:\\publish\\" + filepath.Base(revel.BasePath) + "_" + time.Now().Format("200601021504") + ".tar.gz"
    destFile := "d:\\publish\\" + filepath.Base(revel.BasePath) + "_" + time.Now().Format("200601021504") + ".zip"
    // Create the zip file.
    //archiveName := mustTarGzDir(destFile, revel.BasePath)
    archiveName := mustZipDir(destFile, revel.BasePath)

    fmt.Println("Your archive is ready:", archiveName)

    return 0
}

func mustZipDir(destFilename, srcDir string) string {
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

        fileWriter, err := zipWriter.Create(revel.ImportPath + "\\" + strings.TrimLeft(srcPath[len(srcDir):], string(os.PathSeparator)))
        panicOnError(err, "Failed to create file writer")

        fileContent, err := ioutil.ReadFile(srcPath)
        panicOnError(err, "Failed to read source file")

        _, err = fileWriter.Write(fileContent)
        panicOnError(err, "Failed to write file content")

        return nil
    })

    return zipFile.Name()
}

func mustTarGzDir(destFilename, srcDir string) string {
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

        err = tarWriter.WriteHeader(&tar.Header{
            Name:    revel.ImportPath + "\\" + strings.TrimLeft(srcPath[len(srcDir):], string(os.PathSeparator)),
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

    return false;
}