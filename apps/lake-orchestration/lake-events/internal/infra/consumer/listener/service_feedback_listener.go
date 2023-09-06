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

func (l *ServiceFeedbackListener) Handle(msg amqp.Delivery) error {
	fmt.Println(string(msg.Body))
	var serviceFeedbackDTO eventsInputDTO.ServiceFeedbackDTO
	err := json.Unmarshal(msg.Body, &serviceFeedbackDTO)
	if err != nil {
		return ErrorInvalidServiceFeedbackDTO
	}
	switch serviceFeedbackDTO.Status.Code {
	case 200:
		l.HandleFeedback200(serviceFeedbackDTO)
	case 400:
		l.HandleFeedback400(serviceFeedbackDTO)
	case 500:
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
	fmt.Println(statusDTO)
	fmt.Println(msg)
}

func (l *ServiceFeedbackListener) HandleFeedback500(msg eventsInputDTO.ServiceFeedbackDTO) {
	statusDTO := getStatusInputDTO(msg)
	fmt.Println(statusDTO)
	fmt.Println(msg)
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
