package main

const ROUTES_TEMPLATE = `package routes

import "github.com/revel/revel"

{{range $i, $c := .Controllers}}
type t{{.StructName}} struct {}
var {{.StructName}} t{{.StructName}}

{{range .MethodSpecs}}
func (_ t{{$c.StructName}}) {{.Name}}({{range .Args}}
        {{.Name}} {{if .ImportPath}}interface{}{{else}}{{.TypeExpr.TypeName ""}}{{end}},{{end}}
        ) string {
    args := make(map[string]string)
    {{range .Args}}
    revel.Unbind(args, "{{.Name}}", {{.Name}}){{end}}
    return revel.MainRouter.Reverse("{{$c.StructName}}.{{.Name}}", args).Url
}
{{end}}
{{end}}

var ActionList = []string{
{{range $i, $c := .Controllers}}{{range .MethodSpecs}}
    "{{$c.StructName}}.{{.Name}}",{{end}}
{{end}}
}
`
