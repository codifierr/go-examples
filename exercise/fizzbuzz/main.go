package main

import "fmt"

func main() {
	// loop for number 1 to 20
	for i := 1; i <= 20; i++ {
		//fmt.Printf("Number : %d\n", i)
		switch {
		// check i is divisible by 3 and 5
		case i%3 == 0 && i%5 == 0:
			{
				fmt.Println("fizz buzz")
			}
		// check if is divisible by 3
		case i%3 == 0:
			{
				fmt.Println("fizz")
			}
		// check if i divisible by 5
		case i%5 == 0:
			{
				fmt.Println("buzz")
			}
		// print number if non of the conditions matches
		default:
			fmt.Println(i)
		}
	}
}
