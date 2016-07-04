package template

var IndexHtmlTemplate = `
{%header "{{.entity.EntityChinese}}管理" .%}

<div class="breadcrumbs" id="breadcrumbs">
    <ul class="breadcrumb">
        <li>
            <a href='{%url "Home.Index"%}'><i class="ace-icon fa fa-home home-icon"></i></a>
        </li>
        <li class="parent" data-menu="{{.entity.ModuleLowerCase}}-menu">{{.entity.ModuleChinese}}</li>
        <li class="active" data-menu="{{.entity.ModuleLowerCase}}-{{.entity.EntityLowerCase}}-menu">{{.entity.EntityChinese}}管理</li>
    </ul>
</div>
<div class="page-content">
    <div class="row">
        <div class="col-xs-12" style="padding: 0;">
            <div class="gridToolBar clearfix">
                <button class="btn btn-primary btn-sm" onclick="reloadList('#{{.entity.EntityLowerCase}}List')">
                    <i class="ace-icon fa fa-refresh"></i>
                    刷新
                </button>
                <button class="btn btn-success btn-sm" onclick="show{{.entity.EntityTitleName}}Detail(0)">
                    <i class="ace-icon fa fa-plus"></i>
                    新增
                </button>
                <button class="btn btn-danger btn-sm" onclick="delete{{.entity.EntityTitleName}}()">
                    <i class="ace-icon fa fa-trash-o"></i>
                    删除
                </button>
            </div>
            <div>
                <table id="{{.entity.EntityLowerCase}}List" class="table table-striped table-bordered table-hover"></table>
            </div>
            <div id="{{.entity.EntityLowerCase}}DetailDialogWrapper"></div>
            <script type="text/javascript">
                function show{{.entity.EntityTitleName}}Detail(id) {
                    $.ajax({
                               url: '{%url "{{.entity.ModuleTitleName}}{{.entity.EntityTitleName}}.Detail" %}',
                               type: "POST",
                               data: {
                                   "id": id
                               },
                               success: function (result) {
                                   if (result.success == false) {
                                       $.msg.error(result.message);
                                       return;
                                   }
                                   $("#{{.entity.EntityLowerCase}}DetailDialogWrapper").html(result);
                                   $('#{{.entity.EntityLowerCase}}DetailDialog').modal(modalOptions);
                               }
                           }
                    );
                }

                function delete{{.entity.EntityTitleName}}() {
                    var id_list = [];
                    $('#{{.entity.EntityLowerCase}}List tbody tr>td:first-child input:checkbox').each(function () {
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
                                   url: '{%url "{{.entity.ModuleTitleName}}{{.entity.EntityTitleName}}.Delete"%}',
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
                                       reloadList('#{{.entity.EntityLowerCase}}List');
                                   }
                               });
                    });
                }

                $(function () {
                    initDataTable(
                            '#{{.entity.EntityLowerCase}}List',
                            {
                                "sAjaxSource": '{%url "{{.entity.ModuleTitleName}}{{.entity.EntityTitleName}}.ListData"%}',
                                "aaSorting": [[1, "asc"]],
                                "aoColumns": [
                                    {
                                        "mDataProp": "id", "bSearchable": false, "orderable": false, "sClass": "center", "width": 50,
                                        "title": '<input type="checkbox" class="ace checkall"/><span class="lbl"></span>',
                                        "render": function (data, type, row) {
                                            return '<input type="checkbox" class="ace" value="' + data + '"/><span class="lbl"></span>';
                                        }
                                    },
                                    {"mDataProp": "{{.entity.EntityLowerCase}}_name", "title": "{{.entity.EntityChinese}}名", "bSearchable": true},
                                    {"mDataProp": "create_time", "title": "创建时间", "bSearchable": false, "width": 120},
                                    {"mDataProp": "update_time", "title": "修改时间", "bSearchable": false, "width": 120}
                                ]
                            }
                    );

                    $('#{{.entity.EntityLowerCase}}List tbody').on('dblclick', 'tr', function () {
                        var id = $(this).find('td:first-child input:checkbox').val();
                        show{{.entity.EntityTitleName}}Detail(id);
                    });
                });
            </script>

        </div>
    </div>
</div>
{%footer%}

`