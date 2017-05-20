package main

import (
	"fmt"
)

func main() {
	c := make(chan string)
	go SayHelloMiultipleTimes(c, 5)

	for s := range c {
		fmt.Println(s)
	}

	v, ok := <-c
	fmt.Println("Channel close?", !ok, " value ", v)
}

func SayHelloMiultipleTimes(c chan string, n int) {
	for i := 0; i <= n; i++ {
		c <- "Hello"
	}
	close(c)
}
