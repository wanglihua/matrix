package weixin

import (
	"github.com/antonholmquist/jason"
	"github.com/revel/revel"
	"github.com/revel/revel/testing"
	"github.com/yasuyuky/jsonpath"
	"matrix/core"
	"reflect"
)

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

type JsonPathTest struct {
	testing.TestSuite
}

func (t *JsonPathTest) TestJsonPath() {
	jsonStr := `{"errcode":46003,"errmsg":"menu no exist hint: [mh2CUA0174vr21]"}`
	jsonData, err := jsonpath.DecodeString(jsonStr)
	core.HandleError(err)

	errcode, err := jsonpath.Get(jsonData, []interface{}{"errcode"}, "")
	core.HandleError(err)

	t.AssertEqual(float64(46003.0), errcode.(float64))
}

func (t *JsonPathTest) TestJasonReadInt() {
	jsonData, err := jason.NewObjectFromBytes([]byte(`{"errcode":46003,"errmsg":"menu no exist hint: [mh2CUA0174vr21]"}`))
	core.HandleError(err)

	errcode, err := jsonData.GetInt64("errcode")
	core.HandleError(err)

	//errcodeNumber, err := errcode.Int64()
	core.HandleError(err)
	t.AssertEqual(46003, errcode)
}

func (t *JsonPathTest) TestJasonRead() {
	var weixinMenuJson = `
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
    `
	jsonData, _ := jason.NewObjectFromBytes([]byte(weixinMenuJson))
	revel.TRACE.Println(reflect.TypeOf(jsonData))
	revel.TRACE.Println(jsonData)

	buttons, _ := jsonData.GetObjectArray("menu", "button")
	revel.TRACE.Println(reflect.TypeOf(buttons))
	revel.TRACE.Println(buttons)
	revel.TRACE.Println(len(buttons))
	revel.TRACE.Println(buttons == nil)

	for _, button := range buttons {
		btnType, _ := button.GetString("type")
		revel.TRACE.Println(reflect.TypeOf(btnType))
		revel.TRACE.Println(btnType)
		//revel.TRACE.Println(btnType == nil)
		revel.TRACE.Println(btnType == "")
	}

}
