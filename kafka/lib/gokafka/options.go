package gokafka

import (
	"context"
	"errors"
)

type options struct {
	topic      string
	kafkaNodes string
	groupId    string
	ctx        context.Context
	autoAck    bool
}
type Option func(*options) error

func defaultOptions() *options {
	return &options{
		kafkaNodes: "localhost:9092",
		autoAck:    true,
	}
}

func SetTopic(topic string) Option {
	return func(o *options) error {
		if topic != "" {
			o.topic = topic
			return nil
		}
		return errors.New("subscription topic is mandatory")
	}
}
func SetKafkaNodes(kafkaNodes string) Option {
	return func(o *options) error {
		if kafkaNodes != "" {
			o.kafkaNodes = kafkaNodes
			return nil
		}
		return errors.New("kafka nodes is mandatory param")
	}
}
func SetGroupId(groupId string) Option {
	return func(o *options) error {
		o.groupId = groupId
		return nil
	}
}
func SetContext(ctx context.Context) Option {
	return func(o *options) error {
		if isCancellableContext(ctx) {
			o.ctx = ctx
			return nil
		}
		return errors.New("cancellable context is required")
	}
}

func SetAutoAck(autoAck bool) Option {
	return func(o *options) error {
		o.autoAck = autoAck
		return nil
	}
}

//IsCancellableContext func checks if this context is cancellable
func isCancellableContext(ctx context.Context) bool {
	if ctx != nil {
		return ctx.Done() != nil
	}
	return false
}

// setup defaults
func setup(options ...Option) (*options, error) {
	opts := defaultOptions()
	for _, opt := range options {
		opt(opts)
	}
	return opts, nil
}
