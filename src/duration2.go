package main

import (
	"fmt"
	"time"
)

func main() {
	go SlowCounter(2)
	time.Sleep(15 * time.Second)
}

func SlowCounter(n int) {
	i := 0
	// create a duration of n Seconds
	d := time.Duration(n) * time.Second

	for {
		// create a timer for this duration
		<-time.After(d)
		i++
		fmt.Println(i)
	}
}
