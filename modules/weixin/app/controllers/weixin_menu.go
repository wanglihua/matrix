package controllers

import (
    "strconv"
    "github.com/revel/revel"
    "matrix/core"
    "matrix/core/requests"
    "matrix/modules/weixin/service"
    "matrix/modules/weixin/models"
    "github.com/antonholmquist/jason"
    "bytes"
    "bufio"
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

    filter, _, offset, limit := core.GetGridRequestParam(c.Request)
    query := session.Where(filter)

    //query extra filter here

    dataQuery := *query
    dataQuery = *dataQuery.Asc("menu_order")

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
    Menu                   models.Menu
    MenuLevelSelectOptions []core.SelectOption
    MenuTypeSelectOptions  []core.SelectOption
}

func (f *MenuForm) InitData() {
    menuLevelSelectOptions := make([]core.SelectOption, 0)
    menuLevelSelectOptions = append(menuLevelSelectOptions, core.SelectOption{Value:"1", Text:"一级菜单"})
    menuLevelSelectOptions = append(menuLevelSelectOptions, core.SelectOption{Value:"2", Text:"二级菜单"})
    f.MenuLevelSelectOptions = menuLevelSelectOptions

    menuTypeSelectOptions := make([]core.SelectOption, 0)
    menuTypeSelectOptions = append(menuTypeSelectOptions, core.SelectOption{Value:"", Text:"父菜单"})
    menuTypeSelectOptions = append(menuTypeSelectOptions, core.SelectOption{Value:"click", Text:"单击事件"})
    menuTypeSelectOptions = append(menuTypeSelectOptions, core.SelectOption{Value:"view", Text:"跳转URL"})
    f.MenuTypeSelectOptions = menuTypeSelectOptions
}

func (f *MenuForm) IsCreate() bool {
    return f.Menu.Id == 0
}

func (f *MenuForm) Valid(validation *revel.Validation) bool {
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
        validation.MaxSize(f.Menu.Data, 250).Message("菜单数据长度不能大于250！")
    }

    validation.Required(f.Menu.Level).Message("层级不能为空！")
    validation.Min(f.Menu.Level, 1).Message("层级不能小于1！")
    validation.Max(f.Menu.Level, 2).Message("层级不能大于2！")

    //validation.Required(f.Menu.Order).Message("排序不能为空！")
    validation.Min(f.Menu.Order, 1).Message("排序不能小于1！")

    return validation.HasErrors() == false
}

