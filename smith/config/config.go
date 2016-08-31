package config

import (
	entity_models "matrix/modules/st/models"
)

var OutputBaseDir = "D:/projects/rc_sales_target/src"

var ModuleTitleName = entity_models.ModuleTitleName
var ModuleLowerCase = entity_models.ModuleLowerCase
var ModuleChinese = entity_models.ModuleChinese
var TablePrefix = entity_models.TablePrefix
var ModelList = []interface{}{
	entity_models.SalesMonthReport{},
}
