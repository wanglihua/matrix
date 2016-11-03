package menu

import (
	"bytes"
	"github.com/go-xorm/xorm"
	"github.com/revel/revel"
	"matrix/core"
)

func GetSettingMenu(db_session *xorm.Session, web_session revel.Session) string {

	//if _, found := revel.ModuleByName("help"); found == false {
	//	return ""
	//}

	login_user := core.GetLoginUser(web_session)
	is_admin := false
	if login_user != nil {
		is_admin = login_user.IsAdmin
	}

	var buffer bytes.Buffer

	if is_admin {
		buffer.WriteString(`
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
