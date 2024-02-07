package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/viniciusvasti/golang-rabbit-mq-channel-goroutine-demo/internal/mqpublisher"
	"github.com/viniciusvasti/golang-rabbit-mq-channel-goroutine-demo/internal/services"
	"github.com/viniciusvasti/golang-rabbit-mq-channel-goroutine-demo/internal/util"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		util.FailOnError(err, "Error loading .env file")
	}

	rabbitMQHost := os.Getenv("RABBITMQ_HOST")
	rabbitMQPort := os.Getenv("RABBITMQ_PORT")
	rabbitMQUser := os.Getenv("RABBITMQ_USER")
	rabbitMQPass := os.Getenv("RABBITMQ_PASS")
	connString := fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitMQUser, rabbitMQPass, rabbitMQHost, rabbitMQPort)

	conn := services.NewRabbitMQConnection(connString)
	defer conn.Close()

	log.Println("Publishing message to RabbitMQ")
	rabbitMQPublisher := mqpublisher.NewRabbitMQPublisher(conn)
	rabbitMQPublisher.Publish()
	log.Println("Message published")
}
