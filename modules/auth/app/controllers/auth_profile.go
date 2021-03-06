package controllers

import (
	"github.com/revel/revel"
	"matrix/core"
	"matrix/modules/auth/forms"
	"matrix/modules/auth/models"
	"strings"
)

type AuthProfile struct {
	*revel.Controller
}

func (c AuthProfile) Index() revel.Result {
	session := c.DbSession

	loginUserId := core.GetLoginUser(c.Session).UserId
	loginUser := new(models.UserInfo)
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
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	loginUserId := core.GetLoginUser(c.Session).UserId
	loginUser := new(models.UserInfo)
	_, err := session.Id(loginUserId).Get(loginUser)
	core.HandleError(err)

	if strings.TrimSpace(form.Password) != "" {
		loginUser.Password = core.EncryptPassword(form.Password)

		_, err = session.Id(loginUserId).Cols("password").Update(loginUser)
		core.HandleError(err)
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: "保存成功!"})
}
