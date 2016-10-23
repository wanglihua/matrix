package template

var RouteTemplate = `
{{range $entityIndex, $entity := .entityList}}
GET  /{{$entity.ModuleLowerCase}}/{{LowerCase $entity.EntityCamelCase}}                    {{$entity.ModuleTitleName}}{{$entity.EntityTitleName}}.Index
POST /{{$entity.ModuleLowerCase}}/{{LowerCase $entity.EntityCamelCase}}/data/list          {{$entity.ModuleTitleName}}{{$entity.EntityTitleName}}.ListData
POST /{{$entity.ModuleLowerCase}}/{{LowerCase $entity.EntityCamelCase}}/data/detail        {{$entity.ModuleTitleName}}{{$entity.EntityTitleName}}.DetailData
POST /{{$entity.ModuleLowerCase}}/{{LowerCase $entity.EntityCamelCase}}/save               {{$entity.ModuleTitleName}}{{$entity.EntityTitleName}}.Save
POST /{{$entity.ModuleLowerCase}}/{{LowerCase $entity.EntityCamelCase}}/delete             {{$entity.ModuleTitleName}}{{$entity.EntityTitleName}}.Delete
{{end}}

`
