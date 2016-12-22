package main

import (
	"bytes"
	"fmt"
	"matrix/core"
	entity_models "matrix/modules/itsm/models"
	"matrix/app/models"
	smith_core "matrix/smith/core"
	"os"
	"path"
)
//var output_base_dir = "/home/wanglihua/gopath/src/matrix"
var output_base_dir = "z:/GoPath/src/matrix"
var module_title_name = entity_models.ModuleTitleName
var module_lower_case = entity_models.ModuleLowerCase
var module_chinese = entity_models.ModuleChinese
var table_prefix = entity_models.TablePrefix
var model_list = []interface{}{
	models.ConfigInfo{},
}

func main() {
	entity_list := smith_core.GetEntityList(model_list, module_title_name, module_lower_case, module_chinese, table_prefix, true)
	for _, entity := range entity_list {

		var code_buffer bytes.Buffer

		code_buffer.WriteString("package models\n\n")
		code_buffer.WriteString(fmt.Sprintf("var %sCols = struct {\n", entity.EntityTitleName))

		for _, field := range entity.FieldList {
			code_buffer.WriteString(fmt.Sprintf("\t%s string\n", field.Name))
		}

		code_buffer.WriteString("}{\n")

		for _, field := range entity.FieldList {
			code_buffer.WriteString(fmt.Sprintf("\t%s: \"%s\",\n", field.Name, field.Column))
		}

		code_buffer.WriteString("}\n")

		models_dir := path.Join(output_base_dir, "modules", entity_list[0].ModuleLowerCase, "models")
		err := os.MkdirAll(models_dir, 0777)
		core.HandleError(err)
		meta_path_file := path.Join(models_dir, fmt.Sprintf("%sInfo.Meta.go", entity.EntityTitleName))
		smith_core.WriteToFile(meta_path_file, code_buffer.String())
	}
}
