package main

import (
	"github.com/viniciusvasti/golang-rabbit-mq-channel-goroutine-demo/internal/mqlistener"
	"github.com/viniciusvasti/golang-rabbit-mq-channel-goroutine-demo/internal/services"
)

func main() {
	conn := services.NewRabbitMQConnection("amqp://admin:admin@rabbitmq:5672/")
	defer conn.Close()
	rabbitMQListener := mqlistener.NewRabbitMQListener(conn)
	rabbitMQListener.Listen()
}
