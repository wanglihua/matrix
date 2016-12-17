package controllers

import (
	"github.com/revel/revel"
	"strconv"

	"matrix/core"
	"matrix/modules/auth/models"
	"strings"
)

type AuthGroupUser struct {
	*revel.Controller
}

func (c AuthGroupUser) Index() revel.Result {
	var groupId int64
	c.Params.Bind(&groupId, "groupId")

	c.RenderArgs["groupId"] = groupId
	return c.RenderTemplate("auth/groupuser/groupuser_index.html")
}

type GroupUserView struct {
	models.GroupUserInfo `xorm:"extends" json:"gu"`
	models.GroupInfo     `xorm:"extends" json:"g"`
	models.UserInfo      `xorm:"extends" json:"u"`
}

func (c AuthGroupUser) ListData() revel.Result {
	session := c.DbSession

	var groupId int64
	c.Params.Bind(&groupId, "groupId")

	filter, order, offset, limit := core.GetGridRequestParam(c.Request)
	query := session.
		Select("gu.*, g.*, u.*").
		Table(models.TablePrefix+"group_user").Alias("gu").
		Join("inner", []string{models.TablePrefix + "group", "g"}, "gu.group_id = g.id").
		Join("inner", []string{models.TablePrefix + "user", "u"}, "gu.user_id = u.id").
		Where(filter)

	//query extra filter here
	query = query.Where("gu.group_id = ?", groupId)

	dataQuery := *query
	if order != "" {
		dataQuery = *dataQuery.OrderBy(order)
	} else {
		dataQuery = *dataQuery.Asc("gu.add_time")
	}

	groupUserList := make([]GroupUserView, 0, limit)
	err := dataQuery.Limit(limit, offset).Find(&groupUserList)
	core.HandleError(err)

	countQuery := *query
	count, err := countQuery.Count(new(GroupUserView))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  groupUserList,
		Total: count,
	})
}

func (c AuthGroupUser) Add() revel.Result {
	var groupId int64
	c.Params.Bind(&groupId, "groupId")

	c.RenderArgs["groupId"] = groupId
	return c.RenderTemplate("auth/groupuser/groupuser_add.html")
}

func (c AuthGroupUser) AddListData() revel.Result {
	session := c.DbSession

	var groupId int64
	c.Params.Bind(&groupId, "groupId")

	//get group user id list
	userIdQuery := session.Where("group_id = ?", groupId)
	groupUserList := make([]models.GroupUserInfo, 0)
	userIdQuery.Select("user_id").Find(&groupUserList)
	userIdList := make([]string, 0)
	for _, groupUser := range groupUserList {
		userIdList = append(userIdList, strconv.FormatInt(groupUser.UserId, 10))
	}

	filter, order, offset, limit := core.GetGridRequestParam(c.Request)
	query := session.Where(filter)

	if len(userIdList) != 0 {
		query = query.Where("id not in (" + strings.Join(userIdList, ",") + ")")
	}

	dataQuery := *query
	if order != "" {
		dataQuery = *dataQuery.OrderBy(order)
	} else {
		dataQuery = *dataQuery.Asc("id")
	}

	userList := make([]models.UserInfo, 0, limit)
	err := dataQuery.Limit(limit, offset).Find(&userList)
	core.HandleError(err)

	countQuery := *query
	count, err := countQuery.Count(new(models.UserInfo))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  userList,
		Total: count,
	})
}

func (c AuthGroupUser) AddSave() revel.Result {
	session := c.DbSession

	var groupId int64
	c.Params.Bind(&groupId, "groupId")

	addUserIdList := make([]int64, 0)
	c.Params.Bind(&addUserIdList, "userIdList")

	for _, userId := range addUserIdList {
		groupUser := new(models.GroupUserInfo)
		groupUser.UserId = userId
		groupUser.GroupId = groupId

		session.Insert(groupUser)
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.Itoa(len(addUserIdList)) + "个成员添加成功!"})
}

func (c AuthGroupUser) Remove() revel.Result {
	session := c.DbSession

	groupUserIdList := make([]int64, 0)
	c.Params.Bind(&groupUserIdList, "id_list")

	groupUser := new(models.GroupUserInfo)
	affected, err := session.In("id", groupUserIdList).Delete(groupUser)
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}
