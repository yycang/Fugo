package main

import "fmt"

func main()  {
	var res int
	var arr = [5]int{0, 1, 2, 3, 4}
	res = sum(arr[:])
	fmt.Println(res)

}

func sum(a [] int) int {
	s := 0
	for i:=0; i<len(a); i++ {
		s += a[i]
	}
	return s
}

