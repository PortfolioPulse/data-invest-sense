package listener

import (
	"encoding/json"
	"errors"
	"log"
	"time"

	"apps/ingestors/file-unzipper/internal/usecase"
	eventsInputDTO "libs/dtos/golang/dto-events/input"
	eventsSharedDTO "libs/dtos/golang/dto-events/shared"
	gatewayOutputDTO "libs/dtos/golang/dto-gateway/output"
	"libs/golang/go-rabbitmq/queue"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	ErrorInvalidServiceInputDTO        = errors.New("invalid service input message")
	ErrorCouldNotNotifyServiceFeedback = errors.New("could not notify service feedback")
	statusCodeOK                       = 200
	StatusCodeError                    = 500
	statusDetailOK                     = "OK"
	StatusDetailError                  = "Could not unzip file"
)

type ServiceInputListener struct {
	ContextEnv  string
	ServiceName string
}

func NewServiceInputListener(contextEnv string, serviceName string) *ServiceInputListener {
	return &ServiceInputListener{
		ContextEnv:  contextEnv,
		ServiceName: serviceName,
	}
}

func (l *ServiceInputListener) Handle(rabbitMQ *queue.RabbitMQ, exchange string, msg amqp.Delivery) error {
	var serviceInputDTO gatewayOutputDTO.InputDTO
     log.Printf("ServiceInputListener.Handle: msg.Body=%s", msg.Body)
	err := json.Unmarshal(msg.Body, &serviceInputDTO)
	if err != nil {
		return ErrorInvalidServiceInputDTO
	}
	source := serviceInputDTO.Metadata.Source
	uri := serviceInputDTO.Data["uri"].(string)
	partition := serviceInputDTO.Data["partition"].(string)
     log.Printf("ServiceInputListener.Handle: uri=%s, partition=%s, source=%s", uri, partition, source)

	unzipFileUseCase := usecase.NewUnzipFileUseCase(l.ContextEnv)

	uris, err := unzipFileUseCase.Execute(uri, partition, source)
	var serviceFeedbackDTO eventsInputDTO.ServiceFeedbackDTO
	if err != nil {
          log.Printf("ServiceInputListener.Handle: err=%s", err)
		uriResult := ""
		serviceFeedbackDTO = l.getServiceOutputDTO(uri, uriResult, partition, serviceInputDTO, StatusCodeError, StatusDetailError)
		jsonOutput, _ := json.Marshal(serviceFeedbackDTO)
		err = DispatchOutput(rabbitMQ, exchange, jsonOutput)
		if err != nil {
			return ErrorCouldNotNotifyServiceFeedback
		}

	} else {
          log.Printf("ServiceInputListener.Handle: uris=%s", uris)
		for _, uriResult := range uris {
               log.Printf("ServiceInputListener.Handle: uriResult=%s", uriResult)
			serviceFeedbackDTO = l.getServiceOutputDTO(uri, uriResult, partition, serviceInputDTO, statusCodeOK, statusDetailOK)
			jsonOutput, _ := json.Marshal(serviceFeedbackDTO)
               log.Printf("ServiceInputListener.Handle: jsonOutput=%s", jsonOutput)
			err = DispatchOutput(rabbitMQ, exchange, jsonOutput)
			if err != nil {
				return ErrorCouldNotNotifyServiceFeedback
			}
		}
	}

	return nil
}

func DispatchOutput(rabbitMQ *queue.RabbitMQ, exchange string, jsonOutput []byte) error {
	err := rabbitMQ.Notify(
		jsonOutput,
		"application/json",
		exchange,
		"feedback",
	)
	if err != nil {
		return err
	}
     return nil
}

func (l *ServiceInputListener) getServiceOutputDTO(uriOrigin string, uriResult string, partition string, serviceInputDTO gatewayOutputDTO.InputDTO, statusCode int, StatusDetail string) eventsInputDTO.ServiceFeedbackDTO {
	return eventsInputDTO.ServiceFeedbackDTO{
		Data: map[string]interface{}{
			"uri":       uriResult,
			"partition": partition,
		},
		Metadata: l.getJobMetadataDTO(uriOrigin, serviceInputDTO),
		Status: eventsSharedDTO.Status{
			Code:   statusCode,
			Detail: StatusDetail,
		},
	}
}

func (l *ServiceInputListener) getJobMetadataDTO(uriOrigin string, serviceInputDTO gatewayOutputDTO.InputDTO) eventsSharedDTO.Metadata {
	return eventsSharedDTO.Metadata{
		Input: eventsSharedDTO.MetadataInput{
			ID:                  serviceInputDTO.ID,
			Data:                serviceInputDTO.Data,
			ProcessingId:        serviceInputDTO.Metadata.ProcessingId,
			ProcessingTimestamp: serviceInputDTO.Metadata.ProcessingTimestamp,
			Source: eventsSharedDTO.MetadataInputOrigin{
				Gateway:    serviceInputDTO.Metadata.Source,
				Controller: serviceInputDTO.Metadata.Source, // inputs in file-unzipper are always from controller
			},
		},
		Service: eventsSharedDTO.MetadataInputOrigin{
			Gateway:    serviceInputDTO.Metadata.Service,
			Controller: l.ServiceName,
		},
		ProcessingId:        serviceInputDTO.Metadata.ProcessingId,
		ProcessingTimestamp: time.Now().Format(time.RFC3339),
	}
}
