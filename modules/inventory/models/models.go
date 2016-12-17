package models

import (
	"matrix/core"
)

var ModuleTitleName = "Inventory"
var ModuleLowerCase = "inventory"
var ModuleChinese = "仓储管理"

var TablePrefix = "inventory_"

//---------------------------------------------------------------------------------------------------------------

type Supplier struct {
	Id int64 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	Name        string `xorm:"nvarchar(255) notnull unique index 'sup_name'" json:"sup_name"`
	Code        string `xorm:"nvarchar(100) notnull unique index 'sup_code'" json:"sup_code"`
	Contact     string `xorm:"nvarchar(100) null 'contact'" json:"contact"`
	Phone       string `xorm:"nvarchar(100) null 'phone'" json:"phone"`
	Address     string `xorm:"nvarchar(300) null 'address'" json:"address"`
	Bank        string `xorm:"nvarchar(255) null 'bank'" json:"bank"`
	BankAccount string `xorm:"nvarchar(100) null 'bank_account'" json:"bank_account"`
	Remark      string `xorm:"nvarchar(500) null 'remark'" json:"remark"`

	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e Supplier) TableName() string {
	return TablePrefix + "supplier"
}

//---------------------------------------------------------------------------------------------------------------

type Client struct {
	Id int64 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	Name        string `xorm:"nvarchar(255) notnull unique index 'client_name'" json:"client_name"`
	Code        string `xorm:"nvarchar(100) notnull unique index 'client_code'" json:"client_code"`
	Contact     string `xorm:"nvarchar(100) null 'contact'" json:"contact"`
	Phone       string `xorm:"nvarchar(100) null 'phone'" json:"phone"`
	Address     string `xorm:"nvarchar(300) null 'address'" json:"address"`
	Bank        string `xorm:"nvarchar(255) null 'bank'" json:"bank"`
	BankAccount string `xorm:"nvarchar(100) null 'bank_account'" json:"bank_account"`
	Remark      string `xorm:"nvarchar(500) null 'remark'" json:"remark"`

	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e Client) TableName() string {
	return TablePrefix + "client"
}

//---------------------------------------------------------------------------------------------------------------

type ProductCate struct {
	Id int64 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	Name string          `xorm:"nvarchar(255) notnull unique index 'cate_name'" json:"cate_name" smith:"verbose_name=名称,min=2,max=100"`
	Desc core.NullString `xorm:"nvarchar(500) null 'cate_desc'" json:"cate_desc" smith:"verbose_name=描述,max=300"`

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
	Id int64 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	CateId int64           `xorm:"bigint notnull index 'cate_id'" json:"cate_id" smith:"verbose_name=类别"` //对应于货品类别表中Id
	Code   string          `xorm:"nvarchar(50) notnull unique index 'code'" json:"code" smith:"verbose_name=编码,min=2,max=20"`
	Name   string          `xorm:"nvarchar(255) notnull unique index 'name'" json:"name" smith:"verbose_name=品名,min=2,max=100"`
	Model  core.NullString `xorm:"nvarchar(255) null 'model'" json:"model" smith:"verbose_name=规格型号,max=200"`

	CurrentStock    core.NullInt `xorm:"bigint null 'current_stock'" json:"current_stock" smith:"verbose_name=当前库存,min=0"`
	StockLowerLimit core.NullInt `xorm:"bigint null 'stock_lower_limit'" json:"stock_lower_limit" smith:"verbose_name=库存下限,min=0"`
	StockUpperLimit core.NullInt `xorm:"bigint null 'stock_upper_limit'" json:"stock_upper_limit" smith:"verbose_name=库存上限,min=0"`
	Unit            string       `xorm:"nvarchar(50) notnull 'unit'" json:"unit" smith:"verbose_name=单位,min=1,max=10"`
	CostPrice       core.NullInt `xorm:"bigint null 'cost_price'" json:"cost_price" smith:"verbose_name=成本价,min=0"`
	RetailPrice     int          `xorm:"bigint notnull 'retail_price'" json:"retail_price" smith:"verbose_name=零售价,min=0"`
	State           string       `xorm:"nvarchar(10) notnull 'state'" json:"state" smith:"verbose_name=状态,min=1"`

	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e Product) TableName() string {
	return TablePrefix + "product"
}

func (e Product) ModelDesc() string {
	return "verbose_name=货品信息,entity_json=product"
}

//---------------------------------------------------------------------------------------------------------------

