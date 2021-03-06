package controllers

import (
	"github.com/revel/revel"
	"matrix/app/layout"
	"matrix/core"
	"matrix/modules/auth/forms"
	"matrix/modules/auth/models"
)

type Login struct {
	*revel.Controller
}

func (c Login) Index() revel.Result {
	c.RenderArgs["sys_name"] = layout.GetSysName(c.DbSession)
	return c.RenderTemplate("login/login_index.html")
}

func (c Login) Login() revel.Result {
	session := c.DbSession

	c.makePreviousSessionExpire()

	form := new(forms.LoginForm)
	c.Params.Bind(form, "form")

	if form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	user := new(models.UserInfo)
	has, err := session.Where("login_name = ?", form.LoginName).Get(user)
	core.HandleError(err)

	if has == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: "登录名不存在！"})
	}

	if core.CompareHashAndPassword(user.Password, form.Password) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: "密码错误！"})
	}

	//将当前登录用户信息放入Session中
	var loginUser core.LoginUser
	loginUser.UserId = user.Id
	loginUser.LoginName = user.LoginName
	loginUser.NickName = user.NickName

	admin_count, err := session.Where("user_id = ?", user.Id).Count(new(models.AdminInfo))
	core.HandleError(err)
	loginUser.IsAdmin = false
	if admin_count != 0 {
		loginUser.IsAdmin = true
	}

	core.PutLoginUserToSession(c.Session, loginUser)

	return c.RenderJson(core.JsonResult{Success: true, Message: "登陆成功！"})
}

func (c Login) Logout() revel.Result {
	core.RemoveLoginUserInSession(c.Session)
	c.makePreviousSessionExpire()
	return c.RenderJson(core.JsonResult{Success: true, Message: "注销成功！"})
}

func (c Login) makePreviousSessionExpire() {
	delete(c.Session, revel.SESSION_ID_KEY)
}
