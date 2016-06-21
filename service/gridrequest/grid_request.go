package GridRequest

import (
    "github.com/revel/revel"
    "strconv"
    "strings"
)

func GetParam(request *revel.Request) (filter string, orderList []GridOrder, offset int, limit int) {

    searchValue := request.Form["sSearch"][0]

    condition := ""
    colCount, _ := strconv.Atoi(request.Form["iColumns"][0])
    gridColumnList := make([]gridColumn, 0, colCount)
    for colIndex := 0; colIndex < colCount; colIndex++ {
        colName := request.Form["mDataProp_" + strconv.Itoa(colIndex)][0]
        colSearchable := request.Form["bSearchable_" + strconv.Itoa(colIndex)][0]
        colSortable := request.Form["bSortable_" + strconv.Itoa(colIndex)][0]

        gridCol := gridColumn{
            ColName :colName,
            ColSearchable : colSearchable,
            ColSortable : colSortable,
        }
        gridColumnList = append(gridColumnList, gridCol)

        if (searchValue != "" && colSearchable == "true") {
            condition += " or (" + colName + " like '%" + searchValue + "%'" + ") "
        }
    }
    if condition != "" {
        condition = " ((1=0) " + condition + " ) "
    }

    orderCount, _ := strconv.Atoi(request.Form["iSortingCols"][0])
    gridOrderList := make([]GridOrder, 0, orderCount)
    for orderIndex := 0; orderIndex < orderCount; orderIndex++ {
        orderColIndex, _ := strconv.Atoi(request.Form["iSortCol_" + strconv.Itoa(orderIndex)][0])
        orderDir := strings.ToLower(request.Form["sSortDir_" + strconv.Itoa(orderIndex)][0])

        gridCol := gridColumnList[orderColIndex]
        if gridCol.ColSortable == "true" {
            gridOrder := GridOrder{
                ColName:  strings.Replace(strings.Replace(strings.Replace(gridCol.ColName, "--", "", -1), "#", "", -1), ";", "", -1),
                OrderAsc: strings.ToLower(orderDir) != "desc",
            }

            gridOrderList = append(gridOrderList, gridOrder)
        }
    }

    filter = strings.Replace(strings.Replace(strings.Replace(condition, "--", "", -1), "#", "", -1), ";", "", -1) // return
    orderList = gridOrderList // return

    offset, _ = strconv.Atoi(request.Form["iDisplayStart"][0]) // return
    limit, _ = strconv.Atoi(request.Form["iDisplayLength"][0]) // return

    return
}

type gridColumn struct {
    ColName       string
    ColSearchable string
    ColSortable   string
}

type GridOrder struct {
    ColName  string // column name
    OrderAsc bool   // asc true, desc false
}
