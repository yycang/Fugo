package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Topologicalgenus interface {
	Rank() int
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (sq *Square) Rank() int  {
	return 1
}

type Rectangle struct {
	length, width float32
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func (r Rectangle) Rank() int  {
	return 2
}

func main() {
	r := Rectangle{5, 3}
	q := &Square{5}
	shapes := []Shaper{r, q}

	fmt.Println("looping through shapes for area")
	for n, _ := range shapes {
		fmt.Println("shape detail:", shapes[n])
		fmt.Println("area of this shape is ", shapes[n].Area())
	}
	topgen := []Topologicalgenus{r, q}
	fmt.Println("looping through topgen for rank")
	for n, _ := range topgen {
		fmt.Println("shape details: ", topgen[n])
		fmt.Println("topological genus of this shape is: ", topgen[n].Rank())
	}
}

