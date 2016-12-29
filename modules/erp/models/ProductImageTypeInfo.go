package models

import (
	"matrix/core"
	"github.com/go-xorm/xorm"
)

const (
	PRODUCT_IMAGE_TYPE_BIG_ID = 100
)

type ProductImageTypeInfo struct {
	Id         int64                 `xorm:"bigint notnull pk 'id'" json:"id"`
	Name       string                `xorm:"nvarchar(200) notnull index 'name'" json:"name" smith:"verbose_name=名称,min=2,max=100"`
	Desc       core.NullString       `xorm:"nvarchar(500) 'desc'" json:"desc" smith:"verbose_name=描述,min=2,max=400"`
	CreateTime core.Time             `xorm:"created notnull 'create_time'" json:"create_time"`
	UpdateTime core.Time             `xorm:"updated notnull 'update_time'" json:"update_time"`
	Version    int                   `xorm:"version notnull 'version'" json:"version"`
}

func (e ProductImageTypeInfo) TableName() string {
	return TablePrefix + "product_image_type"
}

func (e ProductImageTypeInfo) ModelDesc() string {
	return "verbose_name=商品图片类型,entity_json=product_image_type,route=product/image/type"
}

func (e ProductImageTypeInfo) InitData(db_session *xorm.Session) {
	//商品图片类型

	product_image_type := new(ProductImageTypeInfo)
	product_image_type.Id = PRODUCT_IMAGE_TYPE_BIG_ID
	product_image_type.Name = "大图"
	product_image_type.Desc = core.NewNullString("图片类型 大图", true)
	db_session.Insert(product_image_type)
}
