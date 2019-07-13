// 反转一个数组
package main

import "fmt"

func rev(s []int)  {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1{
		s[i], s[j] = s[j], s[i]
	}
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	rev(s)
	fmt.Println(s)
}