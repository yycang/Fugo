package main

import (
	"fmt"
	"time"
)

var values = [5]int{10, 11, 12, 13, 14}

func main() {
	for ix := range values {
		func() {
			fmt.Print(ix, " ")
		}()
	}
	fmt.Println()

	// 版本b
	for ix := range values {
		go func() {
			fmt.Print(ix, " ")
		}()
	}
	// 输出4， 4， 4， 4， 4。 闭包只绑定到一个变量上，每次循环都打印最后一个
	// 索引，因为协程在循环结束后还没有开始执行

	fmt.Println()
	time.Sleep(5e9)

	// 版本c
	for ix := range values {
		go func(ix interface{}) {
			fmt.Print(ix, " ")
		}(ix)
	}
	// 正确都写法， ix在每次循环都会被重新赋值
	fmt.Println()
	time.Sleep(5e9)

	// 版本d
	for ix := range values {
		val := values[ix]
		go func() {
			fmt.Print(val, " ")
		}()
	}
	// 在循环体里声明都变量协程之间不共享，所以这些变量可以单独被每个闭包使用
	time.Sleep(1e9)
}