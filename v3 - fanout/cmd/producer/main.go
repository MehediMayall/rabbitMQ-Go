package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/mehedimayall/rabbitmq-go/internal/models"
	"github.com/mehedimayall/rabbitmq-go/internal/rabbitmq"
)

func print[T any](values T) {
	fmt.Println(values)
}

const queueName = "user_created"
const userExchange = "user_events"
const hostName = "localhost:5672"
const userName = "mehedi"
const password = "mehedi007"
const virtualhost = "booking"

func main() {
	// Create connection
	conn, err := rabbitmq.CreateConnection(hostName, userName, password, virtualhost)

	if err != nil {
		log.Fatalln(err.Error())
	}

	// defer conn.Close()

	// Create Client
	client, err := rabbitmq.CreateClient(conn)
	if err != nil {
		log.Fatalln(err.Error())
	}

	defer client.Close()

	// Create Queue
	queue, err := client.CreateQueue(queueName, false, false)
	if err != nil {
		log.Fatalln(err)
	}

	// Bind Queue
	if err := client.BindQueue(queue.Name, "user.created.*", userExchange); err != nil {
		log.Fatalln(err)
	}

	// Create timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Create Message

	var users = models.GetUsers()

	// Send Message and wait for the confirmation

	for _, user := range users {

		usersInJson, err := json.Marshal(user)
		if err != nil {
			continue
		}

		var userInBytes = []byte(usersInJson)

		isPublished, err := client.SendAndGetConfirmed(ctx, userExchange, "user.created.nj", *client.CreateOptionsPersistent(userInBytes))

		if err != nil {
			log.Println(err)
		}

		if isPublished {
			print("Sent message successfully!")
		} else {
			print("Attempted to send the message but failed")
		}
	}

}
