package main

import "fmt"

type IDuck interface {
	Quack()
	Walk()
}

func DuckDance(duck IDuck) {
	for i:=1; i<=3; i++ {
		duck.Quack()
		duck.Walk()
	}
}

type Bird struct {

}

func (b *Bird) Quack() {
	fmt.Println("i am quacking")
}

func (b *Bird) Walk() {
	fmt.Println("i am walking")
}

func main()  {
	b := new(Bird)
	DuckDance(b)
}