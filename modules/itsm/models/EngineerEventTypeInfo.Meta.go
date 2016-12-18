package models

var EngineerEventTypeCols = struct {
	Id          string
	EventTypeId string
	EngineerId  string
	CreateTime  string
	UpdateTime  string
	Version     string
}{
	Id:          "id",
	EventTypeId: "type_id",
	EngineerId:  "engineer_id",
	CreateTime:  "create_time",
	UpdateTime:  "update_time",
	Version:     "version",
}
