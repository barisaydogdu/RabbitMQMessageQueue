package main

import (
	"context"
	"flag"
	"log"
	"os"
	"os/signal"
	"syscall"

	messaging "github.com/barisaydogdu/MessageQueuesRabbitMQ/internal/messaging"
	rabbitMQ "github.com/barisaydogdu/MessageQueuesRabbitMQ/pkg/rabbitMQ"
)

var (
	typ   string
	qType string
	rKey  string
	qName string
	msg   string
)

func main() {

	flag.StringVar(&typ, "typ", "", "Application running type")
	flag.StringVar(&qType, "qType", "", "")          //Qeueu Type
	flag.StringVar(&qName, "qName", "", "")          //QName
	flag.StringVar(&rKey, "rKey", "", "Routing Key") //Routing Key
	flag.StringVar(&msg, "msg", "", "")              //Message
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
	rabbitMQClient, err := rabbitMQ.ConnectToRabbitMQ(url)
	rabbitMQ.FailOnError(err, "Failed to connect to RabbitMQ")
	defer rabbitMQClient.CloseConnection()

	err = rabbitMQClient.OpenChannel()
	rabbitMQ.FailOnError(err, "Failed to open a channel")
	defer rabbitMQClient.CloseChannel()

	//Messaging struct
	messagingClient := messaging.NewMessaging(ctx, *rabbitMQClient)

	switch typ {
	case "consumer":

		if qType == "" || qName == "" || rKey == "" {
			log.Fatal("Queue type name and routing key cannot be empty")
		}
		go messagingClient.ReceiveMessage(qType, qName, rKey)
		break

	case "publisher":
		if qName == "" || rKey == "" || msg == "" {
			log.Fatal("Queue type,name, Routing key and message cannot be empty")
		}
		messagingClient.SendMessage(qName, msg, rKey)
		signal.Stop(ch)
		cancel()
		break
	}

	go func() {
		select {
		case <-ch:
			_ = rabbitMQClient.CloseChannel()
			rabbitMQClient.CloseConnection()
			cancel()
		}
	}()

	println("application started...")

	select {
	case <-ctx.Done():
		//
		println(" application shutting down...")
	}
}
