package models

type EventTypeCols struct {
	Id         string
	Name       string
	Desc       string
	CreateTime string
	UpdateTime string
	Version    string
}

func (c EventTypeCols) init() {
	c.Id = "id"
	c.Name = "type_name"
	c.Desc = "type_desc"
	c.CreateTime = "create_time"
	c.UpdateTime = "update_time"
	c.Version = "version"
}

var event_type_cols EventTypeCols

func GetEventTypeCols() EventTypeCols {
	if event_type_cols.Id == "" {
		event_type_cols.init()
	}
	return event_type_cols
}
