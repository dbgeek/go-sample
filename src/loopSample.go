package main

import "fmt"

func main() {
    fmt.Println("counting")
    for i := 1; i <= 4; i++ {
        fmt.Println(i)
    }
    fmt.Println("done")
    
    fmt.Println("defer counting")
    for i := 1; i <= 4; i++ {
        defer fmt.Println(i)
    }
    fmt.Println("defer done")
}