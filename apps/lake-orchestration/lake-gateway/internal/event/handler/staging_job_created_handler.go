package handler

import (
	"encoding/json"
	"fmt"
	"sync"

	"libs/golang/events"
	"libs/golang/go-rabbitmq/queue"
)

type StagingJobCreatedHandler struct {
	RabbitMQ *queue.RabbitMQ
}

func NewStagingJobCreatedHandler(rabbitMQ *queue.RabbitMQ) *StagingJobCreatedHandler {
	return &StagingJobCreatedHandler{
		RabbitMQ: rabbitMQ,
	}
}

func (si *StagingJobCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup, exchangeName string, routingKey string) {
	defer wg.Done()
	fmt.Printf("StagingJob created: %v", event.GetPayload())
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
