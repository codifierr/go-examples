package events

import (
	"encoding/json"
	"log"

	"github.com/streadway/amqp"
)

var (
	//SubscriptionEvents channel
	subscriptionEvents = make(chan map[string]bool, 20)
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type subscriptionNotification struct {
	Subscriptions []*subscription `json:"subscriptions"`
}

type subscription struct {
	Topic  string `json:"topic"`
	Status bool   `json:"status"`
}

// NewSubscriptionsMap function to create sunscription map
func newSubscriptionsMap() map[string]bool {
	return map[string]bool{
		"monitoring":     false,
		"apprf":          false,
		"presence":       false,
		"cloudguest":     false,
		"usermanagement": false,
		"visualrf":       false,
		"contextengine":  false,
		"ucc":            false,
		"lync":           false,
		"audit":          false,
		"clickstream":    false,
		"location":       false,
		"rapids":         false,
		"syslog":         false,
		"stream-stats":   false,
	}
}

func SubscriptionNotificationPublisher(quit chan bool) {
	// code to create connection with rmq start
	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"channel1", // name
		false,      // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	failOnError(err, "Failed to declare a queue")
	// code to create connection with rmq end
	//connRetryCounter := 0
	for {
		select {
		case subscriptionsMap := <-subscriptionEvents:
			if json, err := json.Marshal(prepareSubscriptionNotificationMessage(subscriptionsMap)); err != nil {
				log.Fatal("Unable to prepare message")

			} else {
				err = ch.Publish(
					"",     // exchange
					q.Name, // routing key
					false,  // mandatory
					false,  // immediate
					amqp.Publishing{
						ContentType: "text/plain",
						Body:        []byte(json),
					})
				if err != nil {
				}
				failOnError(err, "Failed to publish a message")
			}

		case <-quit:
			return
		}
	}

}

func prepareSubscriptionNotificationMessage(subscriptionMap map[string]bool) *subscriptionNotification {
	subscriptions := make([]*subscription, len(subscriptionMap))
	i := 0
	for key, value := range subscriptionMap {
		subscription := &subscription{
			key,
			value,
		}
		subscriptions[i] = subscription
		i++
	}

	return &subscriptionNotification{
		subscriptions,
	}
}
