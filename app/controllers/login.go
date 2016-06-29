package controllers

import (
    "github.com/revel/revel"
    "matrix/core"
    "matrix/modules/auth/models"
    //"time"
    //"net/http"
)

type Login struct {
    *revel.Controller
    core.BaseController
}

func (c Login) Index() revel.Result {
    return c.RenderTemplate("login/login_index.html")
}

type LoginForm struct {
    LoginName string
    Password  string
}

func (c Login) Login() revel.Result {
    session := c.DbSession

    c.makePreviousSessionExpire()

    loginForm := new(LoginForm)
    c.Params.Bind(&loginForm, "login")

    c.Validation.Required(loginForm.LoginName).Message("登录名不能为空！")
    c.Validation.MinSize(loginForm.LoginName, 3).Message("登录名长度不能小于3！")
    c.Validation.MaxSize(loginForm.LoginName, 10).Message("登录名长度不能大于10！")

    c.Validation.Required(loginForm.Password).Message("密码不能为空！")
    c.Validation.MinSize(loginForm.Password, 6).Message("密码长度不能小于6！")
    c.Validation.MaxSize(loginForm.Password, 12).Message("密码长度不能大于12！")

    if c.Validation.HasErrors() {
        return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage() })
    }

    user := new(models.User)
    has, err := session.Where("login_name = ?", loginForm.LoginName).Get(user)
    core.HandleError(err)

    if has == false {
        return c.RenderJson(core.JsonResult{Success: false, Message: "登录名不存在！" })
    }

    if core.CompareHashAndPassword(user.Password, loginForm.Password) == false {
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
