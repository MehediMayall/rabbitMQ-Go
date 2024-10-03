package rabbitmq

import (
	"context"
	"fmt"

	amqp "github.com/rabbitmq/amqp091-go"
)

// To hold connection and channel object
type RabbitClient struct {
	connection *amqp.Connection
	channel    *amqp.Channel
}

// Create Connection
func CreateConnection(hostname, username, password, virtualhost string) (*amqp.Connection, error) {
	return amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s/%s", username, password, hostname, virtualhost))
}

// Create Client
func CreateClient(connection *amqp.Connection) (*RabbitClient, error) {
	channel, err := connection.Channel()
	if err != nil {
		return &RabbitClient{}, err
	}

	return &RabbitClient{
		connection: connection,
		channel:    channel,
	}, nil
}

// Close Channel
func (rc *RabbitClient) Close() error {
	return rc.channel.Close()
}

// Create Queue
func (rc *RabbitClient) CreateQueue(queuename string, durable, autoDelete bool) error {
	_, err := rc.channel.QueueDeclare(queuename, durable, autoDelete, false, false, nil)
	return err
}

// Create Binding
func (rc *RabbitClient) BindQueue(name, binding, exchange string) error {
	return rc.channel.QueueBind(name, binding, exchange, false, nil)
}

// Create Options
func (rc *RabbitClient) CreateOptionsPersistent(body []byte) *amqp.Publishing {
	return &amqp.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: amqp.Persistent,
		Body:         body,
	}
}
func (rc *RabbitClient) CreateOptionsTransient(body []byte) *amqp.Publishing {
	return &amqp.Publishing{
		ContentType:  "text/plain",
		DeliveryMode: amqp.Transient,
		Body:         body,
	}
}

// Send Message
func (rc *RabbitClient) Send(ctx context.Context, exchangeName, routingKey string, options amqp.Publishing) error {
	return rc.channel.PublishWithContext(
		ctx,
		exchangeName,
		routingKey,
		true,
		false,
		options,
	)
}

// Consume
func (rc *RabbitClient) Consume(queue, consumerName string, autoAcknowledge bool) (<-chan amqp.Delivery, error) {
	return rc.channel.Consume(queue, consumerName, autoAcknowledge, false, false, false, nil)
}
