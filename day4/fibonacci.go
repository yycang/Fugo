package main

import "fmt"

func main() {
	result := 0
	for i:=0; i<=10; i++ {
		result = finbonacci(i)
		fmt.Printf("No %d finbonacci is %d\n", result)
	}
}

func finbonacci(n int) (res int)  {
	if n <= 1 {
		res = 1
	} else {
		res = finbonacci(n - 1) + finbonacci(n - 2)
	}
	return
}