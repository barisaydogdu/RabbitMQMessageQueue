package messaging

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/barisaydogdu/MessageQueuesRabbitMQ/util"
	amqp "github.com/rabbitmq/amqp091-go"
)

func SendMessage(ctx context.Context, ch *amqp.Channel) error {
	//5 saniye içinde tamamlanmazsa iptal et
	inctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	err := ch.PublishWithContext(inctx,
		"logs_direct",              //exchange
		util.SeverityFrom(os.Args), // Routing Keys
		false,                      // mandatory -herhangi bir kuyruk bu routing key ile eşleşmiyorsa hata alıp almayacağını belirler. False olduğunda hata fırlatmaz
		false,                      //immediate- Mesajın hemen tüketilip tüketilmemesi gerektiğini belirler. False olduğunda hemen tüketilmeyebilir. Önce bir kuyrukta saklabilir.
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(util.BodyFrom(os.Args)),
		})
	if err != nil {
		return err
	}

	log.Printf("[x] Sent %s", util.BodyFrom(os.Args))

	return nil
}
