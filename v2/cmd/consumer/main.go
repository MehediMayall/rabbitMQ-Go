package main

import (
	"fmt"
	"log"

	"github.com/mehedimayall/rabbitmq-go/internal/rabbitmq"
)

func print[T any](values T) {
	fmt.Println(values)
}

const queueName = "user_created"
const userExchange = "user_events"

func main() {
	conn, err := rabbitmq.CreateConnection("localhost:5672", "mehedi", "mehedi007", "booking")

	if err != nil {
		log.Fatalln(err.Error())
	}
	defer conn.Close()

	client, err := rabbitmq.CreateClient(conn)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer client.Close()

	messageBus, err := client.Consume(queueName, "email-service", true)
	if err != nil {
		log.Fatalln(err.Error())
	}

	var blocking chan struct{}

	go func() {
		for message := range messageBus {
			print(string(message.Body))
		}
	}()

	print("Consuming, press ctrl+c to exit...")
	<-blocking

}
