package main

import "fmt"

type CMember struct {
	name string
	age int
	address string
	rank string
	clearance int
}

func main () {
	// declare a struct

	// cm := CMember{"björn", 39, "Lejonets gata 351", "Oracle DBA", 5}

	// or

	/*
   cm := CMember{
     name: "björn", 
     clearance:  5,
     age: 39,
     address: "Lejonets gata 351",
     rank: "Oracle DBA",
   }
  */

	var cm CMember
	cm.name = "björn"
	cm.age = 39
	cm.address = "LEjonets gatan 351"
	cm.rank = "Oracle DBA"
	cm.clearance = 5

	cmp := &cm
	cmp.clearance = 4

  fmt.Println("variable cm:")
	fmt.Println(cm)
	fmt.Println("variable cmp:")
	
	fmt.Println(cmp)

	var crew []CMember

	crew = append(crew, cm, CMember{"Zaneta", 38, "lejonets gatan 351", "Hud Terapeut", 1})

	fmt.Println("loop crew Slice")
	for i, v := range crew {
		fmt.Println(i,v)
	}

	// maps with struct
	var m map[string]CMember
	m = make(map[string]CMember)

	m["Dominik"] = cm

	//or

	mm := map[string]CMember{
		"Evgin": CMember{name: "Evgin", address: "Brandbergen"},
		"Daniel": CMember{name: "Daniel", address: "huddinge"},
	}

	//add
	mm["kalle"] = CMember{name: "kalle", address: "Stockholm"}

	//retrieve
	element := mm["kalle"]

	fmt.Println("element:", element)
	fmt.Println(mm)

	//Check if key exists
	v, ok := mm["Daniel"]
  fmt.Println(v,ok)
	//delete
	delete(mm, "Daniel")

	for key, value := range mm {
		fmt.Println("key:", key, "value:", value)
	}

	cm.PrintSecurityClearance()
}

func (cm CMember) PrintSecurityClearance() {
		fmt.Println(cm.clearance)
	}
