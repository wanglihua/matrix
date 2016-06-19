package app

import (
    "github.com/revel/revel"
    "runtime"
    "fmt"
    "html/template"
)

func init() {
    runtime.GOMAXPROCS(runtime.NumCPU())

    revel.InterceptFunc(userAuthInterceptor, revel.BEFORE, revel.ALL_CONTROLLERS)

    // Filters is the default set of global filters.
    revel.Filters = []revel.Filter{
        revel.PanicFilter, // Recover from panics and display an error page instead.
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

    revel.TemplateFuncs["test"] = testTemplateFunc

    // register startup functions with OnAppStart
    // ( order dependent )
    // revel.OnAppStart(InitDB)
    // revel.OnAppStart(FillCache)
}

// TODO turn this into revel.HeaderFilter
// should probably also have a filter for CSRF
// not sure if it can go in the same filter or not
var HeaderFilter = func(c *revel.Controller, fc []revel.Filter) {
    // Add some common security headers
    c.Response.Out.Header().Add("X-Frame-Options", "SAMEORIGIN")
    c.Response.Out.Header().Add("X-XSS-Protection", "1; mode=block")
    c.Response.Out.Header().Add("X-Content-Type-Options", "nosniff")

    fc[0](c, fc[1:]) // Execute the next filter stage.
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
        fmt.Println(c.Action)
    }

    return nil
}

func testTemplateFunc(renderArgs map[string]interface{}) template.HTML {
    session := renderArgs["session"].(revel.Session)
    return template.HTML("<span style='color:red;'>Hello World Just a test " + session.Id() + "!</span>")
    //return template.HTML("<span style='color:red;'>Hello World Just a test!</span>")
}