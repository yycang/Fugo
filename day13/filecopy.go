package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	CopyFile("target.txt", "source.txt")
	fmt.Println("copy done")
}

func CopyFile(d, s string) (written int64, err error) {
	src, err := os.Open(s)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(d)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}