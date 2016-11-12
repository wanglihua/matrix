package template

var IndexHtmlTemplate = `{%header "{{.entity.EntityChinese}}" .%}
<div id="page">
   <div class="breadcrumbs" id="breadcrumbs">
       <ul class="breadcrumb">
           <li>
               <a href='{%url "Home.Index"%}'><i class="ace-icon fa fa-home home-icon"></i></a>
           </li>
           <li class="parent" data-menu="{{.entity.ModuleLowerCase}}-menu">{{.entity.ModuleChinese}}</li>
           <li class="active" data-menu="{{.entity.ModuleLowerCase}}-{{LowerCase .entity.EntityCamelCase}}-menu">{{.entity.EntityChinese}}</li>
       </ul>
   </div>
   <div class="page-content">
       <div class="row">
           <div class="col-xs-12" style="padding: 0;">
               <div class="gridToolBar clearfix">
                   <button class="btn btn-primary btn-sm" @click="reloadList('#{{.entity.EntityCamelCase}}List')">
                       <i class="ace-icon fa fa-refresh"></i>
                       刷新
                   </button>
                   <button class="btn btn-success btn-sm" @click="show{{.entity.EntityTitleName}}Detail(0)">
                       <i class="ace-icon fa fa-plus"></i>
                       新增
                   </button>
                   <button class="btn btn-danger btn-sm" @click="delete{{.entity.EntityTitleName}}()">
                       <i class="ace-icon fa fa-trash-o"></i>
                       删除
                   </button>
               </div>
               <div>
                   <table id="{{.entity.EntityCamelCase}}List" class="table table-striped table-bordered table-hover"></table>
               </div>
           </div>
       </div>
   </div>
    <div class="modal fade" id="{{.entity.EntityCamelCase}}DetailDialog">
        <div class="modal-dialog" style="width: 680px;">
            <div class="modal-content">
                <div class="modal-header">
                    <button type="button" class="close" data-dismiss="modal" aria-hidden="True">&times;</button>
                    <h4 class="modal-title">{{.entity.EntityChinese}}详细</h4>
                </div>
                <div class="modal-body clearfix">
                    <div class="row">
                        <div class="col-xs-12">
                            <form class="form-horizontal" role="form" id="{{.entity.EntityCamelCase}}DetailForm">
                                <input id="{{Underscore .entity.EntityCamelCase}}_id" name="detail.{{.entity.EntityJson}}.id" type="hidden" v-model.number="detail.{{.entity.EntityJson}}.id"/>
                                <input id="{{Underscore .entity.EntityCamelCase}}_version" name="detail.{{.entity.EntityJson}}.version" type="hidden" v-model.number="detail.{{.entity.EntityJson}}.version"/>{{range $fieldIndex, $field := .entity.FieldList}}
                                <div class="row">
                                    <div class="col-xs-6">
                                        <div class="form-group">
                                            <label for="{{$field.Column}}" class="col-xs-2 control-label no-padding-right">{{$field.VerboseName}}</label>
                                            <div class="col-xs-10">
                                                <input id="{{$field.Column}}" name="detail.{{$.entity.EntityJson}}.{{$field.Column}}" type="text" class="input-sm form-control" v-model{{if $field.IsNumber}}.number{{end}}="detail.{{$.entity.EntityJson}}.{{$field.Column}}"/>
                                            </div>
                                        </div>
                                    </div>
                                </div>
{{end}}
                            </form>
                        </div>
                    </div>
                </div>
                <div class="modal-footer">
                    <button type="button" class="btn btn-success btn-sm" @click="save{{.entity.EntityTitleName}}()">
                        <i class="ace-icon fa fa-floppy-o"></i>保存
                    </button>
                    <button type="button" class="btn btn-default btn-sm" data-dismiss="modal">
                        <i class="ace-icon fa fa-times"></i>关闭
                    </button>
                </div>
            </div>
        </div>
    </div>
</div>
<script type="text/javascript">
    var vm = new Vue({
        el: "#page",
        data: {
            "detail": {
                "{{.entity.EntityJson}}": {}
            }
        },
        mounted: function () {
            initDataTable(
                    '#{{.entity.EntityCamelCase}}List',
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
                            },{{range $fieldIndex, $field := .entity.FieldList}}
                            {"mDataProp": "{{$field.Column}}", "title": "{{$field.VerboseName}}", "bSearchable": {{FieldSearchable $field}} },{{end}}
                            {"mDataProp": "create_time", "title": "创建时间", "bSearchable": false, "width": 120 },
                            {"mDataProp": "update_time", "title": "修改时间", "bSearchable": false, "width": 120 }
                        ]
                   }
            );

            $('#{{.entity.EntityCamelCase}}List tbody').on('dblclick', 'tr', function () {
                var id = $(this).find('td:first-child input:checkbox').val();
                vm.show{{.entity.EntityTitleName}}Detail(id);
            });
        },
        methods: {
            reloadList: function () {
                reloadList('#{{.entity.EntityCamelCase}}List');
            },
            show{{.entity.EntityTitleName}}Detail: function (id) {
                showMask("加载中...");
                $("#{{.entity.EntityCamelCase}}DetailForm").clearValidation();
                $.ajax({
                           url: '{%url "{{.entity.ModuleTitleName}}{{.entity.EntityTitleName}}.DetailData" %}',
                           type: "POST",
                           data: {
                               "id": id
                           },
                           success: function (result) {
                               hideMask();
                               if (result.success == false) {
                                   $.msg.error(result.message);
                                   return;
                               }
                               vm.detail = result;
                               $('#{{.entity.EntityCamelCase}}DetailDialog').modal(modalOptions);
                           }
                       }
                );
            },
            save{{.entity.EntityTitleName}}: function () {
                showMask("提交中...");
                var form = $("#{{.entity.EntityCamelCase}}DetailForm");
                var valid = validate(form, { {{ $fieldMaxIndex := ListMaxIndex .entity.FieldList}} {{range $fieldIndex, $field := .entity.FieldList}}
                    'detail.{{$.entity.EntityJson}}.{{$field.Column}}': {
                        {{FieldClienValid $field}}
                    }{{if ne $fieldMaxIndex $fieldIndex}},{{end}}{{end}}
                });

                if (!valid) {
                    hideMask();
                } else {
                    $.ajax({
                               url: '{%url "{{.entity.ModuleTitleName}}{{.entity.EntityTitleName}}.Save"%}',
                               type: "POST",
                               processData: false,
                               contentType: 'application/json',
                               data: JSON.stringify(vm.detail),
                               success: function (data) {
                                   hideMask();
                                   if (data.success == false) {
                                       $.msg.error(data.message);
                                       //return;
                                   }
                                   else {
                                       $.msg.show(data.message);
                                       $('#{{.entity.EntityCamelCase}}DetailDialog').modal('hide');
                                       vm.reloadList();
                                       //return;
                                   }
                               }
                           }
                    );
                }
            },
            delete{{.entity.EntityTitleName}}: function () {
                var id_list = [];
                $('#{{.entity.EntityCamelCase}}List tbody tr>td:first-child input:checkbox').each(function () {
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
                                   vm.reloadList();
                               }
                           });
                });
            }
        }
    });
</script>
{%footer%}

`
