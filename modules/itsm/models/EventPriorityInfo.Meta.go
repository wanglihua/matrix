package models

var EventPriorityCols = struct {
	Id         string
	Name       string
	Desc       string
	CreateTime string
	UpdateTime string
	Version    string
}{
	Id:         "id",
	Name:       "priority_name",
	Desc:       "priority_desc",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	Version:    "version",
}
