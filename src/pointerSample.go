package main

import (
    "fmt"
)

func main() {
    X := 5
    ChangeXByValue(X) //Passing value
    fmt.Println("The value of X:", X)
    ChangeXByPointer(&X) //Passing pointer
    fmt.Println("The value of X:", X)
    X = ChangeXByReturn(X) //Function that return value
    fmt.Println("The value of X:", X)
}

func ChangeXByValue(X int){
    X = 10
}


func ChangeXByPointer(X *int){
    *X = 10
}

func ChangeXByReturn(X int)(returnX){
    var returnX int = 15
    return returnX
}