package core

import (
	"fmt"
	xorm_core "github.com/go-xorm/core"
	"matrix/smith/core/fieldtype"
	"reflect"
	"strings"
)

type Field struct {
	VerboseName string
	Name        string
	Column      string
	FieldType   fieldtype.FieldType
	Length      int
	Precision   int
	Scale       int
	Unique      string // "true" or unique name
	Null        bool
	Blank       bool
	Min         string //如果是数字型型就是最小值，如果是字符串型就是最小长度
	Max         string //如果是数字型型就是最大值，如果是字符串型就是最大长度
}

func (f Field) IsNumber() bool {
	if f.FieldType == fieldtype.Int || f.FieldType == fieldtype.BigInt || f.FieldType == fieldtype.Decimal || f.FieldType == fieldtype.Version {
		return true
	} else {
		return false
	}
}

type Entity struct {
	ModuleTitleName string
	ModuleLowerCase string
	ModuleChinese   string
	EntityTitleName string
	EntityCamelCase string
	EntityChinese   string
	EntityJson      string //比如: product
	EntityJsonTag   string //比如: `json:"product"`
	TableName       string
	TablePrefix     string
	Route           string

	FieldList []Field
}

func GetEntityList(model_list []interface{}, module_title_name string, module_lower_case string, module_chinese string, table_prefix string, for_meta bool) []Entity {
	entityList := make([]Entity, 0)
	for _, model := range model_list {
		entity := Entity{}
		entity.ModuleTitleName = module_title_name
		entity.ModuleLowerCase = module_lower_case
		entity.ModuleChinese = module_chinese
		entity.EntityTitleName, entity.EntityCamelCase, entity.TableName, entity.EntityChinese, entity.EntityJson, entity.Route = GetModelInfo(model)
		entity.EntityJsonTag = fmt.Sprintf("`json:\"%s\"`", entity.EntityJson) //在我们的模板定义中输出 ` 不太方便，所以在这里就先把 ` 放进去
		entity.TablePrefix = table_prefix

		if strings.TrimSpace(entity.Route) == "" {
			entity.Route = entity.EntityCamelCase
		}

		entity.FieldList = make([]Field, 0)

		model_type := reflect.TypeOf(model)
		model_field_num := model_type.NumField()
		for model_field_index := 0; model_field_index < model_field_num; model_field_index++ {
			field := Field{}
			model_field := model_type.Field(model_field_index)

			field.Name = model_field.Name
			if for_meta == false && (field.Name == "Id" || field.Name == "CreateTime" || field.Name == "UpdateTime" || field.Name == "Version") {
				continue
			}

			field_xorm_column, unique := GetFieldXormColumn(model_field)

			field.Column = field_xorm_column.FieldName

			switch field_xorm_column.SQLType.Name {
			case xorm_core.Bit, xorm_core.TinyInt, xorm_core.SmallInt, xorm_core.MediumInt, xorm_core.Int, xorm_core.Integer:
				field.FieldType = fieldtype.Int
			case xorm_core.BigInt:
				field.FieldType = fieldtype.BigInt
			case xorm_core.Varchar, xorm_core.NVarchar, xorm_core.Text, xorm_core.TinyText, xorm_core.MediumText, xorm_core.LongText:
				field.FieldType = fieldtype.NVarchar
			case xorm_core.Date, xorm_core.DateTime, xorm_core.Time, xorm_core.TimeStamp, xorm_core.TimeStampz:
				field.FieldType = fieldtype.DateTime
			case xorm_core.Decimal, xorm_core.Float, xorm_core.Double:
				field.FieldType = fieldtype.Decimal
			case xorm_core.Bool:
				field.FieldType = fieldtype.Boolean
			default:
				field.FieldType = fieldtype.NVarchar
			}

			if field_xorm_column.IsCreated {
				field.FieldType = fieldtype.Create
			}

			if field_xorm_column.IsUpdated {
				field.FieldType = fieldtype.Update
			}

			if field_xorm_column.IsVersion {
				field.FieldType = fieldtype.Version
			}

			if field.FieldType == fieldtype.NVarchar {
				field.Length = field_xorm_column.Length
			}

			field.Null = field_xorm_column.Nullable
			field.Blank = field_xorm_column.Nullable
			field.Unique = unique

			field.Precision = field_xorm_column.Length
			field.Scale = field_xorm_column.Length2

			smith_tag := model_field.Tag.Get("smith")

			scan_count := 0

			var field_verbose_name string
			verbose_name_split_str_list := strings.Split(smith_tag, "verbose_name")
			if len(verbose_name_split_str_list) > 1 {
				scan_count, _ = fmt.Sscanf(verbose_name_split_str_list[1], "=%s", &field_verbose_name)
				if scan_count != 0 {
					field_verbose_name = strings.TrimSpace(strings.Split(field_verbose_name, ",")[0])
				}
			}
			field.VerboseName = field_verbose_name

			if strings.TrimSpace(field_verbose_name) == "" {
				field.VerboseName = field.Name
			}

			field_min := "0"
			min_split_str_list := strings.Split(smith_tag, "min")
			if len(min_split_str_list) > 1 {
				scan_count, _ = fmt.Sscanf(min_split_str_list[1], "=%s", &field_min)
				if scan_count != 0 {
					field_min = strings.TrimSpace(strings.Split(field_min, ",")[0])
				}
			}
			field.Min = field_min

			field_max := "0"
			max_split_str_list := strings.Split(smith_tag, "max")
			if len(max_split_str_list) > 1 {
				scan_count, _ = fmt.Sscanf(max_split_str_list[1], "=%s", &field_max)
				if scan_count != 0 {
					field_max = strings.TrimSpace(strings.Split(field_max, ",")[0])
				}
			}
			field.Max = field_max

			entity.FieldList = append(entity.FieldList, field)
		}

		entityList = append(entityList, entity)
	}

	return entityList
}

