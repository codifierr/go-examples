package main

import (
	"fmt"
	"os"
	"time"

	"github.com/memphisdev/memphis.go"
)

func main() {
	conn, err := memphis.Connect("broker.sandbox.memphis.dev", "ssingh", "XrHmszw6rgm8IyOPNNTy")
	if err != nil {
		os.Exit(1)
	}
	defer conn.Close()

	// Send Message

	p, err := conn.CreateProducer("teststation", "myProducer")
	if err != nil {
		fmt.Printf("Producer creation failed: %v", err)
		os.Exit(1)
	}

	err = p.Produce([]byte("You have a message!"))
	if err != nil {
		fmt.Printf("Produce failed: %v", err)
		os.Exit(1)
	}
	fmt.Println("Message Sent")

	// Receive Message

	consumer, err := conn.CreateConsumer("teststation", "myConsumer", memphis.PullInterval(15*time.Second))
	if err != nil {
		fmt.Printf("Consumer creation failed: %v", err)
		os.Exit(1)
	}

	handler := func(msgs []*memphis.Msg, err error) {
		if err != nil {
			fmt.Printf("Fetch failed: %v", err)
			return
		}

		for _, msg := range msgs {
			fmt.Println("Received Message : " + string(msg.Data()))
			msg.Ack()
		}
	}

	consumer.Consume(handler)
	time.Sleep(5 * time.Second)
}
