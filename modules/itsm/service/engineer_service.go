package service

import (
	"github.com/go-xorm/xorm"
	"matrix/core"
)

type EngineerService struct {
	
}

func (_ EngineerService) GetEngineerGroupIdList(db_session xorm.Session, engineer_id int64) []int64 {
	var group_list = make([]struct {
		GroupId int64 `xorm:"'group_id'"`
	}, 0)
	sql := `SELECT DISTINCT s.group_id AS group_id FROM hd_itsm_engineer e INNER JOIN hd_itsm_engineer_group_setting s ON e.id = s.engineer_id WHERE e.id = ?;`
	err := db_session.Sql(sql, engineer_id).Find(&group_list)
	core.HandleError(err)
	
	engineer_group_id_list := make([]int64, 0)
	for _, group := range group_list {
		engineer_group_id_list = append(engineer_group_id_list, group.GroupId)
	}
	return engineer_group_id_list
}

func (_ EngineerService) GetEngineerServiceAreaIdList(db_session xorm.Session, engineer_id int64) []int64 {
	var service_area_list = make([]struct {
		ServiceAreaId int64 `xorm:"'service_area_id'"`
	}, 0)
	sql := `SELECT DISTINCT a.service_area_id AS service_area_id FROM hd_itsm_engineer e INNER JOIN hd_itsm_engineer_service_area a ON e.id = a.engineer_id WHERE e.id = ?;`
	err := db_session.Sql(sql, engineer_id).Find(&service_area_list)
	core.HandleError(err)
	
	service_area_id_list := make([]int64, 0)
	for _, service_area := range service_area_list {
		service_area_id_list = append(service_area_id_list, service_area.ServiceAreaId)
	}
	return service_area_id_list
}

func (_ EngineerService) GetEngineerEventTypeIdList(db_session xorm.Session, engineer_id int64) []int64 {
	var event_type_list = make([]struct {
		EventTypeId int64 `xorm:"'event_type_id'"`
	}, 0)
	sql := `SELECT DISTINCT t.type_id AS event_type_id FROM hd_itsm_engineer e INNER JOIN hd_itsm_engineer_event_type t ON e.id = t.engineer_id WHERE e.id = ?;`
	err := db_session.Sql(sql, engineer_id).Find(&event_type_list)
	core.HandleError(err)
	
	event_type_id_list := make([]int64, 0)
	for _, event_type := range event_type_list {
		event_type_id_list = append(event_type_id_list, event_type.EventTypeId)
	}
	return event_type_id_list
}

func (_ EngineerService) GetEngineerIsManager(db_session xorm.Session, engineer_id int64) int {
	var manager_count struct {
		ManagerCount int `xorm:"'manager_count'"`
	}
	sql := `SELECT count(DISTINCT m.id) AS manager_count FROM hd_itsm_engineer e INNER JOIN hd_itsm_engineer_manager m ON e.id = m.engineer_id WHERE e.id = ?;`
	_, err := db_session.Sql(sql, engineer_id).Get(&manager_count)
	core.HandleError(err)
	if manager_count.ManagerCount == 0 {
		return 0
	} else {
		return 1
	}
}
