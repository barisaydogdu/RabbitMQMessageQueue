package rabbitMQ

import (
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQClient struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

func ConnectToRabbitMQ(url string) (*RabbitMQClient, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}
	return &RabbitMQClient{conn: conn}, nil
}

// Getter for rabbitmqservice channel
func (r *RabbitMQClient) Channel() *amqp.Channel {
	return r.channel
}

func (r *RabbitMQClient) OpenChannel() error {
	ch, err := r.conn.Channel()
	if err != nil {
		return err
	}
	r.channel = ch
	return nil
}

func (r *RabbitMQClient) CloseChannel() error {
	if r.channel != nil {
		return r.channel.Close()
	} else {
		return fmt.Errorf("channel is already closed or nil")
	}
}

// Close connection struct method
func (r *RabbitMQClient) CloseConnection() {
	if r.conn != nil {
		r.conn.Close()
	}
}

func FailOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s:%s", msg, err)
	}
}
