package main

import "fmt"

func main() {
	ch := make(chan int)
	//this will block
	/*<-ch
	fmt.Println("here")
	*/

	go func() {
		//send a number to the channel
		ch <- 43
		ch <- 44
		ch <- 45
		close(ch)
	}()
	for val := range ch {
		fmt.Println(val)
	}

}
