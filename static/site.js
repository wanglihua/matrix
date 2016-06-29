function dtFormatDate(data) {
    if (!data) {
        return "";
    }
    data = data.replace(/T|Z/g, ' ');
    return stringToDate(data).format("yyyy-MM-dd");
}

function dtFormatDateTime(data) {
    if (!data) {
        return "";
    }
    data = data.replace(/T|Z/g, ' ');
    return stringToDate(data).format("yyyy-MM-dd hh:mm");
}

function dtFormatDateTime2(data) {
    if (!data) {
        return "";
    }
    data = data.replace(/T|Z/g, ' ');
    return stringToDate(data).format("yyyy-MM-dd hh:mm:ss");
}

function dtFormatFloat(data) {
    if (!data) {
        return "";
    }
    return $.formatNumber(data, {format: "#,##0.00", locale: "cn"});
}

function dtFormatFloat4(data) {
    if (!data) {
        return "";
    }
    return $.formatNumber(data, {format: "#,##0.0000", locale: "cn"});
}

function dtFormatInt(data) {
    if (!data) {
        return "";
    }
    return $.formatNumber(data, {format: "#,##0", locale: "cn"});
}

$.messager.model = {
    ok: {text: '<i class="ace-icon fa fa-check"></i>确认', classed: 'btn-success btn-sm'},
    cancel: {text: '<i class="ace-icon fa fa-times"></i>取消', classed: 'btn-default btn-sm'}
};

$.msg = {
    error: function (text) {
        $.messager.alert("错误", '<div class="huaduConfirm"><i class="fa fa-times-circle fa-5x red"></i></br></br>' + text + '</div>');
    },
    //waring: function (text) {
    //    $.messager.alert("警告", '<div class="huaduConfirm"><i class="fa fa-exclamation-triangle fa-5x orange2"></i></br></br>' + text + '</div>');
    //},
    alert: function (text) {
        $.messager.alert("提示", '<div class="huaduConfirm"><i class="fa fa-exclamation-triangle fa-5x orange2"></i></br></br>' + text + '</div>');
    },
    confirm: function (text, fn) {
        $.messager.confirm("确认", '<div class="huaduConfirm"><i class="fa fa-question-circle fa-5x orange2"></i></br></br> ' + text + '</div>', fn);
    },
    show: function (text) {
        $.messager.popup("<span style='color:green;font-weight:bold;'>" + text + "</span>");
    }
};

window.onresize = function () {
    $("#huaduCover").height($(window).height()).css("line-height", $(window).height() + "px").width($(window).width());
};

/*
$(document).ajaxStart(function () {
    showMask("数据加载中，请稍候");
}).ajaxStop(function () {
    hideMask();
});
*/

//显示遮照层
function showMask(text) {
    //$('body').css("overflow", "hidden")
    $("#huaduCover").show();
    $("#huaduProgressBar").html(text);
}
//隐藏遮照层
function hideMask() {
    //$('body').css("overflow", "")
    $("#huaduCover").hide();
}

function handleJQueryAjaxError(xhr, textStatus, error) {
    hideMask();
    if (textStatus === 'timeout') {
        $.msg.error('服务器超时！');
    }
    else {
        $.msg.error('访问服务器出错，请重试！' + error);
    }
}

$.ajaxSetup(
    {
        type: "POST",
        "timeout": 15000,
        "error": handleJQueryAjaxError
    }
);

