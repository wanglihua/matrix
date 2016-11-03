package menu

import (
	"bytes"
	"github.com/go-xorm/xorm"
	"github.com/revel/revel"
)

func GetHelpMenu(db_session *xorm.Session, web_session revel.Session) string {

	//if _, found := revel.ModuleByName("help"); found == false {
	//	return ""
	//}

	var buffer bytes.Buffer

	buffer.WriteString(`
<li id="help-menu" class="">
    <a href="#" class="dropdown-toggle">
        <i class="menu-icon fa fa-book"></i>
        <span class="menu-text">帮助中心</span>
        <b class="arrow fa fa-angle-down"></b>
    </a>
    <b class="arrow"></b>
    <ul class="submenu">
        <li id="help-index-menu" class="">
            <a href="/help">
                <i class="menu-icon fa fa-caret-right"></i>
                帮助中心
            </a>
            <b class="arrow"></b>
        </li>
        <li id="help-faq-menu" class="">
            <a href="/help/faq">
                <i class="menu-icon fa fa-caret-right"></i>
                常见问题
            </a>
            <b class="arrow"></b>
        </li>
        <li id="help-bug-menu" class="">
            <a href="/help/bug">
                <i class="menu-icon fa fa-caret-right"></i>
                Bug提报
            </a>
            <b class="arrow"></b>
        </li>
        <li id="help-requirement-menu" class="">
            <a href="/help/requirement">
                <i class="menu-icon fa fa-caret-right"></i>
                需求建议
            </a>
            <b class="arrow"></b>
        </li>
        <li id="help-license-menu" class="">
            <a href="/help/license">
                <i class="menu-icon fa fa-caret-right"></i>
                使用授权
            </a>
            <b class="arrow"></b>
        </li>
        <li id="help-about-menu" class="">
            <a href="/help/about">
                <i class="menu-icon fa fa-caret-right"></i>
                关于
            </a>
            <b class="arrow"></b>
        </li>
    </ul>
</li>
`)

	return buffer.String()
}
