package main

import "fmt"

func Multiply(a, b int, reply *int)  {
	*reply = a * b
}

// 函数中值传递是改变不了变量
// 传递指针可以节省内存和修改变量
func main() {
	n := 0
	reply := &n
	Multiply(10, 5, reply)
	fmt.Println("Multiply:", *reply)
}