package main

import (
	"fmt"
	"net"
	"time"

	"github.com/jpillora/backoff"
)

// main is the entry point of the program.
func main() {
	// Create a backoff instance with default values.
	b := &backoff.Backoff{
		Min:    100 * time.Millisecond, // Minimum backoff duration
		Max:    10 * time.Second,       // Maximum backoff duration
		Factor: 2,                      // Factor to multiply the backoff duration
		Jitter: false,                  // Whether to add jitter to the backoff duration
	}

	// Print the backoff duration three times.
	printBackoffDuration(b, 3)

	// Reset the backoff instance.
	resetBackoff(b)

	// Create a backoff instance with jitter enabled.
	d := &backoff.Backoff{
		Jitter: true, // Enable jitter
	}

	// Print the backoff duration three times.
	printBackoffDuration(d, 3)

	// Reset the backoff instance.
	resetBackoff(d)

	// Create a backoff instance with a maximum duration of 5 minutes.
	c := &backoff.Backoff{
		Max: 5 * time.Minute, // Maximum backoff duration
	}

	// Continuously attempt to connect to example.com:5309.
	for {
		conn, err := net.Dial("tcp", "example.com:5309")
		if err != nil {
			// If there's an error, calculate the backoff duration and sleep.
			d := c.Duration()
			fmt.Printf("%s, reconnecting in %s", err, d)
			time.Sleep(d)
			continue
		}

		// Once connected, reset the backoff instance.
		resetBackoff(b)

		// Write data to the connection.
		conn.Write([]byte("hello world!"))

		// ... Read ... Write ... etc

		// Close the connection.
		conn.Close()

		// Continue to attempt to connect.
	}
}

// printBackoffDuration prints the backoff duration of the given backoff instance a specified number of times.
func printBackoffDuration(b *backoff.Backoff, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%s\n", b.Duration())
	}
}

// resetBackoff resets the given backoff instance.
func resetBackoff(b *backoff.Backoff) {
	fmt.Printf("Reset!\n")
	b.Reset()
}
