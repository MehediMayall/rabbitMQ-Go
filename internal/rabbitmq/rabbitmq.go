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
	return rc.Close()
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

// Send Message
func (rc *RabbitClient) Send(ctx context.Context, exchangeType, routingKey string, options amqp.Publishing) error {
	return rc.channel.PublishWithContext(
		ctx,
		exchangeType,
		routingKey,
		true,
		false,
		options,
	)
}

// Consume
func (rc *RabbitClient) Consume(queue, consumer string, autoAcknowledge bool) (<-chan amqp.Delivery, error) {
	return rc.channel.Consume(queue, consumer, autoAcknowledge, false, false, false, nil)
}
