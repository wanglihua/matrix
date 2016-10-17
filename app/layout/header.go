package layout

import (
	"github.com/revel/revel"
	"github.com/go-xorm/xorm"
	"bytes"
	"matrix/core"
)

func GetHeader(title string, db_session *xorm.Session, web_session revel.Session) string {
	sys_name := GetSysName(db_session)

	login_user := core.GetLoginUser(web_session)
	login_user_nick_name := "未登录"
	is_admin := false
	if login_user != nil {
		login_user_nick_name = login_user.NickName
		is_admin = login_user.IsAdmin
	}

	var header_buffer bytes.Buffer
	header_buffer.WriteString(`<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1"/>
    <meta charset="utf-8"/>
    <title>`)

	header_buffer.WriteString(title)
	header_buffer.WriteString("-")
	header_buffer.WriteString(sys_name)

	header_buffer.WriteString(`</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0"/>
    <link type="text/css" rel="stylesheet" href="/static/ace/css/all.css" class="ace-main-stylesheet" id="main-ace-style" />
    <!--[if lte IE 9]>
    <link type="text/css" rel="stylesheet" href="/static/ace/css/ace-part2.min.css" class="ace-main-stylesheet"/>
    <![endif]-->
    <!--[if lte IE 9]>
    <link type="text/css" rel="stylesheet" href="/static/ace/css/ace-ie.min.css"/>
    <![endif]-->
    <link type="text/css" rel="stylesheet" href="/static/site.css"/>

    <!--[if !IE]> -->
    <script type="text/javascript">
        window.jQuery || document.write("<script type='text/javascript' src='/static/ace/js/jquery.min.js'>" + "<" + "/script>");
    </script>
    <!-- <![endif]-->
    <!--[if IE]>
    <script type="text/javascript">
        window.jQuery || document.write("<script type='text/javascript' src='/static/ace/js/jquery1x.min.js'>" + "<" + "/script>");
    </script>
    <![endif]-->
    <script type="text/javascript">
        if ('ontouchstart' in document.documentElement) document.write("<script type='text/javascript' src='/static/ace/js/jquery.mobile.custom.min.js'>" + "<" + "/script>");
    </script>`)

	header_buffer.WriteString(`
</head>

<body class="no-skin">
<div id="huaduCover" class="huaduCover" style="display: none;">
    <div class="huaduBg" id="huaduBg">&nbsp;</div>
    <div class="huaduProgressBar" id="huaduProgressBar"></div>
</div>
<div id="navbar" class="navbar navbar-default" xmlns="http://www.w3.org/1999/html">
    <script type="text/javascript">
        try {
            ace.settings.check('navbar', 'fixed')
        } catch (e) {
        }
    </script>
    <div class="navbar-container" id="navbar-container">
        <button type="button" class="navbar-toggle menu-toggler pull-left" id="menu-toggler" data-target="#sidebar">
            <span class="sr-only">Toggle sidebar</span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
            <span class="icon-bar"></span>
        </button>
        <div class="navbar-header pull-left">
            <a href="/" class="navbar-brand">
                <small>`)

	header_buffer.WriteString(sys_name)

	header_buffer.WriteString(`</small>
            </a>
            </div>
        <div class="navbar-buttons navbar-header pull-right" role="navigation">
            <ul class="nav ace-nav">
                <li class="light-blue">
                    <a data-toggle="dropdown" href="#" class="dropdown-toggle">
                        <img class="nav-user-photo" src="/static/ace/avatars/user3.jpg" alt="user photo"/>
                                <span class="user-info"><small>欢迎！</small>`)

	header_buffer.WriteString(login_user_nick_name)

	header_buffer.WriteString(`</span>
                        <i class="ace-icon fa fa-caret-down"></i>
                    </a>
                    <ul class="user-menu dropdown-menu-right dropdown-menu dropdown-yellow dropdown-caret dropdown-closer">
                        <li>
                            <a href="/auth/profile">
                                <i class="ace-icon fa fa-cog"></i>
                                个人设置
                            </a>
                        </li>
                        <li class="divider"></li>
                        <li>
                            <a href="javascript:logout()">
                                <i class="ace-icon fa fa-power-off"></i>
                                注销
                            </a>
                            <script type="text/javascript">
                                function logout() {
                                    $.msg.confirm('确认要退出系统吗？', function () {
                                        $.ajax({
                                                   url: '/logout',
                                                   type: "POST",
                                                   success: function (data) {
                                                       if (data.success == true) {
                                                           window.location.href = '/login';
                                                           //return;
                                                       }
                                                       else {
                                                           $.msg.error(data.message);
                                                           //return;
                                                       }
                                                   }
                                               }
                                        );
                                    });
                                }
                            </script>
                        </li>
                     </ul>
                 </li>
            </ul>
        </div>
     </div>
</div>
<div class="main-container" id="main-container">
    <script type="text/javascript">
        $(function () {
            try {
                var parentMenuName = $("div#breadcrumbs ul.breadcrumb li.parent").data().menu;
                var activeMenuName = $("div#breadcrumbs ul.breadcrumb li.active").data().menu;
                var parentMenu = $("#" + parentMenuName);
                parentMenu.addClass("open");
                parentMenu.find("#" + activeMenuName).addClass("active");

				var subParentBreadcrumb = $("div#breadcrumbs ul.breadcrumb li.subparent");
				if(subParentBreadcrumb.size() != 0) {
	                var subParentMenuName = $("div#breadcrumbs ul.breadcrumb li.subparent").data().menu;
	                var subParentMenu = $("#" + subParentMenuName);
	                subParentMenu.addClass("open");
                }
            }
            catch (e) {
            }
        });
    </script>
    <!--sidebar-->
    <div id="sidebar" class="sidebar responsive">
        <script type="text/javascript">
            try {
                ace.settings.check('sidebar', 'fixed')
            } catch (e) {
            }
        </script>
        <ul class="nav nav-list">
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
                    <li id="inventory-stock-menu" class="">
                        <a href="/inventory/stock">
                            <i class="menu-icon fa fa-caret-right"></i>
                            仓库管理
                        </a>
                        <b class="arrow"></b>
                    </li>
                </ul>
            </li>
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

	if is_admin {
		header_buffer.WriteString(`
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
		header_buffer.WriteString(`
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

	header_buffer.WriteString(`
        <div class="sidebar-toggle sidebar-collapse" id="sidebar-collapse">
            <i class="ace-icon fa fa-angle-double-left" data-icon1="ace-icon fa fa-angle-double-left" data-icon2="ace-icon fa fa-angle-double-right"></i>
        </div>
        <script type="text/javascript">
            try {
                ace.settings.check('sidebar', 'collapsed')
            } catch (e) {
            }
        </script>
    </div>

    <div class="main-content">
`)
	return header_buffer.String()
}