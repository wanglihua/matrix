package GridRequest

import (
    "github.com/revel/revel"
    "strconv"
    "fmt"
)

func GetParam(request *revel.Request) (filter string, order string, offset int, limit int) {

    searchValue := request.Form["sSearch"][0]

    fmt.Println(searchValue)

    colCount, _ := strconv.Atoi(request.Form["iColumns"][0])
    for colIndex := 0; colIndex < colCount; colIndex++ {
        colName := request.Form["mDataProp_" + strconv.Itoa(colIndex)][0]
        colSearchable := request.Form["bSearchable_" + strconv.Itoa(colIndex)][0]
        colSortable := request.Form["bSortable_" + strconv.Itoa(colIndex)][0]

        fmt.Println(colName)
        fmt.Println(colSearchable)
        fmt.Println(colSortable)
    }

    orderCount, _ := strconv.Atoi(request.Form["iSortingCols"][0])
    for orderIndex := 0; orderIndex < orderCount; orderIndex++ {
        orderColIndex, _ := strconv.Atoi(request.Form["iSortCol_" + strconv.Itoa(orderIndex)][0])
        orderDir := request.Form["sSortDir_" + strconv.Itoa(orderIndex)][0]

        fmt.Println(orderColIndex)
        fmt.Println(orderDir)
    }

    filter = ""
    order = ""

    offset, _ = strconv.Atoi(request.Form["iDisplayStart"][0])
    limit, _ = strconv.Atoi(request.Form["iDisplayLength"][0])

    return
}

type GridOrder struct {
    OrderDir    string
    OrderColumn string
}
