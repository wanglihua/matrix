package models

var EventLogCols = struct {
	Id         string
	EventId    string
	Reason     string
	Remark     string
	CreateTime string
	UpdateTime string
	Version    string
}{
	Id:         "id",
	EventId:    "event_id",
	Reason:     "reason",
	Remark:     "remark",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	Version:    "version",
}
