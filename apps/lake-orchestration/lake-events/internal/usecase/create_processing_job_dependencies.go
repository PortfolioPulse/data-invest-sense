package usecase

import (
	apiClient "libs/api-client/golang/go-controller"
	controllerInputDTO "libs/dtos/golang/dto-controller/input"
     controllerOutputDTO "libs/dtos/golang/dto-controller/output"
)

type CreateProcessingJobDependenciesUseCase struct {
     ControllerAPIClient apiClient.Client
}

func NewCreateProcessingJobDependenciesUseCase() *CreateProcessingJobDependenciesUseCase {
     return &CreateProcessingJobDependenciesUseCase{
          ControllerAPIClient: *apiClient.NewClient(),
     }
}

func (la *CreateProcessingJobDependenciesUseCase) Execute(jobDependencies controllerInputDTO.ProcessingJobDependenciesDTO) (controllerOutputDTO.ProcessingJobDependenciesDTO, error) {
     jobDependenciesCreated, err := la.ControllerAPIClient.CreateProcessingJobDependencies(jobDependencies)
     if err != nil {
          return controllerOutputDTO.ProcessingJobDependenciesDTO{}, err
     }
     return jobDependenciesCreated, nil
}
