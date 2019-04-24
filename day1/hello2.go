package main

import "fmt"

type Color int

const (
	RED Color = iota  // 0
	ORANGE
	GREEN
	BLUE
	YELLO 	// 4
)

func main() {
	fmt.Println(YELLO)
}

