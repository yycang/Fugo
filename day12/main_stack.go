package main

import (
	"./stack"
	"fmt"
)

var st1 stack.Stack

func main() {
	st1.Push("Brown")
	st1.Push(2233)
	st1.Push(3.14)
	st1.Push([]string{"hello", "world"})
	for {
		item, err := st1.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}

