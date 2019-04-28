package main

import "fmt"


func main() {
	p2 := Add2()
	fmt.Printf("%v\n", p2(3))

	p3 := Adder(2)
	fmt.Printf("%v", p3(3))
}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}


func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}