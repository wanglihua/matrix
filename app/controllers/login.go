package controllers

import (
    "github.com/revel/revel"
    "matrix/core"
    "matrix/modules/auth/models"
    "matrix/modules/auth/forms"
    "matrix/app/layout"
)

type Login struct {
    *revel.Controller
    core.BaseController
}

func (c Login) Index() revel.Result {
    c.RenderArgs["sys_name"] = layout.GetSysName()
    return c.RenderTemplate("login/login_index.html")
}

func (c Login) Login() revel.Result {
    session := c.DbSession

    c.makePreviousSessionExpire()

    form := new(forms.LoginForm)
    c.Params.Bind(form, "form")

    if form.Valid(c.Validation) == false {
        return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage() })
    }

    user := new(models.User)
    has, err := session.Where("login_name = ?", form.LoginName).Get(user)
    core.HandleError(err)

    if has == false {
        return c.RenderJson(core.JsonResult{Success: false, Message: "登录名不存在！" })
    }

    if core.CompareHashAndPassword(user.Password, form.Password) == false {
        return c.RenderJson(core.JsonResult{Success: false, Message: "密码错误！" })
    }

    //将当前登录用户信息放入Session中
    var loginUser core.LoginUser
    loginUser.UserId = user.Id
    loginUser.LoginName = user.LoginName
    loginUser.NickName = user.NickName

    core.PutLoginUserToSession(c.Session, loginUser)

    return c.RenderJson(core.JsonResult{Success:true, Message:"登陆成功！"})
}

func (c Login) Logout() revel.Result {
    core.RemoveLoginUserInSession(c.Session)

    c.makePreviousSessionExpire()

    return c.RenderJson(core.JsonResult{Success:true, Message:"注销成功！"})
}

func (c Login) makePreviousSessionExpire() {
    delete(c.Session, revel.SESSION_ID_KEY)

    //sessionCookie := c.Session.Cookie()
    //sessionCookie.Expires = time.Time{}
    //http.SetCookie(c.Response.Out, sessionCookie)
}
