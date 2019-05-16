package main

import "fmt"

func main() {
	for i:=0; i<=100; i++ {
		fmt.Printf("is the integer %d even? %v\n", i, even.Even(i))
	}
}