func (c WeixinMenu) Detail() revel.Result {
    session := c.DbSession

    var menuId int64
    c.Params.Bind(&menuId, "id")

    menu := new(models.Menu)
    if menuId != 0 {
        has, err := session.Id(menuId).Get(menu)
        core.HandleError(err)
        if has == false {
            panic("指定的微信菜单不存在！")
        }
    }

    form := new(MenuForm)
    form.InitData()

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

    revel.TRACE.Println(c.Request.Form["form.Menu.Type"][0])
    revel.TRACE.Println(menu.Type)

    var affected int64
    var err error
    if form.IsCreate() {
        affected, err = session.Insert(menu)
        core.HandleError(err)
    } else {
        affected, err = session.Id(menu.Id).Update(menu)
        core.HandleError(err)

        if affected == 0 {
            return c.RenderJson(core.JsonResult{Success: false, Message: "数据保存失败，请重试！"})
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

    //jason, err := jason.NewObjectFromBytes([]byte(`{"errcode":46003,"errmsg":"menu no exist hint: [mh2CUA0174vr21]"}`))
    weixinMenuJason, err := jason.NewObjectFromBytes([]byte(responeStr))
    core.HandleError(err)

    errorCode, err := weixinMenuJason.GetInt64("errcode")
    if err == nil && errorCode != 0 {
        revel.TRACE.Println(errorCode)
        errorMsg, err := weixinMenuJason.GetString("errmsg")
        core.HandleError(err)

        return c.RenderJson(core.JsonResult{Success: false, Message: "微信菜单下载失败!详情：" + errorMsg})
    }

    weixinButtonListJason, err := weixinMenuJason.GetObjectArray("menu", "button")
    core.HandleError(err)

    menuList := make([]models.Menu, 0)
    menuOrder := 1
    for _, weixinButtonJason := range (weixinButtonListJason) {
        weixinSubButtonListJason, err := weixinButtonJason.GetObjectArray("sub_button")
        core.HandleError(err)

        if len(weixinSubButtonListJason) == 0 {
            //没有子菜单
            menu := getWeixinMenuFromJason(weixinButtonJason)
            menu.Order = menuOrder
            menuOrder++
            menu.Level = 1
            menuList = append(menuList, menu)
        } else {
            //有子菜单
            var menu models.Menu
            menuName, err := weixinButtonJason.GetString("name")
            core.HandleError(err)

            menu.Name = menuName
            menu.Order = menuOrder
            menuOrder++
            menu.Level = 1
            menuList = append(menuList, menu)

            for _, weixinSubButtonJason := range (weixinSubButtonListJason) {
                menu := getWeixinMenuFromJason(weixinSubButtonJason)
                menu.Order = menuOrder
                menuOrder++
                menu.Level = 2
                menuList = append(menuList, menu)
            }
        }
    }

    err = session.Begin()
    core.HandleError(err)

    /*
    下载时采用不删除原来菜单的方式，用户可以自行先删除本地不再需要的菜单
    _, err = session.Where("id <> 0").Delete(new(models.Menu))
    if err != nil {
        core.HandleError(session.Rollback())
        return c.RenderJson(core.JsonResult{Success: false, Message: "微信菜单下载失败!详情：" + err.Error()})
    }
    */

    revel.TRACE.Println("menu length: " + strconv.Itoa(len(menuList)))

    for _, menu := range (menuList) {
        affected, err := session.Insert(menu)
        if affected != 1 || err != nil {
            core.HandleError(session.Rollback())

            return c.RenderJson(core.JsonResult{Success: false, Message: "微信菜单下载失败!详情：" + err.Error()})
        }
    }

    err = session.Commit()
    core.HandleError(err)

    return c.RenderJson(core.JsonResult{Success: true, Message: "微信菜单下载成功!"})
}

func getWeixinMenuFromJason(jasonObject *jason.Object) models.Menu {
    var menu models.Menu

    var err error
    menu.Name, err = jasonObject.GetString("name")
    core.HandleError(err)

    menu.Type, err = jasonObject.GetString("type")
    core.HandleError(err)

    if menu.Type == "click" {
        menu.Data, err = jasonObject.GetString("key")
        core.HandleError(err)
    } else if menu.Type == "view" {
        menu.Data, err = jasonObject.GetString("url")
        core.HandleError(err)
    }

    return menu
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

    var err error
    menuList := make([]models.Menu, 0)
    err = session.Asc("menu_order").Find(&menuList)
    core.HandleError(err)

    /*
        {"button":[
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
                    }
                ]
            }
        ]}
    */

    menuJsonBuffer := bytes.NewBuffer(make([]byte, 0))
    bw := bufio.NewWriter(menuJsonBuffer)
    bw.WriteString(`{"button":[`)

    prevLevel := 0
    currLevel := 1
    for _, menu := range (menuList) {
        currLevel = menu.Level

        if prevLevel == 0 && currLevel == 1 {
            bw.WriteString(`{`)
        } else if prevLevel == 1 && currLevel == 1 {
            bw.WriteString(`},{`)
        } else if prevLevel == 1 && currLevel == 2 {
            bw.WriteString(`,"sub_button":[{`)
        } else if prevLevel == 2 && currLevel == 1 {
            bw.WriteString(`}]},{`)
        } else if prevLevel == 2 && currLevel == 2 {
            bw.WriteString(`},{`)
        }

        //type
        if menu.Type != "" {
            bw.WriteString(`"type":"`)
            bw.WriteString(menu.Type)
            bw.WriteString(`",`)
        }

        //data
        if menu.Type == "click" {
            bw.WriteString(`"key": "`)
            bw.WriteString(menu.Data)
            bw.WriteString(`",`)
        } else if menu.Type == "view" {
            bw.WriteString(`"url": "`)
            bw.WriteString(menu.Data)
            bw.WriteString(`",`)
        }

        //name
        bw.WriteString(`"name":"`)
        bw.WriteString(menu.Name)
        bw.WriteString(`"`) //这里可以不需要 ","

        prevLevel = currLevel
    }

    if currLevel == 1 {
        bw.WriteString(`}`)
    } else if currLevel == 2 {
        bw.WriteString(`}]}`)
    }

    bw.WriteString("]}")
    bw.Flush()

    menuJson := menuJsonBuffer.String()

    // https://api.weixin.qq.com/cgi-bin/menu/create?access_token=ACCESS_TOKEN
    accessToken := service.GetAccessToken(session)
    url := "https://api.weixin.qq.com/cgi-bin/menu/create?access_token=" + accessToken

    /*
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
                    }
                ]
            }]
        }
    */

    revel.TRACE.Println(menuJson)

    var responeStr = requests.PostJson(url, menuJson)

    /*
    {"errcode":0,"errmsg":"ok"}
    {"errcode":40018,"errmsg":"invalid button name size"}
     */

    resultJason, err := jason.NewObjectFromBytes([]byte(responeStr))
    core.HandleError(err)
    errCode, err := resultJason.GetInt64("errcode")
    core.HandleError(err)
    if errCode != 0 {
        errMsg, err := resultJason.GetString("errmsg")
        core.HandleError(err)

        return c.RenderJson(core.JsonResult{Success: true, Message: "菜单上传微信失败!详情：" + errMsg})
    }

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
