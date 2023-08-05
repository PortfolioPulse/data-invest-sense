package schemas

import (
	"encoding/json"
	"fmt"
	"sync"

	"apps/lake-manager/pkg/events"

	"github.com/streadway/amqp"
)

type SchemaInputCreatedHandler struct {
	RabbitMQChannel *amqp.Channel
}

func NewOrderCreatedHandler(rabbitMQChannel *amqp.Channel) *SchemaInputCreatedHandler {
	return &SchemaInputCreatedHandler{
		RabbitMQChannel: rabbitMQChannel,
	}
}

func (si *SchemaInputCreatedHandler) Handle(event events.EventInterface, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Printf("Schema Input created: %v", event.GetPayload())
	jsonOutput, _ := json.Marshal(event.GetPayload())

	msgRabbitmq := amqp.Publishing{
		ContentType: "application/json",
		Body:        jsonOutput,
	}

	si.RabbitMQChannel.Publish(
		"amq.direct", // exchange
		"",           // key name
		false,        // mandatory
		false,        // immediate
		msgRabbitmq,  // message to publish
	)
}
