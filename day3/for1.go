package main

import "fmt"

const (
	Fizz = 3
	Buzz = 5
	FizzBuzz = 15
)

func main()  {
	/*
	for i:=1; i<=25; i++ {
		for j:=1; j<=i; j++ {
			fmt.Printf("G")
		}
		fmt.Println()
	}

	for i:=0; i<=10; i++ {
		fmt.Printf("%b:%b\n", i, ^i)
	}

	 */
	/*
	for i:=1; i<=100; i++ {
		switch {
		case i%FizzBuzz == 0:
			fmt.Println("FizzBuzz")
		case i%Fizz == 0:
			fmt.Println("Fizz")
		case i%Buzz == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}

	 */
	for i:=1; i<=10; i++ {
		for j:=1; j<=20; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}