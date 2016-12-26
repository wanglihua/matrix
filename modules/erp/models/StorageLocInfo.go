package models

import (
	"matrix/core"
)

type StorageLocInfo struct {
	Id         int64                 `xorm:"bigint notnull pk autoincr 'id'" json:"id"`
	StockId    int64                 `xorm:"bigint notnull index 'stock_id'" json:"stock_id" smith:"verbose_name=仓库"`
	Name       string                `xorm:"nvarchar(200) notnull index 'name'" json:"name" smith:"verbose_name=名称,min=2,max=100"`
	Desc       core.NullString       `xorm:"nvarchar(500) 'desc'" json:"desc" smith:"verbose_name=描述,max=500"`
	XAxes      string                `xorm:"nvarchar(100) notnull 'x_axes'" json:"x_axes" smith:"verbose_name=X轴坐标,max=100"`
	YAxes      string                `xorm:"nvarchar(100) notnull 'y_axes'" json:"y_axes" smith:"verbose_name=Y轴坐标,max=100"`
	CreateTime core.Time             `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time             `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int                   `xorm:"version notnull 'version'" json:"version"`
}

func (e StorageLocInfo) TableName() string {
	return TablePrefix + "storage_loc"
}

func (e StorageLocInfo) ModelDesc() string {
	return "verbose_name=仓位,entity_json=storage_loc,route=storage/loc"
}
