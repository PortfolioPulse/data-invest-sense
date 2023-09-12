package usecase

import (
	"apps/lake-orchestration/lake-gateway/internal/entity"
)

type CreateStagingJobUseCase struct {
	StagingJobRepository entity.StagingJobInterface
}

func NewCreateStagingJobUseCase(
	repository entity.StagingJobInterface,
) *CreateStagingJobUseCase {
	return &CreateStagingJobUseCase{
		StagingJobRepository: repository,
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
		ID:           string(stagingJobEntity.ID),
		InputId:      stagingJobEntity.InputId,
		Input:        stagingJobEntity.Input,
		Source:       stagingJobEntity.Source,
		Service:      stagingJobEntity.Service,
		ProcessingId: stagingJobEntity.ProcessingId,
	}

	return dto, nil

}
