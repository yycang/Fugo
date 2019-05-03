package main

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	c float32
	int			// anonymous
	innerS		// anonymous
}

func main() {
	outer := new(outerS)
	outer.b = 22
	outer.c = 33.0
	outer.int = 66
	outer.in1 = 5
	outer.in2 = 10

	fmt.Printf("%d\n", outer.b)
	fmt.Printf("%f\n", outer.c)
	fmt.Printf("%d\n", outer.int)
	fmt.Printf("%d\n", outer.in1)
	fmt.Printf("%d\n", outer.in2)

	outer2 := outerS{4, 5.0, 50, innerS{5, 10}}
	fmt.Println(outer2)
}