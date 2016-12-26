package main

import (
	"fmt"
	"matrix/core"
	"matrix/modules/erp/models"
	smith_core "matrix/smith/core"
	"matrix/smith/core/template"
	"os"
)

//var output_base_dir = "/home/wanglihua/gopath/src/matrix"
var output_base_dir = "z:/GoPath/src/matrix"
var import_path = "matrix"
var module_title_name = models.ModuleTitleName
var module_lower_case = models.ModuleLowerCase
var module_chinese = models.ModuleChinese
var table_prefix = models.TablePrefix
var model_list = []interface{}{
	models.SupplierInfo{},
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

	entityList := smith_core.GetEntityList(model_list, module_title_name, module_lower_case, module_chinese, table_prefix, false)

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
