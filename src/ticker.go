package main

import (
	"fmt"
	"time"
)

func main() {
	go tickerCounter(1)
	time.Sleep(5 * time.Second)
}

func tickerCounter(n int) {
	ticker := time.NewTicker(time.Duration(n) * time.Second)
	i := 0
	for t := range ticker.C {
		i++
		fmt.Println("Count ", i, " at ", t)
	}
}
