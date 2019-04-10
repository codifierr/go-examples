package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	//Slices are like references to arrays
	
	a := names[0:2]
	b := names[1:3]

	fmt.Println(a,b)
	
	// change of value in slice will reflect in original array
	b[0] = "XXX"

	fmt.Println(a,b)

	fmt.Println(names)

}