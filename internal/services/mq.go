package services

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/viniciusvasti/golang-rabbit-mq-channel-goroutine-demo/internal/util"
)

func NewRabbitMQConnection(connString string) *amqp.Connection {
	log.Println("Connecting to RabbitMQ")
	conn, err := amqp.Dial(connString)
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	return conn
}

func CreateQueue(ch *amqp.Channel) amqp.Queue {
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
	return q
}
