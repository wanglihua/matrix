package template

var ControllerTemplate = `
package controllers

import (
	"encoding/json"
	"strconv"
	"github.com/revel/revel"

	"{{.import_path}}/core"
	"{{.import_path}}/modules/{{.entity.ModuleLowerCase}}/models"
)

type {{.entity.ModuleTitleName}}{{.entity.EntityTitleName}} struct {
	*revel.Controller
}

func (c {{.entity.ModuleTitleName}}{{.entity.EntityTitleName}}) Index() revel.Result {
	return c.RenderTemplate("{{.entity.ModuleLowerCase}}/{{Underscore .entity.EntityCamelCase}}/{{Underscore .entity.EntityCamelCase}}_index.html")
}

func (c {{.entity.ModuleTitleName}}{{.entity.EntityTitleName}}) ListData() revel.Result {
	db_session := c.DbSession

	filter, order, offset, limit := core.GetGridRequestParam(c.Request)
	query := db_session.Where(filter)

	//query extra filter here

	data_query := *query
	if order != "" {
		data_query = *data_query.OrderBy(order)
	} else {
		data_query = *data_query.Asc("id")
	}

	{{Underscore .entity.EntityCamelCase}}_list := make([]models.{{.entity.EntityTitleName}}, 0, limit)
	err := data_query.Limit(limit, offset).Find(&{{Underscore .entity.EntityCamelCase}}_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(models.{{.entity.EntityTitleName}}))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  {{Underscore .entity.EntityCamelCase}}_list,
		Total: count,
	})
}

type {{.entity.EntityTitleName}}DetailForm struct {
	{{.entity.EntityTitleName}} models.{{.entity.EntityTitleName}} {{.entity.EntityJsonTag}}
}

func (f *{{.entity.EntityTitleName}}DetailForm) IsCreate() bool {
	return f.{{.entity.EntityTitleName}}.Id == 0
}

func (f *{{.entity.EntityTitleName}}DetailForm) Valid(validation *revel.Validation) bool { {{range $fieldIndex, $field := .entity.FieldList}}
{{FieldValid $.entity $field}}{{end}}
	return validation.HasErrors() == false
}

func (c {{.entity.ModuleTitleName}}{{.entity.EntityTitleName}}) DetailData() revel.Result {
	db_session := c.DbSession

	var {{Underscore .entity.EntityCamelCase}}_id int64
	c.Params.Bind(&{{Underscore .entity.EntityCamelCase}}_id, "id")

	var {{Underscore .entity.EntityCamelCase}} models.{{.entity.EntityTitleName}}
	if {{Underscore .entity.EntityCamelCase}}_id != 0 {
		has, err := db_session.Id({{Underscore .entity.EntityCamelCase}}_id).Get(&{{Underscore .entity.EntityCamelCase}})
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的{{.entity.EntityChinese}}不存在！"})
		}
	}

	var detail_form {{.entity.EntityTitleName}}DetailForm
	detail_form.{{.entity.EntityTitleName}} = {{Underscore .entity.EntityCamelCase}}

	return c.RenderJson(detail_form)
}

func (c {{.entity.ModuleTitleName}}{{.entity.EntityTitleName}}) Save() revel.Result {
	db_session := c.DbSession

	var detail_form {{.entity.EntityTitleName}}DetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	{{Underscore .entity.EntityCamelCase}} := detail_form.{{.entity.EntityTitleName}}

	var affected int64
	if detail_form.IsCreate() { {{range $fieldIndex, $field := .entity.FieldList}}{{if ne $field.Unique "false"}}
{{CheckUniqueCreate $.entity $field}}{{end}}{{end}}
		affected, err = db_session.Insert(&{{Underscore .entity.EntityCamelCase}})
		core.HandleError(err)
	} else { {{range $fieldIndex, $field := .entity.FieldList}}{{if ne $field.Unique "false"}}
{{CheckUniqueUpdate $.entity $field}}{{end}}{{end}}
		affected, err = db_session.Id({{Underscore .entity.EntityCamelCase}}.Id).Update(&{{Underscore .entity.EntityCamelCase}})
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c {{.entity.ModuleTitleName}}{{.entity.EntityTitleName}}) Delete() revel.Result {
	db_session := c.DbSession

	{{Underscore .entity.EntityCamelCase}}_id_list := make([]int64, 0)
	c.Params.Bind(&{{Underscore .entity.EntityCamelCase}}_id_list, "id_list")

	affected, err := db_session.In("id", {{Underscore .entity.EntityCamelCase}}_id_list).Delete(new(models.{{.entity.EntityTitleName}}))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}

`
