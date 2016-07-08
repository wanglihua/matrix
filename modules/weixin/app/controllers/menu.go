package controllers

import (
    "strconv"
    "github.com/revel/revel"
    "matrix/core"
    "matrix/core/requests"
    "matrix/modules/weixin/service"
    "matrix/modules/weixin/models"
)

type WeixinMenu struct {
    *revel.Controller
    core.BaseController
}

func (c WeixinMenu) Index() revel.Result {
    return c.RenderTemplate("weixin/menu/menu_index.html")
}

func (c WeixinMenu) ListData() revel.Result {
    session := c.DbSession

    filter, order, offset, limit := core.GetGridRequestParam(c.Request)
    query := session.Where(filter)

    //query extra filter here

    dataQuery := *query
    if order != "" {
        dataQuery = *dataQuery.OrderBy(order)
    } else {
        dataQuery = *dataQuery.Asc("id")
    }

    menuList := make([]models.Menu, 0, limit)
    err := dataQuery.Limit(limit, offset).Find(&menuList)
    core.HandleError(err)

    countQuery := *query
    count, err := countQuery.Count(new(models.Menu))
    core.HandleError(err)

    return c.RenderJson(core.GridResult{
        Data:  menuList,
        Total: count,
    })
}

type MenuForm struct {
    Menu models.Menu
}

func (f MenuForm) IsCreate() bool {
    return f.Menu.Id == 0
}

func (f MenuForm) Valid(validation *revel.Validation) bool { 
    validation.Required(f.Menu.Name).Message("菜单名不能为空！")
    if f.Menu.Name != "" {
        validation.MinSize(f.Menu.Name, 0).Message("菜单名长度不能小于0！")
    }
    if f.Menu.Name != "" {
        validation.MaxSize(f.Menu.Name, 12).Message("菜单名长度不能大于12！")
    }

    if f.Menu.Type != "" {
        validation.MinSize(f.Menu.Type, 0).Message("类型长度不能小于0！")
    }
    if f.Menu.Type != "" {
        validation.MaxSize(f.Menu.Type, 12).Message("类型长度不能大于12！")
    }

    if f.Menu.Data != "" {
        validation.MinSize(f.Menu.Data, 2).Message("菜单数据长度不能小于2！")
    }
    if f.Menu.Data != "" {
        validation.MaxSize(f.Menu.Data, 12).Message("菜单数据长度不能大于12！")
    }

    validation.Required(f.Menu.Level).Message("层级不能为空！")
    validation.Min(f.Menu.Level, 1).Message("层级不能小于1！")
    validation.Max(f.Menu.Level, 2).Message("层级不能大于2！")

    validation.Required(f.Menu.Order).Message("排序不能为空！")

    return validation.HasErrors() == false
}

func (c WeixinMenu) Detail() revel.Result {
    session := c.DbSession

    menuId := core.GetInt64FromRequest(c.Request, "id")

    menu := new(models.Menu)
    if menuId != 0 {
        has, err := session.Id(menuId).Get(menu)
        core.HandleError(err)
        if has == false {
            panic("指定的微信菜单不存在！")
        }
    }

    form := new(MenuForm)
    form.Menu = *menu

    c.RenderArgs["form"] = form
    return c.RenderTemplate("weixin/menu/menu_detail.html")
}

func (c WeixinMenu) Save() revel.Result {
    session := c.DbSession

    form := new(MenuForm)
    c.Params.Bind(form, "form")

    if form.Valid(c.Validation) == false {
        return c.RenderJson(core.JsonResult{Success: false, Message: c.GetValidationErrorMessage() })
    }

    menu := &form.Menu

    var affected int64
    var err error
    if form.IsCreate() { 
        affected, err = session.Insert(menu)
        core.HandleError(err)
    } else { 
        affected, err = session.Id(menu.Id).Update(menu)
        core.HandleError(err)

        if affected == 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！" })
        }
    }

    return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据保存成功!"})
}

//查看服务器端菜单
func (c WeixinMenu) Server() revel.Result {
    return c.RenderTemplate("weixin/menu/menu_server.html")
}

func (c WeixinMenu) Download() revel.Result {
    session := c.DbSession

    //https://api.weixin.qq.com/cgi-bin/menu/get?access_token=ACCESS_TOKEN
    accessToken := service.GetAccessToken(session)
    url := "https://api.weixin.qq.com/cgi-bin/menu/get?access_token=" + accessToken
    var responeStr = requests.Get(url)
    revel.TRACE.Println(responeStr)

    // {"errcode":46003,"errmsg":"menu no exist hint: [mh2CUA0174vr21]"}

    /*
    {
        "menu": {
            "button": [
                {
                    "type": "click",
                    "name": "今日歌曲",
                    "key": "V1001_TODAY_MUSIC",
                    "sub_button": [ ]
                },
                {
                    "type": "click",
                    "name": "歌手简介",
                    "key": "V1001_TODAY_SINGER",
                    "sub_button": [ ]
                },
                {
                    "name": "菜单",
                    "sub_button": [
                        {
                            "type": "view",
                            "name": "搜索",
                            "url": "http://www.soso.com/",
                            "sub_button": [ ]
                        },
                        {
                            "type": "view",
                            "name": "视频",
                            "url": "http://v.qq.com/",
                            "sub_button": [ ]
                        },
                        {
                            "type": "click",
                            "name": "赞一下我们",
                            "key": "V1001_GOOD",
                            "sub_button": [ ]
                        }
                    ]
                }
            ]
        }
    }
    */

    //data, err := jsonpath.DecodeString(responeStr)


    return c.RenderJson(core.JsonResult{Success: true, Message: "菜单微信下载成功!"})
}

func (c WeixinMenu) Upload() revel.Result {
    session := c.DbSession

    //是不是需要先删除，在创建。在菜单不存在的时候删除会怎样
    /*
        请求说明
        http请求方式：GET
        https://api.weixin.qq.com/cgi-bin/menu/delete?access_token=ACCESS_TOKEN
        返回说明
        对应创建接口，正确的Json返回结果:
        {"errcode":0,"errmsg":"ok"}
    */

    // https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN
    accessToken := service.GetAccessToken(session)
    url := "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=" + accessToken

    json :=
    `
     {
         "button":[
         {
              "type":"click",
              "name":"今日歌曲",
              "key":"V1001_TODAY_MUSIC"
          },
          {
               "name":"菜单",
               "sub_button":[
               {
                   "type":"view",
                   "name":"搜索",
                   "url":"http://www.soso.com/"
                },
                {
                   "type":"view",
                   "name":"视频",
                   "url":"http://v.qq.com/"
                },
                {
                   "type":"click",
                   "name":"赞一下我们",
                   "key":"V1001_GOOD"
                }]
           }]
     }
    `

    var responeStr = requests.PostJson(url, json)

    /*
    {"errcode":0,"errmsg":"ok"}
    {"errcode":40018,"errmsg":"invalid button name size"}
     */

    revel.TRACE.Println(responeStr)

    return c.RenderJson(core.JsonResult{Success: true, Message: "菜单上传微信成功!"})
}

func (c WeixinMenu) Delete() revel.Result {
    session := c.DbSession

    menuIdList := make([]int64, 0)
    c.Params.Bind(&menuIdList, "id_list")

    menu := new(models.Menu)
    affected, err := session.In("id", menuIdList).Delete(menu)
    core.HandleError(err)

    return c.RenderJson(core.JsonResult{Success: true, Message: strconv.FormatInt(affected, 10) + "条数据删除成功!"})
}
