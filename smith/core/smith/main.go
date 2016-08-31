package main

import (
    "os"
    "fmt"
    "matrix/core"
    "matrix/smith/core/template"
    "strings"
    "path/filepath"
    "log"
    "io"
    "gopkg.in/xmlpath"
    "bytes"
    "strconv"
    "matrix/smith/core/fieldtype"
	smith_core "matrix/smith/core"
)

var outputBaseDir = "d:\\codegen_output"

func main() {

    if (len(os.Args) < 2) {
        fmt.Println("smith subdir/file.xml")
    }

    workingDir, err := filepath.Abs(filepath.Dir(os.Args[0]))
    if err != nil {
        log.Fatal(err)
    }

    xmlFileName := os.Args[1]
    xmlFullFileName := filepath.Join(workingDir, xmlFileName)
    fmt.Println("xml file: " + xmlFullFileName)

    xmlFile, err := os.Open(xmlFullFileName)
    core.HandleError(err)
    defer xmlFile.Close()

    xmlFileContentBytes := make([]byte, 0, 10240)
    fileReadBuffer := make([]byte, 1024)
    for {
        readCount, err := xmlFile.Read(fileReadBuffer)
        if err != nil && err != io.EOF {
            panic(err)
        }

        if 0 == readCount {
            break
        }

        xmlFileContentBytes = append(xmlFileContentBytes, fileReadBuffer[:readCount]...)
    }

    xmlNode, err := xmlpath.Parse(bytes.NewBuffer(xmlFileContentBytes))
    entityPath := xmlpath.MustCompile("/xml/entity")

    entityList := make([]smith_core.Entity, 0)
    entityIter := entityPath.Iter(xmlNode)
    entityIterCount := 1
    for entityIter.Next() {
        entity := smith_core.Entity{}
        entity.ModuleTitleName, _ = xmlpath.MustCompile(fmt.Sprintf("/xml/entity[%d]/moduleTitleName", entityIterCount)).String(entityIter.Node())
        entity.ModuleLowerCase, _ = xmlpath.MustCompile(fmt.Sprintf("/xml/entity[%d]/moduleLowerCase", entityIterCount)).String(entityIter.Node())
        entity.ModuleChinese, _ = xmlpath.MustCompile(fmt.Sprintf("/xml/entity[%d]/moduleChinese", entityIterCount)).String(entityIter.Node())
        entity.EntityTitleName, _ = xmlpath.MustCompile(fmt.Sprintf("/xml/entity[%d]/entityTitleName", entityIterCount)).String(entityIter.Node())
        entity.EntityCamelCase, _ = xmlpath.MustCompile(fmt.Sprintf("/xml/entity[%d]/entityCamelCase", entityIterCount)).String(entityIter.Node())
        entity.EntityChinese, _ = xmlpath.MustCompile(fmt.Sprintf("/xml/entity[%d]/entityChinese", entityIterCount)).String(entityIter.Node())
        entity.TableName, _ = xmlpath.MustCompile(fmt.Sprintf("/xml/entity[%d]/tableName", entityIterCount)).String(entityIter.Node())
        entity.TablePrefix, _ = xmlpath.MustCompile(fmt.Sprintf("/xml/entity[%d]/tablePrefix", entityIterCount)).String(entityIter.Node())

        fieldPath := xmlpath.MustCompile(fmt.Sprintf("/xml/entity[%d]/fieldList/field", entityIterCount))
        fieldIter := fieldPath.Iter(entityIter.Node())
        entity.FieldList = make([]smith_core.Field, 0)
        fieldIterCount := 1
        for fieldIter.Next() {
            field := smith_core.Field{}

            field.VerboseName, _ = xmlpath.MustCompile(fmt.Sprintf("/xml/entity[%d]/fieldList/field[%d]/verboseName", entityIterCount, fieldIterCount)).String(fieldIter.Node())
            field.Name, _ = xmlpath.MustCompile(fmt.Sprintf("/xml/entity[%d]/fieldList/field[%d]/name", entityIterCount, fieldIterCount)).String(fieldIter.Node())
            field.Column, _ = xmlpath.MustCompile(fmt.Sprintf("/xml/entity[%d]/fieldList/field[%d]/column", entityIterCount, fieldIterCount)).String(fieldIter.Node())

            fieldFieldType, _ := xmlpath.MustCompile(fmt.Sprintf("/xml/entity[%d]/fieldList/field[%d]/fieldType", entityIterCount, fieldIterCount)).String(fieldIter.Node())
            field.FieldType = strToFieldType(fieldFieldType)

            fieldLength, _ := xmlpath.MustCompile(fmt.Sprintf("/xml/entity[%d]/fieldList/field[%d]/length", entityIterCount, fieldIterCount)).String(fieldIter.Node())
            field.Length, err = strconv.Atoi(fieldLength)
            core.HandleError(err)

            fieldPrecision, _ := xmlpath.MustCompile(fmt.Sprintf("/xml/entity[%d]/fieldList/field[%d]/precision", entityIterCount, fieldIterCount)).String(fieldIter.Node())
            field.Precision, err = strconv.Atoi(fieldPrecision)
            core.HandleError(err)

            fieldScale, _ := xmlpath.MustCompile(fmt.Sprintf("/xml/entity[%d]/fieldList/field[%d]/scale", entityIterCount, fieldIterCount)).String(fieldIter.Node())
            field.Scale, err = strconv.Atoi(fieldScale)
            core.HandleError(err)

            field.Unique, _ = xmlpath.MustCompile(fmt.Sprintf("/xml/entity[%d]/fieldList/field[%d]/unique", entityIterCount, fieldIterCount)).String(fieldIter.Node())

            fieldIndex, _ := xmlpath.MustCompile(fmt.Sprintf("/xml/entity[%d]/fieldList/field[%d]/index", entityIterCount, fieldIterCount)).String(fieldIter.Node())
            field.Index = strToBool(fieldIndex)

            fieldNull, _ := xmlpath.MustCompile(fmt.Sprintf("/xml/entity[%d]/fieldList/field[%d]/null", entityIterCount, fieldIterCount)).String(fieldIter.Node())
            field.Null = strToBool(fieldNull)

            fieldBlank, _ := xmlpath.MustCompile(fmt.Sprintf("/xml/entity[%d]/fieldList/field[%d]/blank", entityIterCount, fieldIterCount)).String(fieldIter.Node())
            field.Blank = strToBool(fieldBlank)

            field.Min, _ = xmlpath.MustCompile(fmt.Sprintf("/xml/entity[%d]/fieldList/field[%d]/min", entityIterCount, fieldIterCount)).String(fieldIter.Node())
            field.Max, _ = xmlpath.MustCompile(fmt.Sprintf("/xml/entity[%d]/fieldList/field[%d]/max", entityIterCount, fieldIterCount)).String(fieldIter.Node())

            entity.FieldList = append(entity.FieldList, field)
            fieldIterCount++
        }

        //fmt.Println(iter.Node().String())
        entityList = append(entityList, entity)
        entityIterCount++
    }

    var outputDir = outputBaseDir + "\\matrix"

    err = os.MkdirAll(outputDir, 0777) //创建目录，如果目录已经存在就不会创建，这里是保证目录存在
    core.HandleError(err)

    err = os.RemoveAll(outputDir) // 删除之前生成的目录和文件，这时会把 matrix 目录也删除
    core.HandleError(err)

    err = os.MkdirAll(outputDir, 0777) // 把删除了的 matrix 目录再创建出来
    core.HandleError(err)

    controllersDir := outputDir + "\\modules\\" + entityList[0].ModuleLowerCase + "\\app\\controllers"
    err = os.MkdirAll(controllersDir, 0777)
    core.HandleError(err)

    for _, entity := range (entityList) {
        controllerCode := smith_core.RenderCodeTemplate("controller", template.ControllerTemplate, map[string]interface{}{
            "entity": entity,
        })

        smith_core.WriteToFile(controllersDir + "\\" + entityList[0].ModuleLowerCase + "_" + smith_core.ToUnderscore(entity.EntityCamelCase) + ".go", controllerCode)
    }

    viewsDir := outputDir + "\\modules\\" + entityList[0].ModuleLowerCase + "\\app\\views\\" + entityList[0].ModuleLowerCase + "\\" + smith_core.ToUnderscore(entityList[0].EntityCamelCase) + "\\"
    err = os.MkdirAll(viewsDir, 0777)
    core.HandleError(err)

    for _, entity := range (entityList) {

        indexhtmlCode := smith_core.RenderCodeTemplate("indexhtml", template.IndexHtmlTemplate, map[string]interface{}{
            "entity": entity,
        })

        smith_core.WriteToFile(viewsDir + "\\" + smith_core.ToUnderscore(entity.EntityCamelCase) + "_index.html", indexhtmlCode)

        detailhtmlCode := smith_core.RenderCodeTemplate("detailhtml", template.DetailHtmlTemplate, map[string]interface{}{
            "entity": entity,
        })

        smith_core.WriteToFile(viewsDir + "\\" + smith_core.ToUnderscore(entity.EntityCamelCase) + "_detail.html", detailhtmlCode)
    }

    confDir := outputDir + "\\modules\\" + entityList[0].ModuleLowerCase + "\\conf"
    err = os.MkdirAll(confDir, 0777)
    core.HandleError(err)
    routeCode := smith_core.RenderCodeTemplate("route", template.RouteTemplate, map[string]interface{}{
        "entityList": entityList,
    })

    smith_core.WriteToFile(confDir + "\\route", routeCode)

    menuDir := outputDir + "\\modules\\" + entityList[0].ModuleLowerCase
    err = os.MkdirAll(menuDir, 0777)
    core.HandleError(err)
    menuCode := smith_core.RenderCodeTemplate("menu", template.MenuTemplate, map[string]interface{}{
        "entityList": entityList,
    })

    smith_core.WriteToFile(menuDir + "\\menu.html", menuCode)

    fmt.Println("Code Generated in " + outputDir)
}

func strToBool(val string) bool {
    if val == "true" {
        return true
    } else {
        return false
    }
}

func strToFieldType(val string) fieldtype.FieldType {
    if val == "Int" {
        return fieldtype.Int
    } else if (strings.ToLower(val) == strings.ToLower("BigInt")) {
        return fieldtype.BigInt
    } else if (strings.ToLower(val) == strings.ToLower("Decimal")) {
        return fieldtype.Decimal
    } else if (strings.ToLower(val) == strings.ToLower("NVarchar")) {
        return fieldtype.NVarchar
    } else if (strings.ToLower(val) == strings.ToLower("DateTime")) {
        return fieldtype.DateTime
    } else if (strings.ToLower(val) == strings.ToLower("Boolean")) {
        return fieldtype.Boolean
    } else if (strings.ToLower(val) == strings.ToLower("Create")) {
        return fieldtype.Create
    } else if (strings.ToLower(val) == strings.ToLower("Update")) {
        return fieldtype.Update
    } else if (strings.ToLower(val) == strings.ToLower("Version")) {
        return fieldtype.Version
    }

    return fieldtype.Int
}
