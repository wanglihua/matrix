package controllers

import (
	"encoding/json"
	"github.com/revel/revel"
	"strconv"

	"fmt"
	"matrix/core"
	auth_models "matrix/modules/auth/models"
	"matrix/modules/itsm/models"
	"matrix/modules/itsm/service/engineer_service"
)

type ItsmEngineer struct {
	*revel.Controller
}

func (c ItsmEngineer) Index() revel.Result {
	return c.RenderTemplate("itsm/engineer/engineer_index.html")
}

type EngineerViewItem struct {
	auth_models.UserInfo `xorm:"extends" json:"u"`
	models.EngineerInfo  `xorm:"extends" json:"e"`
}

func (c ItsmEngineer) ListData() revel.Result {
	db_session := c.DbSession

	filter, order, offset, limit := core.GetGridRequestParam(c.Request)
	query := db_session.
		Select("u.*, e.*").
		Table(models.TablePrefix+"engineer").Alias("e").
		Join("inner", []string{auth_models.TablePrefix + "user", "u"}, "u.id=e.user_id").
		Where(filter)

	//query extra filter here

	data_query := *query
	if order != "" {
		data_query = *data_query.OrderBy(order)
	} else {
		data_query = *data_query.Asc("e.id")
	}

	engineer_list := make([]EngineerViewItem, 0, limit)
	err := data_query.Limit(limit, offset).Find(&engineer_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(EngineerViewItem))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  engineer_list,
		Total: count,
	})
}

func (c ItsmEngineer) AddListData() revel.Result {
	session := c.DbSession

	filter, order, offset, limit := core.GetGridRequestParam(c.Request)
	query := session.Where(filter).
		And(fmt.Sprintf("id NOT IN (SELECT user_id FROM %sengineer)", models.TablePrefix))

	//query extra filter here

	dataQuery := *query
	if order != "" {
		dataQuery = *dataQuery.OrderBy(order)
	} else {
		dataQuery = *dataQuery.Asc("id")
	}

	userList := make([]auth_models.UserInfo, 0, limit)
	err := dataQuery.Limit(limit, offset).Find(&userList)
	core.HandleError(err)

	countQuery := *query
	count, err := countQuery.Count(new(auth_models.UserInfo))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  userList,
		Total: count,
	})
}

func (c ItsmEngineer) AddSave() revel.Result {
	session := c.DbSession

	userIdList := make([]int64, 0)
	c.Params.Bind(&userIdList, "id_list")

	if len(userIdList) == 0 {
		return c.RenderJson(core.JsonResult{Success: false, Message: "添加失败，请选择用户再添加!"})
	}

	for _, userId := range userIdList {
		engineer := new(models.EngineerInfo)
		engineer.UserId = userId
		_, err := session.Insert(engineer)
		core.HandleError(err)
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: "添加成功!"})
}

func (c ItsmEngineer) Delete() revel.Result {
	db_session := c.DbSession

	engineer_id_list := make([]int64, 0)
	c.Params.Bind(&engineer_id_list, "id_list")

	affected, err := db_session.In("id", engineer_id_list).Delete(new(models.EngineerInfo))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}

type EngineerDetailForm struct {
	ServiceAreaList   []models.ServiceAreaInfo   `json:"service_area_list"`
	EventTypeList     []models.EventTypeInfo     `json:"event_type_list"`
	EngineerGroupList []models.EngineerGroupInfo `json:"engineer_group_list"`

	EngineerId int64 `json:"engineer_id"`

	ServiceArea   []int64 `json:"service_area"`
	EventType     []int64 `json:"event_type"`
	EngineerGroup []int64 `json:"engineer_group"`
	IsManager     int     `json:"is_manager"`
}

func (c ItsmEngineer) DetailData() revel.Result {
	db_session := c.DbSession

	var engineer_id int64
	c.Params.Bind(&engineer_id, "id")

	var detail_form EngineerDetailForm

	detail_form.ServiceAreaList = make([]models.ServiceAreaInfo, 0)
	err := db_session.Find(&(detail_form.ServiceAreaList))
	core.HandleError(err)

	detail_form.EventTypeList = make([]models.EventTypeInfo, 0)
	err = db_session.Find(&(detail_form.EventTypeList))
	core.HandleError(err)

	detail_form.EngineerGroupList = make([]models.EngineerGroupInfo, 0)
	err = db_session.Find(&(detail_form.EngineerGroupList))
	core.HandleError(err)

	if engineer_id != 0 {
		detail_form.EngineerGroup = engineer_service.GetEngineerGroupIdList(db_session, engineer_id)
		detail_form.ServiceArea = engineer_service.GetEngineerServiceAreaIdList(db_session, engineer_id)
		detail_form.EventType = engineer_service.GetEngineerEventTypeIdList(db_session, engineer_id)
		detail_form.IsManager = engineer_service.GetEngineerIsManager(db_session, engineer_id)
	}

	return c.RenderJson(detail_form)
}

func (c ItsmEngineer) Save() revel.Result {
	db_session := c.DbSession

	var detail_form EngineerDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	engineer_id := detail_form.EngineerId

	group_id_list_db := engineer_service.GetEngineerGroupIdList(db_session, engineer_id)
	group_id_list_ui := detail_form.EngineerGroup
	add_group_id_list := engineer_service.GetIntListInThisNotInThat(group_id_list_ui, group_id_list_db)
	delete_group_id_list := engineer_service.GetIntListInThisNotInThat(group_id_list_db, group_id_list_ui)
	engineer_service.AddEngineerToGroups(db_session, engineer_id, add_group_id_list)
	engineer_service.RemoveEngineerFromGroups(db_session, engineer_id, delete_group_id_list)

	service_area_id_list_db := engineer_service.GetEngineerServiceAreaIdList(db_session, engineer_id)
	service_area_id_list_ui := detail_form.ServiceArea
	add_service_area_id_list := engineer_service.GetIntListInThisNotInThat(service_area_id_list_ui, service_area_id_list_db)
	delete_service_area_id_list := engineer_service.GetIntListInThisNotInThat(service_area_id_list_db, service_area_id_list_ui)
	engineer_service.AddEngineerToServiceAreas(db_session, engineer_id, add_service_area_id_list)
	engineer_service.RemoveEngineerFromServiceAreas(db_session, engineer_id, delete_service_area_id_list)

	event_type_id_list_db := engineer_service.GetEngineerEventTypeIdList(db_session, engineer_id)
	event_type_id_list_ui := detail_form.EventType
	add_event_type_id_list := engineer_service.GetIntListInThisNotInThat(event_type_id_list_ui, event_type_id_list_db)
	delete_event_type_id_list := engineer_service.GetIntListInThisNotInThat(event_type_id_list_db, event_type_id_list_ui)
	engineer_service.AddEngineerToEventTypes(db_session, engineer_id, add_event_type_id_list)
	engineer_service.RemoveEngineerFromEventTypes(db_session, engineer_id, delete_event_type_id_list)

	if detail_form.IsManager == 0 {
		engineer_service.RemoveEngineerFromManager(db_session, engineer_id)
	} else if engineer_service.IsEngineerManager(db_session, engineer_id) == true {
		engineer_service.RemoveEngineerFromManager(db_session, engineer_id)
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: "保存成功!"})
}
