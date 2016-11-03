package menu

import (
	"bytes"
	"github.com/go-xorm/xorm"
	"github.com/revel/revel"
)

func GetWeixinMenu(db_session *xorm.Session, web_session revel.Session) string {
	if _, found := revel.ModuleByName("weixin"); found == false {
		return ""
	}

	var buffer bytes.Buffer

	buffer.WriteString(`
<li id="weixin-menu" class="">
    <a href="#" class="dropdown-toggle">
        <i class="menu-icon fa fa-wechat"></i>
        <span class="menu-text">微信管理</span>
        <b class="arrow fa fa-angle-down"></b>
    </a>
    <b class="arrow"></b>
    <ul class="submenu">
        <li id="weixin-config-menu" class="">
            <a href="/weixin/config">
                <i class="menu-icon fa fa-caret-right"></i>
                微信配置
            </a>
            <b class="arrow"></b>
        </li>
        <li id="weixin-menu-menu" class="">
            <a href="/weixin/menu">
                <i class="menu-icon fa fa-caret-right"></i>
                微信菜单
            </a>
            <b class="arrow"></b>
        </li>
    </ul>
</li>
`)

	return buffer.String()
}
