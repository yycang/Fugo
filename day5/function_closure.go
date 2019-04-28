package main

import "fmt"

func main()  {
	var f = Adder()
	fmt.Print(f(1), "-")
	fmt.Print(f(10), "-")
	fmt.Print(f(100), "\n")
}

func Adder() func(int) int {
	var x int
	return func(i int) int {
		x += i
		return x
	}
}