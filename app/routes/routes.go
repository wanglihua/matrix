package routes

import "github.com/revel/revel"


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


type tHome struct {}
var Home tHome


func (_ tHome) Index(
        ) string {
    args := make(map[string]string)
    
    return revel.MainRouter.Reverse("Home.Index", args).Url
}


type tTestRunner struct {}
var TestRunner tTestRunner


func (_ tTestRunner) Index(
        ) string {
    args := make(map[string]string)
    
    return revel.MainRouter.Reverse("TestRunner.Index", args).Url
}

func (_ tTestRunner) Suite(
        suite string,
        ) string {
    args := make(map[string]string)
    
    revel.Unbind(args, "suite", suite)
    return revel.MainRouter.Reverse("TestRunner.Suite", args).Url
}

func (_ tTestRunner) Run(
        suite string,
        test string,
        ) string {
    args := make(map[string]string)
    
    revel.Unbind(args, "suite", suite)
    revel.Unbind(args, "test", test)
    return revel.MainRouter.Reverse("TestRunner.Run", args).Url
}

func (_ tTestRunner) List(
        ) string {
    args := make(map[string]string)
    
    return revel.MainRouter.Reverse("TestRunner.List", args).Url
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



var ActionList = []string{

    "Login.Index",
    "Login.Login",
    "Login.Logout",

    "Home.Index",

    "TestRunner.Index",
    "TestRunner.Suite",
    "TestRunner.Run",
    "TestRunner.List",

    "Static.Serve",
    "Static.ServeModule",

}