func GetModelInfo(model interface{}) (title_name, camel_case_name, table_name, chinese_name, entity_json, route string) {
	model_type := reflect.TypeOf(model)
	title_name = model_type.Name()
	if strings.HasSuffix(title_name, "Info") {
		title_name = strings.TrimSuffix(title_name, "Info")
	}
	camel_case_name = strings.ToLower(title_name)[0:1] + title_name[1:]

	model_value := reflect.ValueOf(model)

	table_name = ""
	if tn, ok := model_value.Interface().(TableName); ok {
		table_name = tn.TableName()
	}

	chinese_name = title_name
	entity_json = ToUnderscore(camel_case_name)
	route = camel_case_name
	if md, ok := model_value.Interface().(ModelDesc); ok {
		model_desc := md.ModelDesc()
		verbose_name_split_str_list := strings.Split(model_desc, "verbose_name")
		if len(verbose_name_split_str_list) == 2 {
			scan_count, _ := fmt.Sscanf(verbose_name_split_str_list[1], "=%s", &chinese_name)
			if scan_count != 0 {
				chinese_name = strings.TrimSpace(strings.Split(chinese_name, ",")[0])
			}
		}
		entity_json_split_str_list := strings.Split(model_desc, "entity_json")
		if len(entity_json_split_str_list) == 2 {
			scan_count, _ := fmt.Sscanf(entity_json_split_str_list[1], "=%s", &entity_json)
			if scan_count != 0 {
				entity_json = strings.TrimSpace(strings.Split(entity_json, ",")[0])
			}
		}
		route_split_str_list := strings.Split(model_desc, "route")
		if len(route_split_str_list) > 1 {
			scan_count, _ := fmt.Sscanf(route_split_str_list[1], "=%s", &route)
			if scan_count != 0 {
				route = strings.TrimSpace(strings.Split(route, ",")[0])
			}
		}
	}

	return
}

type TableName interface {
	TableName() string
}

type ModelDesc interface {
	ModelDesc() string
}
