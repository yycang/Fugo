package main

import "fmt"

func main() {
	seasons := []string{"spring", "summery", "autumn", "winter"}
	for ix, season := range seasons {
		fmt.Printf("season %d is %s\n", ix, season)
	}

	var season string
	for _, season = range seasons {
		fmt.Printf("%s\n", season)
	}

	items := [...]int{10, 20, 30, 40, 50}
	for ix := range items {
		items[ix] *= 2
	}
	fmt.Print(items)
}