package controllers

import (
    "github.com/revel/revel"
    "matrix/service"
    "strings"
    "matrix/modules/auth/models"
)

type AuthProfile struct {
    *revel.Controller
    service.BaseController
}

func (c AuthProfile) Index() revel.Result {
    session := c.DbSession

    loginUserId := c.User.UserId
    loginUser := new(models.User)
    _, err := session.Id(loginUserId).Get(loginUser)
    service.HandleError(err)

    c.RenderArgs["user"] = loginUser

    return c.RenderTemplate("auth/profile/profile_index.html")
}

func (c AuthProfile) Save() revel.Result {
    session := c.DbSession
    var password, passwordAgain string
    c.Params.Bind(&password, "Password")
    c.Params.Bind(&passwordAgain, "PasswordAgain")

    if strings.Trim(password, " ") != "" || strings.Trim(passwordAgain, " ") != "" {
        c.Validation.Required(password).Message("密码不能为空！")
        c.Validation.MinSize(password, 6).Message("密码长度不能小于6！")
        c.Validation.MaxSize(password, 12).Message("密码长度不能大于12！")

        if password != passwordAgain {
            c.Validation.Errors = append(c.Validation.Errors, &revel.ValidationError{
                Key:"password_again_not_match",
                Message:"两次输入的密码不一致！",
            })
        }
    }

    if c.Validation.HasErrors() {
        return c.RenderJson(service.JsonResult{Success: false, Message: c.GetValidationErrorMessage() })
    }

    loginUserId := c.User.UserId
    loginUser := new(models.User)
    _, err := session.Id(loginUserId).Get(loginUser)
    service.HandleError(err)

    loginUser.Password = service.EncryptPassword(password)

    _, err = session.Id(loginUserId).Cols("password").Update(loginUser)
    service.HandleError(err)

    return c.RenderJson(service.JsonResult{Success: true, Message: "保存成功!"})
}
