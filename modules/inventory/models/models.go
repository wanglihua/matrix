
package models

import (
    "matrix/core"
)

var TablePrefix = "hd_inventory_"


//---------------------------------------------------------------------------------------------------------------

type Supplier struct {
    Id             int64             `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

    Name           string            `xorm:"nvarchar(255) notnull unique index 'sup_name'" json:"sup_name"`
    Code           string            `xorm:"nvarchar(100) notnull unique index 'sup_code'" json:"sup_code"`
    Contact        string            `xorm:"nvarchar(100) null 'contact'" json:"contact"`
    Phone          string            `xorm:"nvarchar(100) null 'phone'" json:"phone"`
    Address        string            `xorm:"nvarchar(300) null 'address'" json:"address"`
    Bank           string            `xorm:"nvarchar(255) null 'bank'" json:"bank"`
    BankAccount    string            `xorm:"nvarchar(100) null 'bank_account'" json:"bank_account"`
    Remark         string            `xorm:"nvarchar(500) null 'remark'" json:"remark"`

    CreateTime     core.Time         `xorm:"created notnull 'create_time'" json:"create_time"`
    UpdateTime     core.Time         `xorm:"updated notnull 'update_time'" json:"update_time"`
    Version        int               `xorm:"version notnull 'version'" json:"version"`
}

func (e Supplier) TableName() string {
    return TablePrefix + "supplier"
}


