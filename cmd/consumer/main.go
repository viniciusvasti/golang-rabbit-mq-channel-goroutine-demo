package main

import (
	"fmt"
	"os"

	"github.com/viniciusvasti/golang-rabbit-mq-channel-goroutine-demo/internal/mqlistener"
	"github.com/viniciusvasti/golang-rabbit-mq-channel-goroutine-demo/internal/services"
)

func main() {
	rabbitMQHost := os.Getenv("RABBITMQ_HOST")
	rabbitMQPort := os.Getenv("RABBITMQ_PORT")
	rabbitMQUser := os.Getenv("RABBITMQ_USER")
	rabbitMQPass := os.Getenv("RABBITMQ_PASS")
	connString := fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitMQUser, rabbitMQPass, rabbitMQHost, rabbitMQPort)

	conn := services.NewRabbitMQConnection(connString)
	defer conn.Close()
	rabbitMQListener := mqlistener.NewRabbitMQListener(conn)
	rabbitMQListener.Listen()
}
