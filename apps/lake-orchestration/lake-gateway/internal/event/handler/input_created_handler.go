package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"libs/golang/events"
     "libs/golang/go-rabbitmq/queue"
)

type InputCreatedHandler struct {
	RabbitMQ *queue.RabbitMQ
}

func NewInputCreatedHandler(rabbitMQ *queue.RabbitMQ) *InputCreatedHandler {
	return &InputCreatedHandler{
		RabbitMQ: rabbitMQ,
	}
}

func (si *InputCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup, exchangeName string, routingKey string) {
	defer wg.Done()
	fmt.Printf("Input created: %v", event.GetPayload())
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
