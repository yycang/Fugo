// 十进制数字的字符串每三位插一个逗号
package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return comma(s[:n-3]) + "," + s[n-3:]
}

func comma2(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	for i, v := range s {
		if i % 3 == 2 {
			buf.WriteString(",")
		}
		buf.WriteString(string(v))
	}
	return buf.String()
}

func main() {
	fmt.Println(comma2("123456789"))
}