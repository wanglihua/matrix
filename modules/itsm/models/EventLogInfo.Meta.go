package models

type event_log_cols_type struct {
	Id         string
	EventId    string
	Reason     string
	Remark     string
	CreateTime string
	UpdateTime string
	Version    string
}

var EventLogCols event_log_cols_type

func init() {
	EventLogCols.Id = "id"
	EventLogCols.EventId = "event_id"
	EventLogCols.Reason = "reason"
	EventLogCols.Remark = "remark"
	EventLogCols.CreateTime = "create_time"
	EventLogCols.UpdateTime = "update_time"
	EventLogCols.Version = "version"
}
