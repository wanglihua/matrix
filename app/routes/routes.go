package routes

import "github.com/revel/revel"


type tHome struct {}
var Home tHome


func (_ tHome) Index(
        ) string {
    args := make(map[string]string)
    
    return revel.MainRouter.Reverse("Home.Index", args).Url
}


type tLogin struct {}
var Login tLogin


func (_ tLogin) Index(
        ) string {
    args := make(map[string]string)
    
    return revel.MainRouter.Reverse("Login.Index", args).Url
}

func (_ tLogin) Login(
        ) string {
    args := make(map[string]string)
    
    return revel.MainRouter.Reverse("Login.Login", args).Url
}

func (_ tLogin) Logout(
        ) string {
    args := make(map[string]string)
    
    return revel.MainRouter.Reverse("Login.Logout", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
        prefix string,
        filepath string,
        ) string {
    args := make(map[string]string)
    
    revel.Unbind(args, "prefix", prefix)
    revel.Unbind(args, "filepath", filepath)
    return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
        moduleName string,
        prefix string,
        filepath string,
        ) string {
    args := make(map[string]string)
    
    revel.Unbind(args, "moduleName", moduleName)
    revel.Unbind(args, "prefix", prefix)
    revel.Unbind(args, "filepath", filepath)
    return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


