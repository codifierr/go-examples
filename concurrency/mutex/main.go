// Files contains example of autorelease lock
package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var mutex sync.RWMutex

func main() {
	count := int32(0)

	fmt.Println("hi this is lock playground")
	mutex.Lock()
	fmt.Println("got lock")

	go func() {
		time.Sleep(50 * time.Millisecond)
		fmt.Println("count is ", count)
		if count < 1 {
			fmt.Println("Now is", atomic.AddInt32(&count, 1)) // Now is 1
			mutex.Unlock()
		}
	}()

	time.Sleep(50 * time.Millisecond)

	if count < 1 {
		mutex.Unlock()
		atomic.AddInt32(&count, 1)
		fmt.Println("released lock")
	}

	if count < 1 {
		mutex.Unlock()
		atomic.AddInt32(&count, 1)
		fmt.Println("released lock")
	}

	time.Sleep(50 * time.Millisecond)

}
