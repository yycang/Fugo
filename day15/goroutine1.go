package main

import (
	"fmt"
	"time"
)

// main 函数返回，程序退出， 不会等待任何非main协程的结束
func main() {
	fmt.Println("in main")

	go longWait()
	go shortWait()

	fmt.Println("about to sleep in main")

	time.Sleep(10 * 1e9)
	fmt.Println("at the end of main")
}

func longWait()  {
	fmt.Println("beginning longwait")
	time.Sleep(5 * 1e9)
	fmt.Println("end of longwait")
}

func shortWait()  {
	fmt.Println("beginning shortWait")
	time.Sleep(2 * 1e9)
	fmt.Println("end of shortWait")
}