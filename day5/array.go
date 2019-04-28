package main

import "fmt"

func main() {
	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}

	// 声明数组
	var arr1 [5]string
	var arr2 = new([5]int)

	// 数组常量
	var arr3 = [10]int{1,2,3}	// 前三数是1，2，3。后面全是0
	var arr4 = [...]int{5,6,7,8,22}		// 切片
	var arr5 = [5]string{3: "a", 4: "b"}	// 只有索引3，4被赋予值
}



