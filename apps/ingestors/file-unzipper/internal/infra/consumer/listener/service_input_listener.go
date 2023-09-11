package listener

import (
	"encoding/json"
	"errors"
	"log"

	"apps/ingestors/file-unzipper/internal/usecase"
	gatewayInputDTO "libs/dtos/golang/dto-gateway/input"
	gatewayOutputDTO "libs/dtos/golang/dto-gateway/output"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	ErrorInvalidServiceInputDTO = errors.New("invalid service input message")
)

type ServiceInputListener struct {
}

func NewServiceInputListener() *ServiceInputListener {
	return &ServiceInputListener{}
}

func (l *ServiceInputListener) Handle(msg amqp.Delivery) error {
	
}
