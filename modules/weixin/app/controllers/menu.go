package controllers

import (
    "github.com/revel/revel"
    "matrix/core"
    "matrix/modules/weixin/service"
    "net/http"
    //"io/ioutil"
    //"github.com/yasuyuky/jsonpath"
    "strings"
    "io/ioutil"
)

type WeixinMenu struct {
    *revel.Controller
    core.BaseController
}

func (c WeixinMenu) Index() revel.Result {
    return c.RenderTemplate("weixin/menu/menu_index.html")
}

func (c WeixinMenu) Detail() revel.Result{
    return c.RenderTemplate("weixin/menu/menu_detail.html")
}

func (c WeixinMenu) Save() revel.Result {
    return c.RenderJson(core.JsonResult{Success: true, Message: "保存成功!"})
}

func (c WeixinMenu) Download() revel.Result {
    session := c.DbSession

    //https://api.weixin.qq.com/cgi-bin/menu/get?access_token=ACCESS_TOKEN
    accessToken := service.GetAccessToken(session)
    url := "https://api.weixin.qq.com/cgi-bin/menu/get?access_token=" + accessToken
    httpResp, err := http.Get(url)
    core.HandleError(err)
    defer httpResp.Body.Close()

    //responeBytes, err := ioutil.ReadAll(httpResp.Body)
    //core.HandleError(err)
    //var responeStr = string(responeBytes)
    //revel.TRACE.Println(responeStr)

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

    httpResp, err := http.Post(url, "application/json", strings.NewReader(json))
    core.HandleError(err)
    defer httpResp.Body.Close()

    responeBytes, err := ioutil.ReadAll(httpResp.Body)
    core.HandleError(err)

    /*
    {"errcode":0,"errmsg":"ok"}
    {"errcode":40018,"errmsg":"invalid button name size"}
     */

    var responeStr = string(responeBytes)
    revel.TRACE.Println(responeStr)

    //responeBytes, err := ioutil.ReadAll(httpResp.Body)

    return c.RenderJson(core.JsonResult{Success: true, Message: "菜单上传微信成功!"})
}

func (c WeixinMenu)Delete() revel.Result {
    return c.RenderJson(core.JsonResult{Success: true, Message: "菜单删除成功!"})
}
