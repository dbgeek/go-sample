package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	go tickerCounter(ticker)
	time.Sleep(5 * time.Second)
	ticker.Stop()
	time.Sleep(10 * time.Second)
	fmt.Println("Exiting....")
}

func tickerCounter(ticker *time.Ticker) {

	i := 0
	for t := range ticker.C {
		i++
		fmt.Println("Count ", i, " at ", t)
	}
}
