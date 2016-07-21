package controllers

import (
    "github.com/revel/revel"
    "matrix/core"
    "matrix/modules/auth/models"
    "matrix/modules/auth/forms"
)

type AuthProfile struct {
    *revel.Controller
    core.BaseController
}

func (c AuthProfile) Index() revel.Result {
    session := c.DbSession

    loginUserId := c.User.UserId
    loginUser := new(models.User)
    _, err := session.Id(loginUserId).Get(loginUser)
    core.HandleError(err)

    c.UnbindToRenderArgs(loginUser, "user")

    return c.RenderTemplate("auth/profile/profile_index.html")
}

func (c AuthProfile) Save() revel.Result {
    session := c.DbSession

    form := new(forms.ProfileForm)
    c.Params.Bind(form, "form")

    if form.Valid(c.Validation) == false {
        return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage() })
    }

    loginUserId := c.User.UserId
    loginUser := new(models.User)
    _, err := session.Id(loginUserId).Get(loginUser)
    core.HandleError(err)

    loginUser.Password = core.EncryptPassword(form.Password)

    _, err = session.Id(loginUserId).Cols("password").Update(loginUser)
    core.HandleError(err)

    return c.RenderJson(core.JsonResult{Success: true, Message: "保存成功!"})
}
