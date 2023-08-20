package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"libs/golang/events"
	"libs/golang/go-rabbitmq/queue"
)

type SchemaCreatedHandler struct {
	RabbitMQ *queue.RabbitMQ
}

func NewSchemaCreatedHandler(rabbitMQ *queue.RabbitMQ) *SchemaCreatedHandler {
	return &SchemaCreatedHandler{
		RabbitMQ: rabbitMQ,
	}
}

func (si *SchemaCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup, exchangeName string, routingKey string) {
	defer wg.Done()
	fmt.Printf("Schema created: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())
	err := si.RabbitMQ.Notify(
		jsonOutput,
		"application/json",
		exchangeName,
		routingKey,
	)
	if err != nil {
		fmt.Println(err)
	}
}
