package models

import (
	"matrix/core"
)

type ProductInfo struct {
	Id            int64                 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	Sku           string                `xorm:"nvarchar(200) notnull index 'sku'" json:"sku" smith:"verbose_name=SKU,min=2,max=100"`
	BrandId       core.NullInt          `xorm:"bigint 'brand_id'" json:"brand_id" smith:"verbose_name=品牌"`
	Name          string                `xorm:"nvarchar(200) notnull index 'name'" json:"name" smith:"verbose_name=名称,min=2,max=100"`
	EnglishName   string                `xorm:"nvarchar(200) notnull index 'english_name'" json:"english_name" smith:"verbose_name=英文名称,min=2,max=100"`
	RetailPrice   float64               `xorm:"float notnull 'retail_price'" json:"retail_price" smith:"verbose_name=零售价,min=0"`
	CostPrice     float64               `xorm:"float notnull 'cost_price'" json:"cost_price" smith:"verbose_name=成本价,min=0"`
	Quantity      int64                 `xorm:"bigint notnull 'quantity'" json:"quantity" smith:"verbose_name=库存,min=0"`
	Weight        float64               `xorm:"float notnull 'weight'" json:"weight" smith:"verbose_name=重量,min=0"`
	StockId       core.NullInt          `xorm:"bigint 'stock_id'" json:"stock_id" smith:"verbose_name=仓库"`
	StorageLocId  core.NullInt          `xorm:"bigint 'storage_loc_id'" json:"storage_loc_id" smith:"verbose_name=仓位"`
	ProducingArea core.NullString       `xorm:"nvarchar(200) 'producing_area'" json:"producing_area" smith:"verbose_name=产地,max=100"`
	Desc          core.NullString       `xorm:"nvarchar(500) 'desc'" json:"desc" smith:"verbose_name=描述,max=500"`
	StatusId      int64                 `xorm:"bigint notnull index 'status_id'" json:"status_id" smith:"verbose_name=状态"`
	CreateTime    core.Time             `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime    core.Time             `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version       int                   `xorm:"version notnull 'version'" json:"version"`
}

func (e ProductInfo) TableName() string {
	return TablePrefix + "product"
}

func (e ProductInfo) ModelDesc() string {
	return "verbose_name=商品,entity_json=product,route=product"
}
