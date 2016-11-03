package menu

import (
	"bytes"
	"github.com/go-xorm/xorm"
	"github.com/revel/revel"
	"matrix/core"
)

func GetAuthMenu(db_session *xorm.Session, web_session revel.Session) string {
	if _, found := revel.ModuleByName("oa"); found == false {
		return ""
	}

	login_user := core.GetLoginUser(web_session)
	is_admin := false
	if login_user != nil {
		is_admin = login_user.IsAdmin
	}

	var buffer bytes.Buffer

	if is_admin {
		buffer.WriteString(`
<li id="auth-menu" class="">
    <a href="#" class="dropdown-toggle">
        <i class="menu-icon fa fa-users"></i>
        <span class="menu-text">权限管理</span>
        <b class="arrow fa fa-angle-down"></b>
    </a>
    <b class="arrow"></b>
    <ul class="submenu">
        <li id="auth-user-menu" class="">
            <a href="/auth/user">
                <i class="menu-icon fa fa-caret-right"></i>
                用户管理
            </a>
            <b class="arrow"></b>
        </li>
        <li id="auth-admin-menu" class="">
            <a href="/auth/admin">
                <i class="menu-icon fa fa-caret-right"></i>
                管理员组
            </a>
            <b class="arrow"></b>
        </li>
        <li id="auth-group-menu" class="">
            <a href="/auth/group">
                <i class="menu-icon fa fa-caret-right"></i>
                群组管理
            </a>
            <b class="arrow"></b>
        </li>
    </ul>
</li>
<li id="settings-menu" class="">
    <a href="#" class="dropdown-toggle">
        <i class="menu-icon fa fa-cog"></i>
        <span class="menu-text">系统配置</span>
        <b class="arrow fa fa-angle-down"></b>
    </a>
    <b class="arrow"></b>
    <ul class="submenu">
        <li id="settings-config-menu" class="">
            <a href="/config">
                <i class="menu-icon fa fa-caret-right"></i>
                参数配置
            </a>
            <b class="arrow"></b>
        </li>
    </ul>
</li>
`)
	}
	return buffer.String()
}
