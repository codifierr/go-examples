package main

import (
	"fmt"
	"math/rand"
	"net"
	"time"

	"github.com/jpillora/backoff"
)

func main() {
	b := &backoff.Backoff{
		//These are the defaults
		Min:    100 * time.Millisecond,
		Max:    10 * time.Second,
		Factor: 2,
		Jitter: false,
	}

	fmt.Printf("%s\n", b.Duration())
	fmt.Printf("%s\n", b.Duration())
	fmt.Printf("%s\n", b.Duration())

	fmt.Printf("Reset!\n")
	b.Reset()

	fmt.Printf("%s\n", b.Duration())

	d := &backoff.Backoff{
		Jitter: true,
	}

	rand.Seed(42)

	fmt.Printf("%s\n", d.Duration())
	fmt.Printf("%s\n", d.Duration())
	fmt.Printf("%s\n", d.Duration())

	fmt.Printf("Reset!\n")
	d.Reset()

	fmt.Printf("%s\n", d.Duration())
	fmt.Printf("%s\n", d.Duration())
	fmt.Printf("%s\n", d.Duration())

	c := &backoff.Backoff{
		Max: 5 * time.Minute,
	}

	for {
		conn, err := net.Dial("tcp", "example.com:5309")
		if err != nil {
			d := c.Duration()
			fmt.Printf("%s, reconnecting in %s", err, d)
			time.Sleep(d)
			continue
		}
		//connected
		b.Reset()
		conn.Write([]byte("hello world!"))
		// ... Read ... Write ... etc
		conn.Close()
		//disconnected
	}

}
