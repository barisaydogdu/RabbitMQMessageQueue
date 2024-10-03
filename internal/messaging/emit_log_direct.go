package messaging

import (
	"context"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func (m *Messaging) SendMessage(qName string, message string, rKey string) error {
	//5 saniye içinde tamamlanmazsa iptal et
	inctx, cancel := context.WithTimeout(m.ctx, 5*time.Second)
	defer cancel()

	err := m.rabbitMQ.Channel().PublishWithContext(inctx,
		qName+"_exchange", //exchange
		rKey,              // Routing Keys
		false,             // mandatory -herhangi bir kuyruk bu routing key ile eşleşmiyorsa hata alıp almayacağını belirler. False olduğunda hata fırlatmaz
		false,             //immediate- Mesajın hemen tüketilip tüketilmemesi gerektiğini belirler. False olduğunda hemen tüketilmeyebilir. Önce bir kuyrukta saklabilir.
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	if err != nil {
		return err
	}

	log.Printf("[x] Sent %s %s", message, rKey)

	return nil
}
