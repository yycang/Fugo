package main

import (
	"./person"
	"fmt"
)

func main() {
	p := new(person.Person)

	p.SetFirstName("cang")
	fmt.Println(p.FirstName())
}