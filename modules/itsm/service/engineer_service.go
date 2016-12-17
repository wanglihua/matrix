package service

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"matrix/core"
	itsm_models "matrix/modules/itsm/models"
	"strings"
)

type EngineerService struct {
}

func (_ EngineerService) GetEngineerGroupIdList(db_session *xorm.Session, engineer_id int64) []int64 {
	var group_list = make([]struct {
		GroupId int64 `xorm:"'group_id'"`
	}, 0)
	sql := `SELECT DISTINCT s.group_id AS group_id FROM itsm_engineer e INNER JOIN itsm_engineer_group_setting s ON e.id = s.engineer_id WHERE e.id = ?;`
	err := db_session.Sql(sql, engineer_id).Find(&group_list)
	core.HandleError(err)

	engineer_group_id_list := make([]int64, 0)
	for _, group := range group_list {
		engineer_group_id_list = append(engineer_group_id_list, group.GroupId)
	}
	return engineer_group_id_list
}

func (_ EngineerService) GetEngineerServiceAreaIdList(db_session *xorm.Session, engineer_id int64) []int64 {
	var service_area_list = make([]struct {
		ServiceAreaId int64 `xorm:"'service_area_id'"`
	}, 0)
	sql := `SELECT DISTINCT a.service_area_id AS service_area_id FROM itsm_engineer e INNER JOIN itsm_engineer_service_area a ON e.id = a.engineer_id WHERE e.id = ?;`
	err := db_session.Sql(sql, engineer_id).Find(&service_area_list)
	core.HandleError(err)

	service_area_id_list := make([]int64, 0)
	for _, service_area := range service_area_list {
		service_area_id_list = append(service_area_id_list, service_area.ServiceAreaId)
	}
	return service_area_id_list
}

func (_ EngineerService) GetEngineerEventTypeIdList(db_session *xorm.Session, engineer_id int64) []int64 {
	var event_type_list = make([]struct {
		EventTypeId int64 `xorm:"'event_type_id'"`
	}, 0)
	sql := `SELECT DISTINCT t.type_id AS event_type_id FROM itsm_engineer e INNER JOIN itsm_engineer_event_type t ON e.id = t.engineer_id WHERE e.id = ?;`
	err := db_session.Sql(sql, engineer_id).Find(&event_type_list)
	core.HandleError(err)

	event_type_id_list := make([]int64, 0)
	for _, event_type := range event_type_list {
		event_type_id_list = append(event_type_id_list, event_type.EventTypeId)
	}
	return event_type_id_list
}

func (_ EngineerService) GetEngineerIsManager(db_session *xorm.Session, engineer_id int64) int {
	var manager_count struct {
		ManagerCount int `xorm:"'manager_count'"`
	}
	sql := `SELECT count(DISTINCT m.id) AS manager_count FROM itsm_engineer e INNER JOIN itsm_engineer_manager m ON e.id = m.engineer_id WHERE e.id = ?;`
	_, err := db_session.Sql(sql, engineer_id).Get(&manager_count)
	core.HandleError(err)
	if manager_count.ManagerCount == 0 {
		return 0
	} else {
		return 1
	}
}

func (_ EngineerService) GetIntListInThisNotInThat(this_list []int64, that_list []int64) []int64 {
	result_list := make([]int64, 0)
	for _, this_int := range this_list {
		not_in_that := true
		for _, that_int := range that_list {
			if this_int == that_int {
				not_in_that = false
			}
		}
		if not_in_that {
			result_list = append(result_list, this_int)
		}
	}
	return result_list
}

func (_ EngineerService) AddEngineerToGroups(db_session *xorm.Session, engineer_id int64, group_id_list []int64) {
	if group_id_list == nil || len(group_id_list) == 0 {
		return
	}
	egineer_group_setting_list := make([]itsm_models.EngineerGroupSetting, 0)
	for _, group_id := range group_id_list {
		var engineer_group_setting itsm_models.EngineerGroupSetting
		engineer_group_setting.EngineerId = engineer_id
		engineer_group_setting.GroupId = group_id
		egineer_group_setting_list = append(egineer_group_setting_list, engineer_group_setting)
	}
	db_session.Insert(&egineer_group_setting_list)
}

