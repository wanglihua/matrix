package main

import (
    "os"
    "strconv"
    "bytes"
    "strings"

    textTemplate "text/template"

    "matrix/core"
    "matrix/smith"
    "matrix/smith/template"
    "matrix/smith/models"
    "matrix/smith/fieldtype"
    "fmt"
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
        models.GroupEntity,
    }

    modelsCode := RenderCodeTemplate("models", template.ModelsTemplate, map[string]interface{}{
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

        WriteToFile(controllersDir + "\\" + entity.EntityLowerCase + ".go", controllerCode)
    }

    viewsDir := outputDir + "\\modules\\" + entityList[0].ModuleLowerCase + "\\app\\views\\" + entityList[0].ModuleLowerCase + "\\"
    err = os.MkdirAll(viewsDir, 0777)
    core.HandleError(err)

    for _, entity := range (entityList) {

        indexhtmlCode := RenderCodeTemplate("indexhtml", template.IndexHtmlTemplate, map[string]interface{}{
            "entity": entity,
        })

        WriteToFile(viewsDir + "\\" + entity.EntityLowerCase + "_index.html", indexhtmlCode)

        detailhtmlCode := RenderCodeTemplate("detailhtml", template.DetailHtmlTemplate, map[string]interface{}{
            "entity": entity,
        })

        WriteToFile(viewsDir + "\\" + entity.EntityLowerCase + "_detail.html", detailhtmlCode)
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

func WriteToFile(fileName string, code string) {

    checkFileIsExist := func(filename string) bool {
        var exist = true;
        if _, err := os.Stat(filename); os.IsNotExist(err) {
            exist = false;
        }

        return exist;
    }

    var file *os.File
    var err error
    if checkFileIsExist(fileName) {
        file, err = os.OpenFile(fileName, os.O_APPEND, 0666)  //打开文件
        core.HandleError(err)
    } else {
        file, err = os.Create(fileName)  //创建文件
        core.HandleError(err)
    }

    _, err = file.WriteString(code)
    core.HandleError(err)

    file.Sync()
    file.Close()
}

func RenderCodeTemplate(tplName string, tplContent string, args  map[string]interface{}) string {
    template, err := textTemplate.New(tplName).Funcs(textTemplate.FuncMap{
        "RenderField": TemplateFuncRenderField,
        "FirstEntity": TemplateFuncFirstEntity,
    }).Parse(tplContent)

    if err != nil {
        panic(err)
    }

    var buffer bytes.Buffer
    template.Execute(&buffer, args)
    return buffer.String()
}

func TemplateFuncRenderField(field smith.Field) string {
    //field name
    fieldName := "    "
    nameCharLength := 15 //15 是字段名称在代码编辑器中占的字符数
    fieldName += field.Name
    fieldName += strings.Repeat(" ", nameCharLength - len(field.Name))


    //field type
    fieldType := ""
    //field column
    fieldColumnType := ""
    typeLength := 18 //18 是字段类型在代码编辑器中占的字符数
    if field.FieldType == fieldtype.Int {
        fieldType = "int"
        fieldType += strings.Repeat(" ", typeLength - len(fieldType))

        fieldColumnType = "int"
    } else if field.FieldType == fieldtype.BigInt {
        fieldType = "int64"
        fieldType += strings.Repeat(" ", typeLength - len(fieldType))

        fieldColumnType = "bigint"
    } else if field.FieldType == fieldtype.Decimal {
        fieldType = "float64"
        fieldType += strings.Repeat(" ", typeLength - len(fieldType))

        fieldColumnType = "decimal(" + strconv.Itoa(field.Precision) + "," + strconv.Itoa(field.Scale) + ")"

    } else if field.FieldType == fieldtype.NVarchar {
        //GroupName  string           `xorm:"nvarchar(255) notnull unique index 'group_name'" json:"group_name"`
        fieldType = "string"
        fieldType += strings.Repeat(" ", typeLength - len(fieldType))

        fieldColumnType += "nvarchar(" + strconv.Itoa(field.Length) + ")"

    } else if field.FieldType == fieldtype.DateTime {
        fieldType = "core.Time"
        fieldType += strings.Repeat(" ", typeLength - len(fieldType))

        fieldColumnType += "datetime"

    } else if field.FieldType == fieldtype.Boolean {
        fieldType = "bool"
        fieldType += strings.Repeat(" ", typeLength - len(fieldType))

        fieldColumnType += "bool"

    } else if field.FieldType == fieldtype.Create {
        fieldType = "core.Time"
        fieldType += strings.Repeat(" ", typeLength - len(fieldType))

        fieldColumnType += "created"

    } else if field.FieldType == fieldtype.Update {
        fieldType = "core.Time"
        fieldType += strings.Repeat(" ", typeLength - len(fieldType))

        fieldColumnType += "updated"

    } else if field.FieldType == fieldtype.Version {
        fieldType = "int"
        fieldType += strings.Repeat(" ", typeLength - len(fieldType))

        fieldColumnType += "version"
    }

    fieldNull := ""
    if field.Null {
        fieldNull += "null"
    } else {
        fieldNull += "notnull"
    }

    fieldPrimaryKey := ""
    if field.Primarykey {
        fieldPrimaryKey = "pk autoincr "
    }

    fieldUnique := ""
    if field.Unique {
        fieldUnique = "unique "
    }

    fieldIndex := ""
    if field.Index {
        fieldIndex = "index "
    }

    fieldColumn := field.Column

    fieldCode := core.FormatText("field", `{{.fieldName}}{{.fieldType}}{{.tagchar}}xorm:"{{.fieldColumnType}} {{.fieldNull}} {{.fieldPrimaryKey}}{{.fieldUnique}}{{.fieldIndex}}'{{.fieldColumn}}'" json:"{{.fieldColumn}}"{{.tagchar}}`, map[string]interface{}{
        "tagchar": "`",
        "fieldName": fieldName,
        "fieldType": fieldType,
        "fieldColumnType": fieldColumnType,
        "fieldNull": fieldNull,
        "fieldPrimaryKey": fieldPrimaryKey,
        "fieldUnique": fieldUnique,
        "fieldIndex": fieldIndex,
        "fieldColumn": fieldColumn,
    })

    return fieldCode
}

func TemplateFuncFirstEntity(entityList []smith.Entity) smith.Entity {
    if len(entityList) != 0 {
        return entityList[0]
    } else {
        return smith.Entity{} //返回个空的，防止出错
    }
}