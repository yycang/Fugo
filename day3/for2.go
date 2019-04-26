package main

import "fmt"

var i int = 5

func main() {
	for i >= 0 {
		fmt.Println(i)
		i --
	}

	for i:=0; i<5; i++ {
		var v int
		fmt.Printf("%d", v)
		v = 5
	}

	for i:=0; i<3; {
		fmt.Println(i)
		if i == 2 {
			break
		}
		i++
	}
}