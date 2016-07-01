package app

import (
    "github.com/revel/revel"
    "runtime"
    "matrix/core"
)

func init() {
    runtime.GOMAXPROCS(runtime.NumCPU())

    revel.TimeFormats = append(revel.TimeFormats, "2006-01-02 15:04:05")

    revel.InterceptMethod((*core.BaseController).Before, revel.BEFORE)
    revel.InterceptMethod((*core.BaseController).After, revel.AFTER)
    revel.InterceptMethod((*core.BaseController).Panic, revel.PANIC)

    //revel.InterceptFunc(userAuthInterceptor, revel.BEFORE, revel.ALL_CONTROLLERS)

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

    revel.OnAppStart(func() {
        revel.LoadMimeConfig()
    })
}
