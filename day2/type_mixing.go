package main

import (
	"fmt"
	"math"
)

func main()  {
	var a int
	var b int32
	a = 15
	// b = a + a 	error

	var n int16 = 33
	var m int32

	m = int32(n)
	fmt.Printf("32bit int为%d\n", m)
	fmt.Printf("16bit int为%d\n", n)

	res, res1 := Uint8FromInt(a)
	fmt.Printf("%d", res)
}


func Uint8FromInt(n int) (uint8, error) {
	if 0 <= n && n <= math.MaxUint8 {
		return uint8(n), nil
	}
	return 0, fmt.Errorf("%d is out of the uint8 range", n)
}

func IntFromFloat64(x float64) int {
	if math.MinInt32 <= x && x <= math.MaxInt32 {
		whole, fraction := math.Modf(x)
		if fraction >= 0.5 {
			whole ++
		}
		return int(whole)
	}
	panic(fmt.Sprintf("%g is out of the int32 range", x))
}