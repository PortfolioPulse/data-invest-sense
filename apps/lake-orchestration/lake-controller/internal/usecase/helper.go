package usecase

import (
	"apps/lake-orchestration/lake-controller/internal/entity"
)

func ConvertEntityToUsecaseDependencies(entityDeps []entity.JobDependencies) []JobDependencies {
	usecaseDeps := make([]JobDependencies, len(entityDeps))
	for i, dep := range entityDeps {
		usecaseDeps[i] = JobDependencies{
			Service: dep.Service,
			Source:  dep.Source,
		}
	}
	return usecaseDeps
}
