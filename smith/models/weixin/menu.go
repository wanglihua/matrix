package weixin

import (
    "matrix/smith"
    "matrix/smith/fieldtype"
)

var MenuEntity = smith.Entity{
    ModuleTitleName     : "Weixin",
    ModuleLowerCase     : "weixin",
    ModuleChinese       : "微信管理",
    EntityTitleName     : "Menu",
    EntityCamelCase     : "menu",
    EntityChinese       : "微信菜单",
    TableName           : "menu",
    TablePrefix         : "hd_weixin_",

    FieldList: []smith.Field{
        smith.Field{
            VerboseName : "菜单名",
            Name        : "Name",
            Column      : "menu_name",
            FieldType   : fieldtype.NVarchar,
            Length      : 30,
            Precision   : 0,
            Scale       : 0,
            Unique      : "false",
            Index       : false,
            Null        : false,
            Blank       : false,
            Min         : "0",
            Max         : "12",
        },
        smith.Field{
            VerboseName : "类型",
            Name        : "Type",
            Column      : "menu_type",
            FieldType   : fieldtype.NVarchar,
            Length      : 30,
            Precision   : 0,
            Scale       : 0,
            Unique      : "false",
            Index       : false,
            Null        : true,
            Blank       : true,
            Min         : "0",
            Max         : "12",
        },
        smith.Field{
            VerboseName : "菜单数据",
            Name        : "Data",
            Column      : "menu_data",
            FieldType   : fieldtype.NVarchar,
            Length      : 300,
            Precision   : 0,
            Scale       : 0,
            Unique      : "false",
            Index       : false,
            Null        : true,
            Blank       : true,
            Min         : "2",
            Max         : "250",
        },
        smith.Field{
            VerboseName : "层级",
            Name        : "Level",
            Column      : "menu_level",
            FieldType   : fieldtype.Int,
            Length      : 0,
            Precision   : 0,
            Scale       : 0,
            Unique      : "false",
            Index       : false,
            Null        : false,
            Blank       : false,
            Min         : "1",
            Max         : "2",
        },
        smith.Field{
            VerboseName : "排序",
            Name        : "Order",
            Column      : "menu_order",
            FieldType   : fieldtype.Int,
            Length      : 0,
            Precision   : 0,
            Scale       : 0,
            Unique      : "false",
            Index       : false,
            Null        : false,
            Blank       : false,
            Min         : "",
            Max         : "",
        },
    },
}
