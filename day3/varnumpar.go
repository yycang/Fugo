package main

import "fmt"

func main() {
	x := min(4, 3, 12, 1, 55)
	fmt.Printf("minimum is %d\n", x)
	slice := []int{131, 414, 5, 11, 231, 59}
	x = min(slice...)
	fmt.Printf("minimum is %d\n", x)

}

func min(s ... int) int {
	if len(s) == 0 { return 0 }
	min := s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}