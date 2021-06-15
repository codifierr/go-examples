package main

import (
	"context"
	"fmt"
	"gokafka"
	"log"
	"time"

	"github.com/google/uuid"
)

func main() {
	//init config
	kafkaURL := "localhost:9092"
	topic := "test-topic"
	groupID := "test1"

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go runProducer(ctx, kafkaURL, topic)
	runConsumer(ctx, kafkaURL, topic, groupID)
}

func runConsumer(ctx context.Context, kafkaURL, topic, groupID string) {
	fmt.Println("init consumer")
	c, err := gokafka.NewConsumer(gokafka.SetKafkaNodes(kafkaURL), gokafka.SetAutoAck(false), gokafka.SetContext(ctx), gokafka.SetTopic(topic), gokafka.SetGroupId(groupID))
	go func() {
		for {
			err := <-c.ErrChan
			log.Println("Error in Consuming", err)
		}
	}()
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("starts consuming")
	for {
		m := <-c.MessageChan
		fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))

		c.CommitMessageChan <- m
	}
}

func runProducer(ctx context.Context, kafkaURL, topic string) {
	fmt.Println("init producer")
	producer, err := gokafka.NewProducer(gokafka.SetContext(ctx), gokafka.SetKafkaNodes(kafkaURL), gokafka.SetTopic(topic))
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println("start producing ... !!")
	for i := 0; ; i++ {
		key := fmt.Sprintf("Key-%d", i)
		msg := gokafka.Message{
			Key:   []byte(key),
			Value: []byte(fmt.Sprint(uuid.New())),
		}
		err := producer.SendMessage(msg)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("produced", key)
		}
		time.Sleep(1 * time.Second)
	}
}