func (_ EngineerService) RemoveEngineerFromGroups(db_session *xorm.Session, engineer_id int64, group_id_list []int64) {
	if group_id_list == nil || len(group_id_list) == 0 {
		return
	}
	sql := `
DELETE FROM itsm_engineer_group_setting WHERE engineer_id = ? AND group_id IN (%s)
`
	group_id_str_list := make([]string, len(group_id_list))
	for _, group_id := range group_id_list {
		group_id_str_list = append(group_id_str_list, fmt.Sprint(group_id))
	}
	sql = fmt.Sprintf(sql, strings.Join(group_id_str_list, ","))
	db_session.Exec(sql, engineer_id)
}

func (_ EngineerService) AddEngineerToServiceAreas(db_session *xorm.Session, engineer_id int64, service_area_id_list []int64) {
	if service_area_id_list == nil || len(service_area_id_list) == 0 {
		return
	}
	egineer_service_area_list := make([]itsm_models.EngineerServiceArea, 0)
	for _, service_area_id := range service_area_id_list {
		var engineer_service_area itsm_models.EngineerServiceArea
		engineer_service_area.EngineerId = engineer_id
		engineer_service_area.ServiceAreaId = service_area_id
		egineer_service_area_list = append(egineer_service_area_list, engineer_service_area)
	}
	db_session.Insert(&egineer_service_area_list)
}

func (_ EngineerService) RemoveEngineerFromServiceAreas(db_session *xorm.Session, engineer_id int64, service_area_id_list []int64) {
	if service_area_id_list == nil || len(service_area_id_list) == 0 {
		return
	}
	sql := `
DELETE FROM itsm_engineer_service_area WHERE engineer_id = ? AND service_area_id IN (%s)
`
	service_area_id_str_list := make([]string, len(service_area_id_list))
	for _, service_area_id := range service_area_id_list {
		service_area_id_str_list = append(service_area_id_str_list, fmt.Sprint(service_area_id))
	}
	sql = fmt.Sprintf(sql, strings.Join(service_area_id_str_list, ","))
	db_session.Exec(sql, engineer_id)
}

func (_ EngineerService) AddEngineerToEventTypes(db_session *xorm.Session, engineer_id int64, event_type_id_list []int64) {
	if event_type_id_list == nil || len(event_type_id_list) == 0 {
		return
	}
	egineer_event_type_list := make([]itsm_models.EngineerEventType, 0)
	for _, event_type_id := range event_type_id_list {
		var engineer_event_type itsm_models.EngineerEventType
		engineer_event_type.EngineerId = engineer_id
		engineer_event_type.EventTypeId = event_type_id
		egineer_event_type_list = append(egineer_event_type_list, engineer_event_type)
	}
	db_session.Insert(&egineer_event_type_list)
}

func (_ EngineerService) RemoveEngineerFromEventTypes(db_session *xorm.Session, engineer_id int64, event_type_id_list []int64) {
	if event_type_id_list == nil || len(event_type_id_list) == 0 {
		return
	}
	sql := `
DELETE FROM itsm_engineer_event_type WHERE engineer_id = ? AND type_id IN (%s);
`
	event_type_id_str_list := make([]string, len(event_type_id_list))
	for _, event_type_id := range event_type_id_list {
		event_type_id_str_list = append(event_type_id_str_list, fmt.Sprint(event_type_id))
	}
	sql = fmt.Sprintf(sql, strings.Join(event_type_id_str_list, ","))
	db_session.Exec(sql, engineer_id)
}

func (_ EngineerService) IsEngineerManager(db_session *xorm.Session, engineer_id int64) bool {
	sql := `
SELECT count(*) as count FROM itsm_engineer_manager WHERE engineer_id = ?
`
	var count_result struct {
		Count int64 `xorm:"'count'"`
	}
	_, err := db_session.Sql(sql, engineer_id).Get(&count_result)
	core.HandleError(err)
	return count_result.Count != 0
}

func (_ EngineerService) AddEngineerToManager(db_session *xorm.Session, engineer_id int64) {
	var engineer_manager itsm_models.EngineerManager
	engineer_manager.EngineerId = engineer_id
	db_session.Insert(&engineer_manager)
}

func (_ EngineerService) RemoveEngineerFromManager(db_session *xorm.Session, engineer_id int64) {
	sql := `
DELETE FROM itsm_engineer_manager WHERE engineer_id = ?
`
	db_session.Exec(sql, engineer_id)
}
