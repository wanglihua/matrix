package template

var RouteTemplate = `
{{range $entityIndex, $entity := .entityList}}
GET  /{{$entity.ModuleLowerCase}}/{{$entity.EntityLowerCase}}                    {{$entity.ModuleTitleName}}{{$entity.EntityTitleName}}.Index
POST /{{$entity.ModuleLowerCase}}/{{$entity.EntityLowerCase}}/listdata           {{$entity.ModuleTitleName}}{{$entity.EntityTitleName}}.ListData
POST /{{$entity.ModuleLowerCase}}/{{$entity.EntityLowerCase}}/detail             {{$entity.ModuleTitleName}}{{$entity.EntityTitleName}}.Detail
POST /{{$entity.ModuleLowerCase}}/{{$entity.EntityLowerCase}}/save               {{$entity.ModuleTitleName}}{{$entity.EntityTitleName}}.Save
POST /{{$entity.ModuleLowerCase}}/{{$entity.EntityLowerCase}}/delete             {{$entity.ModuleTitleName}}{{$entity.EntityTitleName}}.Delete
{{end}}

`
