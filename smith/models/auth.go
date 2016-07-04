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
    EntityLowerCase     : "group",
    EntityChinese       : "群组",
    TableName           : "group",
    TablePrefix         : "hd_auth_",

    FieldList: []smith.Field{
        smith.Field{
            VerboseName : "Id",
            Name        : "Id",
            Column      : "id",
            FieldType   : fieldtype.BigInt,
            Length      : 0,
            Precision   : 0,
            Scale       : 0,
            Primarykey  : true,
            Unique      : true,
            Index       : false,
            Null        : false,
            Blank       : false,
            Min         : "",
            Max         : "",
        },
        smith.Field{
            VerboseName : "群组名称",
            Name        : "GroupName",
            Column      : "group_name",
            FieldType   : fieldtype.NVarchar,
            Length      : 255,
            Precision   : 0,
            Scale       : 0,
            Primarykey  : false,
            Unique      : true,
            Index       : true,
            Null        : false,
            Blank       : false,
            Min         : "",
            Max         : "",
        },
        smith.Field{
            VerboseName : "创建时间",
            Name        : "CreateTime",
            Column      : "create_time",
            FieldType   : fieldtype.Create,
            Null        : false,

        },
        smith.Field{
            VerboseName : "修改时间",
            Name        : "UpdateTime",
            Column      : "update_time",
            FieldType   : fieldtype.Update,
            Null        : false,
        },
        smith.Field{
            VerboseName : "Version",
            Name        : "Version",
            Column      : "version",
            FieldType   : fieldtype.Version,
            Null        : false,
        },
    },
}
