package main

import (
    "os"
    "matrix/core"
)

var outputBaseDir = "d:\\codegen_output"

func main() {

    var outputDir = outputBaseDir + "\\matrix"

    err := os.MkdirAll(outputDir, 0777) //创建目录，如果目录已经存在就不会创建，这里是保证目录存在
    core.HandleError(err)

    err = os.RemoveAll(outputDir) // 删除之前生成的目录和文件，这时会把 matrix 目录也删除
    core.HandleError(err)
}
