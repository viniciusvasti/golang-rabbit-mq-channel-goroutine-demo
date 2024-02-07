package main

import (
	"log"

	"github.com/viniciusvasti/golang-rabbit-mq-channel-goroutine-demo/internal/mqpublisher"
	"github.com/viniciusvasti/golang-rabbit-mq-channel-goroutine-demo/internal/services"
)

func main() {
	conn := services.NewRabbitMQConnection("amqp://admin:admin@localhost:5672/")
	defer conn.Close()

	log.Println("Publishing message to RabbitMQ")
	rabbitMQPublisher := mqpublisher.NewRabbitMQPublisher(conn)
	rabbitMQPublisher.Publish()
	log.Println("Message published")
}
