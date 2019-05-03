package main

import (
	"fmt"
	"./struct_pack"
)

func main() {
	struct1 := new(structPack.ExpStruct)
	struct1.Mi1 = 10
	struct1.Mf1 = 14.1

	fmt.Printf("mi1 = %d\n", struct1.Mi1)
	fmt.Printf("mf1 = %f\n", struct1.Mf1)
}

