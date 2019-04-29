package main

import "fmt"

func main() {
	var value int
	var isPresent bool
	map1 := make(map[string]int)
	map1["new delhi"] = 55
	map1["beijing"] = 20
	map1["washinton"] = 25

	if v, ok := map1["beijing"]; ok {
		fmt.Printf("the value is %d\n", v)
	} else {
		fmt.Printf("doesn't have key")
	}

	value, isPresent = map1["paris"]
	fmt.Printf("is paris in map1? %t\n", isPresent)
	fmt.Printf("value is %d\n", value)


	delete(map1, "washington")
	if value2, ok3 := map1["washington"]; ok3 {
		fmt.Printf("value is %d\n", value2)
	} else {
		fmt.Println("doesn't have key")
	}
}