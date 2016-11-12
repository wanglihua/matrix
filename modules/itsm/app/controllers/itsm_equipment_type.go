package controllers

import (
	"github.com/revel/revel"
)

type ItsmEquipmentType struct {
	*revel.Controller
}

func (c ItsmEquipmentType) Index() revel.Result {
	return c.RenderTemplate("itsm/equipment_type/equipment_type_index.html")
}
