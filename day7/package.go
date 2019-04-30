package main

import (
	"fmt"
	"./pack1"
)

func main() {
	var test1 string
	test1 = pack1.ReturnStr()

	fmt.Printf("returnStr %s\n", test1)
	fmt.Printf("%d\n", pack1.Pack1Int)
	fmt.Printf("%f\n", pack1.PackFloat)
}