//扩展dataTable默认值
$.extend($.fn.dataTable.defaults, {
    "sDom": "<'row'<'col-xs-6'l><'col-xs-6'f>r><'row'<'col-xs-12't>><'row'<'col-xs-6'i><'col-xs-6'p>>",
    "bAutoWidth": true,
    "bFilter": true,
    "bSort": true,
    "bInfo": true,
    "bLengthChange": true,
    "bServerSide": true,
    "sServerMethod": "POST",
    "bProcessing": false,
    "bPaginate": true,
    "lengthMenu": [[10, 15, 25, 50, 100], [10, 15, 25, 50, 100]],
    "iDisplayLength": 10,
    "sPaginationType": "full_numbers",
    "oLanguage": {
        "sLengthMenu": "每页显示 _MENU_ 条记录",
        "sZeroRecords": "没有匹配数据",
        "sInfo": "当前显示第 _START_ 至 _END_ 条，共 _TOTAL_ 条记录",
        "sInfoEmpty": "显示第 0 至 0 项结果，共 0 项",
        "sInfoFiltered": "(从 _MAX_ 条记录中过滤)",
        "sInfoPostFix": "",
        "sSearch": "搜索:",
        "sUrl": "",
        "sEmptyTable": "没有查询到任何数据！",
        "sLoadingRecords": "载入中...",
        //"sInfoThousands": ',',
        "sThousands": '',
        "oPaginate": {
            "sFirst": "首页",
            "sPrevious": "上页",
            "sNext": "下页",
            "sLast": "末页"
        },
        "oAria": {
            "sSortAscending": ": 以升序排列此列",
            "sSortDescending": ": 以降序排列此列"
        }
    },
    "fnServerData": function (sSource, aoData, fnCallback) {
        showMask("数据加载中，请稍候");
        $.ajax(
            {
                "dataType": 'json',
                "url": sSource,
                "data": aoData,
                "success": function (data) {
                    hideMask();
                    if (data.success == false) {
                        $.msg.error(data.message);
                        return;
                    }
                    fnCallback(data);
                }
            });
    }
});

function initDataTable(table, options) {

    var htmlTable = $(table);
    var dataTable = htmlTable.dataTable(options);

    function checkAll() {
        var thCheckBox = this;
        $(table).find('tbody tr>td:first-child input:checkbox').each(function () {
            var tdCheckBox = this;
            tdCheckBox.checked = thCheckBox.checked;
            var tableRow = $(this).closest('tr');
            if (tdCheckBox.checked) {
                tableRow.addClass('selected');
            } else {
                tableRow.removeClass('selected');
            }
        });
    }

    htmlTable.on('click', 'thead th input.checkall', checkAll);

    $(table + '_wrapper').on('click', 'div.dataTables_scroll div.dataTables_scrollHead thead th input.checkall', checkAll);

    htmlTable.on('click', 'tbody tr', function () {
        var tableRow = $(this);
        var checkbox = tableRow.find('td:first-child input:checkbox')[0];
        if (checkbox != undefined) {
            checkbox.checked = !checkbox.checked;
            if (checkbox.checked) {
                tableRow.addClass('selected');
            } else {
                tableRow.removeClass('selected');
            }
        }
    });

    htmlTable.on('click', 'tbody td:first-child input:checkbox', function (event) {
        event.stopPropagation();
        var checkbox = $(this)[0];
        var tableRow = $(this).closest('tr');
        if (checkbox.checked) {
            tableRow.addClass('selected');
        } else {
            tableRow.removeClass('selected');
        }
    });

    return dataTable;
}

function initSingleSelectDataTable(table, options) {
    var htmlTable = $(table);
    htmlTable.dataTable(options);

    htmlTable.on('click', 'tbody tr', function () {
        var tableRow = $(this);
        var checkbox = tableRow.find('td:first-child input:checkbox')[0];
        if (checkbox != undefined) {
            clearSelected(htmlTable);
            checkbox.checked = true;
            if (checkbox.checked) {
                tableRow.addClass('selected');
            }
        }
    });

    htmlTable.on('click', 'tbody td:first-child input:checkbox', function (event) {
        event.stopPropagation();
        var checkbox = $(this)[0];
        clearSelected(htmlTable, checkbox);
        var tableRow = $(this).closest('tr');
        if (checkbox.checked) {
            tableRow.addClass('selected');
        }
    });

    function clearSelected(htmlTable, exceptedCheckbox) {
        htmlTable.find('tbody tr').removeClass('selected');
        htmlTable.find('tbody td:first-child input:checkbox').each(function () {
            var checkbox = $(this)[0];
            if (exceptedCheckbox != checkbox) {
                checkbox.checked = false;
            }
        });
    }
}

