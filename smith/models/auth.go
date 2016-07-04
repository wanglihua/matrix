package models

import (
    "matrix/smith"
    "matrix/smith/fieldtype"
)

var GroupEntity = smith.Entity{
    ModuleTitleName : "",
    ModuleLowerCase: "",
    ModuleChinese: "",
    EntityTitleName:"",
    EntityLowerCase:"",
    EntityChinese :"",
    TableName       :"",
    TablePrefix     :"",

    FieldList: []smith.Field{
        smith.Field{
            VerboseName : "",
            Name        : "",
            Column      : "",
            FieldType   : fieldtype.BigInt,
            Primarykey  : true,
            Unique      : true,
            Null        : false,
            Blank       : false,
            Min         : "",
            Max         : "",
            Precision   : 0,
            Scale       : 0,
        },
        smith.Field{
            VerboseName : "",
            Name        : "",
            Column      : "",
            FieldType   : fieldtype.BigInt,
            Primarykey  : true,
            Unique      : true,
            Null        : false,
            Blank       : false,
            Min         : "",
            Max         : "",
            Precision   : 0,
            Scale       : 0,
        },
    },
}
