package main

import (
	"fmt"
)

// Pi constant
const Pi = 3.14

func main() {
	const World = "World"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")
	const Truth = true
	fmt.Println("Go Rules?", Truth)
}