function reloadList(list) {
    var tables = $(list).dataTable().api();
    tables.draw(false);
}

function trim(str) { //删除左右两端的空格
    return str.replace(/(^\s*)|(\s*$)/g, "");
}

//bootstrap modal options
var modalOptions = {backdrop: 'static', keyboard: false};

$.extend($.validator.messages, {
    required: "必填",
    remote: "请修正该字段",
    email: "请输入正确格式的电子邮件",
    url: "请输入合法的网址",
    date: "请输入合法的日期",
    dateISO: "请输入合法的日期(ISO).",
    number: "请输入合法的数字",
    digits: "只能输入整数",
    creditcard: "请输入合法的信用卡号",
    equalTo: "请再次输入相同的值",
    accept: "请输入拥有合法后缀名的字符串",
    maxlength: jQuery.validator.format("请输入一个长度最多是 {0} 的字符串"),
    minlength: jQuery.validator.format("请输入一个长度最少是 {0} 的字符串"),
    rangelength: jQuery.validator.format("请输入一个长度介于 {0} 和 {1} 之间的字符串"),
    range: jQuery.validator.format("请输入一个介于 {0} 和 {1} 之间的值"),
    max: jQuery.validator.format("请输入一个最大为 {0} 的值"),
    min: jQuery.validator.format("请输入一个最小为 {0} 的值")
});

function validate(form, rules, errorPlacement) {
    if (errorPlacement == undefined) {
        form.validate(
            {
                submitHandler: function () {
                },
                rules: rules,
                highlight: function (e) {
                    $(e).closest('.form-group').addClass('has-error');
                },
                success: function (e) {
                    $(e).closest('.form-group').removeClass('has-error');
                    $(e).remove();
                },
                errorElement: 'div',
                errorClass: 'help-block inline',
                focusInvalid: false,
                ignore: "",
                errorPlacement: function (error, element) {
                    //var parent = element.parent("div");
                    var parent = element.parent("div[class!='input-group']");
                    //if(!parent){
                    //    parent = element.parent().parent();
                    //}
                    error.appendTo(parent);
                }
            }
        );
    }
    else {
        form.validate(
            {
                submitHandler: function () {
                },
                rules: rules,
                highlight: function (e) {
                    $(e).closest('.form-group').addClass('has-error');
                },
                success: function (e) {
                    $(e).closest('.form-group').removeClass('has-error');
                    $(e).remove();
                },
                errorElement: 'div',
                errorClass: 'help-block inline',
                focusInvalid: false,
                ignore: "",
                errorPlacement: errorPlacement
            }
        );
    }

    return form.valid();
}


/*
 function getCookie(name) {
 var cookieValue = null;
 if (document.cookie && document.cookie != '') {
 var cookies = document.cookie.split(';');
 for (var i = 0; i < cookies.length; i++) {
 var cookie = jQuery.trim(cookies[i]);
 // Does this cookie string begin with the name we want?
 if (cookie.substring(0, name.length + 1) == (name + '=')) {
 cookieValue = decodeURIComponent(cookie.substring(name.length + 1));
 break;
 }
 }
 }
 return cookieValue;
 }

 var csrftoken = getCookie('csrftoken');

 var csrftoken = Cookies.get('csrftoken');
 */


/*
 function csrfSafeMethod(method) {
 // these HTTP methods do not require CSRF protection
 return (/^(GET|HEAD|OPTIONS|TRACE)$/.test(method));
 }

 $.ajaxSetup({
 beforeSend: function (xhr, settings) {
 if (!csrfSafeMethod(settings.type) && !this.crossDomain) {
 var csrftoken = Cookies.get('csrftoken');
 xhr.setRequestHeader("X-CSRFToken", csrftoken);
 }
 }
 });

 */

function isEmptyStr(str) {
    if (!str || str.length === 0 || str === "" || typeof str == 'undefined' || !/[^\s]/.test(str) || /^\s*$/.test(str) || str.replace(/\s/g, "") === "") {
        return true;
    } else {
        return false;
    }
}