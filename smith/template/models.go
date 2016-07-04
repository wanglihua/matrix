package template

var ModelsTemplate = `
package models

import (
    "matrix/core"
)

var TablePrefix = "{{.tablePrefix}}"

{{range $entityIndex, $entity := .entityList}}
//---------------------------------------------------------------------------------------------------------------

type {{$entity.EntityTitleName}} struct {
{{range $fieldIndex, $field := $entity.FieldList}}{{RenderField $field}}
{{end}}
}

func (e {{$entity.EntityTitleName}}) TableName() string {
    return TablePrefix + "{{$entity.EntityLowerCase}}"
}

{{end}}

`
