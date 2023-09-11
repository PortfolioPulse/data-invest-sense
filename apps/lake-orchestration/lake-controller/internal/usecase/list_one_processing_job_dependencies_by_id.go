package usecase

import (
     "apps/lake-orchestration/lake-controller/internal/entity"
     outputDTO "libs/dtos/golang/dto-controller/output"
)

type ListOneProcessingJobDependenciesByIdUseCase struct {
     ProcessingJobDependenciesRepository entity.ProcessingJobDependenciesInterface
}

func NewListOneProcessingJobDependenciesByIdUseCase(
     repository entity.ProcessingJobDependenciesInterface,
) *ListOneProcessingJobDependenciesByIdUseCase {
     return &ListOneProcessingJobDependenciesByIdUseCase{
          ProcessingJobDependenciesRepository: repository,
     }
}

func (la *ListOneProcessingJobDependenciesByIdUseCase) Execute(id string) (outputDTO.ProcessingJobDependenciesDTO, error) {
     item, err := la.ProcessingJobDependenciesRepository.FindOneById(id)
     if err != nil {
          return outputDTO.ProcessingJobDependenciesDTO{}, err
     }
     dto := outputDTO.ProcessingJobDependenciesDTO{
          ID:              item.ID,
          Service:         item.Service,
          Source:          item.Source,
          JobDependencies: ConvertEntityToUsecaseJobDependenciesWithProcessingData(item.JobDependencies),
     }
     return dto, nil
}

