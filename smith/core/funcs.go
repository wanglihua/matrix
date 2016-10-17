package core

import (
	"os"
	"bytes"
	"reflect"
	"strings"
	"time"
	"strconv"
	matrix_core "matrix/core"
	textTemplate "text/template"
	xorm_core "github.com/go-xorm/core"
)

func WriteToFile(fileName string, code string) {

	checkFileIsExist := func(filename string) bool {
		var exist = true;
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			exist = false;
		}

		return exist;
	}

	if checkFileIsExist(fileName) {
		return //文件存在就返回，不做任何事情
	} else {
		var file *os.File
		var err error

		file, err = os.Create(fileName)  //创建文件
		matrix_core.HandleError(err)

		_, err = file.WriteString(code)
		matrix_core.HandleError(err)

		file.Sync()
		file.Close()
	}
}

func AppendToFile(fileName string, code string) {

	checkFileIsExist := func(filename string) bool {
		var exist = true;
		if _, err := os.Stat(filename); os.IsNotExist(err) {
			exist = false;
		}

		return exist;
	}

	if checkFileIsExist(fileName) {
		var file *os.File
		var err error
		file, err = os.OpenFile(fileName, os.O_APPEND | os.O_WRONLY, 0666)  //打开文件
		matrix_core.HandleError(err)

		_, err = file.WriteString(code)
		matrix_core.HandleError(err)

		file.Sync()
		file.Close()
	} else {
		return //如果文件不存在就返回, 不做任何事
	}
}

func RenderCodeTemplate(tplName string, tplContent string, args  map[string]interface{}) string {
	template, err := textTemplate.New(tplName).Funcs(textTemplate.FuncMap{
		"FirstEntity": TemplateFuncFirstEntity,
		"FieldSearchable": TemplateFuncFieldSearchable,
		"ListMaxIndex": TemplateFuncListMaxIndex,
		"FieldClienValid": TemplateFuncFieldClienValid,
		"FieldValid": TemplateFuncFieldValid,
		"CheckUniqueCreate": TemplateFuncCheckUniqueCreate,
		"CheckUniqueUpdate": TemplateFuncCheckUniqueUpdate,
		"LowerCase": TemplateFuncLowerCase,
		"Underscore": TemplateFuncUnderscore,
	}).Parse(tplContent)

	if err != nil {
		panic(err)
	}

	var buffer bytes.Buffer
	template.Execute(&buffer, args)
	return buffer.String()
}

