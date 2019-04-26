package main

import (
	"fmt"
	"runtime"
)

var prompt = "enter a digit, or %s to quit."

func init()  {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z")
	} else {
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
}

func main()  {
	var first int = 10
	var cond int

	if first <= 0 {
		fmt.Println("first less than or equal zero")
	} else if first >0 && first <= 20 {
		fmt.Println("first between 0 and 20")
	} else {
		fmt.Println("others")
	}

	if cond = 5; cond > 10 {
		fmt.Println("cond greater than 10")
	} else {
		fmt.Println("cond is not greater than 10")
	}
}