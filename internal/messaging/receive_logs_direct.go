package messaging

import (
	"context"
	"errors"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func ReceiveMessage(ctx context.Context, ch *amqp.Channel, typ string, qName string) error {
	if qName == "" {
		return errors.New("qName is nil")
	}

	err := ch.ExchangeDeclare(
		qName+"_exchange", //name
		typ,               //exchange type
		true,              //durable- yeniden başlatıldığında exhange kaybolmaz
		false,             //auto-deleted
		false,             //internal - içsel
		false,             //nowait
		nil,
	)
	if err != nil {
		return err
	}

	q, err := ch.QueueDeclare(
		qName, //name
		false, //durable
		false, //delete when unused
		true,  //exclusive- true olduğunda kuyruk sadece bu kanal ile kullanılabilir.
		false, //no-wait
		nil,   //arguments
	)
	if err != nil {
		return err
	}

	log.Printf("Binding queue %s to exchange %s with routing key", q.Name, "logs_direct")
	err = ch.QueueBind(
		q.Name,            //Queue name
		"",                //Routing key
		qName+"_exchange", //exchange
		false,             //no-wait
		nil)               //args)
	if err != nil {
		return err
	}

	//tüketici
	msgs, err := ch.Consume(
		q.Name, //queue
		"",     //consumer
		true,   // auto-ack- Mesajı otomatik olarak onaylamak
		false,  //exclusive- sadece bu tüketici mi dinleyecek
		false,  //nolocal- bu bağlantı üzerinden gönderilen messjların bu bağlatıdan dinleyen bir tüketiciye iletilip iletilmediğini kontrol eder.
		false,  //nowait- RabbitMQ'dan bir onay mesajı beklenir ve bağlantı kurulduğu zaman bir yanıt alınır.
		nil,
	)
	if err != nil {
		return err
	}

	inctx, cancel := context.WithCancel(ctx)
	defer cancel()

	go func() {
		for d := range msgs {
			log.Printf("[x] %s", d.Body)
		}
	}()

	log.Printf("[*] Waiting for logs. To Exit press CTRL+C")
	<-inctx.Done()

	return nil
}
