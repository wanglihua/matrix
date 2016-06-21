package GridRequest

import (
    "github.com/revel/revel"
    "strconv"
    "fmt"
    "strings"
)

func GetParam(request *revel.Request) (filter string, orderList []GridOrder, offset int, limit int) {

    searchValue := request.Form["sSearch"][0]

    fmt.Println(searchValue)

    colCount, _ := strconv.Atoi(request.Form["iColumns"][0])
    gridColumnList := make([]gridColumn, colCount)
    for colIndex := 0; colIndex < colCount; colIndex++ {
        colName := request.Form["mDataProp_" + strconv.Itoa(colIndex)][0]
        colSearchable := request.Form["bSearchable_" + strconv.Itoa(colIndex)][0]
        colSortable := request.Form["bSortable_" + strconv.Itoa(colIndex)][0]

        gridCol := gridColumn{
            ColName: colName,
            ColSearchable: colSearchable,
            ColSortable: colSortable,
        }

        gridColumnList = append(gridColumnList, gridCol)

        fmt.Println(colName)
        fmt.Println(colSearchable)
        fmt.Println(colSortable)
    }

    orderCount, _ := strconv.Atoi(request.Form["iSortingCols"][0])
    fmt.Println("OrderCount: ", orderCount)
    gridOrderList := make([]GridOrder, orderCount)
    for orderIndex := 0; orderIndex < orderCount; orderIndex++ {
        orderColIndex, _ := strconv.Atoi(request.Form["iSortCol_" + strconv.Itoa(orderIndex)][0])
        orderDir := strings.ToLower(request.Form["sSortDir_" + strconv.Itoa(orderIndex)][0])

        gridCol := gridColumnList[orderColIndex]
        if gridCol.ColSortable == "true" {
            gridOrder := GridOrder{
                OrderColumn: gridCol.ColName,
                OrderDir: orderDir,
            }

            gridOrderList = append(gridOrderList, gridOrder)
        }

        fmt.Println(orderColIndex)
        fmt.Println(orderDir)
    }

    filter = ""
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
    OrderColumn string // column name
    OrderDir    string // asc desc
}
