package core

import "matrix/smith/core/fieldtype"

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

type Entity struct {
	ModuleTitleName string
	ModuleLowerCase string
	ModuleChinese   string
	EntityTitleName string
	EntityCamelCase string
	EntityChinese   string
	TableName       string
	TablePrefix     string

	FieldList []Field
}
