package main

import "fmt"

func main()  {
	var slice1 []int = make([]int, 10)
	for i:=0; i<len(slice1); i++ {
		slice1[i] = 5 * i
	}

	for i:=0; i<len(slice1); i++ {
		fmt.Printf("slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("length of slice1 %d\n", len(slice1))
	fmt.Printf("capacity of slice1 %d\n", cap(slice1))

	s := make([]byte, 10)
	s = s[2:4]

	fmt.Printf("length of slice1 %d\n", len(s))
	fmt.Printf("capacity of slice1 %d\n", cap(s))

	s1 := []string{"p", "o", "m", "e"}
	s2 := s1[2:]
	s2[1] = "t"
	fmt.Print(s1)
	fmt.Print(s2)
}