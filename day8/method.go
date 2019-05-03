package main

import "fmt"

type TwoInts struct {
	a int
	b int
}

type IntVector []int

func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 24

	fmt.Printf("the sum is %d\n", two1.AddThem())

	fmt.Printf("the sum add 10 is %d\n", two1.AddParams(10))

	two2 := TwoInts{10, 20}
	fmt.Printf("the sum is %d\n", two2.AddThem())

	fmt.Println(IntVector{1, 2, 3}.Sum())

}

func (a *TwoInts) AddThem() int {
	return a.a + a.b
}

func (b *TwoInts) AddParams(param int) int {
	return b.a + b.b + param
}

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}
