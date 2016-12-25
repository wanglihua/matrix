package menu

import (
	"bytes"
	"github.com/go-xorm/xorm"
	"github.com/revel/revel"
)

func GetErpMenu(db_session *xorm.Session, web_session revel.Session) string {
	if _, found := revel.ModuleByName("erp"); found == false {
		return ""
	}

	var buffer bytes.Buffer

	buffer.WriteString(`
<li id="erp-menu" class="">
    <a href="#" class="dropdown-toggle">
        <i class="menu-icon fa fa-bar-chart-o"></i>
        <span class="menu-text">ERP</span>
        <b class="arrow fa fa-angle-down"></b>
    </a>
    <b class="arrow"></b>
    <ul class="submenu">

        <li id="erp-brand-menu" class="">
            <a href="/erp/brand">
                <i class="menu-icon fa fa-caret-right"></i>
                品牌
            </a>
            <b class="arrow"></b>
        </li>

    </ul>
</li>
`)

	return buffer.String()
}
