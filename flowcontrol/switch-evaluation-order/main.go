package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println("Today : %V", today)
	fmt.Println("time.Saturday : %V", time.Saturday)
	fmt.Println("time.Now().UTC : %V", time.Now().UTC())
	switch time.Saturday {
	case today:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	case today + 3:
		fmt.Println("In three days.")
	default:
		fmt.Println("Too far away.")
	}
}
