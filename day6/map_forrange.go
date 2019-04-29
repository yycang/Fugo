package main

import "fmt"

func main() {
	// A
	items := make([]map[int]int, 5)
	for i:= range items {
		items[i] = make(map[int]int, 1)
		items[i][1] = 2
	}
	fmt.Printf("A: %v\n", items)

	// B	仅是map值的一个拷贝而已，map元素并没有得到初始化
	items2 := make([]map[int]int, 5)
	for _, item := range items2 {
		item = make(map[int]int, 1)
		item[1] = 2
	}
	fmt.Printf("B: %v\n", items2)
}