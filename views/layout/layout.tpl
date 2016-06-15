<!DOCTYPE html>
<html>
<head>
    <meta http-equiv="X-UA-Compatible" content="IE=edge,chrome=1"/>
    <meta charset="utf-8"/>
    <title>Home - 系统</title>
    <meta name="viewport" content="width=device-width, initial-scale=1.0, maximum-scale=1.0"/>
    <link type="text/css" rel="stylesheet" href="/static/ace/css/bootstrap.min.css"/>
    <link type="text/css" rel="stylesheet" href="/static/ace/css/font-awesome.min.css"/>
    <link type="text/css" rel="stylesheet" href="/static/ace/css/ace.min.css" class="ace-main-stylesheet" id="main-ace-style"/>
    <link type="text/css" rel="stylesheet" href="/static/ace/css/ace-fonts.min.css"/>
    <!--[if lte IE 9]>
        <link type="text/css" rel="stylesheet" href="/static/ace/css/ace-part2.min.css" class="ace-main-stylesheet" />
    <![endif]-->
    <!--[if lte IE 9]>
        <link type="text/css" rel="stylesheet" href="/static/ace/css/ace-ie.min.css" />
    <![endif]-->
    <link type="text/css" rel="stylesheet" href="/static/ace/css/bootstrap-datetimepicker.min.css"/>
    <link type="text/css" rel="stylesheet" href="/static/ace/css/bootstrap-multiselect.min.css"/>
    <link type="text/css" rel="stylesheet" href="/static/site.css"/>
    <!--[if !IE]> -->
    <script type="text/javascript">
        window.jQuery || document.write("<script type='text/javascript' src='/static/ace/js/jquery.min.js'%}'>" + "<" + "/script>");
    </script>
    <!-- <![endif]-->
    <!--[if IE]>
    <script type="text/javascript">
        window.jQuery || document.write("<script type='text/javascript' src='/static/ace/js/jquery1x.min.js'>"+"<"+"/script>");
    </script>
    <![endif]-->
    <script type="text/javascript">
        if ('ontouchstart' in document.documentElement) document.write("<script type='text/javascript' src='/static/ace/js/jquery.mobile.custom.min.js'%}'>" + "<" + "/script>");
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
            <a href="{{urlfor "Home.Index"}}" class="navbar-brand">
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
                                        window.location.href = '#';
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
    {{template "layout/sidebar.tpl" .}}
    <div class="main-content">
        <div class="breadcrumbs" id="breadcrumbs">
            <script type="text/javascript">
                try {
                    ace.settings.check('breadcrumbs', 'fixed')
                } catch (e) {
                }
            </script>
            <ul class="breadcrumb">
                <li>
                    <a href="/"><i class="ace-icon fa fa-home home-icon"></i></a>
                </li>
                <li class="active">Home</li>
            </ul>
        </div>
        <div class="page-content">
            <div class="row">
                <div class="col-xs-12" style="padding: 0;">
                    <!-- main content -->
                    {{.LayoutContent}}
                </div>
            </div>
        </div>
    </div>
</div>
</body>
</html>
