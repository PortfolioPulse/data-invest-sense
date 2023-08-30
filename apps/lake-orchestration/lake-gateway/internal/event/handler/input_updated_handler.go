package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"libs/golang/events"

     "libs/golang/go-rabbitmq/queue"
)

type InputUpdatedHandler struct {
	RabbitMQ *queue.RabbitMQ
}

func NewInputUpdatedHandler(rabbitMQ *queue.RabbitMQ) *InputUpdatedHandler {
	return &InputUpdatedHandler{
		RabbitMQ: rabbitMQ,
	}
}

func (iu *InputUpdatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup, exchangeName string, routingKey string) {
	defer wg.Done()
	fmt.Printf("Input created: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())
     err := iu.RabbitMQ.Notify(
          jsonOutput,
          "application/json",
          exchangeName,
          routingKey,
     )
     if err != nil {
          fmt.Println(err)
     }
}
