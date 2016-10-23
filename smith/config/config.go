package config

import (
	entity_models "matrix/modules/inventory/models"
)

//var OutputBaseDir = "/home/wanglihua/gopath/src/matrix"
var OutputBaseDir = "/home/wanglihua/code_gen"

var ModuleTitleName = entity_models.ModuleTitleName
var ModuleLowerCase = entity_models.ModuleLowerCase
var ModuleChinese = entity_models.ModuleChinese
var TablePrefix = entity_models.TablePrefix
var ModelList = []interface{}{
	entity_models.Product{},
}
