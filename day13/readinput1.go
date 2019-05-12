package main

import "fmt"

var (
	firstName, lastName, s string
	i int
	f float32
	input = "2233 / 22.33 / go"
	format = "%d / %f / %s"
)

func main() {
	fmt.Println("please enter your name:")
	fmt.Scanln(&firstName, &lastName)

	fmt.Printf("hi %s %s\n", firstName, lastName)
	fmt.Sscanf(input, format, &i, &f, &s)
	fmt.Println(f, i, s)
}