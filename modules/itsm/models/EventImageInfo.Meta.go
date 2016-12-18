package models

var EventImageCols = struct {
	Id         string
	EventId    string
	ImageUrl   string
	CreateTime string
	UpdateTime string
	Version    string
}{
	Id:         "id",
	EventId:    "event_id",
	ImageUrl:   "image_url",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	Version:    "version",
}
