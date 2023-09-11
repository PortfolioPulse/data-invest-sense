package listener

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

type MessageHandlerInterface interface {
	Handle(msg amqp.Delivery) error
}


