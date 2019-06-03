// 如何遍历一个通道
ch := make(chan type,buf)
for v := range ch {

}

// 如何检测一个通道是否关闭
for {
	if input, open := <-ch; !open {
		break
	}
	fmt.Printf("%s", input)
}

// 如何通过一个通道让主程序等待直到协程完成
ch := make(chan int)

go func() {
	ch <- 1
}()
doSomething()
<-ch


// 通道的工厂模板
func pump() chan int {
	ch := make(chan int)
	go func() {
		for i := 0; i++ {
			ch <- i
		}
	}()
	return ch
}

// 超时模板
timeout := make(chan bool, 1)
go func() {
	time.Sleep(1e9)
	timeout <- true
}()
select {
	case <-ch:
		// has occurred
	case <-timeout:
		// has time out
}

// 如何使用输入通道和输出通道代理锁
func Worker(in, out chan *Task) {
	for {
		t := <-in
		process(t)
		out <- t
	}
}


