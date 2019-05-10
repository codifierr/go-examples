package events

import (
	"fmt"
	"net/http"
)

//SubscribeEventsHandler func
func SubscribeEventsHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprint(rw, "Processing subscription request.")
	topicsToSubscribe := []string{"cloudguest", "rapids", "presence", "visualrf"}

	subscriptionMap := newSubscriptionsMap()
	for _, topics := range topicsToSubscribe {
		subscriptionMap[topics] = true
	}
	go func() {
		subscriptionEvents <- subscriptionMap
	}()
}
