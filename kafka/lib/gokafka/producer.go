package gokafka

import (
	"context"
	"strings"

	kafka "github.com/segmentio/kafka-go"
)

type Producer struct {
	writer *kafka.Writer
	ctx    context.Context
}

type Message struct {
	Key    []byte
	Value  []byte
	Offset int64
}

func NewProducer(options ...Option) (*Producer, error) {
	ops, err := setup(options...)
	if err != nil {
		return nil, err
	}
	writer := getKafkaWriter(ops.kafkaNodes, ops.topic)
	p := &Producer{
		writer: writer,
		ctx:    ops.ctx,
	}
	go p.implicitClose()
	return p, nil
}

func (p Producer) SendMessage(message Message) error {
	kafkaMessage := kafka.Message{
		Key:   message.Key,
		Value: message.Value,
	}
	err := p.writer.WriteMessages(p.ctx, kafkaMessage)
	if err != nil {
		return err
	}
	return nil
}
func (p Producer) implicitClose() {
	<-p.ctx.Done()
	p.writer.Close()
}

func getKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	brokers := strings.Split(kafkaURL, ",")
	return &kafka.Writer{
		Addr:     kafka.TCP(brokers...),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}
