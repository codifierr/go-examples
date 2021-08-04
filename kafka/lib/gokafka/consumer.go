package gokafka

import (
	"context"
	"strings"

	kafka "github.com/segmentio/kafka-go"
)

type Consumer struct {
	MessageChan       chan kafka.Message
	CommitMessageChan chan kafka.Message
	ErrChan           chan error
	reader            *kafka.Reader
	autoAck           bool
	ctx               context.Context
}

func NewConsumer(options ...Option) (*Consumer, error) {
	messageChannel := make(chan kafka.Message, 2)
	commitMessageChan := make(chan kafka.Message, 10)
	errChan := make(chan error)
	ops, err := setup(options...)
	if err != nil {
		return nil, err
	}
	reader := getKafkaReader(ops.kafkaNodes, ops.topic, ops.groupId)
	c := &Consumer{
		MessageChan:       messageChannel,
		CommitMessageChan: commitMessageChan,
		ErrChan:           errChan,
		reader:            reader,
		autoAck:           ops.autoAck,
		ctx:               ops.ctx,
	}
	go c.start()
	go c.implicitStop()
	return c, nil
}
func (c Consumer) start() {
	if c.autoAck {
		for {
			m, err := c.reader.ReadMessage(c.ctx)
			if err != nil {
				c.ErrChan <- err
			}
			c.MessageChan <- m
		}
	}
	// fire a gopher to read messages and send on message channel
	go func() {
		for {
			m, err := c.reader.FetchMessage(c.ctx)
			if err != nil {
				c.ErrChan <- err
			}
			c.MessageChan <- m
		}
	}()

	for {
		select {
		case m := <-c.CommitMessageChan:
			if err := c.reader.CommitMessages(c.ctx, m); err != nil {
				c.ErrChan <- err
			}
		case <-c.ctx.Done():
			return
		}
	}
}

func (c Consumer) implicitStop() {
	<-c.ctx.Done()
	c.Stop()
}

func (c Consumer) Stop() {
	c.reader.Close()
	close(c.MessageChan)
	close(c.ErrChan)
	close(c.CommitMessageChan)
}

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}
