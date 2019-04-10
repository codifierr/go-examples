package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}
	
	//a[low : high]
	var a []int = primes[1:4]
	b := primes[2:4]

	fmt.Println(a)
	fmt.Println(b)
}