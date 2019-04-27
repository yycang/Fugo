package main

import "fmt"

func main() {
	fmt.Printf("%d is even : %t\n", 5, even(5))
	fmt.Printf("%d is odd : %t\n", 6, odd(6))
	fmt.Printf("%d is odd : %t\n", 7, odd(7))
	a(10)
}

func a(n int) {
	if n == 0 {
		return
	}
	println(n)
	a(n - 1)
}

func even(n int) bool {
	if n == 0 {
		return true
	}
	return odd(RevSign(n) - 1)
}

func odd(n int) bool  {
	if n == 0 {
		return false
	}
	return even(RevSign(n) - 1)
}

func RevSign(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
