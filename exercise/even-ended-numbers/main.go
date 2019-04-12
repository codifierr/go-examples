package main

import "fmt"

func main() {

	count := 1

	for i := 1000; i < 9999; i++ {

		for j := i; j < 9999; j++ { // i already provided number before

			n := i * j

			s := fmt.Sprintf("%d", n)

			if s[0] == s[len(s)-1] {
				count++
			}

		}

	}

	println(count)

}
