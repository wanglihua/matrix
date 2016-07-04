package template

var DetailHtmlTemplate = `
<div class="modal fade" id="groupDetailDialog">
    <div class="modal-dialog" style="width: 440px;">
        <div class="modal-content">
            <div class="modal-header">
                <button type="button" class="close" data-dismiss="modal" aria-hidden="True">&times;</button>
                <h4 class="modal-title">群组详细</h4>
            </div>
            <div class="modal-body clearfix">
                <div class="row">
                    <div class="col-xs-12">
                        <form class="form-horizontal" role="form" id="groupDetailForm">
                            <input id="group_id" name="form.Group.Id" type="hidden" value="{%.form.Group.Id%}"/>
                            <input id="group_version" name="form.Group.Version" type="hidden" value="{%.form.Group.Version%}"/>
                            <div class="form-group">
                                <label for="group_name" class="col-xs-2 control-label no-padding-right">群组名</label>
                                <div class="col-xs-10">
                                    <input id="group_name" name="form.Group.GroupName" type="text" class="input-sm form-control" value="{%.form.Group.GroupName%}"/>
                                </div>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
            <div class="modal-footer">
                <button type="button" class="btn btn-success btn-sm" onclick="saveGroup()">
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
function saveGroup() {
    showMask("提交中...");
    var form = $("#groupDetailForm");
    var valid = validate(form, {
        'form.Group.GroupName': {
            required: true,
            minlength: 3,
            maxlength: 20
        }
    });

    if (!valid) {
        hideMask();
    } else {
        $.ajax({
                   url: '{%url "AuthGroup.Save"%}',
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
                           $('#groupDetailDialog').modal('hide');
                           reloadList('#groupList');
                           //return;
                       }
                   }
               }
        );
    }
}
</script>

`
