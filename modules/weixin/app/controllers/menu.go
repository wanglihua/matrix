package controllers

import (
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
    //session := c.DbSession


    return c.RenderJson(core.GridResult{
        Data:  make([]models.Menu, 0),
        Total: 0,
    })
}

func (c WeixinMenu) Detail() revel.Result {
    return c.RenderTemplate("weixin/menu/menu_detail.html")
}

func (c WeixinMenu) Save() revel.Result {
    return c.RenderJson(core.JsonResult{Success: true, Message: "保存成功!"})
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

func (c WeixinMenu)Delete() revel.Result {

    return c.RenderJson(core.JsonResult{Success: true, Message: "菜单删除成功!"})
}
