package weixin

import (
	"github.com/revel/revel"
	"github.com/revel/revel/testing"
	"matrix/modules/weixin/models"
	"sort"
)

type MenuSortTest struct {
	testing.TestSuite
}

func (t *MenuSortTest) TestSort() {
	var menuList models.MenuList
	menuList = append(menuList, models.MenuInfo{Id: 1, Name: "menu1", Order: 1})
	menuList = append(menuList, models.MenuInfo{Id: 3, Name: "menu3", Order: 3})
	menuList = append(menuList, models.MenuInfo{Id: 2, Name: "menu2", Order: 2})
	menuList = append(menuList, models.MenuInfo{Id: 4, Name: "menu4", Order: 4})
	menuList = append(menuList, models.MenuInfo{Id: 6, Name: "menu6", Order: 6})
	menuList = append(menuList, models.MenuInfo{Id: 5, Name: "menu5", Order: 5})

	sort.Sort(menuList)

	revel.TRACE.Println(menuList)
}
