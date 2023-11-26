package main

import (
	"testing"
	"time"

	"github.com/jpillora/backoff"
)

func TestMain(t *testing.T) {
	t.Run("TestPrintBackoffDuration", func(t *testing.T) {
		b := &backoff.Backoff{
			Min:    100 * time.Millisecond,
			Max:    10 * time.Second,
			Factor: 2,
			Jitter: false,
		}
		printBackoffDuration(b, 3)
		// Add assertions here
	})

	t.Run("TestResetBackoff", func(t *testing.T) {
		b := &backoff.Backoff{
			Min:    100 * time.Millisecond,
			Max:    10 * time.Second,
			Factor: 2,
			Jitter: false,
		}
		resetBackoff(b)
		// Add assertions here
	})

	t.Run("TestPrintBackoffDurationWithJitter", func(t *testing.T) {
		d := &backoff.Backoff{
			Jitter: true,
		}
		printBackoffDuration(d, 3)
		// Add assertions here
	})

	t.Run("TestResetBackoffWithMaxDuration", func(t *testing.T) {
		c := &backoff.Backoff{
			Max: 5 * time.Minute,
		}
		resetBackoff(c)
		// Add assertions here
	})

	t.Run("TestMain", func(t *testing.T) {
		// Add test cases here
	})
}
