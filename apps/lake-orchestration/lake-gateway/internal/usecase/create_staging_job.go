package usecase

import (
     "apps/lake-orchestration/lake-gateway/internal/entity"
     "fmt"
     "libs/golang/events"
)

type CreateStagingJobUseCase struct {
     StagingJobRepository entity.StagingJobInterface
     StagingJobCreated    events.EventInterface
     EventDispatcher      events.EventDispatcherInterface
}

func NewCreateStagingJobUseCase(
     repository entity.StagingJobInterface,
     StagingJobCreated events.EventInterface,
     EventDispatcher events.EventDispatcherInterface,
) *CreateStagingJobUseCase {
     return &CreateStagingJobUseCase{
          StagingJobRepository: repository,
          StagingJobCreated:    StagingJobCreated,
          EventDispatcher:      EventDispatcher,
     }
}

func (csju *CreateStagingJobUseCase) Execute(stagingJob StagingJobInputDTO) (StagingJobOutputDTO, error) {
     stagingJobEntity, err := entity.NewStagingJob(
          stagingJob.InputId,
          stagingJob.Input,
          stagingJob.Source,
          stagingJob.Service,
          stagingJob.ProcessingId,
     )

     if err != nil {
          return StagingJobOutputDTO{}, err
     }

     err = csju.StagingJobRepository.SaveStagingJob(stagingJobEntity)
     if err != nil {
          return StagingJobOutputDTO{}, err
     }

     dto := StagingJobOutputDTO{
          ID:   string(stagingJobEntity.ID),
          InputId: stagingJobEntity.InputId,
          Input: stagingJobEntity.Input,
          Source: stagingJobEntity.Source,
          Service: stagingJobEntity.Service,
          ProcessingId: stagingJobEntity.ProcessingId,
     }
     csju.StagingJobCreated.SetPayload(dto)
     csju.EventDispatcher.Dispatch(csju.StagingJobCreated, "services", fmt.Sprintf("%s.staging_jobs.%s", dto.Service, dto.Source))

     return dto, nil

}
