package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctxb := context.Background()
	_, cancel := context.WithCancel(ctxb)
	defer cancel()
	_, ok := ctxb.Deadline()
	fmt.Println(ok)
	ctx1, cancel1 := context.WithTimeout(ctxb, 5*time.Second)
	defer cancel1()
	deadline, okd := ctx1.Deadline()
	getTimeOutFromContext(ctx1)
	fmt.Println(deadline)
	fmt.Println(okd)
	// go func() {
	// 	s := bufio.NewScanner(os.Stdin)
	// 	s.Scan()
	// 	time.Sleep(15 * time.Second)
	// 	cancel()
	// }()

	// time.AfterFunc(2*time.Second, cancel)
	if d, err := getTimeOutFromContext(ctx1); err == nil {
		sleepAndTalk(ctxb, d, "Hello")
	}

}

func sleepAndTalk(ctx context.Context, d time.Duration, msg string) {
	for {
		select {
		case <-time.After(d):
			fmt.Println("Hello")
		case <-ctx.Done():
			fmt.Println("bye bye")
			return
		}
	}
}

func getTimeOutFromContext(ctx context.Context) (time.Duration, error) {
	if t, ok := ctx.Deadline(); ok {
		return t.Sub(time.Now()), nil
	}
	return 0 * time.Second, fmt.Errorf("Context with deadline is required")
}
