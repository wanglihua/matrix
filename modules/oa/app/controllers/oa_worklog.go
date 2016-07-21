package controllers

import (
    "strconv"
    "github.com/revel/revel"

    "matrix/core"
    "matrix/modules/oa/models"
)

type OaWorklog struct {
    *revel.Controller
    core.BaseController
}

func (c OaWorklog) Index() revel.Result {
    return c.RenderTemplate("oa/worklog/worklog_index.html")
}

func (c OaWorklog) ListData() revel.Result {
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

    worklogList := make([]models.Worklog, 0, limit)
    err := dataQuery.Limit(limit, offset).Find(&worklogList)
    core.HandleError(err)

    countQuery := *query
    count, err := countQuery.Count(new(models.Worklog))
    core.HandleError(err)

    return c.RenderJson(core.GridResult{
        Data:  worklogList,
        Total: count,
    })
}

type WorklogForm struct {
    Worklog models.Worklog
}

func (f *WorklogForm) IsCreate() bool {
    return f.Worklog.Id == 0
}

func (f *WorklogForm) Valid(validation *revel.Validation) bool {
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

func (c OaWorklog) Detail() revel.Result {
    session := c.DbSession

    var worklogId int64
    c.Params.Bind(&worklogId, "id")

    worklog := new(models.Worklog)
    if worklogId != 0 {
        has, err := session.Id(worklogId).Get(worklog)
        core.HandleError(err)
        if has == false {
            panic("指定的工作日志不存在！")
        }
    }

    form := new(WorklogForm)
    form.Worklog = *worklog

    c.UnbindToRenderArgs(form, "form")

    return c.RenderTemplate("oa/worklog/worklog_detail.html")
}

func (c OaWorklog) Save() revel.Result {
    session := c.DbSession

    form := new(WorklogForm)
    c.Params.Bind(form, "form")

    if form.Valid(c.Validation) == false {
        return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage() })
    }

    worklog := &form.Worklog

    var affected int64
    var err error
    if form.IsCreate() {
        affected, err = session.Insert(worklog)
        core.HandleError(err)
    } else {
        affected, err = session.Id(worklog.Id).Update(worklog)
        core.HandleError(err)

        if affected == 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
        }
    }

    return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c OaWorklog) Delete() revel.Result {
    session := c.DbSession

    worklogIdList := make([]int64, 0)
    c.Params.Bind(&worklogIdList, "id_list")

    worklog := new(models.Worklog)
    affected, err := session.In("id", worklogIdList).Delete(worklog)
    core.HandleError(err)

    return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}

