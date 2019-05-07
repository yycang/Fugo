package main

import "fmt"

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInto(a Appender, start, end int) {
	for i:= start; i<=end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

// lst 上无法调用指针，CountInto因为方法定义在指针上，但是可以调用LongEnough， 因为其方法定义在值上
func main() {
	var lst List

	if LongEnough(lst) {
		fmt.Printf("- lst is long enough\n")
	}
	plst := new(List)
	CountInto(plst, 1, 10)
	if LongEnough(plst) {
		fmt.Printf("- plst is long enough\n")
	}
}