package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// 打开链接
	conn, err := net.Dial("tcp", "localhost:50000")
	if err != nil {
		fmt.Println("error dialing", err.Error())
		return
	}

	inputReader := bufio.NewReader(os.Stdin)
	fmt.Println("first what is your name")
	clientName, _ := inputReader.ReadString('\n')
	trimmedClient := strings.Trim(clientName, "\n")
	// 给服务器发送信息直到程序退出

	for {
		fmt.Println("type q to quit")
		input, _ := inputReader.ReadString('\n')
		trimmedInput := strings.Trim(input, "\n")

		if trimmedInput == "q" {
			return
		}
		_, err = conn.Write([]byte(trimmedClient + " says:" + trimmedInput))
	}
}