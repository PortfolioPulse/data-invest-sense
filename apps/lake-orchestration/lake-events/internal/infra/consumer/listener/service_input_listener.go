package listener

import (
	"encoding/json"
	"errors"
	"log"

	"apps/lake-orchestration/lake-events/internal/usecase"
	gatewayInputDTO "libs/dtos/golang/dto-gateway/input"
	gatewayOutputDTO "libs/dtos/golang/dto-gateway/output"
	gatewaySharedDTO "libs/dtos/golang/dto-gateway/shared"

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
	var serviceInputDTO gatewayOutputDTO.InputDTO
	err := json.Unmarshal(msg.Body, &serviceInputDTO)
	if err != nil {
		return ErrorInvalidServiceInputDTO
	}
	source := serviceInputDTO.Metadata.Source
	service := serviceInputDTO.Metadata.Service
	statusInputDTO := setStatusFlagToProcessing(serviceInputDTO.ID)
	log.Println(statusInputDTO)

	stagingJobDTO := gatewayInputDTO.StagingJobDTO{
		InputId:      serviceInputDTO.ID,
		Input:        serviceInputDTO.Data,
		Service:      service,
		Source:       source,
		ProcessingId: serviceInputDTO.Metadata.ProcessingId,
	}

	createStagingJobUseCase := usecase.NewCreateStagingJobUseCase()
	updateInputUseCase := usecase.NewUpdateStatusInputUseCase()

	stagingJob, err := createStagingJobUseCase.Execute(stagingJobDTO)
	if err != nil {
		return err
	}

	log.Println(stagingJob)

	statusInput, err := updateInputUseCase.Execute(statusInputDTO, service, source)
	if err != nil {
		return err
	}

	log.Println(statusInput)

	return nil
}

func setStatusFlagToProcessing(id string) gatewayInputDTO.InputStatusDTO {
	return gatewayInputDTO.InputStatusDTO{
		ID: id,
		Status: gatewaySharedDTO.Status{
			Code:   1,
			Detail: "Processing",
		},
	}
}
