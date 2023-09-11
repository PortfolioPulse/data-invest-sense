package listener

import (
	"encoding/json"
	"errors"
	"fmt"

	"apps/lake-orchestration/lake-events/internal/usecase"
	controllerInputDTO "libs/dtos/golang/dto-controller/input"
	controllerSharedDTO "libs/dtos/golang/dto-controller/shared"
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

	stagingJobDTO := gatewayInputDTO.StagingJobDTO{
		InputId:      serviceInputDTO.ID,
		Input:        serviceInputDTO.Data,
		Service:      service,
		Source:       source,
		ProcessingId: serviceInputDTO.Metadata.ProcessingId,
	}

	createStagingJobUseCase := usecase.NewCreateStagingJobUseCase()
	updateInputUseCase := usecase.NewUpdateStatusInputUseCase()
	findAllDependentJobUseCase := usecase.NewListAllConfigsByDependentJobUseCase()
	createProcessingJobDependenciesUseCase := usecase.NewCreateProcessingJobDependenciesUseCase()

	_, err = createStagingJobUseCase.Execute(stagingJobDTO)
	if err != nil {
		return err
	}

	_, err = updateInputUseCase.Execute(statusInputDTO, service, source)
	if err != nil {
		return err
	}

	dependentJobs, err := findAllDependentJobUseCase.Execute(service, source)
	if err != nil {
		fmt.Println(err)
	}

	for _, dependentJob := range dependentJobs {
		jobDeps := make([]controllerSharedDTO.ProcessingJobDependencies, len(dependentJob.DependsOn))
		for i, dep := range dependentJob.DependsOn {
			jobDeps[i] = controllerSharedDTO.ProcessingJobDependencies{
				Service: dep.Service,
				Source:  dep.Source,
			}
		}

		processingJobDependency := controllerInputDTO.ProcessingJobDependenciesDTO{
			Service:         dependentJob.Service,
			Source:          dependentJob.Source,
			JobDependencies: jobDeps,
		}

		_, err = createProcessingJobDependenciesUseCase.Execute(processingJobDependency)
		if err != nil {
			fmt.Println(err)
		}

	}

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
