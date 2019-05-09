package main

import "fmt"

var i = 5
var str = "ABC"

type Person struct {
	name string
	age int
}

type Any interface {}

func main() {
	var val Any
	val = 5
	fmt.Printf("val has the value %v\n", val)
	val = str
	fmt.Printf("val has the value %v\n", val)
	pers1 := new(Person)
	pers1.name = "hulk"
	pers1.age = 55
	val = pers1
	fmt.Printf("val has the value: %v\n", val)
	switch t := val.(type) {
	case int:
		fmt.Printf("int %T\n", t)
	case string:
		fmt.Printf("string %T\n", t)
	case bool:
		fmt.Printf("boolean %T\n", t)
	case *Person:
		fmt.Printf("type pointer to person %T\n", t)
	default:
		fmt.Printf("unexpected type %T", t)
	}
}