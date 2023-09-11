package usecase

import (
	apiClient "libs/api-client/golang/go-controller"
	controllerOutputDTO "libs/dtos/golang/dto-controller/output"
	controllerSharedDTO "libs/dtos/golang/dto-controller/shared"
	"log"
)

type UpdateProcessingJobDependenciesUseCase struct {
     ControllerAPIClient apiClient.Client
}

func NewUpdateProcessingJobDependenciesUseCase() *UpdateProcessingJobDependenciesUseCase {
     return &UpdateProcessingJobDependenciesUseCase{
          ControllerAPIClient: *apiClient.NewClient(),
     }
}

func (la *UpdateProcessingJobDependenciesUseCase) Execute(id string, jobDependencies controllerSharedDTO.ProcessingJobDependencies) (controllerOutputDTO.ProcessingJobDependenciesDTO, error) {
     log.Printf("UpdateProcessingJobDependenciesUseCase.Execute: id=%s, jobDependencies=%v", id, jobDependencies)
     jobDependenciesUpdated, err := la.ControllerAPIClient.UpdateProcessingJobDependencies(id, jobDependencies)
     log.Printf("UpdateProcessingJobDependenciesUseCase.Execute: jobDependenciesUpdated=%v, err=%v", jobDependenciesUpdated, err)
     if err != nil {
          return controllerOutputDTO.ProcessingJobDependenciesDTO{}, err
     }
     return jobDependenciesUpdated, nil
}
