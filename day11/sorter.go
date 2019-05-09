package main

import (
	"./sort"
	"fmt"
)

func ints() {
	data := []int{75, 22, 33, -10, 455, -250, 0, 0, 1234}
	a := sort.IntArray(data)
	sort.Sort(a)
	if !sort.IsSorted(a) {
		panic("fails")
	}
	fmt.Printf("the sorted array is %v\n", a)
}

func strings() {
	data := []string{"monday", "aaa", "bbb", "ccc"}
	a := sort.StringArray(data)
	sort.Sort(a)
	if !sort.IsSorted(a) {
		panic("fails")
	}
	fmt.Printf("the sorted string array is %v\n", a)
}

func main() {
	ints()
	strings()
}
