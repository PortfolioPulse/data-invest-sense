package usecase

import (
	apiClient "libs/api-client/golang/go-controller"
	controllerOutputDTO "libs/dtos/golang/dto-controller/output"
)

type ListAllConfigsByDependentJobUseCase struct {
	ControllerAPIClient apiClient.Client
}

func NewListAllConfigsByDependentJobUseCase() *ListAllConfigsByDependentJobUseCase {
	return &ListAllConfigsByDependentJobUseCase{
		ControllerAPIClient: *apiClient.NewClient(),
	}
}

func (la *ListAllConfigsByDependentJobUseCase) Execute(service string, source string) ([]controllerOutputDTO.ConfigDTO, error) {
	configs, err := la.ControllerAPIClient.ListAllConfigsByDependentJob(service, source)
	if err != nil {
		return []controllerOutputDTO.ConfigDTO{}, err
	}
	return configs, nil
}
