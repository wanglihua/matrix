package models

var EventCols = struct {
	Id string
	Code string
	TypeId string
	PriorityId string
	ApplyDepartmentId string
	ApplyUserId string
	Contact string
	EngineerId string
	AreaId string
	EquipmentTypeId string
	ApplicationTypeId string
	Location string
	Description string
	Solution string
	SrcEventId string
	StatusId string
	CreateTime string
	UpdateTime string
	Version string
}{
	Id: "id",
	Code: "code",
	TypeId: "type_id",
	PriorityId: "priority_id",
	ApplyDepartmentId: "apply_department_id",
	ApplyUserId: "apply_user_id",
	Contact: "contact",
	EngineerId: "engineer_id",
	AreaId: "area_id",
	EquipmentTypeId: "equipment_type_id",
	ApplicationTypeId: "application_type_id",
	Location: "location",
	Description: "description",
	Solution: "solution",
	SrcEventId: "src_event_id",
	StatusId: "status_id",
	CreateTime: "create_time",
	UpdateTime: "update_time",
	Version: "version",
}
