package main

import (
	"fmt"
)

func main() {

	pow := []int{1, 2, 3, 4, 5, 6}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}
