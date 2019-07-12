// 数组的比较
package main

import (
	"crypto/sha256"
	"fmt"
)

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))

	fmt.Printf("%X\n%X\n%t\n%T\n", c1, c2, c1==c2, c1)
}


/* 数组元素类型可比较，那么数组则可比较，数组比较时长度也要相同，不然编译错误。 */
