package models

var EventTypeCols = struct {
	Id         string
	Name       string
	Desc       string
	CreateTime string
	UpdateTime string
	Version    string
}{
	Id:         "id",
	Name:       "type_name",
	Desc:       "type_desc",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	Version:    "version",
}