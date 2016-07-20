package requests

import (
    "net/http"
    "io/ioutil"
    "strings"
    "net/url"

    "matrix/core"
)

func Get(url string) string {
    httpResp, err := http.Get(url)
    core.HandleError(err)
    if httpResp != nil {
        defer httpResp.Body.Close()
    }

    responeBytes, err := ioutil.ReadAll(httpResp.Body)
    core.HandleError(err)

    return string(responeBytes)
}

func Post(url string) string {
    httpResp, err := http.Post(url, "application/x-www-form-urlencoded", strings.NewReader(""))
    core.HandleError(err)
    if httpResp != nil {
        defer httpResp.Body.Close()
    }

    responeBytes, err := ioutil.ReadAll(httpResp.Body)
    core.HandleError(err)

    return string(responeBytes)
}

func PostForm(url string, data url.Values) string {
    httpResp, err := http.PostForm(url, data)
    core.HandleError(err)
    if httpResp != nil {
        defer httpResp.Body.Close()
    }

    responeBytes, err := ioutil.ReadAll(httpResp.Body)
    core.HandleError(err)

    return string(responeBytes)
}

func PostJson(url string, json string) string {
    return postWithContentType(url, "application/json", json)
}

func postWithContentType(url string, bodyType string, body string) string {
    httpResp, err := http.Post(url, bodyType, strings.NewReader(body))
    core.HandleError(err)
    if httpResp != nil {
        defer httpResp.Body.Close()
    }

    responeBytes, err := ioutil.ReadAll(httpResp.Body)
    core.HandleError(err)

    return string(responeBytes)
}

/*

// "net/http"的使用

func httpGet() {
    resp, err := http.Get("http://www.01happy.com/demo/accept.php?id=1")
    if err != nil {
        // handle error
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }

    fmt.Println(string(body))
}

func httpPost() {
    resp, err := http.Post("http://www.01happy.com/demo/accept.php",
        "application/x-www-form-urlencoded",
        strings.NewReader("name=cjb"))
    if err != nil {
        fmt.Println(err)
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }

    fmt.Println(string(body))
}

func httpPostForm() {
    resp, err := http.PostForm("http://www.01happy.com/demo/accept.php",
        url.Values{"key": {"Value"}, "id": {"123"}})

    if err != nil {
        // handle error
    }

    defer resp.Body.Close()
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }

    fmt.Println(string(body))
}

func httpDo() {
    client := &http.Client{}

    req, err := http.NewRequest("POST", "http://www.01happy.com/demo/accept.php", strings.NewReader("name=cjb"))
    if err != nil {
        // handle error
    }

    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
    req.Header.Set("Cookie", "name=anny")

    resp, err := client.Do(req)

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        // handle error
    }

    fmt.Println(string(body))
}

*/