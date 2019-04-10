package main

import (
	"fmt"
	"math/cmplx"
)

var (
	toBe = false
	maxInt uint64 = 1 << 64 - 1
	z = cmplx.Sqrt(-5 + 12i)
)

func main(){
	fmt.Printf("Type: %T Value: %v\n",toBe,toBe)
	fmt.Printf("Type: %T Value: %v\n", maxInt, maxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}