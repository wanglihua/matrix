package GridRequest

import (
    "github.com/revel/revel"
    "strconv"
)

func GetParam(request *revel.Request) (filter string, order string, offset int, limit int) {

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
