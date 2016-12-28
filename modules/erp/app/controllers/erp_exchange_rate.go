package controllers

import (
	"encoding/json"
	"strconv"
	"github.com/revel/revel"

	"matrix/core"
	"matrix/modules/erp/models"
	"matrix/modules/erp/service/status_service"
)

type ErpExchangeRate struct {
	*revel.Controller
}

func (c ErpExchangeRate) Index() revel.Result {
	return c.RenderTemplate("erp/exchange_rate/exchange_rate_index.html")
}

type ErpExchangeRateViewItem struct {
	models.ExchangeRateInfo `xorm:"extends" json:"r"`
	models.StatusInfo       `xorm:"extends" json:"s"`
}

func (c ErpExchangeRate) ListData() revel.Result {
	db_session := c.DbSession

	filter, order, offset, limit := core.GetGridRequestParam(c.Request)

	/*
		SELECT r.*, s.*
		FROM erp_exchange_rate r
		    INNER JOIN erp_status s ON r.status_id = s.id
	 */

	query := db_session.
	Select("r.*, s.*").
		Table("erp_exchange_rate").Alias("r").
		Join("INNER", []string{"erp_status", "s"}, "r.status_id = s.id").
		Where(filter)

	//query extra filter here

	data_query := *query
	if order != "" {
		data_query = *data_query.OrderBy(order)
	} else {
		data_query = *data_query.Asc("r.id")
	}

	exchange_rate_list := make([]ErpExchangeRateViewItem, 0, limit)
	err := data_query.Limit(limit, offset).Find(&exchange_rate_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(ErpExchangeRateViewItem))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  exchange_rate_list,
		Total: count,
	})
}

type ExchangeRateDetailForm struct {
	ExchangeRate models.ExchangeRateInfo `json:"exchange_rate"`
	StatusList   []models.StatusInfo     `json:"status_list"`
}

func (f *ExchangeRateDetailForm) IsCreate() bool {
	return f.ExchangeRate.Id == 0
}

func (f *ExchangeRateDetailForm) Valid(validation *revel.Validation) bool {
	validation.Required(f.ExchangeRate.Code).Message("符号不能为空！")
	if f.ExchangeRate.Code != "" {
		validation.MinSize(f.ExchangeRate.Code, 1).Message("符号长度不能小于1！")
	}
	if f.ExchangeRate.Code != "" {
		validation.MaxSize(f.ExchangeRate.Code, 10).Message("符号长度不能大于10！")
	}

	validation.Required(f.ExchangeRate.Name).Message("名称不能为空！")
	if f.ExchangeRate.Name != "" {
		validation.MinSize(f.ExchangeRate.Name, 1).Message("名称长度不能小于1！")
	}
	if f.ExchangeRate.Name != "" {
		validation.MaxSize(f.ExchangeRate.Name, 20).Message("名称长度不能大于20！")
	}

	validation.Required(f.ExchangeRate.Rate).Message("汇率不能为空！")

	validation.Required(f.ExchangeRate.StatusId).Message("状态不能为空！")

	return validation.HasErrors() == false
}

func (c ErpExchangeRate) DetailData() revel.Result {
	db_session := c.DbSession

	var exchange_rate_id int64
	c.Params.Bind(&exchange_rate_id, "id")

	var exchange_rate models.ExchangeRateInfo
	if exchange_rate_id != 0 {
		has, err := db_session.Id(exchange_rate_id).Get(&exchange_rate)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的货币汇率不存在！"})
		}
	}

	status_list := status_service.GetStatusList(db_session, models.STATUS_TYPE_EXCHANGE_RATE)

	var detail_form ExchangeRateDetailForm
	detail_form.ExchangeRate = exchange_rate
	detail_form.StatusList = status_list

	return c.RenderJson(detail_form)
}

func (c ErpExchangeRate) Save() revel.Result {
	db_session := c.DbSession

	var detail_form ExchangeRateDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	exchange_rate := detail_form.ExchangeRate

	var affected int64
	if detail_form.IsCreate() {
		affected, err = db_session.Insert(&exchange_rate)
		core.HandleError(err)
	} else {
		affected, err = db_session.Id(exchange_rate.Id).Update(&exchange_rate)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c ErpExchangeRate) Delete() revel.Result {
	db_session := c.DbSession

	exchange_rate_id_list := make([]int64, 0)
	c.Params.Bind(&exchange_rate_id_list, "id_list")

	affected, err := db_session.In("id", exchange_rate_id_list).Delete(new(models.ExchangeRateInfo))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}
