package usecase

import (
	apiClient "libs/api-client/golang/go-controller"
     controllerOutputDTO "libs/dtos/golang/dto-controller/output"
)

type ListOneProcessingJobDependenciesByIdUseCase struct {
     ControllerAPIClient apiClient.Client
}

func NewListOneProcessingJobDependenciesByIdUseCase() *ListOneProcessingJobDependenciesByIdUseCase {
     return &ListOneProcessingJobDependenciesByIdUseCase{
          ControllerAPIClient: *apiClient.NewClient(),
     }
}

func (la *ListOneProcessingJobDependenciesByIdUseCase) Execute(id string) (controllerOutputDTO.ProcessingJobDependenciesDTO, error) {
     jobDependencies, err := la.ControllerAPIClient.ListOneProcessingJobDependenciesById(id)
     if err != nil {
          return controllerOutputDTO.ProcessingJobDependenciesDTO{}, err
     }
     return jobDependencies, nil
}
