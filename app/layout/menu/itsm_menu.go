package menu

import (
	"bytes"
	"github.com/go-xorm/xorm"
	"github.com/revel/revel"
)

func GetItsmMenu(db_session *xorm.Session, web_session revel.Session) string {
	if _, found := revel.ModuleByName("itsm"); found == false {
		return ""
	}

	var buffer bytes.Buffer

	buffer.WriteString(`
<li id="itsm-menu" class="">
    <a href="#" class="dropdown-toggle">
    <i class="menu-icon fa fa-table"></i>
    <span class="menu-text">ITSM</span>
    <b class="arrow fa fa-angle-down"></b>
</a>
<b class="arrow"></b>
<ul class="submenu"p>

    <li id="itsm-eventtype-menu" class="">
        <a href="/itsm/event/type">
            <i class="menu-icon fa fa-caret-right"></i>
            事件类型
        </a>
        <b class="arrow"></b>
    </li>

    <li id="itsm-servicearea-menu" class="">
        <a href="/itsm/service/area">
            <i class="menu-icon fa fa-caret-right"></i>
            服务区域
        </a>
        <b class="arrow"></b>
    </li>
    <li id="itsm-engineergroup-menu" class="">
        <a href="/itsm/engineer/group">
            <i class="menu-icon fa fa-caret-right"></i>
            工程师分组
        </a>
        <b class="arrow"></b>
    </li>
	<li id="itsm-engineer-menu" class="">
		<a href="/itsm/engineer">
			<i class="menu-icon fa fa-caret-right"></i>
			工程师
		</a>
		<b class="arrow"></b>
	</li>
</ul>
</li>
`)

	return buffer.String()
}
