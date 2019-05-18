package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("starting the server")
	listener, err := net.Listen("tcp", "localhost:50000")

	if err != nil {
		fmt.Println("error", err.Error())
		return
	}
	// 监听并接受客户端链接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("error accepting", err.Error())
			return
		}
		go doServerStuff(conn)
	}
}

func doServerStuff(conn net.Conn)  {
	for {
		// 使用512字节缓冲data来读取客户端发送的数据
		buf := make([]byte, 512)
		len, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error reading", err.Error())
			return
		}
		fmt.Printf("received data: %v", string(buf[:len]))
	}
}