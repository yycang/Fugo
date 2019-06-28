// cf将数值参数转为摄氏温度和华式温度
package main

import (
	"fmt"
	"os"
	"strconv"
	"Fugo/ch2/tempconv"
)

func main() {
	for _, arg := range(os.Args[1:]) {
		t, err := strconv.ParseFloat(arg, 64)	// 类型转换64float
		if err != nil {
			fmt.Fprintf(os.Stderr, "cf:%v\n", err)
			os.Exit(0)
		}
		f := tempconv.Fahrenheit(t)
		c := tempconv.Celsius(t)

		fmt.Printf("%s = %s, %s = %s\n",
			f, tempconv.FToC(f), c, tempconv.CToF(c))
	}
}