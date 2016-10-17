package models

import (
	"matrix/core"
)

var ModuleTitleName = "Inventory"
var ModuleLowerCase = "inventory"
var ModuleChinese = "仓储管理"

var TablePrefix = "hd_inventory_"


//---------------------------------------------------------------------------------------------------------------

type Supplier struct {
	Id          int64             `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	Name        string            `xorm:"nvarchar(255) notnull unique index 'sup_name'" json:"sup_name"`
	Code        string            `xorm:"nvarchar(100) notnull unique index 'sup_code'" json:"sup_code"`
	Contact     string            `xorm:"nvarchar(100) null 'contact'" json:"contact"`
	Phone       string            `xorm:"nvarchar(100) null 'phone'" json:"phone"`
	Address     string            `xorm:"nvarchar(300) null 'address'" json:"address"`
	Bank        string            `xorm:"nvarchar(255) null 'bank'" json:"bank"`
	BankAccount string            `xorm:"nvarchar(100) null 'bank_account'" json:"bank_account"`
	Remark      string            `xorm:"nvarchar(500) null 'remark'" json:"remark"`

	CreateTime  core.Time         `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime  core.Time         `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version     int               `xorm:"version notnull 'version'" json:"version"`
}

func (e Supplier) TableName() string {
	return TablePrefix + "supplier"
}

//---------------------------------------------------------------------------------------------------------------

type Client struct {
	Id          int64             `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	Name        string            `xorm:"nvarchar(255) notnull unique index 'client_name'" json:"client_name"`
	Code        string            `xorm:"nvarchar(100) notnull unique index 'client_code'" json:"client_code"`
	Contact     string            `xorm:"nvarchar(100) null 'contact'" json:"contact"`
	Phone       string            `xorm:"nvarchar(100) null 'phone'" json:"phone"`
	Address     string            `xorm:"nvarchar(300) null 'address'" json:"address"`
	Bank        string            `xorm:"nvarchar(255) null 'bank'" json:"bank"`
	BankAccount string            `xorm:"nvarchar(100) null 'bank_account'" json:"bank_account"`
	Remark      string            `xorm:"nvarchar(500) null 'remark'" json:"remark"`

	CreateTime  core.Time         `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime  core.Time         `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version     int               `xorm:"version notnull 'version'" json:"version"`
}

func (e Client) TableName() string {
	return TablePrefix + "client"
}

//---------------------------------------------------------------------------------------------------------------

type ProductCate struct {
	Id         int64             `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	Name       string            `xorm:"nvarchar(255) notnull unique index 'cate_name'" json:"cate_name" smith:"verbose_name=名称,min=2,max=100"`
	Desc       core.NullString   `xorm:"nvarchar(500) null 'cate_desc'" json:"cate_desc" smith:"verbose_name=描述,min=2,max=300"`

	CreateTime core.Time         `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time         `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int               `xorm:"version notnull 'version'" json:"version"`
}

func (e ProductCate) TableName() string {
	return TablePrefix + "product_cate"
}

func (e ProductCate) ModelDesc() string {
	return "货品类别"
}

//---------------------------------------------------------------------------------------------------------------

type Stock struct {
	Id         int64             `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	Code       string            `xorm:"nvarchar(100) notnull unique index 'stock_code'" json:"stock_code"`
	Name       string            `xorm:"nvarchar(255) notnull unique index 'stock_name'" json:"stock_name"`
	Address    string            `xorm:"nvarchar(300) null 'address'" json:"address"`
	Keeper     string            `xorm:"nvarchar(100) null 'keeper'" json:"keeper"`
	Phone      string            `xorm:"nvarchar(100) null 'phone'" json:"phone"`
	Remark     string            `xorm:"nvarchar(500) null 'remark'" json:"remark"`

	CreateTime core.Time         `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time         `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int               `xorm:"version notnull 'version'" json:"version"`
}

func (e Stock) TableName() string {
	return TablePrefix + "stock"
}

//---------------------------------------------------------------------------------------------------------------

type StorageLoc struct {
	Id         int64             `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	Code       string            `xorm:"nvarchar(100) notnull unique(uq_storage_loc_code) 'stock_code'" json:"stock_code"`
	Name       string            `xorm:"nvarchar(255) notnull unique(uq_storage_loc_name) 'stock_name'" json:"stock_name"`
	StockId    int64             `xorm:"bigint notnull unique(uq_storage_loc_code) unique(uq_storage_loc_name) index 'stock_id'" json:"stock_id"`
	Remark     string            `xorm:"nvarchar(500) null 'remark'" json:"remark"`

	CreateTime core.Time         `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time         `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int               `xorm:"version notnull 'version'" json:"version"`
}

func (e StorageLoc) TableName() string {
	return TablePrefix + "storage_loc"
}


