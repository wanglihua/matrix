package controllers

import (
    "github.com/revel/revel"

    "matrix/modules/auth/models"
    "matrix/service"
    "strconv"
    "strings"
)

type AuthUser struct {
    *revel.Controller
    service.BaseController
}

func (c AuthUser) Index() revel.Result {
    return c.RenderTemplate("user/user_index.html")
}

func (c AuthUser) ListData() revel.Result {
    session := c.DbSession

    filter, order, offset, limit := service.GetGridRequestParam(c.Request)
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
    service.HandleError(err)

    countQuery := *query
    count, err := countQuery.Count(new(models.User))
    service.HandleError(err)

    return c.RenderJson(service.GridResult{
        Data:  userList,
        Total: count,
    })
}

func (c AuthUser) Detail() revel.Result {
    session := c.DbSession

    userId := service.GetInt64FromRequest(c.Request, "id")

    user := new(models.User)
    if userId != 0 {
        has, err := session.Id(userId).Get(user)
        service.HandleError(err)
        if has == false {
            panic("指定的用户不存在！")
        }
    }

    c.RenderArgs["is_create"] = userId == 0
    c.RenderArgs["user"] = user
    return c.RenderTemplate("user/user_detail.html")
}

func (c AuthUser) Save() revel.Result {
    session := c.DbSession

    user := new(models.User)
    c.Params.Bind(&user, "user")

    isCreate := user.Id == 0

    c.Validation.Required(user.LoginName).Message("登录名不能为空！")
    c.Validation.MinSize(user.LoginName, 3).Message("登录名长度不能小于3！")
    c.Validation.MaxSize(user.LoginName, 10).Message("登录名长度不能大于10！")

    c.Validation.Required(user.NickName).Message("用户名不能为空！")
    c.Validation.MinSize(user.NickName, 3).Message("用户名长度不能小于3！")
    c.Validation.MaxSize(user.NickName, 10).Message("用户名长度不能大于10！")

    if isCreate {
        var passwordAgain string
        c.Params.Bind(&passwordAgain, "PasswordAgain")

        c.Validation.Required(user.Password).Message("密码不能为空！")
        c.Validation.MinSize(user.Password, 6).Message("密码长度不能小于6！")
        c.Validation.MaxSize(user.Password, 12).Message("密码长度不能大于12！")

        if user.Password != passwordAgain {
            c.Validation.Errors = append(c.Validation.Errors, &revel.ValidationError{
                Key:"password_again_not_match",
                Message:"两次输入的密码不一致！",
            })
        }

        user.Password = service.EncryptPassword(user.Password)

    } else {
        var newPassword, newPasswordAgain string
        c.Params.Bind(&newPassword, "NewPassword")
        c.Params.Bind(&newPasswordAgain, "NewPasswordAgain")

        if newPassword != "" || newPasswordAgain != "" {
            c.Validation.MinSize(newPassword, 6).Message("密码长度不能小于6！")
            c.Validation.MaxSize(newPassword, 12).Message("密码长度不能大于12！")

            if strings.Trim(newPassword, " ") != "" || strings.Trim(newPasswordAgain, " ") != "" {
                c.Validation.Errors = append(c.Validation.Errors, &revel.ValidationError{
                    Key:"password_again_not_match",
                    Message:"两次输入的密码不一致！",
                })
            }

            user.Password = service.EncryptPassword(newPassword)
        }
    }

    if c.Validation.HasErrors() {
        return c.RenderJson(service.JsonResult{Success: false, Message: c.GetValidationErrorMessage() })
    }

    var affected int64
    if isCreate {
        count, err := session.Where("login_name = ?", user.LoginName).Count(new(models.User))
        service.HandleError(err)
        if count != 0 {
            return c.RenderJson(service.JsonResult{Success: false, Message: "保存失败，登录名已存在！" })
        }

        affected, err = session.Insert(user)
        service.HandleError(err)
    } else {
        count, err := session.Where("id <> ? and login_name = ?", user.Id, user.LoginName).Count(new(models.User))
        service.HandleError(err)
        if count != 0 {
            return c.RenderJson(service.JsonResult{Success: false, Message: "保存失败，登录名已存在！" })
        }

        affected, err = session.Id(user.Id).Update(user)
        //affected, err := session.Id(user.Id).Cols("nick_name").Update(user)
        //affected, err := session.Table(new(User)).Id(user.Id).Update(map[string]interface{}{"password":"123456"})
        service.HandleError(err)

        if affected == 0 {
            return c.RenderJson(service.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
        }
    }

    return c.RenderJson(service.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c AuthUser) Delete() revel.Result {
    session := c.DbSession

    userIdList := make([]int64, 0)
    c.Params.Bind(&userIdList, "id_list")

    user := new(models.User)
    affected, err := session.In("id", userIdList).Delete(user)
    service.HandleError(err)

    return c.RenderJson(service.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}
