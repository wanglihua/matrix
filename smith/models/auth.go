package models

import (
    "matrix/smith"
    "matrix/smith/fieldtype"
)

var GroupEntity = smith.Entity{
    ModuleTitleName     : "Auth",
    ModuleLowerCase     : "auth",
    ModuleChinese       : "权限管理",
    EntityTitleName     : "Group",
    EntityCamelCase     : "group",
    EntityChinese       : "群组",
    TableName           : "group",
    TablePrefix         : "hd_auth_",

    FieldList: []smith.Field{
        smith.Field{
            VerboseName : "群组名称",
            Name        : "GroupName",
            Column      : "group_name",
            FieldType   : fieldtype.NVarchar,
            Length      : 255,
            Precision   : 0,
            Scale       : 0,
            Unique      : true,
            Index       : true,
            Null        : false,
            Blank       : false,
            Min         : "3",
            Max         : "20",
        },
    },
}
