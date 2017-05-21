package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	done := make(chan bool)
	go tickerCounter(ticker, done)
	time.Sleep(5 * time.Second)
	ticker.Stop()
	done <- true
	time.Sleep(10 * time.Second)
	fmt.Println("Exiting....")
}

func tickerCounter(ticker *time.Ticker, done chan bool) {

	i := 0
Loop:
	for {
		select {
		case t := <-ticker.C:
			i++
			fmt.Println("Count ", i, " at ", t)
		case <-done:
			fmt.Println("Got Done Signal")
			break Loop
		}
	}
	fmt.Println("Exiting the ticker counter Method")
}
