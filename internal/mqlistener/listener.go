package mqlistener

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

var mockMessages = []string{}

var pattern = "amount: %d; price: %d"

func init() {
	for i := 0; i < 100; i++ {
		mockMessages = append(mockMessages, fmt.Sprintf(pattern, (i+1)*100, (i+1)*10))
	}
}

type RabbitMQListener struct {
}

func (rl RabbitMQListener) Listen() {
	fmt.Println("Listening to RabbitMQ")
	// Create a channel to write messages to
	dataChannel := make(chan string)
	// Define the amount of workers (concurrent processes) to be used
	workersAmount := 3
	for i := 0; i < workersAmount; i++ {
		// Start the workers (concurrent processes) which will process the messages in the channel
		go worker(dataChannel)
	}
	for _, message := range mockMessages {
		// Write the message to the channel
		dataChannel <- message
	}
	fmt.Println("Finished listening to RabbitMQ")
}

func worker(dataChannel chan string) {
	for message := range dataChannel {
		processMessage(message)
	}
}

func processMessage(message string) {
	time.Sleep(1 * time.Second)
	amount := extractAmount(message)
	price := extractPrice(message)
	fmt.Printf("Total: %s\n", calculateTotal(amount, price))
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
