package service

import "encoding/json"

type JsonResult struct {
    Success bool        `json:"success"`
    Message string      `json:"message"`
}

type GridResult struct {
    Data  interface{}
    Total int64
}

func (r GridResult) MarshalJSON() ([]byte, error) {
    resultMap := map[string]interface{}{
        "aaData": r.Data,
        "iTotalDisplayRecords":r.Total,
        "iTotalRecords":r.Total,
    }

    return  json.Marshal(resultMap)
}
