package controllers

import (
	"github.com/revel/revel"

	"matrix/core"
	"matrix/service"
	"matrix/modules/auth/models"
	"fmt"
)

type AuthAdmin struct {
	*revel.Controller
	service.BaseController
}

func (c AuthAdmin) Index() revel.Result {
	return c.RenderTemplate("auth/admin/admin_index.html")
}

type AdminView struct {
	models.Admin    `xorm:"extends" json:"a"`
	LoginName string `xorm:"'login_name'" json:"login_name"`
	NickName  string `xorm:"'nick_name'" json:"nick_name"`
}

func (c AuthAdmin) ListData() revel.Result {
	session := c.DbSession

	filter, order, offset, limit := core.GetGridRequestParam(c.Request)
	query := session.
	Select("a.*, u.login_name, u.nick_name").
			Table(models.TablePrefix + "user").Alias("u").
			Join("inner", []string{models.TablePrefix + "admin", "a"}, "u.id = a.user_id").
			Where(filter)

	//query extra filter here

	dataQuery := *query
	if order != "" {
		dataQuery = *dataQuery.OrderBy(order)
	} else {
		dataQuery = *dataQuery.Asc("a.id")
	}

	adminList := make([]AdminView, 0, limit)
	err := dataQuery.Limit(limit, offset).Find(&adminList)
	core.HandleError(err)

	countQuery := *query
	count, err := countQuery.Count(new(AdminView))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  adminList,
		Total: count,
	})
}

func (c AuthAdmin) Add() revel.Result {
	//session := c.DbSession
	return c.RenderTemplate("auth/admin/admin_add.html")
}

func (c AuthAdmin) AddListData() revel.Result {
	session := c.DbSession

	filter, order, offset, limit := core.GetGridRequestParam(c.Request)
	query := session.Where(filter).
			And(fmt.Sprintf("id NOT IN (SELECT user_id FROM %sadmin)", models.TablePrefix))

	//query extra filter here

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

func (c AuthAdmin) AddSave() revel.Result {
	session := c.DbSession

	userIdList := make([]int64, 0)
	c.Params.Bind(&userIdList, "id_list")

	if len(userIdList) == 0 {
		return c.RenderJson(core.JsonResult{Success: false, Message: "添加失败，请选择用户再添加!"})
	}

	for _, userId := range userIdList {
		admin := new(models.Admin)
		admin.UserId = userId
		_, err := session.Insert(admin)
		core.HandleError(err)
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: "添加成功!"})
}

func (c AuthAdmin) Remove() revel.Result {
	session := c.DbSession

	adminIdList := make([]int, 0)
	c.Params.Bind(&adminIdList, "id_list")

	admin := new(models.Admin)
	count, err := session.Count(admin)
	if int64(len(adminIdList)) == count {
		return c.RenderJson(core.JsonResult{Success: false, Message: "移除失败，管理员不能全部移除!"})
	}

	_, err = session.In("id", adminIdList).Delete(admin)
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: "移除成功!"})
}
