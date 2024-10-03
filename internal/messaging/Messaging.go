package messaging

import (
	"context"

	rabbitMQ "github.com/barisaydogdu/MessageQueuesRabbitMQ/pkg/rabbitMQ"
)

type Messaging struct {
	ctx      context.Context
	rabbitMQ rabbitMQ.RabbitMQClient
}

func NewMessaging(ctx context.Context, rabbitMQ rabbitMQ.RabbitMQClient) *Messaging {
	return &Messaging{
		ctx:      ctx,
		rabbitMQ: rabbitMQ,
	}
}

func (m *Messaging) Context() context.Context {
	return m.ctx
}