type Stock struct {
	Id int64 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	Code    string `xorm:"nvarchar(100) notnull unique index 'stock_code'" json:"stock_code"`
	Name    string `xorm:"nvarchar(255) notnull unique index 'stock_name'" json:"stock_name"`
	Address string `xorm:"nvarchar(300) null 'address'" json:"address"`
	Keeper  string `xorm:"nvarchar(100) null 'keeper'" json:"keeper"`
	Phone   string `xorm:"nvarchar(100) null 'phone'" json:"phone"`
	Remark  string `xorm:"nvarchar(500) null 'remark'" json:"remark"`

	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e Stock) TableName() string {
	return TablePrefix + "stock"
}

//---------------------------------------------------------------------------------------------------------------

type StorageLoc struct {
	Id int64 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	Code    string `xorm:"nvarchar(100) notnull unique(uq_storage_loc_code) 'stock_code'" json:"stock_code"`
	Name    string `xorm:"nvarchar(255) notnull unique(uq_storage_loc_name) 'stock_name'" json:"stock_name"`
	StockId int64  `xorm:"bigint notnull unique(uq_storage_loc_code) unique(uq_storage_loc_name) index 'stock_id'" json:"stock_id"`
	Remark  string `xorm:"nvarchar(500) null 'remark'" json:"remark"`

	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e StorageLoc) TableName() string {
	return TablePrefix + "storage_loc"
}

//---------------------------------------------------------------------------------------------------------------

type CapitalAccount struct {
	Id int64 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	Name   string          `xorm:"nvarchar(255) notnull unique 'account_name'" json:"account_name" smith:"verbose_name=名称,min=2,max=200"`
	Remark core.NullString `xorm:"nvarchar(500) null 'remark'" json:"remark" smith:"verbose_name=描述,max=300"`

	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e CapitalAccount) TableName() string {
	return TablePrefix + "capital_account"
}

func (e CapitalAccount) ModelDesc() string {
	return "verbose_name=资金账户,entity_json=ca,route=capital/account"
}

//---------------------------------------------------------------------------------------------------------------

//stock in out type
type StockIoType struct {
	Id int64 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	Cate   string          `xorm:"nvarchar(50) notnull 'type_cate'" json:"type_cate" smith:"verbose_name=类别,min=2,max=20"` //出库 或者 入库
	Name   string          `xorm:"nvarchar(255) notnull unique 'type_name'" json:"type_name" smith:"verbose_name=名称,min=2,max=200"`
	Code   string          `xorm:"nvarchar(50) notnull unique 'type_code'" json:"type_code" smith:"verbose_name=代码,min=2,max=20"`
	Footer core.NullString `xorm:"nvarchar(255) null 'bill_footer'" json:"bill_footer" smith:"verbose_name=单据页脚,max=200"`
	RecPay core.NullBool   `xorm:"nvarchar(255) null 'rec_pay'" json:"rec_pay" smith:"verbose_name=形成应收应付"`
	State  string          `xorm:"nvarchar(255) notnull 'type_state'" json:"type_state" smith:"verbose_name=状态,min=1"`

	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e StockIoType) TableName() string {
	return TablePrefix + "stock_io_type"
}

func (e StockIoType) ModelDesc() string {
	return "verbose_name=出入库类型,entity_json=iotype,route=stock/io/type"
}

//---------------------------------------------------------------------------------------------------------------

//payment type
type PayType struct {
	Id int64 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	Cate   string          `xorm:"nvarchar(50) notnull 'type_cate'" json:"type_cate" smith:"verbose_name=类别,min=2,max=20"` //收款 或者 付款
	Name   string          `xorm:"nvarchar(255) notnull unique 'type_name'" json:"type_name" smith:"verbose_name=名称,min=2,max=200"`
	Prefix string          `xorm:"nvarchar(20) notnull unique 'prefix'" json:"prefix" smith:"verbose_name=单据前缀,min=2,max=10"`
	Remark core.NullString `xorm:"nvarchar(500) null 'remark'" json:"remark" smith:"verbose_name=描述,max=300"`
	State  string          `xorm:"nvarchar(255) notnull 'type_state'" json:"type_state" smith:"verbose_name=状态,min=1"`

	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e PayType) TableName() string {
	return TablePrefix + "pay_type"
}

func (e PayType) ModelDesc() string {
	return "verbose_name=收付款类型,entity_json=paytype,route=pay/type"
}

//---------------------------------------------------------------------------------------------------------------

type Handler struct {
	Id int64 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	Name  string          `xorm:"nvarchar(255) notnull unique index 'handler_name'" json:"handler_name" smith:"verbose_name=名称,min=2,max=100"`
	Desc  core.NullString `xorm:"nvarchar(500) null 'handler_desc'" json:"handler_desc" smith:"verbose_name=描述,max=300"`
	State string          `xorm:"nvarchar(255) notnull 'handler_state'" json:"handler_state" smith:"verbose_name=状态,min=1"`

	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e Handler) TableName() string {
	return TablePrefix + "handler"
}

