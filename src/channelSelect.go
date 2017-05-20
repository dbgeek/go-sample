package main

import (
	"fmt"
	"math/rand"
	"time"
)

var scMapping = map[string]int{
	"James": 5,
	"Kevin": 10,
	"Rahul": 9,
}

func findSC(name, server string, c chan int) {
	//Simulate searching
	time.Sleep(time.Duration(rand.Intn(70)) * time.Second)

	// return security clearance from map
	c <- scMapping[name]
}

func main() {
	rand.Seed(time.Now().UnixNano())

	c1 := make(chan int)
	c2 := make(chan int)

	name := "James"

	go findSC(name, "Server 1", c1)
	go findSC(name, "server 2", c2)

	select {
	case sc := <-c1:
		fmt.Println(name, " has a security clearance of ", sc, "found in server1")
	case sc := <-c2:
		fmt.Println(name, " has a security clearance of ", sc, "found in server2")
	case <-time.After(1 * time.Minute):
		fmt.Println("Search time out!!")
	}
}
