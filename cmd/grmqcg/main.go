package main

import "github.com/viniciusvasti/golang-rabbit-mq-channel-goroutine-demo/internal/mqlistener"

func main() {
	rabbitMQListener := mqlistener.RabbitMQListener{}
	rabbitMQListener.Listen()
}
