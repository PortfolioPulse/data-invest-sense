package usecase

import (
	apiClient "libs/api-client/golang/go-controller"
     controllerOutputDTO "libs/dtos/golang/dto-controller/output"
)

type RemoveProcessingJobDependenciesUseCase struct {
     ControllerAPIClient apiClient.Client
}

func NewRemoveProcessingJobDependenciesUseCase() *RemoveProcessingJobDependenciesUseCase {
     return &RemoveProcessingJobDependenciesUseCase{
          ControllerAPIClient: *apiClient.NewClient(),
     }
}

func (la *RemoveProcessingJobDependenciesUseCase) Execute(id string) (controllerOutputDTO.ProcessingJobDependenciesDTO, error) {
     _, err := la.ControllerAPIClient.RemoveProcessingJobDependencies(id)
     if err != nil {
          return controllerOutputDTO.ProcessingJobDependenciesDTO{}, err
     }
     return controllerOutputDTO.ProcessingJobDependenciesDTO{}, nil
}
