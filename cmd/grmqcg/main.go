package main

import (
	"time"

	"github.com/viniciusvasti/golang-rabbit-mq-channel-goroutine-demo/internal/mqlistener"
)

func main() {
	startTime := time.Now()
	rabbitMQListener := mqlistener.RabbitMQListener{}
	rabbitMQListener.Listen()
	elapsedTime := time.Since(startTime)
	println("Elapsed time: ", elapsedTime / time.Second, " seconds")
}
