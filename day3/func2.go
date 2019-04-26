package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	if n, err := MySqrt(-5); err != nil {
		fmt.Println("return an error value", n, err)
	} else {
		fmt.Println("it's right", n, err)
	}

	fmt.Println(MySqrt(5))

}

func MySqrt(n float64)(res float64, err error) {
	if n < 0 {
		res = float64(math.NaN())
		err = errors.New("must greater than zero")
	} else {
		res = math.Sqrt(n)
		// err gets default value nil
	}
	return
}