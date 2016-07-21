package main

import (
    "fmt"
    "time"
)

func main() {
    //当前时间戳
    fmt.Println(time.Now().Unix())

    //当前时间戳格式化
    fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

    //时间戳格式化
    str_time := time.Unix(1389058332, 0).Format("2006-01-02 15:04:05")
    fmt.Println(str_time)

    //时间反格式化成时间戳（月份必须是month）
    the_time := time.Date(2016, 1, 5, 0, 0, 0, 0, time.Local)
    unix_time1 := the_time.Unix()
    fmt.Println(unix_time1)

    //时间反格式化成时间戳
    the_time, err := time.ParseInLocation("2006-01-02", "2015-12-01", time.Local)
    //the_time, err := time.Parse("2006-01-02", "2015-12-01")
    if err != nil {
        fmt.Println(err)
    }

    unix_time2 := the_time.Unix()
    fmt.Println(unix_time2)

    fmt.Printf("%t", true)
}
