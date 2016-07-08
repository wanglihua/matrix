package main

import (
    "os"
    "fmt"
    "matrix/core"
    "matrix/smith"
    "matrix/smith/template"
    "matrix/smith/models/weixin"
    "strings"
)

var outputBaseDir = "d:\\codegen_output"

func main() {

    var outputDir = outputBaseDir + "\\matrix"

    err := os.MkdirAll(outputDir, 0777) //创建目录，如果目录已经存在就不会创建，这里是保证目录存在
    core.HandleError(err)

    err = os.RemoveAll(outputDir) // 删除之前生成的目录和文件，这时会把 matrix 目录也删除
    core.HandleError(err)

    err = os.MkdirAll(outputDir, 0777) // 把删除了的 matrix 目录再创建出来
    core.HandleError(err)

    entityList := []smith.Entity{
        //inventory.SupplierEntity,
        //inventory.StockEntity,
        //inventory.StorageLocEntity,
        weixin.MenuEntity,
    }

    modelsCode := RenderCodeTemplate("models", template.ModelsTemplate, map[string]interface{}{
        "tagchar": "`",
        "tablePrefix": entityList[0].TablePrefix,
        "entityList": entityList,
    })

    modelsDir := outputDir + "\\modules\\" + entityList[0].ModuleLowerCase + "\\models"
    err = os.MkdirAll(modelsDir, 0777)
    core.HandleError(err)
    WriteToFile(modelsDir + "\\models.go", modelsCode)

    controllersDir := outputDir + "\\modules\\" + entityList[0].ModuleLowerCase + "\\app\\controllers"
    err = os.MkdirAll(controllersDir, 0777)
    core.HandleError(err)

    for _, entity := range (entityList) {
        controllerCode := RenderCodeTemplate("controller", template.ControllerTemplate, map[string]interface{}{
            "entity": entity,
        })

        WriteToFile(controllersDir + "\\" + strings.ToLower(entity.EntityCamelCase) + ".go", controllerCode)
    }

    viewsDir := outputDir + "\\modules\\" + entityList[0].ModuleLowerCase + "\\app\\views\\" + entityList[0].ModuleLowerCase + "\\" + strings.ToLower(entityList[0].EntityCamelCase) + "\\"
    err = os.MkdirAll(viewsDir, 0777)
    core.HandleError(err)

    for _, entity := range (entityList) {

        indexhtmlCode := RenderCodeTemplate("indexhtml", template.IndexHtmlTemplate, map[string]interface{}{
            "entity": entity,
        })

        WriteToFile(viewsDir + "\\" + strings.ToLower(entity.EntityCamelCase) + "_index.html", indexhtmlCode)

        detailhtmlCode := RenderCodeTemplate("detailhtml", template.DetailHtmlTemplate, map[string]interface{}{
            "entity": entity,
        })

        WriteToFile(viewsDir + "\\" + strings.ToLower(entity.EntityCamelCase) + "_detail.html", detailhtmlCode)
    }

    confDir := outputDir + "\\modules\\" + entityList[0].ModuleLowerCase + "\\conf"
    err = os.MkdirAll(confDir, 0777)
    core.HandleError(err)
    routeCode := RenderCodeTemplate("route", template.RouteTemplate, map[string]interface{}{
        "entityList": entityList,
    })

    WriteToFile(confDir + "\\route", routeCode)

    menuDir := outputDir + "\\modules\\" + entityList[0].ModuleLowerCase
    err = os.MkdirAll(menuDir, 0777)
    core.HandleError(err)
    menuCode := RenderCodeTemplate("menu", template.MenuTemplate, map[string]interface{}{
        "entityList": entityList,
    })

    WriteToFile(menuDir + "\\menu.html", menuCode)

    fmt.Println("Code Generated in " + outputDir)
}
