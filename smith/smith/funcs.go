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
        "LowerCase": TemplateFuncLowerCase,
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

    fieldUnique := ""
    if field.Unique != "false" {
        if field.Unique == "true" {
            fieldUnique = "unique "
        } else if strings.Index(field.Unique, ",") == -1 {
            fieldUnique = "unique(" + field.Unique + ") "
        } else {
            uniqueList := strings.Split(field.Unique, ",")
            for _, unique := range (uniqueList) {
                fieldUnique += "unique(" + strings.Trim(unique, " ") + ") "
            }
        }
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

    if field.FieldType == fieldtype.Int || field.FieldType == fieldtype.BigInt {
        validRules += ", digits: true"

        if field.Min != "" && field.Min != "0"   {
            validRules += ", min: " + field.Min
        }

        if field.Max != "" && field.Max != "0"   {
            validRules += ", max: " + field.Max
        }
    }

    if field.FieldType == fieldtype.Decimal {
        validRules += ", number: true"

        if field.Min != "" && field.Min != "0"   {
            validRules += ", min: " + field.Min
        }

        if field.Max != "" && field.Max != "0"   {
            validRules += ", max: " + field.Max
        }
    }

    if field.FieldType == fieldtype.NVarchar {
        if field.Min != "" && field.Min != "0"   {
            validRules += ", minlength: " + field.Min
        }

        if field.Max != "" && field.Max != "0"   {
            validRules += ", maxlength: " + field.Max
        }
    }

    return validRules
}

func TemplateFuncFieldValid(entity smith.Entity, field smith.Field) string {
    validRules := ""
    if field.Blank == false {
        validRules += core.FormatText("valid_required", `    validation.Required(f.{{.entityTitleName}}.{{.fieldName}}).Message("{{.fieldVerboseName}}不能为空！")`, map[string]interface{}{
            "entityTitleName": entity.EntityTitleName,
            "fieldName": field.Name,
            "fieldVerboseName": field.VerboseName,
        })

        validRules += "\r\n"
    }

    //if field.FieldType == fieldtype.Int || field.FieldType == fieldtype.BigInt {
    if field.FieldType == fieldtype.Int {
        if field.Min != "" && field.Min != "0"   {
            validRules += core.FormatText("valid_min", `    validation.Min(f.{{.entityTitleName}}.{{.fieldName}}, {{.minValue}}).Message("{{.fieldVerboseName}}不能小于{{.minValue}}！")`, map[string]interface{}{
                "entityTitleName": entity.EntityTitleName,
                "fieldName": field.Name,
                "fieldVerboseName": field.VerboseName,
                "minValue": field.Min,
            })

            validRules += "\r\n"
        }

        if field.Max != "" && field.Max != "0"   {
            validRules += core.FormatText("valid_max", `    validation.Max(f.{{.entityTitleName}}.{{.fieldName}}, {{.maxValue}}).Message("{{.fieldVerboseName}}不能大于{{.maxValue}}！")`, map[string]interface{}{
                "entityTitleName": entity.EntityTitleName,
                "fieldName": field.Name,
                "fieldVerboseName": field.VerboseName,
                "maxValue": field.Max,
            })

            validRules += "\r\n"
        }
    }

    if field.FieldType == fieldtype.BigInt {
        if field.Min != "" && field.Min != "0"   {
            validRules += core.FormatText("valid_min", `    validation.Min(int(f.{{.entityTitleName}}.{{.fieldName}}), {{.minValue}}).Message("{{.fieldVerboseName}}不能小于{{.minValue}}！")`, map[string]interface{}{
                "entityTitleName": entity.EntityTitleName,
                "fieldName": field.Name,
                "fieldVerboseName": field.VerboseName,
                "minValue": field.Min,
            })

            validRules += "\r\n"
        }

        if field.Max != "" && field.Max != "0"   {
            validRules += core.FormatText("valid_max", `    validation.Max(int(f.{{.entityTitleName}}.{{.fieldName}}), {{.maxValue}}).Message("{{.fieldVerboseName}}不能大于{{.maxValue}}！")`, map[string]interface{}{
                "entityTitleName": entity.EntityTitleName,
                "fieldName": field.Name,
                "fieldVerboseName": field.VerboseName,
                "maxValue": field.Max,
            })

            validRules += "\r\n"
        }
    }

    if field.FieldType == fieldtype.Decimal {
        if field.Min != "" && field.Min != "0"   {
            validRules += core.FormatText("valid_min", `    validation.Min(int(f.{{.entityTitleName}}.{{.fieldName}}), {{.minValue}}).Message("{{.fieldVerboseName}}不能小于{{.minValue}}！")`, map[string]interface{}{
                "entityTitleName": entity.EntityTitleName,
                "fieldName": field.Name,
                "fieldVerboseName": field.VerboseName,
                "minValue": field.Min,
            })

            validRules += "\r\n"
        }

        if field.Max != "" && field.Max != "0"   {
            validRules += core.FormatText("valid_max", `    validation.Max(int(f.{{.entityTitleName}}.{{.fieldName}}), {{.maxValue}}).Message("{{.fieldVerboseName}}不能大于{{.maxValue}}！")`, map[string]interface{}{
                "entityTitleName": entity.EntityTitleName,
                "fieldName": field.Name,
                "fieldVerboseName": field.VerboseName,
                "maxValue": field.Max,
            })

            validRules += "\r\n"
        }
    }

    if field.FieldType == fieldtype.NVarchar {
        minSizeTemplate :=
        `    if f.{{.entityTitleName}}.{{.fieldName}} != "" {
        validation.MinSize(f.{{.entityTitleName}}.{{.fieldName}}, {{.minLength}}).Message("{{.fieldVerboseName}}长度不能小于{{.minLength}}！")
    }`

        if field.Min != "" && field.Min != "0"   {
            validRules += core.FormatText("valid_min", minSizeTemplate, map[string]interface{}{
                "entityTitleName": entity.EntityTitleName,
                "fieldName": field.Name,
                "fieldVerboseName": field.VerboseName,
                "minLength": field.Min,
            })

            validRules += "\r\n"
        }

        maxSizeTemplate :=
        `    if f.{{.entityTitleName}}.{{.fieldName}} != "" {
        validation.MaxSize(f.{{.entityTitleName}}.{{.fieldName}}, {{.maxLength}}).Message("{{.fieldVerboseName}}长度不能大于{{.maxLength}}！")
    }`

        if field.Max != "" && field.Max != "0"   {
            validRules += core.FormatText("valid_max", maxSizeTemplate, map[string]interface{}{
                "entityTitleName": entity.EntityTitleName,
                "fieldName": field.Name,
                "fieldVerboseName": field.VerboseName,
                "maxLength": field.Max,
            })

            validRules += "\r\n"
        }
    }

    return validRules
}

