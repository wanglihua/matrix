package template

var RouteTemplate = `
{{range $entityIndex, $entity := .entityList}}
GET  /{{$entity.ModuleLowerCase}}/{{LowerCase $entity.Route}}                    {{$entity.ModuleTitleName}}{{$entity.EntityTitleName}}.Index
POST /{{$entity.ModuleLowerCase}}/{{LowerCase $entity.Route}}/data/list          {{$entity.ModuleTitleName}}{{$entity.EntityTitleName}}.ListData
POST /{{$entity.ModuleLowerCase}}/{{LowerCase $entity.Route}}/data/detail        {{$entity.ModuleTitleName}}{{$entity.EntityTitleName}}.DetailData
POST /{{$entity.ModuleLowerCase}}/{{LowerCase $entity.Route}}/save               {{$entity.ModuleTitleName}}{{$entity.EntityTitleName}}.Save
POST /{{$entity.ModuleLowerCase}}/{{LowerCase $entity.Route}}/delete             {{$entity.ModuleTitleName}}{{$entity.EntityTitleName}}.Delete
{{end}}`

