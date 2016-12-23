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
<ul class="submenu">
    <li id="itsm-event_apply-menu" class="">
        <a href="/itsm/event/apply">
            <i class="menu-icon fa fa-caret-right"></i>
            事件提报
        </a>
        <b class="arrow"></b>
    </li>
    <li id="itsm-event_grab-menu" class="">
        <a href="/itsm/event/grab">
            <i class="menu-icon fa fa-caret-right"></i>
            我要抢单
        </a>
        <b class="arrow"></b>
    </li>
    <li id="itsm-event_process-menu" class="">
        <a href="/itsm/event/process">
            <i class="menu-icon fa fa-caret-right"></i>
            事件处理
        </a>
        <b class="arrow"></b>
    </li>	
    <li id="itsm-event_manage-menu" class="">
        <a href="/itsm/event/manage">
            <i class="menu-icon fa fa-caret-right"></i>
            事件管理
        </a>
        <b class="arrow"></b>
    </li>
    <li id="itsm-eventtype-menu" class="">
        <a href="/itsm/event/type">
            <i class="menu-icon fa fa-caret-right"></i>
            事件类型
        </a>
        <b class="arrow"></b>
    </li>
    <li id="itsm-eventstatus-menu" class="">
        <a href="/itsm/event/status">
            <i class="menu-icon fa fa-caret-right"></i>
            事件状态
        </a>
        <b class="arrow"></b>
    </li>
    <li id="itsm-eventpriority-menu" class="">
        <a href="/itsm/event/priority">
            <i class="menu-icon fa fa-caret-right"></i>
            事件优先级
        </a>
        <b class="arrow"></b>
    </li>

    <li id="itsm-eventlog-menu" class="">
        <a href="/itsm/event/log">
            <i class="menu-icon fa fa-caret-right"></i>
            事件日志
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
    <li id="itsm-eventapplydepartment-menu" class="">
        <a href="/itsm/event/apply/department">
            <i class="menu-icon fa fa-caret-right"></i>
            事件提报部门
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
    <li id="itsm-equipmentstatus-menu" class="">
        <a href="/itsm/equipment/status">
            <i class="menu-icon fa fa-caret-right"></i>
            设备状态
        </a>
        <b class="arrow"></b>
    </li>

    <li id="itsm-equipmenttype-menu" class="">
        <a href="/itsm/equipment/type">
            <i class="menu-icon fa fa-caret-right"></i>
            设备类型
        </a>
        <b class="arrow"></b>
    </li>

    <li id="itsm-equipment-menu" class="">
        <a href="/itsm/equipment">
            <i class="menu-icon fa fa-caret-right"></i>
            设备
        </a>
        <b class="arrow"></b>
    </li>

    <li id="itsm-applicationstatus-menu" class="">
        <a href="/itsm/application/status">
            <i class="menu-icon fa fa-caret-right"></i>
            应用状态
        </a>
        <b class="arrow"></b>
    </li>

    <li id="itsm-applicationtype-menu" class="">
        <a href="/itsm/application/type">
            <i class="menu-icon fa fa-caret-right"></i>
            应用类型
        </a>
        <b class="arrow"></b>
    </li>

    <li id="itsm-application-menu" class="">
        <a href="/itsm/application">
            <i class="menu-icon fa fa-caret-right"></i>
            应用
        </a>
        <b class="arrow"></b>
    </li>	
</ul>
</li>
`)

	return buffer.String()
}
