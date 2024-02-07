package mqpublisher

import (
	"context"
	"fmt"
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/viniciusvasti/golang-rabbit-mq-channel-goroutine-demo/internal/util"
)

var mockMessages = []string{}

var pattern = "amount: %d; price: %d"

func init() {
	for i := 0; i < 100; i++ {
		mockMessages = append(mockMessages, fmt.Sprintf(pattern, (i+1)*100, (i+1)*10))
	}
}

type RabbitMQPublisher struct {
	conn *amqp.Connection
}

func NewRabbitMQPublisher(conn *amqp.Connection) RabbitMQPublisher {
	return RabbitMQPublisher{conn: conn}
}

func (rp RabbitMQPublisher) Publish() {
	ch, err := rp.conn.Channel()
	util.FailOnError(err, "Failed to open a channel")
	defer ch.Close()

	log.Println("Declaring a queue if not exists")
	q, err := ch.QueueDeclare(
		"cart", // name
		false,  // durable
		false,  // delete when unused
		false,  // exclusive
		false,  // no-wait
		nil,    // arguments
	)
	util.FailOnError(err, "Failed to declare a queue")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Publishing to RabbitMQ")
	for _, message := range mockMessages {
		err = ch.PublishWithContext(
			ctx,
			"",     // exchange
			q.Name, // routing key
			false,  // mandatory
			false,  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(message),
			},
		)
	}
	util.FailOnError(err, "Failed to publish a message")
	log.Println("Message published")
}
