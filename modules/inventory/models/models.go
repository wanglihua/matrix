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
	Id          int64 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	Name        string `xorm:"nvarchar(255) notnull unique index 'sup_name'" json:"sup_name"`
	Code        string `xorm:"nvarchar(100) notnull unique index 'sup_code'" json:"sup_code"`
	Contact     string `xorm:"nvarchar(100) null 'contact'" json:"contact"`
	Phone       string `xorm:"nvarchar(100) null 'phone'" json:"phone"`
	Address     string `xorm:"nvarchar(300) null 'address'" json:"address"`
	Bank        string `xorm:"nvarchar(255) null 'bank'" json:"bank"`
	BankAccount string `xorm:"nvarchar(100) null 'bank_account'" json:"bank_account"`
	Remark      string `xorm:"nvarchar(500) null 'remark'" json:"remark"`

	CreateTime  core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime  core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version     int       `xorm:"version notnull 'version'" json:"version"`
}

func (e Supplier) TableName() string {
	return TablePrefix + "supplier"
}

//---------------------------------------------------------------------------------------------------------------

type Client struct {
	Id          int64 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	Name        string `xorm:"nvarchar(255) notnull unique index 'client_name'" json:"client_name"`
	Code        string `xorm:"nvarchar(100) notnull unique index 'client_code'" json:"client_code"`
	Contact     string `xorm:"nvarchar(100) null 'contact'" json:"contact"`
	Phone       string `xorm:"nvarchar(100) null 'phone'" json:"phone"`
	Address     string `xorm:"nvarchar(300) null 'address'" json:"address"`
	Bank        string `xorm:"nvarchar(255) null 'bank'" json:"bank"`
	BankAccount string `xorm:"nvarchar(100) null 'bank_account'" json:"bank_account"`
	Remark      string `xorm:"nvarchar(500) null 'remark'" json:"remark"`

	CreateTime  core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime  core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version     int       `xorm:"version notnull 'version'" json:"version"`
}

func (e Client) TableName() string {
	return TablePrefix + "client"
}

//---------------------------------------------------------------------------------------------------------------

type ProductCate struct {
	Id         int64 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	Name       string          `xorm:"nvarchar(255) notnull unique index 'cate_name'" json:"cate_name" smith:"verbose_name=名称,min=2,max=100"`
	Desc       core.NullString `xorm:"nvarchar(500) null 'cate_desc'" json:"cate_desc" smith:"verbose_name=描述,min=2,max=300"`

	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e ProductCate) TableName() string {
	return TablePrefix + "product_cate"
}

func (e ProductCate) ModelDesc() string {
	return "verbose_name=货品类别"
}


//---------------------------------------------------------------------------------------------------------------

type Product struct {
	Id              int64 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	Code            string          `xorm:"nvarchar(50) notnull unique index 'code'" json:"code" smith:"verbose_name=编码,min=2,max=20"`
	Name            string          `xorm:"nvarchar(255) notnull unique index 'name'" json:"name" smith:"verbose_name=品名,min=2,max=100"`
	Model           core.NullString `xorm:"nvarchar(255) null 'model'" json:"model" smith:"verbose_name=规格型号,min=2,max=200"`

	CurrentStock    core.NullInt    `xorm:"bigint null 'current_stock'" json:"current_stock" smith:"verbose_name=当前库存,min=0"`
	StockLowerLimit core.NullInt    `xorm:"bigint null 'stock_lower_limit'" json:"stock_lower_limit" smith:"verbose_name=库存下限,min=0"`
	StockUpperLimit core.NullInt    `xorm:"bigint null 'stock_upper_limit'" json:"stock_upper_limit" smith:"verbose_name=库存上限,min=0"`
	Unit            string          `xorm:"nvarchar(50) notnull 'unit'" json:"unit" smith:"verbose_name=单位,min=1,max=10"`
	CostPrice       core.NullInt    `xorm:"bigint null 'cost_price'" json:"cost_price" smith:"verbose_name=成本价,min=0"`
	RetailPrice     int    `xorm:"bigint null 'retail_price'" json:"retail_price" smith:"verbose_name=零售价,min=0"`
	State           string          `xorm:"nvarchar(10) notnull 'state'" json:"state" smith:"verbose_name=状态,min=1"`

	CreateTime      core.Time       `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime      core.Time       `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version         int             `xorm:"version notnull 'version'" json:"version"`
}

func (e Product) TableName() string {
	return TablePrefix + "product"
}

func (e Product) ModelDesc() string {
	return "verbose_name=货品信息"
}

//---------------------------------------------------------------------------------------------------------------

type Stock struct {
	Id         int64 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	Code       string `xorm:"nvarchar(100) notnull unique index 'stock_code'" json:"stock_code"`
	Name       string `xorm:"nvarchar(255) notnull unique index 'stock_name'" json:"stock_name"`
	Address    string `xorm:"nvarchar(300) null 'address'" json:"address"`
	Keeper     string `xorm:"nvarchar(100) null 'keeper'" json:"keeper"`
	Phone      string `xorm:"nvarchar(100) null 'phone'" json:"phone"`
	Remark     string `xorm:"nvarchar(500) null 'remark'" json:"remark"`

	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e Stock) TableName() string {
	return TablePrefix + "stock"
}

//---------------------------------------------------------------------------------------------------------------

type StorageLoc struct {
	Id         int64 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	Code       string `xorm:"nvarchar(100) notnull unique(uq_storage_loc_code) 'stock_code'" json:"stock_code"`
	Name       string `xorm:"nvarchar(255) notnull unique(uq_storage_loc_name) 'stock_name'" json:"stock_name"`
	StockId    int64  `xorm:"bigint notnull unique(uq_storage_loc_code) unique(uq_storage_loc_name) index 'stock_id'" json:"stock_id"`
	Remark     string `xorm:"nvarchar(500) null 'remark'" json:"remark"`

	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e StorageLoc) TableName() string {
	return TablePrefix + "storage_loc"
}
