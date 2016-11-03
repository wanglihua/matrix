package menu

import (
	"bytes"
	"github.com/go-xorm/xorm"
	"github.com/revel/revel"
)

func GetInventoryMenu(db_session *xorm.Session, web_session revel.Session) string {
	if _, found := revel.ModuleByName("inventory"); found == false {
		return ""
	}

	var buffer bytes.Buffer

	buffer.WriteString(`
<li id="inventory-menu" class="">
    <a href="#" class="dropdown-toggle">
        <i class="menu-icon fa fa-truck"></i>
        <span class="menu-text">仓储管理</span>
        <b class="arrow fa fa-angle-down"></b>
    </a>
    <b class="arrow"></b>
    <ul class="submenu">
        <li id="inventory-relparty-menu" class="">
            <a href="#" class="dropdown-toggle">
                <i class="menu-icon fa fa-caret-right"></i>
                <span class="menu-text">往来单位</span>
                <b class="arrow fa fa-angle-down"></b>
            </a>
            <b class="arrow"></b>
            <ul class="submenu">
                <li id="inventory-supplier-menu">
                    <a href="/inventory/supplier">
                        <i class="menu-icon fa fa-caret-right"></i>
                        供应商管理
                    </a>
                    <b class="arrow"></b>
                </li>
                <li id="inventory-client-menu">
                    <a href="/inventory/client">
                        <i class="menu-icon fa fa-caret-right"></i>
                        客户管理
                    </a>
                    <b class="arrow"></b>
                </li>
            </ul>
        </li>
        <li id="inventory-prdmgr-menu" class="">
            <a href="#" class="dropdown-toggle">
                <i class="menu-icon fa fa-caret-right"></i>
                <span class="menu-text">货品管理</span>
                <b class="arrow fa fa-angle-down"></b>
            </a>
            <b class="arrow"></b>
            <ul class="submenu">
                <li id="inventory-prdcate-menu">
                    <a href="/inventory/product/cate">
                        <i class="menu-icon fa fa-caret-right"></i>
                        货品类别
                    </a>
                    <b class="arrow"></b>
                </li>
                <li id="inventory-product-menu">
                    <a href="/inventory/product">
                        <i class="menu-icon fa fa-caret-right"></i>
                        货品信息
                    </a>
                    <b class="arrow"></b>
                </li>
            </ul>
        </li>
        <li id="inventory-stock-menu" class="">
            <a href="/inventory/stock">
                <i class="menu-icon fa fa-caret-right"></i>
                仓库信息
            </a>
            <b class="arrow"></b>
        </li>
        <li id="inventory-capitalaccount-menu" class="">
            <a href="/inventory/capital/account">
                <i class="menu-icon fa fa-caret-right"></i>
                资金账户
            </a>
            <b class="arrow"></b>
        </li>
        <li id="inventory-stockiotype-menu" class="">
            <a href="/inventory/stock/io/type">
                <i class="menu-icon fa fa-caret-right"></i>
                出入库类型
            </a>
            <b class="arrow"></b>
        </li>
        <li id="inventory-paytype-menu" class="">
            <a href="/inventory/pay/type">
                <i class="menu-icon fa fa-caret-right"></i>
                收付款类型
            </a>
            <b class="arrow"></b>
        </li>
        <li id="inventory-handler-menu" class="">
            <a href="/inventory/handler">
                <i class="menu-icon fa fa-caret-right"></i>
                经手人
            </a>
            <b class="arrow"></b>
        </li>
        <li id="inventory-stockbill-menu" class="">
            <a href="/inventory/stock/bill">
                <i class="menu-icon fa fa-caret-right"></i>
                出入库单
            </a>
            <b class="arrow"></b>
        </li>
        <li id="inventory-stockbilldetail-menu" class="">
            <a href="/inventory/stock/bill/detail">
                <i class="menu-icon fa fa-caret-right"></i>
                出入库单详细
            </a>
            <b class="arrow"></b>
        </li>
        <li id="inventory-config-menu" class="">
            <a href="/inventory/config">
                <i class="menu-icon fa fa-caret-right"></i>
                系统配置
            </a>
            <b class="arrow"></b>
        </li>
    </ul>
</li>
`)

	return buffer.String()
}
