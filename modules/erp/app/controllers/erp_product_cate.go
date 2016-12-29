package controllers

import (
	"encoding/json"
	"strconv"
	"github.com/revel/revel"

	"matrix/core"
	"matrix/modules/erp/models"
	"matrix/modules/erp/service/status_service"
	"fmt"
)

type ErpProductCate struct {
	*revel.Controller
}

func (c ErpProductCate) Index() revel.Result {
	return c.RenderTemplate("erp/product_cate/product_cate_index.html")
}

type ErpProductCateViewItem struct {
	models.ProductCateInfo          `xorm:"extends" json:"c"`
	Parent models.ProductCateInfo   `xorm:"extends" json:"p"`
	models.StatusInfo               `xorm:"extends" json:"s"`
}

func (c ErpProductCate) ListData() revel.Result {
	db_session := c.DbSession

	filter, order, offset, limit := core.GetGridRequestParam(c.Request)

	/*
		SELECT c.*, p.*, s.*
		FROM erp_product_cate c
		    INNER JOIN erp_product_cate p ON c.parent_id = p.id
		    INNER JOIN erp_status s ON c.status_id = s.id
	 */

	query := db_session.
	Select("c.*, p.*, s.*").
		Table("erp_product_cate").Alias("c").
		Join("LEFT", []string{"erp_product_cate", "p"}, "c.parent_id = p.id").
		Join("INNER", []string{"erp_status", "s"}, "c.status_id = s.id").
		Where(filter)

	//query extra filter here

	data_query := *query
	if order != "" {
		data_query = *data_query.OrderBy(order)
	} else {
		data_query = *data_query.Asc("c.id")
	}

	product_cate_list := make([]ErpProductCateViewItem, 0, limit)
	err := data_query.Limit(limit, offset).Find(&product_cate_list)
	core.HandleError(err)

	count_query := *query
	count, err := count_query.Count(new(ErpProductCateViewItem))
	core.HandleError(err)

	return c.RenderJson(core.GridResult{
		Data:  product_cate_list,
		Total: count,
	})
}

type ProductCateDetailForm struct {
	ProductCate    models.ProductCateInfo   `json:"product_cate"`
	ParentCateList []models.ProductCateInfo `json:"parent_cate_list"`
	StatusList     []models.StatusInfo      `json:"status_list"`
}

func (f *ProductCateDetailForm) IsCreate() bool {
	return f.ProductCate.Id == 0
}

func (f *ProductCateDetailForm) Valid(validation *revel.Validation) bool {
	validation.Required(f.ProductCate.Name).Message("名称不能为空！")
	if f.ProductCate.Name != "" {
		validation.MinSize(f.ProductCate.Name, 2).Message("名称长度不能小于2！")
	}
	if f.ProductCate.Name != "" {
		validation.MaxSize(f.ProductCate.Name, 100).Message("名称长度不能大于100！")
	}

	validation.Required(f.ProductCate.EnglishName).Message("英文名不能为空！")
	if f.ProductCate.EnglishName != "" {
		validation.MinSize(f.ProductCate.EnglishName, 2).Message("英文名长度不能小于2！")
	}
	if f.ProductCate.EnglishName != "" {
		validation.MaxSize(f.ProductCate.EnglishName, 100).Message("英文名长度不能大于100！")
	}

	if f.ProductCate.Desc.Valid == true {
		validation.MaxSize(f.ProductCate.Desc.String, 500).Message("描述长度不能大于500！")
	}

	if f.ProductCate.HsCode.Valid == true {
		validation.MinSize(f.ProductCate.HsCode.String, 2).Message("报关编码长度不能小于2！")
	}
	if f.ProductCate.HsCode.Valid == true {
		validation.MaxSize(f.ProductCate.HsCode.String, 20).Message("报关编码长度不能大于20！")
	}

	validation.Required(f.ProductCate.StatusId).Message("状态不能为空！")

	return validation.HasErrors() == false
}

func (c ErpProductCate) DetailData() revel.Result {
	db_session := c.DbSession

	var product_cate_id int64
	c.Params.Bind(&product_cate_id, "id")

	var product_cate models.ProductCateInfo
	if product_cate_id != 0 {
		has, err := db_session.Id(product_cate_id).Get(&product_cate)
		core.HandleError(err)
		if has == false {
			return c.RenderJson(core.JsonResult{Success: false, Message: "指定的商品目录不存在！"})
		}
	}

	var parent_cate_list []models.ProductCateInfo
	parent_cate_cols := models.ProductCateCols
	if product_cate_id == 0 {
		err := db_session.Where(fmt.Sprintf("%[1]s is null or %[1]s = 0", parent_cate_cols.ParentId)).Find(&parent_cate_list)
		core.HandleError(err)
	} else {
		err := db_session.
		Where(fmt.Sprintf("(%[1]s is null or %[1]s = 0) and (%[2]s <> %[3]d)", parent_cate_cols.ParentId, parent_cate_cols.Id, product_cate_id)).
			Find(&parent_cate_list)
		core.HandleError(err)
	}

	status_list := status_service.GetStatusList(db_session, models.STATUS_TYPE_PRODUCT_CATE)

	var detail_form ProductCateDetailForm
	detail_form.ProductCate = product_cate
	detail_form.StatusList = status_list
	detail_form.ParentCateList = parent_cate_list

	return c.RenderJson(detail_form)
}

func (c ErpProductCate) Save() revel.Result {
	db_session := c.DbSession

	var detail_form ProductCateDetailForm
	err := json.Unmarshal(c.GetRequestBody(), &detail_form)
	core.HandleError(err)

	if detail_form.Valid(c.Validation) == false {
		return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage()})
	}

	product_cate := detail_form.ProductCate

	var affected int64
	if detail_form.IsCreate() {
		affected, err = db_session.Insert(&product_cate)
		core.HandleError(err)
	} else {
		affected, err = db_session.Id(product_cate.Id).Update(&product_cate)
		core.HandleError(err)

		if affected == 0 {
			return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
		}
	}

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

func (c ErpProductCate) Delete() revel.Result {
	db_session := c.DbSession

	product_cate_id_list := make([]int64, 0)
	c.Params.Bind(&product_cate_id_list, "id_list")

	affected, err := db_session.In("id", product_cate_id_list).Delete(new(models.ProductCateInfo))
	core.HandleError(err)

	return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}
