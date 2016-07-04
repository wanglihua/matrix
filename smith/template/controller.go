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
    return c.RenderTemplate("{{.entity.ModuleLowerCase}}/{{.entity.EntityLowerCase}}/{{.entity.EntityLowerCase}}_index.html")
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

    {{.entity.EntityLowerCase}}List := make([]models.{{.entity.EntityTitleName}}, 0, limit)
    err := dataQuery.Limit(limit, offset).Find(&{{.entity.EntityLowerCase}}List)
    core.HandleError(err)

    countQuery := *query
    count, err := countQuery.Count(new(models.{{.entity.EntityTitleName}}))
    core.HandleError(err)

    return c.RenderJson(core.GridResult{
        Data:  {{.entity.EntityLowerCase}}List,
        Total: count,
    })
}

type {{.entity.EntityTitleName}}Form struct {
    {{.entity.EntityTitleName}} models.{{.entity.EntityTitleName}}
}

func (f {{.entity.EntityTitleName}}Form) IsCreate() bool {
    return f.{{.entity.EntityTitleName}}.Id == 0
}

func (f {{.entity.EntityTitleName}}Form) Valid(validation *revel.Validation) bool {
    validation.Required(f.{{.entity.EntityTitleName}}.{{.entity.EntityTitleName}}Name).Message("群组名不能为空！")
    validation.MinSize(f.{{.entity.EntityTitleName}}.{{.entity.EntityTitleName}}Name, 3).Message("群组名长度不能小于3！")
    validation.MaxSize(f.{{.entity.EntityTitleName}}.{{.entity.EntityTitleName}}Name, 20).Message("群组名长度不能大于20！")

    return validation.HasErrors() == false
}

func (c {{.entity.ModuleTitleName}}{{.entity.EntityTitleName}}) Detail() revel.Result {
    session := c.DbSession

    {{.entity.EntityLowerCase}}Id := core.GetInt64FromRequest(c.Request, "id")

    {{.entity.EntityLowerCase}} := new(models.{{.entity.EntityTitleName}})
    if {{.entity.EntityLowerCase}}Id != 0 {
        has, err := session.Id({{.entity.EntityLowerCase}}Id).Get({{.entity.EntityLowerCase}})
        core.HandleError(err)
        if has == false {
            panic("指定的{{.entity.EntityChinese}}不存在！")
        }
    }

    form := new({{.entity.EntityTitleName}}Form)
    form.{{.entity.EntityTitleName}} = *{{.entity.EntityLowerCase}}

    c.RenderArgs["form"] = form
    return c.RenderTemplate("{{.entity.ModuleLowerCase}}/{{.entity.EntityLowerCase}}/{{.entity.EntityLowerCase}}_detail.html")
}

func (c {{.entity.ModuleTitleName}}{{.entity.EntityTitleName}}) Save() revel.Result {
    session := c.DbSession

    form := new({{.entity.EntityTitleName}}Form)
    c.Params.Bind(form, "form")

    if form.Valid(c.Validation) == false {
        return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage() })
    }

    {{.entity.EntityLowerCase}} := &form.{{.entity.EntityTitleName}}

    var affected int64
    if form.IsCreate() {
        count, err := session.Where("{{.entity.EntityLowerCase}}_name = ?", {{.entity.EntityLowerCase}}.{{.entity.EntityTitleName}}Name).Count(new(models.{{.entity.EntityTitleName}}))
        core.HandleError(err)
        if count != 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，群组名已存在！" })
        }

        affected, err = session.Insert({{.entity.EntityLowerCase}})
        core.HandleError(err)
    } else {
        count, err := session.Where("id <> ? and {{.entity.EntityLowerCase}}_name = ?", {{.entity.EntityLowerCase}}.Id, {{.entity.EntityLowerCase}}.{{.entity.EntityTitleName}}Name).Count(new(models.{{.entity.EntityTitleName}}))
        core.HandleError(err)
        if count != 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，群组名已存在！" })
        }

        affected, err = session.Id({{.entity.EntityLowerCase}}.Id).Update({{.entity.EntityLowerCase}})
        core.HandleError(err)

        if affected == 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
        }
    }

    return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c {{.entity.ModuleTitleName}}{{.entity.EntityTitleName}}) Delete() revel.Result {
    session := c.DbSession

    userIdList := make([]int64, 0)
    c.Params.Bind(&userIdList, "id_list")

    {{.entity.EntityLowerCase}} := new(models.{{.entity.EntityTitleName}})
    affected, err := session.In("id", userIdList).Delete({{.entity.EntityLowerCase}})
    core.HandleError(err)

    return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}

`