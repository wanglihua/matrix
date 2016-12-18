package models

var KnowledgeCateCols = struct {
	Id         string
	Name       string
	Desc       string
	ParentId   string
	CreateTime string
	UpdateTime string
	Version    string
}{
	Id:         "id",
	Name:       "cate_name",
	Desc:       "cate_desc",
	ParentId:   "parent_id",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	Version:    "version",
}
