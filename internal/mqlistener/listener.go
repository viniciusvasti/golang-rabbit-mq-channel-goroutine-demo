package mqlistener

import (
	"fmt"
	"strconv"
	"strings"
)

var mockMessages = []string{
	"amount: 100; price: 10",
	"amount: 200; price: 20",
	"amount: 300; price: 30",
	"amount: 400; price: 40",
	"amount: 500; price: 50",
	"amount: 600; price: 60",
	"amount: 700; price: 70",
	"amount: 800; price: 80",
	"amount: 900; price: 90",
	"amount: 1000; price: 100",
}

type RabbitMQListener struct {
}

func (rl RabbitMQListener) Listen() {
	fmt.Println("Listening to RabbitMQ")
	for _, message := range mockMessages {
		processMessage(message)
	}
	fmt.Println("Finished listening to RabbitMQ")
}

func processMessage(message string) {
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
