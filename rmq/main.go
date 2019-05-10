package main

import (
	"fmt"
	"net/http"

	events "github.com/codifierr/go-examples/rmq/events"
)

func main() {
	quit := make(chan bool, 1)
	defer close(quit)
	go events.SubscriptionNotificationPublisher(quit)
	http.HandleFunc("/subscribe", events.SubscribeEventsHandler)

	fmt.Print("Handler subscribe initialized")
	if err := http.ListenAndServe(":8090", nil); err != nil {
		fmt.Println(err)
	}
}
