package app

import (
    "github.com/revel/revel"
    "runtime"
    "fmt"
    "matrix/service"
)

func init() {
    runtime.GOMAXPROCS(runtime.NumCPU())

    revel.TimeFormats = append(revel.TimeFormats, "2006-01-02 15:04:05")

    revel.InterceptMethod((*service.BaseController).Before, revel.BEFORE)
    revel.InterceptMethod((*service.BaseController).After, revel.AFTER)
    revel.InterceptMethod((*service.BaseController).Panic, revel.PANIC)

    revel.InterceptFunc(userAuthInterceptor, revel.BEFORE, revel.ALL_CONTROLLERS)

    // Filters is the default set of global filters.
    revel.Filters = []revel.Filter{
        PanicFilter,
        revel.RouterFilter, // Use the routing table to select the right Action
        revel.FilterConfiguringFilter, // A hook for adding or removing per-Action filters.
        revel.ParamsFilter, // Parse parameters into Controller.Params.
        revel.SessionFilter, // Restore and write the session cookie.
        revel.FlashFilter, // Restore and write the flash cookie.
        revel.ValidationFilter, // Restore kept validation errors and save new ones from cookie.
        revel.I18nFilter, // Resolve the requested language
        HeaderFilter, // Add some security based headers
        revel.InterceptorFilter, // Run interceptors around the action.
        revel.CompressFilter, // Compress the result.
        revel.ActionInvoker, // Invoke the action.
    }

    registerTags()

    // register startup functions with OnAppStart
    // ( order dependent )
    // revel.OnAppStart(InitDB)
    // revel.OnAppStart(FillCache)
}

func userAuthInterceptor(c *revel.Controller) revel.Result {
    //if user := connected(c); user == nil {
    //    c.Flash.Error("请先登录")
    //    return c.Redirect(App.Index)
    //}

    //var excludeCheckActions = []string{};

    /*
    Static.Serve
    Static.ServeModule
    TestRunner.Index
    TestRunner.Run
     */
    if (c.Name != "Static" && c.Name != "TestRunner") {
        fmt.Println("userAuthInterceptor " + c.Action)
    }

    return nil
}
