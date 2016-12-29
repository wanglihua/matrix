package core

import (
	"github.com/revel/revel"
	"strconv"
	"strings"
	"matrix/db"
)

func GetGridRequestParam(request *revel.Request) (filter string, order string, offset int, limit int) {
	filter = "" // return
	order = ""  // return
	offset = 0  // return
	limit = 10  // return

	if len(request.Form["iColumns"]) == 0 {
		return
	}

	colCount, _ := strconv.Atoi(request.Form["iColumns"][0])
	gridColumnList := make([]gridColumn, 0, colCount)

	condition := ""
	if len(request.Form["sSearch"]) != 0 {
		searchValue := request.Form["sSearch"][0]
		for colIndex := 0; colIndex < colCount; colIndex++ {
			colName := request.Form["mDataProp_" + strconv.Itoa(colIndex)][0]
			colSearchable := request.Form["bSearchable_" + strconv.Itoa(colIndex)][0]
			colSortable := request.Form["bSortable_" + strconv.Itoa(colIndex)][0]
			gridCol := gridColumn{
				ColName:       colName,
				ColSearchable: colSearchable,
				ColSortable:   colSortable,
			}
			gridColumnList = append(gridColumnList, gridCol)
			if searchValue != "" && colSearchable == "true" {
				condition += " or (" + db.Engine.Quote(colName) + " like '%" + searchValue + "%'" + ") "
			}
		}
	}
	if condition != "" {
		condition = " ((1=0) " + condition + " ) "
	} else {
		condition = " (1=1) "
	}

	order_str := ""
	if len(request.Form["iSortingCols"]) != 0 {
		orderCount, _ := strconv.Atoi(request.Form["iSortingCols"][0])
		gridOrderList := make([]string, 0, orderCount)
		for orderIndex := 0; orderIndex < orderCount; orderIndex++ {
			orderColIndex, _ := strconv.Atoi(request.Form["iSortCol_" + strconv.Itoa(orderIndex)][0])
			orderDir := strings.ToLower(request.Form["sSortDir_" + strconv.Itoa(orderIndex)][0])

			gridCol := gridColumnList[orderColIndex]
			if gridCol.ColSortable == "true" {
				if strings.ToLower(orderDir) == "desc" {
					gridOrderList = append(gridOrderList, db.Engine.Quote(gridCol.ColName)+" desc")
				} else {
					//不能识别为 desc 的排序，均视为 asc
					gridOrderList = append(gridOrderList, db.Engine.Quote(gridCol.ColName)+" asc")
				}
			}
		}
		order_str = strings.Join(gridOrderList, ",")
	}

	filter = PreventSQLInjection(condition) // return
	order = PreventSQLInjection(order_str)  // return
	if len(request.Form["iDisplayStart"]) != 0 {
		offset, _ = strconv.Atoi(request.Form["iDisplayStart"][0]) // return
	}
	if len(request.Form["iDisplayLength"]) != 0 {
		limit, _ = strconv.Atoi(request.Form["iDisplayLength"][0]) // return
	}

	return
}

type gridColumn struct {
	ColName       string
	ColSearchable string
	ColSortable   string
}
