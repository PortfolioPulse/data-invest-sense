package listener

import (
	"encoding/json"
	"errors"
	"fmt"

	"apps/lake-orchestration/lake-events/internal/usecase"
	controllerSheredDTO "libs/dtos/golang/dto-controller/shared"
	eventsInputDTO "libs/dtos/golang/dto-events/input"
	gatewayInputDTO "libs/dtos/golang/dto-gateway/input"
	gatewaySharedDTO "libs/dtos/golang/dto-gateway/shared"
	configID "libs/golang/goid/config"

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

func extractStatusCodeRange(statusCode int) string {
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
	// fmt.Println(string(msg.Body))
	var serviceFeedbackDTO eventsInputDTO.ServiceFeedbackDTO
	err := json.Unmarshal(msg.Body, &serviceFeedbackDTO)
	if err != nil {
		return ErrorInvalidServiceFeedbackDTO
	}
	statusCodeRange := extractStatusCodeRange(serviceFeedbackDTO.Status.Code)

	statusDTO := getStatusInputDTO(serviceFeedbackDTO)
	service := serviceFeedbackDTO.Metadata.Service.Gateway
	source := serviceFeedbackDTO.Metadata.Input.Source.Gateway

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

	switch statusCodeRange {
	case "2XX":
		l.HandleFeedback200(serviceFeedbackDTO, service, source)
	case "4XX":
		l.HandleFeedback400(serviceFeedbackDTO, service, source)
	case "5XX":
		l.HandleFeedback500(serviceFeedbackDTO, service, source)
	default:
		return ErrorInvalidStatus
	}
	return nil
}

func (l *ServiceFeedbackListener) HandleFeedback200(msg eventsInputDTO.ServiceFeedbackDTO, service string, source string) {
	// log.Printf("ServiceFeedbackListener.HandleFeedback200: msg=%v, service=%s, source=%s", msg, service, source)

	findAllDependentJobUseCase := usecase.NewListAllConfigsByDependentJobUseCase()
	createInputUseCase := usecase.NewCreateInputUseCase()

	updateProcessingJobDependenciesUseCase := usecase.NewUpdateProcessingJobDependenciesUseCase()
	checkAllJobDependenciesStatus200UseCase := usecase.NewCheckAllJobDependenciesStatus200UseCase()
	removeProcessingJobDependenciesUseCase := usecase.NewRemoveProcessingJobDependenciesUseCase()

	dependentJobs, err := findAllDependentJobUseCase.Execute(service, source)
	if err != nil {
		fmt.Println(err)
	}

	inputDTO := gatewayInputDTO.InputDTO{
		Data: map[string]interface{}{
			"uri": msg.Data["documentUri"],
               "partition": msg.Data["partition"],
		},
	}

	jobDep := getProcessingJobDependencies(msg)

	for _, dependentJob := range dependentJobs {
		processingJobDepId := configID.NewID(dependentJob.Service, dependentJob.Source)

		updateProcessingJobDependenciesUseCase.Execute(processingJobDepId, jobDep)
		shouldRun, err := checkAllJobDependenciesStatus200UseCase.Execute(processingJobDepId)
		if err != nil {
			fmt.Println(err)
		}
		if shouldRun {
			_, err := createInputUseCase.Execute(dependentJob.Service, dependentJob.Source, inputDTO)
			if err != nil {
				fmt.Println(err)
			}
			removeProcessingJobDependenciesUseCase.Execute(processingJobDepId)
		}

	}

}

func (l *ServiceFeedbackListener) HandleFeedback400(msg eventsInputDTO.ServiceFeedbackDTO, service string, source string) {

}

func (l *ServiceFeedbackListener) HandleFeedback500(msg eventsInputDTO.ServiceFeedbackDTO, service string, source string) {

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

func getProcessingJobDependencies(msg eventsInputDTO.ServiceFeedbackDTO) controllerSheredDTO.ProcessingJobDependencies {
	return controllerSheredDTO.ProcessingJobDependencies{
		Service:             msg.Metadata.Service.Gateway,
		Source:              msg.Metadata.Input.Source.Gateway,
		ProcessingId:        msg.Metadata.ProcessingId,
		ProcessingTimestamp: msg.Metadata.ProcessingTimestamp,
		StatusCode:          msg.Status.Code,
	}
}
