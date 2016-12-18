package models

var EngineerServiceAreaCols = struct {
	Id            string
	ServiceAreaId string
	EngineerId    string
	CreateTime    string
	UpdateTime    string
	Version       string
}{
	Id:            "id",
	ServiceAreaId: "service_area_id",
	EngineerId:    "engineer_id",
	CreateTime:    "create_time",
	UpdateTime:    "update_time",
	Version:       "version",
}
