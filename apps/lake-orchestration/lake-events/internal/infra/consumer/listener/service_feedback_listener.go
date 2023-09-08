package listener

import (
	"encoding/json"
	"errors"
	"fmt"

	"apps/lake-orchestration/lake-events/internal/usecase"
	eventsInputDTO "libs/dtos/golang/dto-events/input"
	gatewayInputDTO "libs/dtos/golang/dto-gateway/input"
	gatewaySharedDTO "libs/dtos/golang/dto-gateway/shared"

	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	ErrorInvalidServiceFeedbackDTO = errors.New("invalid service feedback message")
	ErrorInvalidStatus             = errors.New("invalid status code")
)

type ServiceFeedbackListener struct {
}

func NewServiceFeedbackListener() *ServiceFeedbackListener {
	return &ServiceFeedbackListener{}
}

func (l *ServiceFeedbackListener) extractStatusCodeRange(statusCode int) string {
     if statusCode >= 200 && statusCode < 300 {
          return "2XX"
     } else if statusCode >= 400 && statusCode < 500 {
          return "4XX"
     } else if statusCode >= 500 && statusCode < 600 {
          return "5XX"
     }
     return "invalid"
}

func (l *ServiceFeedbackListener) Handle(msg amqp.Delivery) error {
	fmt.Println(string(msg.Body))
	var serviceFeedbackDTO eventsInputDTO.ServiceFeedbackDTO
	err := json.Unmarshal(msg.Body, &serviceFeedbackDTO)
	if err != nil {
		return ErrorInvalidServiceFeedbackDTO
	}
     statusCodeRange := l.extractStatusCodeRange(serviceFeedbackDTO.Status.Code)
	switch statusCodeRange {
     case "2XX":
          l.HandleFeedback200(serviceFeedbackDTO)
     case "4XX":
          l.HandleFeedback400(serviceFeedbackDTO)
     case "5XX":
          l.HandleFeedback500(serviceFeedbackDTO)
     default:
          return ErrorInvalidStatus
     }
     return nil
}

func (l *ServiceFeedbackListener) HandleFeedback200(msg eventsInputDTO.ServiceFeedbackDTO) {
	statusDTO := getStatusInputDTO(msg)
	service := msg.Metadata.Service.Gateway
	source := msg.Metadata.Input.Source.Gateway

	updateInputUseCase := usecase.NewUpdateStatusInputUseCase()
	findStagingJobUseCase := usecase.NewListOneStagingJobUsingServiceSourceAndIDUseCase()
	removeStagingJobUseCase := usecase.NewRemoveStagingJobUseCase()

	input, err := updateInputUseCase.Execute(statusDTO, service, source)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(input)

	stagingJob, err := findStagingJobUseCase.Execute(service, source, input.ID)
	if err != nil {
		fmt.Println(err)
	}

	stagingJobIdDTO := gatewayInputDTO.StagingJobIDDTO{
		ID: stagingJob.ID,
	}

	removeStagingJobUseCase.Execute(stagingJobIdDTO)
}

func (l *ServiceFeedbackListener) HandleFeedback400(msg eventsInputDTO.ServiceFeedbackDTO) {
	statusDTO := getStatusInputDTO(msg)
	service := msg.Metadata.Service.Gateway
	source := msg.Metadata.Input.Source.Gateway

	updateInputUseCase := usecase.NewUpdateStatusInputUseCase()
	findStagingJobUseCase := usecase.NewListOneStagingJobUsingServiceSourceAndIDUseCase()
	removeStagingJobUseCase := usecase.NewRemoveStagingJobUseCase()

	input, err := updateInputUseCase.Execute(statusDTO, service, source)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(input)

	stagingJob, err := findStagingJobUseCase.Execute(service, source, input.ID)
	if err != nil {
		fmt.Println(err)
	}

	stagingJobIdDTO := gatewayInputDTO.StagingJobIDDTO{
		ID: stagingJob.ID,
	}

	removeStagingJobUseCase.Execute(stagingJobIdDTO)
}

func (l *ServiceFeedbackListener) HandleFeedback500(msg eventsInputDTO.ServiceFeedbackDTO) {
	statusDTO := getStatusInputDTO(msg)
	service := msg.Metadata.Service.Gateway
	source := msg.Metadata.Input.Source.Gateway

	updateInputUseCase := usecase.NewUpdateStatusInputUseCase()
	findStagingJobUseCase := usecase.NewListOneStagingJobUsingServiceSourceAndIDUseCase()
	removeStagingJobUseCase := usecase.NewRemoveStagingJobUseCase()

	input, err := updateInputUseCase.Execute(statusDTO, service, source)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(input)

	stagingJob, err := findStagingJobUseCase.Execute(service, source, input.ID)
	if err != nil {
		fmt.Println(err)
	}

	stagingJobIdDTO := gatewayInputDTO.StagingJobIDDTO{
		ID: stagingJob.ID,
	}

	removeStagingJobUseCase.Execute(stagingJobIdDTO)
}

func getStatusInputDTO(msg eventsInputDTO.ServiceFeedbackDTO) gatewayInputDTO.InputStatusDTO {
	return gatewayInputDTO.InputStatusDTO{
		ID: msg.Metadata.Input.ID,
		Status: gatewaySharedDTO.Status{
			Code:   msg.Status.Code,
			Detail: msg.Status.Detail,
		},
	}
}
