package controllers

import (
    "github.com/revel/revel"

    "matrix/modules/auth/models"
    "matrix/core"
    "strconv"
)

type AuthGroup struct {
    *revel.Controller
    core.BaseController
}

func (c AuthGroup) Index() revel.Result {
    return c.RenderTemplate("auth/group/group_index.html")
}

func (c AuthGroup) ListData() revel.Result {
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

    groupList := make([]models.Group, 0, limit)
    err := dataQuery.Limit(limit, offset).Find(&groupList)
    core.HandleError(err)

    countQuery := *query
    count, err := countQuery.Count(new(models.Group))
    core.HandleError(err)

    return c.RenderJson(core.GridResult{
        Data:  groupList,
        Total: count,
    })
}

type GroupForm struct {
    Group models.Group
}

func (f *GroupForm) IsCreate() bool {
    return f.Group.Id == 0
}

func (f *GroupForm) Valid(validation *revel.Validation) bool {
    validation.Required(f.Group.GroupName).Message("群组名不能为空！")
    validation.MinSize(f.Group.GroupName, 3).Message("群组名长度不能小于3！")
    validation.MaxSize(f.Group.GroupName, 20).Message("群组名长度不能大于20！")

    return validation.HasErrors() == false
}

func (c AuthGroup) Detail() revel.Result {
    session := c.DbSession

    var groupId int64
    c.Params.Bind(&groupId, "id")

    group := new(models.Group)
    if groupId != 0 {
        has, err := session.Id(groupId).Get(group)
        core.HandleError(err)
        if has == false {
            panic("指定的群组不存在！")
        }
    }

    form := new(GroupForm)
    form.Group = *group

    c.RenderArgs["form"] = form
    return c.RenderTemplate("auth/group/group_detail.html")
}

func (c AuthGroup) Save() revel.Result {
    session := c.DbSession

    form := new(GroupForm)
    c.Params.Bind(form, "form")

    if form.Valid(c.Validation) == false {
        return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage() })
    }

    group := &form.Group

    var affected int64
    if form.IsCreate() {
        count, err := session.Where("group_name = ?", group.GroupName).Count(new(models.Group))
        core.HandleError(err)
        if count != 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，群组名已存在！" })
        }

        affected, err = session.Insert(group)
        core.HandleError(err)
    } else {
        count, err := session.Where("id <> ? and group_name = ?", group.Id, group.GroupName).Count(new(models.Group))
        core.HandleError(err)
        if count != 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，群组名已存在！" })
        }

        affected, err = session.Id(group.Id).Update(group)
        core.HandleError(err)

        if affected == 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
        }
    }

    return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c AuthGroup) Delete() revel.Result {
    session := c.DbSession

    groupIdList := make([]int64, 0)
    c.Params.Bind(&groupIdList, "id_list")

    group := new(models.Group)
    affected, err := session.In("id", groupIdList).Delete(group)
    core.HandleError(err)

    return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}
