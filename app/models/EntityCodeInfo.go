package models

type EntityCodeInfo struct {
	Entity string    `xorm:"nvarchar(200) notnull pk 'entity'" json:"entity" smith:"verbose_name=实体名称"`
	Serial int64     `xorm:"bigint notnull 'serial'" json:"serial" smith:"verbose_name=序列号"`
}

func (e EntityCodeInfo) TableName() string {
	return TablePrefix + "entity_code"
}

func (e EntityCodeInfo) ModelDesc() string {
	return "verbose_name=实体编码,entity_json=entity_code,route=entity_code"
}
