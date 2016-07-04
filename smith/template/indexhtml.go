package template

var IndexHtmlTemplate = `
{%header "群组管理" .%}

<div class="breadcrumbs" id="breadcrumbs">
    <ul class="breadcrumb">
        <li>
            <a href='{%url "Home.Index"%}'><i class="ace-icon fa fa-home home-icon"></i></a>
        </li>
        <li class="parent" data-menu="auth-menu">权限管理</li>
        <li class="active" data-menu="auth-group-menu">群组管理</li>
    </ul>
</div>
<div class="page-content">
    <div class="row">
        <div class="col-xs-12" style="padding: 0;">
            <div class="gridToolBar clearfix">
                <button class="btn btn-primary btn-sm" onclick="reloadList('#groupList')">
                    <i class="ace-icon fa fa-refresh"></i>
                    刷新
                </button>
                <button class="btn btn-success btn-sm" onclick="showGroupDetail(0)">
                    <i class="ace-icon fa fa-plus"></i>
                    新增
                </button>
                <button class="btn btn-danger btn-sm" onclick="deleteGroup()">
                    <i class="ace-icon fa fa-trash-o"></i>
                    删除
                </button>
            </div>
            <div>
                <table id="groupList" class="table table-striped table-bordered table-hover"></table>
            </div>
            <div id="groupDetailDialogWrapper"></div>
            <script type="text/javascript">
                function showGroupDetail(id) {
                    $.ajax({
                               url: '{%url "AuthGroup.Detail" %}',
                               type: "POST",
                               data: {
                                   "id": id
                               },
                               success: function (result) {
                                   if (result.success == false) {
                                       $.msg.error(result.message);
                                       return;
                                   }
                                   $("#groupDetailDialogWrapper").html(result);
                                   $('#groupDetailDialog').modal(modalOptions);
                               }
                           }
                    );
                }

                function deleteGroup() {
                    var id_list = [];
                    $('#groupList tbody tr>td:first-child input:checkbox').each(function () {
                        if (this.checked) {
                            id_list.push($(this).val());
                        }
                    });

                    if (id_list.length == 0) {
                        $.msg.alert("请选择数据！");
                        return;
                    }

                    $.msg.confirm("确认要删除吗？", function () {
                        $.ajax({
                                   url: '{%url "AuthGroup.Delete"%}',
                                   type: "POST",
                                   data: {"id_list": id_list},
                                   success: function (data) {
                                       if (data.success == false) {
                                           $.msg.error(data.message);
                                           return;
                                       }
                                       else {
                                           $.msg.show(data.message);
                                       }
                                       reloadList('#groupList');
                                   }
                               });
                    });
                }

                $(function () {
                    initDataTable(
                            '#groupList',
                            {
                                "sAjaxSource": '{%url "AuthGroup.ListData"%}',
                                "aaSorting": [[1, "asc"]],
                                "aoColumns": [
                                    {
                                        "mDataProp": "id", "bSearchable": false, "orderable": false, "sClass": "center", "width": 50,
                                        "title": '<input type="checkbox" class="ace checkall"/><span class="lbl"></span>',
                                        "render": function (data, type, row) {
                                            return '<input type="checkbox" class="ace" value="' + data + '"/><span class="lbl"></span>';
                                        }
                                    },
                                    {"mDataProp": "group_name", "title": "群组名", "bSearchable": true},
                                    {"mDataProp": "create_time", "title": "创建时间", "bSearchable": false, "width": 120},
                                    {"mDataProp": "update_time", "title": "修改时间", "bSearchable": false, "width": 120}
                                ]
                            }
                    );

                    $('#groupList tbody').on('dblclick', 'tr', function () {
                        var id = $(this).find('td:first-child input:checkbox').val();
                        showGroupDetail(id);
                    });
                });
            </script>

        </div>
    </div>
</div>
{%footer%}

`
