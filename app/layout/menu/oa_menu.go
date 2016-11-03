package menu

import (
	"bytes"
	"github.com/go-xorm/xorm"
	"github.com/revel/revel"
)

func GetOaMenu(db_session *xorm.Session, web_session revel.Session) string {
	if _, found := revel.ModuleByName("oa"); found == false {
		return ""
	}

	var buffer bytes.Buffer

	buffer.WriteString(`
<li id="oa-menu" class="">
    <a href="#" class="dropdown-toggle">
        <i class="menu-icon fa fa-files-o"></i>
        <span class="menu-text">办公系统</span>
        <b class="arrow fa fa-angle-down"></b>
    </a>
    <b class="arrow"></b>
    <ul class="submenu">

        <li id="oa-worklog-menu" class="">
            <a href="/oa/worklog">
                <i class="menu-icon fa fa-caret-right"></i>
                工作日志
            </a>
            <b class="arrow"></b>
        </li>
        <li id="oa-worklogadmin-menu" class="">
            <a href="/oa/worklog/admin">
                <i class="menu-icon fa fa-caret-right"></i>
                日志管理
            </a>
            <b class="arrow"></b>
        </li>

    </ul>
</li>
`)

	return buffer.String()
}
