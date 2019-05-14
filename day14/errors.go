package main

import (
	"errors"
	"fmt"
)

var errNotFound error = errors.New("not found error")

func main() {
	fmt.Printf("error: %v", errNotFound)
}

// 额外等方法定义特定的错误
type Error interface {
	Timeout() bool
	Temporary() bool
}