func (e Handler) ModelDesc() string {
	return "verbose_name=经手人,entity_json=handler,route=handler"
}

//---------------------------------------------------------------------------------------------------------------

type StockBill struct {
	Id int64 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	Code      string          `xorm:"nvarchar(255) notnull unique index 'bill_code'" json:"bill_code" smith:"verbose_name=出入库单号,min=2,max=100"`
	Date      core.Date       `xorm:"datetime notnull 'bill_date'" json:"bill_date" smith:"verbose_name=出入库日期"`
	TypeId    int64           `xorm:"bigint notnull 'type_id'" json:"type_id" smith:"verbose_name=出入库类型"`     //对应于出入库类型表中Id
	StockId   int64           `xorm:"bigint notnull 'stock_id'" json:"stock_id" smith:"verbose_name=仓库"`      //对应于仓库表中Id
	HandlerId int64           `xorm:"bigint notnull 'handler_id'" json:"handler_id" smith:"verbose_name=经手人"` //对应于经手人表中Id
	ObjectId  int64           `xorm:"bigint notnull 'object_id'" json:"object_id" smith:"verbose_name=往来对象"`  //供应商或客户表中Id
	Remark    core.NullString `xorm:"nvarchar(500) null 'remark'" json:"remark" smith:"verbose_name=备注,max=300"`
	State     string          `xorm:"nvarchar(255) notnull 'bill_state'" json:"bill_state" smith:"verbose_name=状态,min=1"`

	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e StockBill) TableName() string {
	return TablePrefix + "stock_bill"
}

func (e StockBill) ModelDesc() string {
	return "verbose_name=出入库单,entity_json=bill,route=stock/bill"
}

//---------------------------------------------------------------------------------------------------------------

type StockBillDetail struct {
	Id int64 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	BillId   int64           `xorm:"bigint notnull 'bill_id'" json:"bill_id" smith:"verbose_name=出入库单"`
	Code     string          `xorm:"nvarchar(50) notnull index 'code'" json:"code" smith:"verbose_name=编码,min=2,max=20"`
	Name     string          `xorm:"nvarchar(255) notnull index 'name'" json:"name" smith:"verbose_name=品名,min=2,max=100"`
	Model    core.NullString `xorm:"nvarchar(255) null 'model'" json:"model" smith:"verbose_name=规格型号,max=200"`
	Unit     string          `xorm:"nvarchar(50) notnull 'unit'" json:"unit" smith:"verbose_name=单位,min=1,max=10"`
	Price    int             `xorm:"bigint notnull 'price'" json:"price" smith:"verbose_name=价格,min=0"`
	Quantity int             `xorm:"bigint notnull 'quantity'" json:"quantity" smith:"verbose_name=数量,min=0"`
	Amount   int             `xorm:"bigint notnull 'amount'" json:"amount" smith:"verbose_name=金额,min=0"`
	Batch    core.NullString `xorm:"nvarchar(100) null 'batch'" json:"batch" smith:"verbose_name=批号,max=50"`
	Remark   core.NullString `xorm:"nvarchar(500) null 'remark'" json:"remark" smith:"verbose_name=备注,max=300"`

	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e StockBillDetail) TableName() string {
	return TablePrefix + "stock_bill_detail"
}

func (e StockBillDetail) ModelDesc() string {
	return "verbose_name=出入库单详细,entity_json=billdetail,route=stock/bill/detail"
}

type Config struct {
	Id int64 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`

	BillInAutoApprove  core.NullInt `xorm:"int null 'bill_in_auto_approve'" json:"bill_in_auto_approve" smith:"verbose_name=入库单自动审核"`
	BillOutAutoApprove core.NullInt `xorm:"int null 'bill_out_auto_approve'" json:"bill_out_auto_approve" smith:"verbose_name=出库单自动审核"`
	PayInAutoApprove   core.NullInt `xorm:"int null 'pay_in_auto_approve'" json:"pay_in_auto_approve" smith:"verbose_name=收款单自动审核"`
	PayOutAutoApprove  core.NullInt `xorm:"int null 'pay_out_auto_approve'" json:"pay_out_auto_approve" smith:"verbose_name=付款单自动审核"`

	CreateTime core.Time `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int       `xorm:"version notnull 'version'" json:"version"`
}

func (e Config) TableName() string {
	return TablePrefix + "config"
}

func (e Config) ModelDesc() string {
	return "verbose_name=系统配置,entity_json=config,route=config"
}
