package main

import (
    "net/http"
    "fmt"
    "io/ioutil"
)

func main() {
    resp, er := http.Get("http://quotes.rest/qod.json")
    if er != nil {
        fmt.Println(er)
        return
    }
    defer resp.Body.Close()
    contents, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println(er)
        return
    }
    fmt.Println(string(contents))
}