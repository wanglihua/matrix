package main

import (
    "matrix/core"
    "matrix/smith"
    "matrix/smith/fieldtype"
    "os"
    "bytes"
    "strings"
    "strconv"
    textTemplate "text/template"
)

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
        "FieldSearchable": TemplateFuncFieldSearchable,
        "ListMaxIndex": TemplateFuncListMaxIndex,
        "FieldClienValid": TemplateFuncFieldClienValid,
        "FieldValid": TemplateFuncFieldValid,
        "CheckUniqueCreate": TemplateFuncCheckUniqueCreate,
        "CheckUniqueUpdate": TemplateFuncCheckUniqueUpdate,
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
        //} else if field.FieldType == fieldtype.BigInt {
        //    fieldType = "int64"
        //    fieldType += strings.Repeat(" ", typeLength - len(fieldType))
        //    fieldColumnType = "bigint"
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

    fieldUnique := ""
    if field.Unique {
        fieldUnique = "unique "
    }

    fieldIndex := ""
    if field.Index {
        fieldIndex = "index "
    }

    fieldColumn := field.Column

    fieldCode := core.FormatText("field", `{{.fieldName}}{{.fieldType}}{{.tagchar}}xorm:"{{.fieldColumnType}} {{.fieldNull}} {{.fieldUnique}}{{.fieldIndex}}'{{.fieldColumn}}'" json:"{{.fieldColumn}}"{{.tagchar}}`, map[string]interface{}{
        "tagchar": "`",
        "fieldName": fieldName,
        "fieldType": fieldType,
        "fieldColumnType": fieldColumnType,
        "fieldNull": fieldNull,
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

func TemplateFuncFieldSearchable(field smith.Field) string {
    if field.FieldType == fieldtype.NVarchar {
        return "true"
    } else {
        return "false"
    }
}

func TemplateFuncListMaxIndex(list []smith.Field) int {
    return len(list) - 1
}

func TemplateFuncFieldClienValid(field smith.Field) string {
    //required: true, minlength: 3, maxlength: 20
    validRules := ""
    if field.Blank {
        validRules += "required: false"
    } else {
        validRules += "required: true"
    }

    //if field.FieldType == fieldtype.Int || field.FieldType == fieldtype.BigInt {
    if field.FieldType == fieldtype.Int {
        validRules += ", digits: true"

        if field.Min != "" {
            validRules += ", min: " + field.Min
        }

        if field.Max != "" {
            validRules += ", max: " + field.Max
        }
    }

    if field.FieldType == fieldtype.Decimal {
        validRules += ", number: true"

        if field.Min != "" {
            validRules += ", min: " + field.Min
        }

        if field.Max != "" {
            validRules += ", max: " + field.Max
        }
    }

    if field.FieldType == fieldtype.NVarchar {
        if field.Min != "" {
            validRules += ", minlength: " + field.Min
        }

        if field.Max != "" {
            validRules += ", maxlength: " + field.Max
        }
    }

    return validRules
}

func TemplateFuncFieldValid(entity smith.Entity, field smith.Field) string {
    validRules := ""
    if field.Blank == false {
        validRules += core.FormatText("valid_required", `    validation.Required(f.{{.entityTitleName}}.{{.fieldName}}).Message("{{.fieldVerboseName}}不能为空！")`, map[string]interface{}{
            "entityTitleName": entity.ModuleTitleName,
            "fieldName": field.Name,
            "fieldVerboseName": field.VerboseName,
        })

        validRules += "\r\n"
    }

    //if field.FieldType == fieldtype.Int || field.FieldType == fieldtype.BigInt {
    if field.FieldType == fieldtype.Int {
        if field.Min != "" {
            validRules += core.FormatText("valid_min", `    validation.Min(f.{{.entityTitleName}}.{{.fieldName}}, {{.minValue}}).Message("{{.fieldVerboseName}}不能小于{{.minValue}}！")`, map[string]interface{}{
                "entityTitleName": entity.ModuleTitleName,
                "fieldName": field.Name,
                "fieldVerboseName": field.VerboseName,
                "minValue": field.Min,
            })

            validRules += "\r\n"
        }

        if field.Max != "" {
            validRules += core.FormatText("valid_max", `    validation.Max(f.{{.entityTitleName}}.{{.fieldName}}, {{.maxValue}}).Message("{{.fieldVerboseName}}不能大于{{.maxValue}}！")`, map[string]interface{}{
                "entityTitleName": entity.ModuleTitleName,
                "fieldName": field.Name,
                "fieldVerboseName": field.VerboseName,
                "maxValue": field.Max,
            })

            validRules += "\r\n"
        }
    }

    if field.FieldType == fieldtype.Decimal {
        if field.Min != "" {
            validRules += core.FormatText("valid_min", `    validation.Min(int(f.{{.entityTitleName}}.{{.fieldName}}), {{.minValue}}).Message("{{.fieldVerboseName}}不能小于{{.minValue}}！")`, map[string]interface{}{
                "entityTitleName": entity.ModuleTitleName,
                "fieldName": field.Name,
                "fieldVerboseName": field.VerboseName,
                "minValue": field.Min,
            })

            validRules += "\r\n"
        }

        if field.Max != "" {
            validRules += core.FormatText("valid_max", `    validation.Max(int(f.{{.entityTitleName}}.{{.fieldName}}), {{.maxValue}}).Message("{{.fieldVerboseName}}不能大于{{.maxValue}}！")`, map[string]interface{}{
                "entityTitleName": entity.ModuleTitleName,
                "fieldName": field.Name,
                "fieldVerboseName": field.VerboseName,
                "maxValue": field.Max,
            })

            validRules += "\r\n"
        }
    }

    if field.FieldType == fieldtype.NVarchar {
        if field.Min != "" {
            validRules += core.FormatText("valid_min", `    validation.MinSize(f.{{.entityTitleName}}.{{.fieldName}}, {{.minLength}}).Message("{{.fieldVerboseName}}长度不能小于{{.minLength}}！")`, map[string]interface{}{
                "entityTitleName": entity.ModuleTitleName,
                "fieldName": field.Name,
                "fieldVerboseName": field.VerboseName,
                "minLength": field.Min,
            })

            validRules += "\r\n"
        }

        if field.Max != "" {
            validRules += core.FormatText("valid_max", `    validation.MaxSize(f.{{.entityTitleName}}.{{.fieldName}}, {{.maxLength}}).Message("{{.fieldVerboseName}}长度不能大于{{.maxLength}}！")`, map[string]interface{}{
                "entityTitleName": entity.ModuleTitleName,
                "fieldName": field.Name,
                "fieldVerboseName": field.VerboseName,
                "maxLength": field.Max,
            })

            validRules += "\r\n"
        }
    }

    return validRules
}

func TemplateFuncCheckUniqueCreate(entity smith.Entity, field smith.Field) string {
    template :=
`        {{.fieldCamelCase}}Count, err := session.Where("{{.field.Column}} = ?", {{.entity.EntityLowerCase}}.{{.field.Name}}).Count(new(models.{{.entity.EntityTitleName}}))
        core.HandleError(err)
        if {{.fieldCamelCase}}Count != 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，{{.field.VerboseName}}已存在！" })
        }
    `
    checkUniqueCode := core.FormatText("check_unique_create", template, map[string]interface{}{
        "entity": entity,
        "field": field,
        "fieldCamelCase": strings.ToLower(field.Name[0:1]) + field.Name[1:],
    })

    return checkUniqueCode
}

func TemplateFuncCheckUniqueUpdate(entity smith.Entity, field smith.Field) string {
    template :=
`        {{.fieldCamelCase}}Count, err := session.Where("id <> ? and {{.field.Column}} = ?", {{.entity.EntityLowerCase}}.Id, {{.entity.EntityLowerCase}}.{{.field.Name}}).Count(new(models.{{.entity.EntityTitleName}}))
        core.HandleError(err)
        if {{.fieldCamelCase}}Count != 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，{{.field.VerboseName}}已存在！" })
        }
    `
    checkUniqueCode := core.FormatText("check_unique_create", template, map[string]interface{}{
        "entity": entity,
        "field": field,
        "fieldCamelCase": strings.ToLower(field.Name[0:1]) + field.Name[1:],
    })

    return checkUniqueCode
}