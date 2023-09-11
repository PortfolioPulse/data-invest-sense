package usecase

import (
     "apps/lake-orchestration/lake-controller/internal/entity"
)

type RemoveProcessingJobDependenciesUseCase struct {
     ProcessingJobDependenciesRepository entity.ProcessingJobDependenciesInterface
}

func NewRemoveProcessingJobDependenciesUseCase(
     repository entity.ProcessingJobDependenciesInterface,
) *RemoveProcessingJobDependenciesUseCase {
     return &RemoveProcessingJobDependenciesUseCase{
          ProcessingJobDependenciesRepository: repository,
     }
}

func (c *RemoveProcessingJobDependenciesUseCase) Execute(id string) error {
     err := c.ProcessingJobDependenciesRepository.DeleteProcessingJobDependencies(id)
     if err != nil {
          return err
     }
     return nil
}
