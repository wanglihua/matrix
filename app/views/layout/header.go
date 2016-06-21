package layout

func GetHeader(title string) string {
    return `<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1"/>
    <meta charset="utf-8"/>
    <title>` + title + `- 系统</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0"/>
    <link type="text/css" rel="stylesheet" href="/static/ace/css/all.css"/>
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
    </script>
    <script type="text/javascript" src="/static/ace/js/all1.js"></script>
    <!--[if lte IE 8]>
    <script type="text/javascript" src="/static/ace/js/html5shiv.min.js"></script>
    <script type="text/javascript" src="/static/ace/js/respond.min.js"></script>
    <![endif]-->
    <script type="text/javascript" src="/static/ace/js/all2.js"></script>
    <script type="text/javascript" src="/static/site.js"></script>
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
                <small>
                    应用系统
                </small>
            </a>
        </div>
        <div class="navbar-buttons navbar-header pull-right" role="navigation">
            <ul class="nav ace-nav">
                <li class="light-blue">
                    <a data-toggle="dropdown" href="#" class="dropdown-toggle">
                        <img class="nav-user-photo" src="/static/ace/avatars/user3.jpg" alt="user photo"/>
                                <span class="user-info">
                                    <small>欢迎！</small>
                                    "UserInfo"
                                </span>
                        <i class="ace-icon fa fa-caret-down"></i>
                    </a>
                    <ul class="user-menu dropdown-menu-right dropdown-menu dropdown-yellow dropdown-caret dropdown-closer">
                        <li>
                            <a href="#">
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
                                        window.location.href = '{{url "Login.Index"}}';
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
            <li id="sync-menu" class="">
                <a href="#" class="dropdown-toggle">
                    <i class="menu-icon fa fa-users"></i>
                    <span class="menu-text">数据同步</span>
                    <b class="arrow fa fa-angle-down"></b>
                </a>
                <b class="arrow"></b>
                <ul class="submenu">
                    <li id="sync-customer-menu" class="">
                        <a href="#">
                            <i class="menu-icon fa fa-caret-right"></i>
                            客户管理
                        </a>
                        <b class="arrow"></b>
                    </li>
                    <li id="sync-supplier-menu" class="">
                        <a href="#">
                            <i class="menu-icon fa fa-caret-right"></i>
                            供应商管理
                        </a>
                        <b class="arrow"></b>
                    </li>
                </ul>
            </li>
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
                </ul>
            </li>
        </ul>
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
`
}

/*
    <link type="text/css" rel="stylesheet" href="/static/ace/css/bootstrap.min.css"/>
    <link type="text/css" rel="stylesheet" href="/static/ace/css/font-awesome.min.css"/>
    <link type="text/css" rel="stylesheet" href="/static/ace/css/ace.min.css" class="ace-main-stylesheet" id="main-ace-style"/>
    <link type="text/css" rel="stylesheet" href="/static/ace/css/ace-fonts.min.css"/>
    <!--[if lte IE 9]>
    <link type="text/css" rel="stylesheet" href="/static/ace/css/ace-part2.min.css" class="ace-main-stylesheet"/>
    <![endif]-->
    <!--[if lte IE 9]>
    <link type="text/css" rel="stylesheet" href="/static/ace/css/ace-ie.min.css"/>
    <![endif]-->
    <link type="text/css" rel="stylesheet" href="/static/ace/css/bootstrap-datetimepicker.min.css"/>
    <link type="text/css" rel="stylesheet" href="/static/ace/css/bootstrap-multiselect.min.css"/>
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
    </script>
    <script type="text/javascript" src="/static/ace/js/jquery.validate.min.js"></script>
    <script type="text/javascript" src="/static/ace/js/jquery-ui.min.js"></script>
    <script type="text/javascript" src="/static/ace/js/jquery-ui.custom.min.js"></script>
    <script type="text/javascript" src="/static/ace/js/bootstrap.min.js"></script>
    <script type="text/javascript" src="/static/ace/js/jquery.bootstrap.min.js"></script>
    <script type="text/javascript" src="/static/ace/js/jquery.ba-resize.js"></script>
    <script type="text/javascript" src="/static/ace/js/ace.min.js"></script>
    <script type="text/javascript" src="/static/ace/js/ace-elements.min.js"></script>
    <script type="text/javascript" src="/static/ace/js/ace-extra.min.js"></script>
    <!--[if lte IE 8]>
    <script type="text/javascript" src="/static/ace/js/html5shiv.min.js"></script>
    <script type="text/javascript" src="/static/ace/js/respond.min.js"></script>
    <![endif]-->
    <script type="text/javascript" src="/static/ace/js/dataTables/jquery.dataTables.min.js"></script>
    <script type="text/javascript" src="/static/ace/js/dataTables/jquery.dataTables.bootstrap.min.js"></script>
    <script type="text/javascript" src="/static/ace/js/dataTables/extensions/buttons/dataTables.buttons.min.js"></script>
    <script type="text/javascript" src="/static/ace/js/dataTables/extensions/buttons/buttons.colVis.min.js"></script>
    <script type="text/javascript" src="/static/ace/js/dataTables/extensions/select/dataTables.select.min.js"></script>
    <script type="text/javascript" src="/static/ace/js/date-time/moment.min.js"></script>
    <script type="text/javascript" src="/static/ace/js/jquery.dataTables.fixedHeader.js"></script>
    <script type="text/javascript" src="/static/ace/js/jquery.dataTables.fixedColumns.js"></script>
    <script type="text/javascript" src="/static/ace/js/extend.datetime.js"></script>
    <script type="text/javascript" src="/static/ace/js/jquery.numberformatter.js"></script>
    <script type="text/javascript" src="/static/ace/js/datetimepicker/bootstrap-datetimepicker.min.js"></script>
    <script type="text/javascript" src="/static/ace/js/datetimepicker/locales/bootstrap-datetimepicker.zh-CN.js"></script>
    <script type="text/javascript" src="/static/ace/js/bootstrap-multiselect.min.js"></script>
    <script type="text/javascript" src="/static/ace/js/fuelux/fuelux.spinner.min.js"></script>
    <script type="text/javascript" src="/static/js/js.cookie.js"></script>
    <script type="text/javascript" src="/static/site.js"></script>
*/