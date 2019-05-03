package main

import "fmt"

type A struct {
	ax, ay int
}

type B struct {
	A
	bx, by float32
}


// 若命名冲突的话，外层名字会覆盖内层，若同级出现则使用时会报错
func main() {
	b := B{A{1,2}, 3, 4}
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)
}