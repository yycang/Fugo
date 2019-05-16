package main

import (
	"fmt"
	"time"
)


// sendData通过通道ch来发送5个字符串，而getData按照顺序接受他们并打印
func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(1e9)
}

func sendData(ch chan string)  {
	ch <- "beijing"
	ch <- "tianjing"
	ch <- "shanghai"
	ch <- "guangdong"
}

func getData(ch chan string)  {
	var input string
	for {
		input = <- ch
		fmt.Printf("%s\n", input)
	}
}