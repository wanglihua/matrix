package models

var EventLogCols = struct {
	Id         string
	EventId    string
	Reason     string
	Remark     string
	CreateTime string
}{
	Id:         "id",
	EventId:    "event_id",
	Reason:     "reason",
	Remark:     "remark",
	CreateTime: "create_time",
}
