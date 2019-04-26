package main

import (
	"fmt"
	"time"
)

func main() {
	k := 6
	switch k {
	case 4:
		fmt.Println("k<=4")
		fallthrough
	case 5:
		fmt.Println("k<=5")
		fallthrough
	case 6:
		fmt.Println("k<=6")
		fallthrough
	case 7:
		fmt.Println("k<=7")
		fallthrough
	case 8:
		fmt.Println("k<=8")
		fallthrough
	default:
		fmt.Println("default")
	}
	month := time.Now().Month()
	fmt.Println(season(int(month)))
}

func season(month int) string {
	switch month {
	case 12, 1, 2:
		return "winter"
	case 3, 4, 5:
		return "spring"
	case 6, 7, 8:
		return "summer"
	case 9, 10, 11:
		return "autumn"
	}
	return "unknown"
}
