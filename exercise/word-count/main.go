package main

import (
	"fmt"
	"strings"
)

func main() {
	text := `Needles and pins
	Needles and pins
	Saw me sails
	To catch winds
	`
	textarr := strings.Fields(strings.ToLower(text))
	count := make(map[string]int) // make a map
	for _, v := range textarr {
		count[v]++
	}

	fmt.Println(count)
}