var checkedUniqueCreateFields map[string]bool = make(map[string]bool)

func TemplateFuncCheckUniqueCreate(entity smith.Entity, field smith.Field) string {
    //进入这个函数说明外面已经检查过 unique 不为 false 的
    if field.Unique == "false" {
        return ""
    }

    singleUniquekey := entity.ModuleTitleName + "_" + entity.EntityTitleName + "_" + field.Name

    if _, has := checkedUniqueCreateFields[singleUniquekey]; has == false && strings.Index(field.Unique, ",") == -1 {
        //含有 "," 的 unique 字段 别的字段会考虑，可以略过，方便处理
        isSingleUnique := true

        uniqueFields := make([]smith.Field, 0)
        if field.Unique == "true" {
            isSingleUnique = true
        } else {
            //这时不可能是 false

            for _, fieldInEntity := range (entity.FieldList) {
                //查找这个 entity 里面的所有包含这个 unique 的 field
                if strings.Contains(fieldInEntity.Unique, field.Unique) {
                    //field.Unique中是不会有 "," 的
                    uniqueFields = append(uniqueFields, fieldInEntity)
                }
            }

            if len(uniqueFields) != 1 {
                isSingleUnique = false
            }
        }

        if isSingleUnique {
            //是 单独 unique

            template :=
            `        {{.fieldCamelCase}}Count, err := session.Where("{{.field.Column}} = ?", {{.entity.EntityCamelCase}}.{{.field.Name}}).Count(new(models.{{.entity.EntityTitleName}}))
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

            checkedUniqueCreateFields[singleUniquekey] = true

            return checkUniqueCode
        } else {
            //是 组合 unique
            fieldsCompare := ""
            fieldsValue := ""
            chineseNames := ""

            isFirst := true
            for _, uniqueField := range (uniqueFields) {
                compositeUniquekey := entity.ModuleTitleName + "_" + entity.EntityTitleName + "_" + uniqueField.Name

                if isFirst {
                    fieldsCompare += uniqueField.Column + " = ?"
                    fieldsValue += entity.EntityCamelCase + "." + uniqueField.Name
                    chineseNames += uniqueField.VerboseName

                    isFirst = false
                } else {
                    fieldsCompare += " and " + uniqueField.Column + " = ?"
                    fieldsValue += ", " + entity.EntityCamelCase + "." + uniqueField.Name
                    chineseNames += "、" + uniqueField.VerboseName
                }

                checkedUniqueCreateFields[compositeUniquekey] = true
            }

            where := fieldsCompare + ", " + fieldsValue

            template :=
            `        {{.fieldCamelCase}}Count, err := session.Where({{.where}}).Count(new(models.{{.entity.EntityTitleName}}))
        core.HandleError(err)
        if {{.fieldCamelCase}}Count != 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，{{.chineseNames}}组合必须唯一！" })
        }
        `
            checkUniqueCode := core.FormatText("check_unique_create", template, map[string]interface{}{
                "entity": entity,
                "field": field,
                "fieldCamelCase": strings.ToLower(field.Name[0:1]) + field.Name[1:],
                "where": where,
                "chineseNames": chineseNames,
            })

            return checkUniqueCode
        }
    } else {
        return ""
    }
}

var checkedUniqueUpdateFields map[string]bool = make(map[string]bool)

func TemplateFuncCheckUniqueUpdate(entity smith.Entity, field smith.Field) string {
    //进入这个函数说明外面已经检查过 unique 不为 false 的
    if field.Unique == "false" {
        return ""
    }

    singleUniquekey := entity.ModuleTitleName + "_" + entity.EntityTitleName + "_" + field.Name

    if _, has := checkedUniqueUpdateFields[singleUniquekey]; has == false && strings.Index(field.Unique, ",") == -1 {
        //含有 "," 的 unique 字段 别的字段会考虑，可以略过，方便处理
        isSingleUnique := true

        uniqueFields := make([]smith.Field, 0)
        if field.Unique == "true" {
            isSingleUnique = true
        } else {
            //这时不可能是 false

            for _, fieldInEntity := range (entity.FieldList) {
                //查找这个 entity 里面的所有包含这个 unique 的 field
                if strings.Contains(fieldInEntity.Unique, field.Unique) {
                    //field.Unique中是不会有 "," 的
                    uniqueFields = append(uniqueFields, fieldInEntity)
                }
            }

            if len(uniqueFields) != 1 {
                isSingleUnique = false
            }
        }

        if isSingleUnique {
            //是 单独 unique

            template :=
            `        {{.fieldCamelCase}}Count, err := session.Where("id <> ? and {{.field.Column}} = ?", {{.entity.EntityCamelCase}}.Id, {{.entity.EntityCamelCase}}.{{.field.Name}}).Count(new(models.{{.entity.EntityTitleName}}))
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

            checkedUniqueUpdateFields[singleUniquekey] = true

            return checkUniqueCode
        } else {
            //是 组合 unique
            fieldsCompare := "id <> ?"
            fieldsValue := entity.EntityCamelCase + ".Id"
            chineseNames := ""

            isFirst := true
            for _, uniqueField := range (uniqueFields) {
                compositeUniquekey := entity.ModuleTitleName + "_" + entity.EntityTitleName + "_" + uniqueField.Name

                fieldsCompare += " and " + uniqueField.Column + " = ?"
                fieldsValue += ", " + entity.EntityCamelCase + "." + uniqueField.Name

                if isFirst {
                    chineseNames += uniqueField.VerboseName
                    isFirst = false
                } else {
                    chineseNames += "、" + uniqueField.VerboseName
                }

                checkedUniqueUpdateFields[compositeUniquekey] = true
            }

            where := fieldsCompare + ", " + fieldsValue

            template :=
            `        {{.fieldCamelCase}}Count, err := session.Where({{.where}}).Count(new(models.{{.entity.EntityTitleName}}))
        core.HandleError(err)
        if {{.fieldCamelCase}}Count != 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，{{.chineseNames}}已存在！" })
        }
        `
            checkUniqueCode := core.FormatText("check_unique_create", template, map[string]interface{}{
                "entity": entity,
                "field": field,
                "fieldCamelCase": strings.ToLower(field.Name[0:1]) + field.Name[1:],
                "where": where,
                "chineseNames": chineseNames,
            })

            return checkUniqueCode
        }
    } else {
        return ""
    }
}

func TemplateFuncLowerCase(name string) string {
    return strings.ToLower(name)
}