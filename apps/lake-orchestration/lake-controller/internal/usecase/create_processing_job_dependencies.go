package usecase

import (
	"apps/lake-orchestration/lake-controller/internal/entity"
	inputDTO "libs/dtos/golang/dto-controller/input"
	outputDTO "libs/dtos/golang/dto-controller/output"
)

type CreateProcessingJobDependenciesUseCase struct {
	ProcessingJobDependenciesRepository entity.ProcessingJobDependenciesInterface
}

func NewCreateProcessingJobDependenciesUseCase(
     repository entity.ProcessingJobDependenciesInterface,
) *CreateProcessingJobDependenciesUseCase {
     return &CreateProcessingJobDependenciesUseCase{
          ProcessingJobDependenciesRepository: repository,
     }
}

func (c *CreateProcessingJobDependenciesUseCase) Execute(input inputDTO.ProcessingJobDependenciesDTO) (outputDTO.ProcessingJobDependenciesDTO, error) {
	entityJobDependencies := make([]entity.JobDependenciesWithProcessingData, len(input.JobDependencies))
	for i, dep := range input.JobDependencies {
		entityJobDependencies[i] = entity.JobDependenciesWithProcessingData{
			Service: dep.Service,
			Source:  dep.Source,
		}
	}

	processingJobDependenciesEntity, err := entity.NewProcessingJobDependencies(
		input.Service,
		input.Source,
		entityJobDependencies,
	)
	if err != nil {
		return outputDTO.ProcessingJobDependenciesDTO{}, err
	}

	err = c.ProcessingJobDependenciesRepository.SaveProcessingJobDependencies(processingJobDependenciesEntity)
	if err != nil {
		return outputDTO.ProcessingJobDependenciesDTO{}, err
	}

	dto := outputDTO.ProcessingJobDependenciesDTO{
		ID:              string(processingJobDependenciesEntity.ID),
		Service:         processingJobDependenciesEntity.Service,
		Source:          processingJobDependenciesEntity.Source,
		JobDependencies: ConvertEntityToUsecaseJobDependenciesWithProcessingData(processingJobDependenciesEntity.JobDependencies),
	}

	return dto, nil
}
