package mqlistener

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/viniciusvasti/golang-rabbit-mq-channel-goroutine-demo/internal/util"
)

var workerPoolSize = 5

type RabbitMQListener struct {
	conn *amqp.Connection
}

func NewRabbitMQListener(conn *amqp.Connection) RabbitMQListener {
	return RabbitMQListener{conn: conn}
}

func (rl RabbitMQListener) Listen() {
	ch, err := rl.conn.Channel()
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

	log.Println("Listening to RabbitMQ")
	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	util.FailOnError(err, "Failed to register a consumer")

	forever := make(chan string)

	for i := 0; i < workerPoolSize; i++ {
		go worker(msgs)
	}

	// makes the listener run forever
	<-forever
}

func worker(dataChannel <-chan amqp.Delivery) {
	for message := range dataChannel {
		processMessage(string(message.Body))
	}
}

func processMessage(message string) {
	time.Sleep(1 * time.Second)
	amount := extractAmount(message)
	price := extractPrice(message)
	log.Printf("Total: %s\n", calculateTotal(amount, price))
}

func calculateTotal(amount string, price string) string {
	amountInt, _ := strconv.Atoi(amount)
	priceInt, _ := strconv.Atoi(price)
	return fmt.Sprintf("%d", amountInt*priceInt)
}

func extractPrice(message string) string {
	price := strings.Trim(strings.Split(strings.Split(message, ";")[1], ":")[1], " ")
	return price
}

func extractAmount(message string) string {
	amount := strings.Trim(strings.Split(strings.Split(message, ";")[0], ":")[1], " ")
	return amount
}
