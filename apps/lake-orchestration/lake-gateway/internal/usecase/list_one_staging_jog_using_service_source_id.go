package usecase

import (
     "apps/lake-orchestration/lake-gateway/internal/entity"
)

type ListOneStagingJobUsingServiceSourceIdUseCase struct {
     StagingJobRepository entity.StagingJobInterface
}

func NewListOneStagingJobUsingServiceSourceIdUseCase(
     repository entity.StagingJobInterface,
) *ListOneStagingJobUsingServiceSourceIdUseCase {
     return &ListOneStagingJobUsingServiceSourceIdUseCase{
          StagingJobRepository: repository,
     }
}

func (lo *ListOneStagingJobUsingServiceSourceIdUseCase) Execute(service string, source string, inputID string) (StagingJobOutputDTO, error) {
     item, err := lo.StagingJobRepository.FindOneStagingJobUsingServiceSourceAndId(service, source, inputID)
     if err != nil {
          return StagingJobOutputDTO{}, err
     }
     dto := StagingJobOutputDTO{
          ID:   string(item.ID),
          InputId: item.InputId,
          Input: item.Input,
          Source: item.Source,
          Service: item.Service,
          ProcessingId: item.ProcessingId,
     }
     return dto, nil
}
