// 输出命令行其参数
package main

import (
	"os"
	"fmt"
	"strings"
)

func main() {
	s, sep := "", ""
	for i:=1; i<len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	//fmt.Println(s)

	s2, sep2 := "", ""

	for _, i := range os.Args[1:] {
		s2 += sep2 + i
		sep2 = " "
	}
	//fmt.Println(s2)

	fmt.Println(strings.Join(os.Args[1:], " "))

	fmt.Println(os.Args[1:])

}