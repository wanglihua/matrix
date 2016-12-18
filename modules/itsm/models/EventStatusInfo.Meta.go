package models

var EventStatusCols = struct {
	Id         string
	Name       string
	Desc       string
	CreateTime string
	UpdateTime string
	Version    string
}{
	Id:         "id",
	Name:       "status_name",
	Desc:       "status_desc",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	Version:    "version",
}
