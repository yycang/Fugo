package main

import "fmt"

func main() {
	ch1 := make(chan int)
	go pump(ch1)
	fmt.Println(<-ch1)
}


// 由于没有接收者，通道会处于阻塞的状态
func pump(ch chan int)  {
	for i:=0; ; i++ {
		ch <- i
	}
}