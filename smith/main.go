package main

import (
	"fmt"
	"matrix/core"
	smith_core "matrix/smith/core"
	"matrix/smith/core/fieldtype"
	"matrix/smith/core/template"
	"os"
	"reflect"
	"strings"

	xorm_core "github.com/go-xorm/core"

	entity_models "matrix/modules/itsm/models"
)

//var output_base_dir = "/home/wanglihua/code_gen"
//var output_base_dir = "z:/GoPath/src/Matrix"

var output_base_dir = "/home/wanglihua/gopath/src/matrix"
var import_path = "matrix"
var module_title_name = entity_models.ModuleTitleName
var module_lower_case = entity_models.ModuleLowerCase
var module_chinese = entity_models.ModuleChinese
var table_prefix = entity_models.TablePrefix
var model_list = []interface{}{
	entity_models.EventPriority{}, entity_models.EventLog{},
	entity_models.EquipmentStatus{}, entity_models.EquipmentType{}, entity_models.Equipment{},
	entity_models.ApplicationStatus{}, entity_models.ApplicationType{}, entity_models.Application{},
}

/*
	sample field definition
	Field1 float64 `xorm:"decimal(18,4) notnull 'field1'" json:"field1" smith:"verbose_name=字段1,min=0,max=100"`
	sample model methods

	func (e Entity) TableName() string {
		return TablePrefix + "table_name"
	}

	func (e Entity) ModelDesc() string {
		return "verbose_name=中文名;entity_json=product_cate;route=product/cate"
	}
*/

func main() {
	var err error

	entityList := make([]smith_core.Entity, 0)
	for _, model := range model_list {
		entity := smith_core.Entity{}
		entity.ModuleTitleName = module_title_name
		entity.ModuleLowerCase = module_lower_case
		entity.ModuleChinese = module_chinese
		entity.EntityTitleName, entity.EntityCamelCase, entity.TableName, entity.EntityChinese, entity.EntityJson, entity.Route = get_model_info(model)
		entity.EntityJsonTag = fmt.Sprintf("`json:\"%s\"`", entity.EntityJson) //在我们的模板定义中输出 ` 不太方便，所以在这里就先把 ` 放进去
		entity.TablePrefix = table_prefix

		if strings.TrimSpace(entity.Route) == "" {
			entity.Route = entity.EntityCamelCase
		}

		entity.FieldList = make([]smith_core.Field, 0)

		model_type := reflect.TypeOf(model)
		model_field_num := model_type.NumField()
		for model_field_index := 0; model_field_index < model_field_num; model_field_index++ {
			field := smith_core.Field{}
			model_field := model_type.Field(model_field_index)

			field.Name = model_field.Name
			if field.Name == "Id" || field.Name == "CreateTime" || field.Name == "UpdateTime" || field.Name == "Version" {
				continue
			}

			field_xorm_column, unique := smith_core.GetFieldXormColumn(model_field)

			field.Column = field_xorm_column.FieldName

			switch field_xorm_column.SQLType.Name {
			case xorm_core.Bit, xorm_core.TinyInt, xorm_core.SmallInt, xorm_core.MediumInt, xorm_core.Int, xorm_core.Integer:
				field.FieldType = fieldtype.Int
			case xorm_core.BigInt:
				field.FieldType = fieldtype.BigInt
			case xorm_core.Varchar, xorm_core.NVarchar, xorm_core.Text, xorm_core.TinyText, xorm_core.MediumInt, xorm_core.LongText:
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

	err = os.MkdirAll(output_base_dir, 0777) //创建目录，如果目录已经存在就不会创建，这里是保证目录存在
	core.HandleError(err)

	controllersDir := output_base_dir + "/modules/" + entityList[0].ModuleLowerCase + "/app/controllers"
	err = os.MkdirAll(controllersDir, 0777)
	core.HandleError(err)

	for _, entity := range entityList {
		controllerCode := smith_core.RenderCodeTemplate("controller", template.ControllerTemplate, map[string]interface{}{
			"entity":      entity,
			"import_path": import_path,
		})

		smith_core.WriteToFile(controllersDir+"/"+entityList[0].ModuleLowerCase+"_"+smith_core.ToUnderscore(entity.EntityCamelCase)+".go", controllerCode)
	}

	viewsDir := output_base_dir + "/modules/" + entityList[0].ModuleLowerCase + "/app/views/" + entityList[0].ModuleLowerCase
	err = os.MkdirAll(viewsDir, 0777)
	core.HandleError(err)

	for _, entity := range entityList {

		indexhtmlCode := smith_core.RenderCodeTemplate("indexhtml", template.IndexHtmlTemplate, map[string]interface{}{
			"entity": entity,
		})
		entity_view_dir := viewsDir + "/" + smith_core.ToUnderscore(entity.EntityCamelCase)
		err = os.MkdirAll(entity_view_dir, 0777)
		core.HandleError(err)
		smith_core.WriteToFile(entity_view_dir+"/"+smith_core.ToUnderscore(entity.EntityCamelCase)+"_index.html", indexhtmlCode)
	}

	confDir := output_base_dir + "/modules/" + entityList[0].ModuleLowerCase + "/conf"
	err = os.MkdirAll(confDir, 0777)
	core.HandleError(err)
	routeCode := smith_core.RenderCodeTemplate("route", template.RouteTemplate, map[string]interface{}{
		"entityList": entityList,
	})

	smith_core.AppendToFile(confDir+"/routes", routeCode)

	menuDir := output_base_dir + "/modules/" + entityList[0].ModuleLowerCase + "/app/views/"
	err = os.MkdirAll(menuDir, 0777)
	core.HandleError(err)
	menuCode := smith_core.RenderCodeTemplate("menu", template.MenuTemplate, map[string]interface{}{
		"entityList": entityList,
	})

	smith_core.WriteToFile(menuDir+"/menu.html", menuCode)

	fmt.Println("Code Generated!")
}

func get_model_info(model interface{}) (title_name, camel_case_name, table_name, chinese_name, entity_json, route string) {
	model_type := reflect.TypeOf(model)
	title_name = model_type.Name()
	camel_case_name = strings.ToLower(title_name)[0:1] + title_name[1:]

	model_value := reflect.ValueOf(model)

	table_name = ""
	if tn, ok := model_value.Interface().(TableName); ok {
		table_name = tn.TableName()
	}

	chinese_name = title_name
	entity_json = smith_core.ToUnderscore(camel_case_name)
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
