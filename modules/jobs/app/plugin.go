package app

import (
    //"fmt"
    "github.com/revel/revel"
    //"time"
    //"matrix/modules/jobs/app/jobs"
)

func init() {
    revel.OnAppStart(func() {
        //fmt.Println("jobs module start!")

        /*
        jobs.Schedule("0 0 0 * * ?",  TestTask{})
        jobs.Schedule("@midnight",    TestTask{})
        jobs.Schedule("@every 24h",   TestTask{})
        jobs.Every(24 * time.Hour,    TestTask{})
        */

        //cron.test = * * * * * ?
        //jobs.Schedule("cron.test", TestTask{}) //测试
    })
}


type TestTask struct {
}

func (e TestTask) Run() {
    revel.TRACE.Println("这是一个定时任务的测试！")
}
