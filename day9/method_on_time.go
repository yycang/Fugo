package main

import (
	"fmt"
	"time"
)

type myTime struct {
	time.Time
}

func (t myTime) first4Char() string {
	return t.Time.String()[0:4]
}

func main() {
	m := myTime{time.Now()}
	fmt.Println("full time now:", m.String())

	fmt.Println("first 3 char", m.first4Char())
}