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
    Id             int64             {{$.tagchar}}xorm:"bigint notnull pk autoincr 'id'" json:"id"{{$.tagchar}}

{{range $fieldIndex, $field := $entity.FieldList}}{{RenderField $field}}
{{end}}
    CreateTime     core.Time         {{$.tagchar}}xorm:"created notnull 'create_time'" json:"create_time"{{$.tagchar}}
    UpdateTime     core.Time         {{$.tagchar}}xorm:"updated notnull 'update_time'" json:"update_time"{{$.tagchar}}
    Version        int               {{$.tagchar}}xorm:"version notnull 'version'" json:"version"{{$.tagchar}}
}

func (e {{$entity.EntityTitleName}}) TableName() string {
    return TablePrefix + "{{$entity.TableName}}"
}

{{end}}
`
