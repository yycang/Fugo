package main

import "fmt"


func main() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5]

	for i:=0; i<len(arr1); i++ {
		arr1[i] = i
	}

	for i:=0; i<len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("length of arr1 %d\n", len(arr1))
	fmt.Printf("length of slice1 %d\n", len(slice1))
	fmt.Printf("capacity of slice1 %d\n", cap(slice1))

	slice1 = slice1[0:4]
	for i:=0; i<len(slice1); i++ {
		fmt.Printf("slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("length of slice1 %d\n", len(slice1))
	fmt.Printf("capacity of slice1 %d\n", cap(slice1))

	b := []string{"g", "o", "l", "a", "n", "g"}
	fmt.Print(b[1:4])
	fmt.Print(b[2:])
	fmt.Print(b[:2])
	fmt.Print(b[:])
}