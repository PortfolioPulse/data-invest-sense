package usecase

import (
	"apps/lake-orchestration/lake-controller/internal/entity"
	sharedDTO "libs/dtos/golang/dto-controller/shared"
)

func ConvertEntityToUsecaseDependencies(entityDeps []entity.JobDependencies) []sharedDTO.JobDependencies {
	usecaseDeps := make([]sharedDTO.JobDependencies, len(entityDeps))
	for i, dep := range entityDeps {
		usecaseDeps[i] = sharedDTO.JobDependencies{
			Service: dep.Service,
			Source:  dep.Source,
		}
	}
	return usecaseDeps
}

func ConvertEntityToUsecaseJobDependenciesWithProcessingData(entityDeps []entity.JobDependenciesWithProcessingData) []sharedDTO.ProcessingJobDependencies {
     usecaseDeps := make([]sharedDTO.ProcessingJobDependencies, len(entityDeps))
     for i, dep := range entityDeps {
          usecaseDeps[i] = sharedDTO.ProcessingJobDependencies{
               Service: dep.Service,
               Source:  dep.Source,
               ProcessingId: dep.ProcessingId,
               ProcessingTimestamp: dep.ProcessingTimestamp,
               StatusCode: dep.StatusCode,
          }
     }
     return usecaseDeps
}
