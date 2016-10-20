package template

var DetailHtmlTemplate = `<div class="modal fade" id="{{.entity.EntityCamelCase}}DetailDialog">
    <div class="modal-dialog" style="width: 440px;">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="True">&times;</button>
                <h4 class="modal-title">{{.entity.EntityChinese}}详细</h4>
            </div>
            <div class="modal-body clearfix">
                <div class="row">
                    <div class="col-xs-12">
                        <form class="form-horizontal" role="form" id="{{.entity.EntityCamelCase}}DetailForm" onsubmit="save{{.entity.EntityTitleName}}();return false;">
                            <input id="{{.entity.EntityCamelCase}}_id" name="form.{{.entity.EntityTitleName}}.Id" type="hidden" value="{%.form_{{.entity.EntityTitleName}}_Id%}"/>
                            <input id="{{.entity.EntityCamelCase}}_version" name="form.{{.entity.EntityTitleName}}.Version" type="hidden" value="{%.form_{{.entity.EntityTitleName}}_Version%}"/>{{range $fieldIndex, $field := .entity.FieldList}}
                            <div class="form-group">
                                <label for="{{$field.Column}}" class="col-xs-2 control-label no-padding-right">{{$field.VerboseName}}</label>
                                <div class="col-xs-10">
                                    <input id="{{$field.Column}}" name="form.{{$.entity.EntityTitleName}}.{{$field.Name}}" type="text" class="input-sm form-control" value="{%.form_{{$.entity.EntityTitleName}}_{{$field.Name}}%}"/>
                                </div>
                            </div>
{{end}}
                        </form>
                    </div>
                </div>
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-success btn-sm" onclick="save{{.entity.EntityTitleName}}()">
                    <i class="ace-icon fa fa-floppy-o"></i>保存
                </button>
                <button type="button" class="btn btn-default btn-sm" data-dismiss="modal">
                    <i class="ace-icon fa fa-times"></i>关闭
                </button>
            </div>
        </div>
    </div>
</div>
<script type='text/javascript'>
function save{{.entity.EntityTitleName}}() {
    showMask("提交中...");
    var form = $("#{{.entity.EntityCamelCase}}DetailForm");
    var valid = validate(form, { {{ $fieldMaxIndex := ListMaxIndex .entity.FieldList}} {{range $fieldIndex, $field := .entity.FieldList}}
        'form.{{$.entity.EntityTitleName}}.{{$field.Name}}': {
            {{FieldClienValid $field}}
        }{{if ne $fieldMaxIndex $fieldIndex}},{{end}}{{end}}
    });

    if (!valid) {
        hideMask();
    } else {
        $.ajax({
                   url: '{%url "{{.entity.ModuleTitleName}}{{.entity.EntityTitleName}}.Save"%}',
                   type: "POST",
                   data: form.serialize(),
                   success: function (data) {
                       hideMask();
                       if (data.success == false) {
                           $.msg.error(data.message);
                           //return;
                       }
                       else {
                           $.msg.show(data.message);
                           $('#{{.entity.EntityCamelCase}}DetailDialog').modal('hide');
                           reloadList('#{{.entity.EntityCamelCase}}List');
                           //return;
                       }
                   }
               }
        );
    }
}
</script>

`
