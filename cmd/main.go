package main

import (
	"context"
	"flag"
	"os"
	"os/signal"
	"syscall"

	messaging "github.com/barisaydogdu/MessageQueuesRabbitMQ/internal/messaging"
	rabbitMQ "github.com/barisaydogdu/MessageQueuesRabbitMQ/pkg/rabbitMQ"
)

var (
	typ   string
	qType string
	qName string
	msg   string
)

func main() {
	flag.StringVar(&typ, "typ", "", "Application running type")
	flag.StringVar(&qType, "qType", "direct", "")
	flag.StringVar(&qName, "qName", "emptyQueue", "")
	flag.StringVar(&msg, "msg", "", "")

	flag.Parse()

	switch typ {
	case "consumer", "publisher":
		break
	default:
		println("invalid typ")
		os.Exit(128)
	}

	ctx, cancel := context.WithCancel(context.Background())
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	url := "amqp://guest:guest@localhost:5672/"
	conn, err := rabbitMQ.ConnectToRabbitMQ(url)
	rabbitMQ.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	publisherChannel, err := rabbitMQ.OpenChannel(conn)
	rabbitMQ.FailOnError(err, "Failed to open a channel")
	defer publisherChannel.Close()

	switch typ {
	case "consumer":
		go messaging.ReceiveMessage(ctx, publisherChannel, qType, qName)
		break
	case "publisher":
		signal.Stop(ch)
		cancel()
		break
	}

	go func() {
		select {
		case <-ch:
			_ = publisherChannel.Close()
			_ = conn.Close()
			cancel()
		}
	}()

	println("application started...")

	select {
	case <-ctx.Done():
		//
		println("applicatin shutting down...")
	}
}