func GetFieldXormColumn(model_field reflect.StructField) (xorm_column *xorm_core.Column, unique string) {

	tag := model_field.Tag

	orm_tag_str := tag.Get("xorm")
	model_field_value := reflect.ValueOf(model_field)
	model_field_type := model_field_value.Type()

	unique = "false"

	var err error
	if orm_tag_str != "" {
		xorm_column = &xorm_core.Column{FieldName: model_field.Name, Nullable: true, IsPrimaryKey: false,
			IsAutoIncrement: false, MapType: xorm_core.TWOSIDES, Indexes: make(map[string]int)}
		tags := split_tag(orm_tag_str)

		if len(tags) > 0 {
			if tags[0] == "-" {
				return nil, unique
			}
			if strings.ToUpper(tags[0]) == "EXTENDS" {
				switch model_field_value.Kind() {
				case reflect.Ptr:
					f := model_field_value.Type().Elem()
					if f.Kind() == reflect.Struct {
						fieldPtr := model_field_value
						model_field_value = model_field_value.Elem()
						if !model_field_value.IsValid() || fieldPtr.IsNil() {
							model_field_value = reflect.New(f).Elem()
						}
					}
					fallthrough
				case reflect.Struct:
					//do nothing
					return nil, unique
				default:
				//do nothing
				}
			}

			var preKey string
			for j, key := range tags {
				k := strings.ToUpper(key)
				switch {
				case k == "<-":
					xorm_column.MapType = xorm_core.ONLYFROMDB
				case k == "->":
					xorm_column.MapType = xorm_core.ONLYTODB
				case k == "PK":
					xorm_column.IsPrimaryKey = true
					xorm_column.Nullable = false
				case k == "NULL":
					if j == 0 {
						xorm_column.Nullable = true
					} else {
						xorm_column.Nullable = (strings.ToUpper(tags[j - 1]) != "NOT")
					}
				case k == "AUTOINCR":
					xorm_column.IsAutoIncrement = true
				case k == "DEFAULT":
					xorm_column.Default = tags[j + 1]
				case k == "CREATED":
					xorm_column.IsCreated = true
				case k == "VERSION":
					xorm_column.IsVersion = true
					xorm_column.Default = "1"
				case k == "UTC":
					xorm_column.TimeZone = time.UTC
				case k == "LOCAL":
					xorm_column.TimeZone = time.Local
				case strings.HasPrefix(k, "LOCALE(") && strings.HasSuffix(k, ")"):
					location := k[len("INDEX") + 1 : len(k) - 1]
					xorm_column.TimeZone, err = time.LoadLocation(location)
					matrix_core.HandleError(err)
				case k == "UPDATED":
					xorm_column.IsUpdated = true
				case k == "DELETED":
					xorm_column.IsDeleted = true
				case strings.HasPrefix(k, "INDEX(") && strings.HasSuffix(k, ")"):
				//do nothing
				case k == "INDEX":
				//do nothing
				case strings.HasPrefix(k, "UNIQUE(") && strings.HasSuffix(k, ")"):
					unique = k[len("UNIQUE") + 1 : len(k) - 1]
				case k == "UNIQUE":
					if unique == "false" {
						unique = "true"
					}
				case k == "NOTNULL":
					xorm_column.Nullable = false
				case k == "CACHE":
				//do nothing
				case k == "NOCACHE":
				//do nothing
				case k == "NOT":
				default:
					if strings.HasPrefix(k, "'") && strings.HasSuffix(k, "'") {
						if preKey != "DEFAULT" {
							//xorm_column.Name = key[1 : len(key) - 1]
							xorm_column.FieldName = key[1 : len(key) - 1]
						}
					} else if strings.Contains(k, "(") && strings.HasSuffix(k, ")") {
						fs := strings.Split(k, "(")

						if _, ok := xorm_core.SqlTypes[fs[0]]; !ok {
							preKey = k
							continue
						}
						xorm_column.SQLType = xorm_core.SQLType{fs[0], 0, 0}
						if fs[0] == xorm_core.Enum && fs[1][0] == '\'' {
							//enum
							options := strings.Split(fs[1][0:len(fs[1]) - 1], ",")
							xorm_column.EnumOptions = make(map[string]int)
							for k, v := range options {
								v = strings.TrimSpace(v)
								v = strings.Trim(v, "'")
								xorm_column.EnumOptions[v] = k
							}
						} else if fs[0] == xorm_core.Set && fs[1][0] == '\'' {
							//set
							options := strings.Split(fs[1][0:len(fs[1]) - 1], ",")
							xorm_column.SetOptions = make(map[string]int)
							for k, v := range options {
								v = strings.TrimSpace(v)
								v = strings.Trim(v, "'")
								xorm_column.SetOptions[v] = k
							}
						} else {
							fs2 := strings.Split(fs[1][0:len(fs[1]) - 1], ",")
							if len(fs2) == 2 {
								xorm_column.Length, err = strconv.Atoi(fs2[0])
								matrix_core.HandleError(err)
								xorm_column.Length2, err = strconv.Atoi(fs2[1])
								matrix_core.HandleError(err)
							} else if len(fs2) == 1 {
								xorm_column.Length, err = strconv.Atoi(fs2[0])
								matrix_core.HandleError(err)
							}
						}
					} else {
						if _, ok := xorm_core.SqlTypes[k]; ok {
							xorm_column.SQLType = xorm_core.SQLType{k, 0, 0}
						} else if key != xorm_column.Default {
							xorm_column.Name = key
						}
					}
				}
				preKey = k
			}
			if xorm_column.SQLType.Name == "" {
				xorm_column.SQLType = xorm_core.Type2SQLType(model_field_type)
			}
			if xorm_column.Length == 0 {
				xorm_column.Length = xorm_column.SQLType.DefaultLength
			}
			if xorm_column.Length2 == 0 {
				xorm_column.Length2 = xorm_column.SQLType.DefaultLength2
			}

			if xorm_column.Name == "" {
				//col.Name = engine.ColumnMapper.Obj2Table(model_field.Name)
				xorm_column.Name = model_field.Name
			}
		}
	} else {
		var sqlType xorm_core.SQLType
		if model_field_value.CanAddr() {
			if _, ok := model_field_value.Addr().Interface().(xorm_core.Conversion); ok {
				sqlType = xorm_core.SQLType{xorm_core.Text, 0, 0}
			}
		}
		if _, ok := model_field_value.Interface().(xorm_core.Conversion); ok {
			sqlType = xorm_core.SQLType{xorm_core.Text, 0, 0}
		} else {
			sqlType = xorm_core.Type2SQLType(model_field_type)
		}
		//col = xorm_core.NewColumn(engine.ColumnMapper.Obj2Table(model_field.Name),
		xorm_column = xorm_core.NewColumn(model_field.Name,
			model_field.Name, sqlType, sqlType.DefaultLength,
			sqlType.DefaultLength2, true)
	}
	if xorm_column.IsAutoIncrement {
		xorm_column.Nullable = false
	}

	return xorm_column, unique
}

func split_tag(tag string) (tags []string) {
	tag = strings.TrimSpace(tag)
	var hasQuote = false
	var lastIdx = 0
	for i, t := range tag {
		if t == '\'' {
			hasQuote = !hasQuote
		} else if t == ' ' {
			if lastIdx < i && !hasQuote {
				tags = append(tags, strings.TrimSpace(tag[lastIdx:i]))
				lastIdx = i + 1
			}
		}
	}
	if lastIdx < len(tag) {
		tags = append(tags, strings.TrimSpace(tag[lastIdx:len(tag)]))
	}
	return
}
