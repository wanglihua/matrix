package template

var DetailHtmlTemplate =
`<div class="modal fade" id="{{.entity.EntityLowerCase}}DetailDialog">
    <div class="modal-dialog" style="width: 440px;">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="True">&times;</button>
                <h4 class="modal-title">{{.entity.EntityChinese}}详细</h4>
            </div>
            <div class="modal-body clearfix">
                <div class="row">
                    <div class="col-xs-12">
                        <form class="form-horizontal" role="form" id="{{.entity.EntityLowerCase}}DetailForm">
                            <input id="{{.entity.EntityLowerCase}}_id" name="form.{{.entity.EntityTitleName}}.Id" type="hidden" value="{%.form.{{.entity.EntityTitleName}}.Id%}"/>
                            <input id="{{.entity.EntityLowerCase}}_version" name="form.{{.entity.EntityTitleName}}.Version" type="hidden" value="{%.form.{{.entity.EntityTitleName}}.Version%}"/>
                            <div class="form-{{.entity.EntityLowerCase}}">
                                <label for="{{.entity.EntityLowerCase}}_name" class="col-xs-2 control-label no-padding-right">{{.entity.EntityChinese}}名</label>
                                <div class="col-xs-10">
                                    <input id="{{.entity.EntityLowerCase}}_name" name="form.{{.entity.EntityTitleName}}.{{.entity.EntityTitleName}}Name" type="text" class="input-sm form-control" value="{%.form.{{.entity.EntityTitleName}}.{{.entity.EntityTitleName}}Name%}"/>
                                </div>
                            </div>
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
    var form = $("#{{.entity.EntityLowerCase}}DetailForm");
    var valid = validate(form, {
        'form.{{.entity.EntityTitleName}}.{{.entity.EntityTitleName}}Name': {
            required: true,
            minlength: 3,
            maxlength: 20
        }
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
                           $('#{{.entity.EntityLowerCase}}DetailDialog').modal('hide');
                           reloadList('#{{.entity.EntityLowerCase}}List');
                           //return;
                       }
                   }
               }
        );
    }
}
</script>

`
