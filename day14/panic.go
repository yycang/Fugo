package main

import "fmt"

func main() {
	fmt.Println("starting")
	panic("error occurred")
	fmt.Println("ending")
}