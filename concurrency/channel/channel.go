package main

import "fmt"

func main() {
	ch := make(chan int)
	ch1 := make(chan int)
	//this will block
	//<-ch
	fmt.Println("here")

	go func() {
		//send a number to the channel
		ch <- 43
		fmt.Println("43")
		ch <- 44
		fmt.Println("44")
		ch <- 45
		fmt.Println("45")
		//close(ch)
	}()
	// for val := range ch {
	// 	fmt.Println(val)
	// }
	<-ch1

}
