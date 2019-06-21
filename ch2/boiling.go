// boiling 输出水的沸点

package main

import "fmt"

const boilingF = 212.0

func main() {
	var f = boilingF
	var c = (f - 32) * 5 / 9
	fmt.Printf("boiling point = %gºF or %gºC\n", f, c)
}

// 输入一个华氏温度，输出摄氏温度
func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
