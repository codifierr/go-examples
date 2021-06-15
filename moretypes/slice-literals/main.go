package main

import "fmt"

func main() {
	q := []int{1, 2, 3, 4}
	fmt.Println(q)

	r := []bool{true, false, true, false}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{1, true},
		{2, true},
		{3, false},
		{4, false},
		{5, true},
	}

	fmt.Println(s)

	fmt.Println(q, r, s)
}
