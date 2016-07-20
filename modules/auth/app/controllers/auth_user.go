package controllers

import (
    "github.com/revel/revel"

    "matrix/modules/auth/models"
    "matrix/modules/auth/forms"
    "matrix/core"
    "strconv"
)

type AuthUser struct {
    *revel.Controller
    core.BaseController
}

func (c AuthUser) Index() revel.Result {
    return c.RenderTemplate("auth/user/user_index.html")
}

func (c AuthUser) ListData() revel.Result {
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

    userList := make([]models.User, 0, limit)
    err := dataQuery.Limit(limit, offset).Find(&userList)
    core.HandleError(err)

    countQuery := *query
    count, err := countQuery.Count(new(models.User))
    core.HandleError(err)

    return c.RenderJson(core.GridResult{
        Data:  userList,
        Total: count,
    })
}

func (c AuthUser) Detail() revel.Result {
    session := c.DbSession

    var userId int64
    c.Params.Bind(&userId, "id")

    user := new(models.User)
    if userId != 0 {
        has, err := session.Id(userId).Get(user)
        core.HandleError(err)
        if has == false {
            panic("指定的用户不存在！")
        }
    }

    form := new(forms.UserDetailForm)
    form.User = *user

    c.RenderArgs["form"] = form
    return c.RenderTemplate("auth/user/user_detail.html")
}

func (c AuthUser) Save() revel.Result {
    session := c.DbSession

    form := new(forms.UserDetailForm)
    c.Params.Bind(form, "form")

    if form.Valid(c.Validation) == false {
        return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage() })
    }

    user := &form.User

    var affected int64
    if form.IsCreate() {
        count, err := session.Where("login_name = ?", user.LoginName).Count(new(models.User))
        core.HandleError(err)
        if count != 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，登录名已存在！" })
        }

        user.Password = core.EncryptPassword(user.Password)

        affected, err = session.Insert(user)
        core.HandleError(err)
    } else {
        count, err := session.Where("id <> ? and login_name = ?", user.Id, user.LoginName).Count(new(models.User))
        core.HandleError(err)
        if count != 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "保存失败，登录名已存在！" })
        }

        user.Password = core.EncryptPassword(form.NewPassword)

        affected, err = session.Id(user.Id).Update(user)
        //affected, err := session.Id(form.Id).Cols("nick_name").Update(user)
        //affected, err := session.Table(new(User)).Id(form.Id).Update(map[string]interface{}{"password":"123456"})
        core.HandleError(err)

        if affected == 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
        }
    }

    return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c AuthUser) Delete() revel.Result {
    session := c.DbSession

    userIdList := make([]int64, 0)
    c.Params.Bind(&userIdList, "id_list")

    count, err := session.In("user_id", userIdList).Count(new(models.Admin))
    core.HandleError(err)
    if count != 0 {
        return c.RenderJson(core.JsonResult{Success: false, Message: "不能删除管理员！" })
    }

    user := new(models.User)
    affected, err := session.In("id", userIdList).Delete(user)
    core.HandleError(err)

    return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}
