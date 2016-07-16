package template

var ControllerTemplate = `
package controllers

import (
    "strconv"
    "github.com/revel/revel"

    "matrix/core"
    "matrix/modules/{{.entity.ModuleLowerCase}}/models"
)

type {{.entity.ModuleTitleName}}{{.entity.EntityTitleName}} struct {
    *revel.Controller
    core.BaseController
}

func (c {{.entity.ModuleTitleName}}{{.entity.EntityTitleName}}) Index() revel.Result {
    return c.RenderTemplate("{{.entity.ModuleLowerCase}}/{{LowerCase .entity.EntityCamelCase}}/{{LowerCase .entity.EntityCamelCase}}_index.html")
}

func (c {{.entity.ModuleTitleName}}{{.entity.EntityTitleName}}) ListData() revel.Result {
    session := c.DbSession

    filter, order, offset, limit := core.GetGridRequestParam(c.Request)
    query := session.Where(filter)

    //query extra filter here

    dataQuery := *query
    if order != "" {
        dataQuery = *dataQuery.OrderBy(order)
    } else {
        dataQuery = *dataQuery.Asc("id")
    }

    {{.entity.EntityCamelCase}}List := make([]models.{{.entity.EntityTitleName}}, 0, limit)
    err := dataQuery.Limit(limit, offset).Find(&{{.entity.EntityCamelCase}}List)
    core.HandleError(err)

    countQuery := *query
    count, err := countQuery.Count(new(models.{{.entity.EntityTitleName}}))
    core.HandleError(err)

    return c.RenderJson(core.GridResult{
        Data:  {{.entity.EntityCamelCase}}List,
        Total: count,
    })
}

type {{.entity.EntityTitleName}}Form struct {
    {{.entity.EntityTitleName}} models.{{.entity.EntityTitleName}}
}

func (f *{{.entity.EntityTitleName}}Form) IsCreate() bool {
    return f.{{.entity.EntityTitleName}}.Id == 0
}

func (f *{{.entity.EntityTitleName}}Form) Valid(validation *revel.Validation) bool { {{range $fieldIndex, $field := .entity.FieldList}}
{{FieldValid $.entity $field}}{{end}}
    return validation.HasErrors() == false
}

func (c {{.entity.ModuleTitleName}}{{.entity.EntityTitleName}}) Detail() revel.Result {
    session := c.DbSession

    {{.entity.EntityCamelCase}}Id := core.GetInt64FromRequest(c.Request, "id")

    {{.entity.EntityCamelCase}} := new(models.{{.entity.EntityTitleName}})
    if {{.entity.EntityCamelCase}}Id != 0 {
        has, err := session.Id({{.entity.EntityCamelCase}}Id).Get({{.entity.EntityCamelCase}})
        core.HandleError(err)
        if has == false {
            panic("指定的{{.entity.EntityChinese}}不存在！")
        }
    }

    form := new({{.entity.EntityTitleName}}Form)
    form.{{.entity.EntityTitleName}} = *{{.entity.EntityCamelCase}}

    c.RenderArgs["form"] = form
    return c.RenderTemplate("{{.entity.ModuleLowerCase}}/{{LowerCase .entity.EntityCamelCase}}/{{LowerCase .entity.EntityCamelCase}}_detail.html")
}

func (c {{.entity.ModuleTitleName}}{{.entity.EntityTitleName}}) Save() revel.Result {
    session := c.DbSession

    form := new({{.entity.EntityTitleName}}Form)
    c.Params.Bind(form, "form")

    if form.Valid(c.Validation) == false {
        return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage() })
    }

    {{.entity.EntityCamelCase}} := &form.{{.entity.EntityTitleName}}

    var affected int64
    if form.IsCreate() { {{range $fieldIndex, $field := .entity.FieldList}}{{if ne $field.Unique "false"}}
{{CheckUniqueCreate $.entity $field}}{{end}}{{end}}
        affected, err = session.Insert({{.entity.EntityCamelCase}})
        core.HandleError(err)
    } else { {{range $fieldIndex, $field := .entity.FieldList}}{{if ne $field.Unique "false"}}
{{CheckUniqueUpdate $.entity $field}}{{end}}{{end}}
        affected, err = session.Id({{.entity.EntityCamelCase}}.Id).Update({{.entity.EntityCamelCase}})
        core.HandleError(err)

        if affected == 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
        }
    }

    return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c {{.entity.ModuleTitleName}}{{.entity.EntityTitleName}}) Delete() revel.Result {
    session := c.DbSession

    {{.entity.EntityCamelCase}}IdList := make([]int64, 0)
    c.Params.Bind(&{{.entity.EntityCamelCase}}IdList, "id_list")

    {{.entity.EntityCamelCase}} := new(models.{{.entity.EntityTitleName}})
    affected, err := session.In("id", {{.entity.EntityCamelCase}}IdList).Delete({{.entity.EntityCamelCase}})
    core.HandleError(err)

    return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}

`
