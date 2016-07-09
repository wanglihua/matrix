package controllers

import (
    "strconv"
    "github.com/revel/revel"

    "matrix/core"
    "matrix/modules/auth/models"
    "strings"
)

type AuthGroupUser struct {
    *revel.Controller
    core.BaseController
}

func (c AuthGroupUser) Index() revel.Result {
    var groupId int64
    c.Params.Bind(&groupId, "groupId")

    c.RenderArgs["groupId"] = groupId
    return c.RenderTemplate("auth/groupuser/groupuser_index.html")
}

func (c AuthGroupUser) ListData() revel.Result {
    session := c.DbSession

    var groupId int64
    c.Params.Bind(&groupId, "groupId")

    filter, order, offset, limit := core.GetGridRequestParam(c.Request)
    query := session.Where(filter)

    //query extra filter here
    query = query.Where("group_id = ?", groupId)

    dataQuery := *query
    if order != "" {
        dataQuery = *dataQuery.OrderBy(order)
    } else {
        dataQuery = *dataQuery.Asc("id")
    }

    groupUserList := make([]models.GroupUser, 0, limit)
    err := dataQuery.Limit(limit, offset).Find(&groupUserList)
    core.HandleError(err)

    countQuery := *query
    count, err := countQuery.Count(new(models.GroupUser))
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
    groupUserList := make([]models.GroupUser, 0)
    userIdQuery.Select("user_id").Find(&groupUserList)
    userIdList := make([]string, 0)
    for _, groupUser := range (groupUserList) {
        userIdList = append(userIdList, strconv.FormatInt(groupUser.UserId, 10))
    }

    filter, order, offset, limit := core.GetGridRequestParam(c.Request)
    query := session.Where(filter)

    //query extra filter here
    query = query.Where("id not in (" + strings.Join(userIdList, ",") + ")")

    dataQuery := *query
    if order != "" {
        dataQuery = *dataQuery.OrderBy(order)
    } else {
        dataQuery = *dataQuery.Asc("id")
    }

    userList := make([]models.User, 0, limit)
    err := dataQuery.Limit(limit, offset).Find(&userList)
    core.HandleError(err)

    countQuery := *query
    count, err := countQuery.Count(new(models.User))
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

    for _, userId := range (addUserIdList) {
        groupUser := new(models.GroupUser)
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

    groupUser := new(models.GroupUser)
    affected, err := session.In("id", groupUserIdList).Delete(groupUser)
    core.HandleError(err)

    return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}

