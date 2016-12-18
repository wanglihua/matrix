package models

type event_type_cols_type struct {
	Id         string
	Name       string
	Desc       string
	CreateTime string
	UpdateTime string
	Version    string
}

var EventTypeCols event_type_cols_type

func init() {
	EventTypeCols.Id = "id"
	EventTypeCols.Name = "type_name"
	EventTypeCols.Desc = "type_desc"
	EventTypeCols.CreateTime = "create_time"
	EventTypeCols.UpdateTime = "update_time"
	EventTypeCols.Version = "version"
}
