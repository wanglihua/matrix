package controllers

import (
	"github.com/revel/revel"

	"matrix/core"
	authModels "matrix/modules/auth/models"
	"matrix/modules/oa/models"
)

type OaWorklogAdmin struct {
	*revel.Controller
}

func (c OaWorklogAdmin) Index() revel.Result {
	return c.RenderTemplate("oa/worklog_admin/worklog_admin_index.html")
}

type WorklogAdminView struct {
	models.WorklogInfo  `xorm:"extends" json:"l"`
	authModels.UserInfo `xorm:"extends" json:"u"`
}

func (c OaWorklogAdmin) ListData() revel.Result {
	session := c.DbSession

	filter, order, offset, limit := core.GetGridRequestParam(c.Request)
	query := session.
		Select("l.*, u.*").
		Table(authModels.TablePrefix+"user").Alias("u").
		Join("inner", []string{models.TablePrefix + "worklog", "l"}, "u.id = l.user_id").
		Where(filter)

	//query extra filter here

	dataQuery := *query
	if order != "" {
		dataQuery = *dataQuery.OrderBy(order)
	} else {
		dataQuery = *dataQuery.Asc("l.id")
	}

	worklogAdminViewList := make([]WorklogAdminView, 0, limit)
	err := dataQuery.Limit(limit, offset).Find(&worklogAdminViewList)
	core.HandleError(err)

	countQuery := *query
	count, err := countQuery.Count(new(models.WorklogInfo))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  worklogAdminViewList,
		Total: count,
	})
}

type WorklogAdminForm struct {
	Worklog models.WorklogInfo
}

func (f *WorklogAdminForm) IsCreate() bool {
	return f.Worklog.Id == 0
}

func (f *WorklogAdminForm) Valid(validation *revel.Validation) bool {
	validation.Required(f.Worklog.UserId).Message("用户不能为空！")

	validation.Required(f.Worklog.Title).Message("日志主题不能为空！")
	if f.Worklog.Title != "" {
		validation.MinSize(f.Worklog.Title, 3).Message("日志主题长度不能小于3！")
	}
	if f.Worklog.Title != "" {
		validation.MaxSize(f.Worklog.Title, 200).Message("日志主题长度不能大于200！")
	}

	validation.Required(f.Worklog.Content).Message("日志内容不能为空！")
	if f.Worklog.Content != "" {
		validation.MinSize(f.Worklog.Content, 3).Message("日志内容长度不能小于3！")
	}
	if f.Worklog.Content != "" {
		validation.MaxSize(f.Worklog.Content, 1000).Message("日志内容长度不能大于1000！")
	}

	if f.Worklog.BeginTime != "" {
		validation.MaxSize(f.Worklog.BeginTime, 15).Message("开始时间长度不能大于15！")
	}

	if f.Worklog.EndTime != "" {
		validation.MaxSize(f.Worklog.EndTime, 15).Message("结束时间长度不能大于15！")
	}

	if f.Worklog.Remark != "" {
		validation.MaxSize(f.Worklog.Remark, 1000).Message("备注长度不能大于1000！")
	}

	return validation.HasErrors() == false
}

func (c OaWorklogAdmin) Detail() revel.Result {
	session := c.DbSession

	var worklogId int64
	c.Params.Bind(&worklogId, "id")

	worklog := new(models.WorklogInfo)
	has, err := session.Id(worklogId).Get(worklog)
	core.HandleError(err)
	if has == false {
		panic("指定的工作日志不存在！")
	}

	form := new(WorklogAdminForm)
	form.Worklog = *worklog

	c.UnbindToRenderArgs(form, "form")

	user := new(authModels.UserInfo)
	_, err = session.Id(worklog.UserId).Get(user)
	core.HandleError(err)

	c.RenderArgs["user_nick_name"] = user.NickName

	return c.RenderTemplate("oa/worklog_admin/worklog_admin_detail.html")
}
