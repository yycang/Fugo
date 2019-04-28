package main

import (
	"fmt"
	"strings"
)

func main() {
	var f = MakeAddSuffix(".jpg")
	fmt.Print(f("file"))
}

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
