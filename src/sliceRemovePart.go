package main

import "fmt"

func main() {
	a := []int{1,2,3,4,5,6,7,8,9,10}
	//remove 2,3
	a = append(a[:2], a[4:]...)
	fmt.Println("remove 3,4 from a;",a)
}
