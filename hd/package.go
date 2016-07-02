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
    "time"
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

    destFile := "d:\\publish\\" + filepath.Base(revel.BasePath) + "_" + time.Now().Format("200601021504") + ".tar.gz"
    // Create the zip file.
    archiveName := mustTarGzDir(destFile, revel.BasePath)

    fmt.Println("Your archive is ready:", archiveName)

    return 0
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

        if strings.Contains(srcPath, ".svn") {
            return nil
        }

        if strings.Contains(srcPath, ".git") {
            return nil
        }

        if strings.Contains(srcPath, ".idea") {
            return nil
        }

        if strings.Contains(srcPath, ".vscode") {
            return nil
        }

        if strings.Contains(srcPath, "\\vendor\\") {
            return nil
        }

        if strings.Contains(srcPath, "\\doc\\") {
            return nil
        }

        if strings.HasSuffix(srcPath, ".exe~") {
            return nil
        }

        if strings.HasSuffix(srcPath, ".go") {
            return nil
        }

        if strings.HasSuffix(srcPath, ".less") {
            return nil
        }

        if strings.HasSuffix(srcPath, ".py") {
            return nil
        }

        if strings.HasSuffix(srcPath, ".log") {
            return nil
        }

        if strings.HasSuffix(srcPath, "debug") {
            return nil
        }

        srcFile, err := os.Open(srcPath)
        panicOnError(err, "Failed to read source file")
        defer srcFile.Close()

        err = tarWriter.WriteHeader(&tar.Header{
            Name:    strings.TrimLeft(srcPath[len(srcDir):], string(os.PathSeparator)),
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