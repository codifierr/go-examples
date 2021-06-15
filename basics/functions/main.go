package main

import (
	"fmt"
)

func main() {
	fmt.Println("Adding of numbers 12 and 13 : ", add(12, 13))
}

func add(x, y int) int {
	return x + y
}
