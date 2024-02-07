package mqlistener

import "fmt"

type RabbitMQListener struct {
}

func (rl RabbitMQListener) Listen()  {
	fmt.Println("Listening to RabbitMQ")
}
