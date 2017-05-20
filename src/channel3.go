package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan string,6)
	go SayHelloMiultipleTimes(c, 15)

	for s := range c {
		fmt.Println("out:",s, " en:", len(c), " cap:", cap(c))
		time.Sleep(2 * time.Second)
	}

	v, ok := <-c
	fmt.Println("Channel close?", !ok, " value ", v)
}

func SayHelloMiultipleTimes(c chan string, n int) {
	for i := 0; i <= n; i++ {
		c <- "Hello"
		fmt.Println("len:", len(c), " cap:", cap(c))
		fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
		time.Sleep(500 * time.Millisecond)
	}
	close(c)
}
