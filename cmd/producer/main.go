package main

import (
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/viniciusvasti/golang-rabbit-mq-channel-goroutine-demo/internal/mqpublisher"
	"github.com/viniciusvasti/golang-rabbit-mq-channel-goroutine-demo/internal/util"
)

func main() {
	log.Println("Connecting to RabbitMQ")
	conn, err := amqp.Dial("amqp://admin:admin@localhost:5672/")
	util.FailOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	log.Println("Publishing message to RabbitMQ")
	rabbitMQPublisher := mqpublisher.NewRabbitMQPublisher(conn)
	rabbitMQPublisher.Publish()
	log.Println("Message published")
}
