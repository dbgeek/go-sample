package main

import (
	"ClassFactoryTutorial/appliances"
	"fmt"
)

func main() {
	fmt.Println("Enter prefered appliance type")
	fmt.Println("0: stove")
	fmt.Println("1: fridge")

	var myType int
	fmt.Scan(&myType)

	myAppliance, err := appliances.CreateAppliance(myType)

	if err == nil {
		myAppliance.Start()
		fmt.Println(myAppliance.GetPurpose())
	} else {
		fmt.Println(err)
	}
}
