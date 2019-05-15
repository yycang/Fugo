package main

import (
	"Fugo/day15/parse"
	"fmt"
)

func main() {
	var examples = []string{
		"1 2 3 4 5",
		"100 13 25 12.5 5.44",
		"2 + 2 = 4",
		"1st class",
		"",
	}
	for _, ex := range examples {
		fmt.Printf("parsing %q:\n", ex)
		nums, err := parse.Parse(ex)
		if err != nil {
			fmt.Println(err)
			continue
		}
		fmt.Println(nums)
	}
}