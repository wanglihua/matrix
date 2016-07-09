package auth

import (
    "matrix/smith"
    "matrix/smith/fieldtype"
)

var GroupUserEntity = smith.Entity{
    ModuleTitleName     : "Auth",
    ModuleLowerCase     : "auth",
    ModuleChinese       : "权限管理",
    EntityTitleName     : "GroupUser",
    EntityCamelCase     : "groupUser",
    EntityChinese       : "群组用户",
    TableName           : "group_user",
    TablePrefix         : "hd_auth_",

    FieldList: []smith.Field{
        smith.Field{
            VerboseName : "群组",
            Name        : "GroupId",
            Column      : "group_id",
            FieldType   : fieldtype.BigInt,
            Length      : 0,
            Precision   : 0,
            Scale       : 0,
            Unique      : "false",
            Index       : true,
            Null        : false,
            Blank       : false,
            Min         : "",
            Max         : "",
        },
        smith.Field{
            VerboseName : "用户",
            Name        : "UserId",
            Column      : "user_id",
            FieldType   : fieldtype.BigInt,
            Length      : 0,
            Precision   : 0,
            Scale       : 0,
            Unique      : "false",
            Index       : true,
            Null        : false,
            Blank       : false,
            Min         : "",
            Max         : "",
        },
    },
}
