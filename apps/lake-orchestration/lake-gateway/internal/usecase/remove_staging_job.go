package usecase

import (
	"apps/lake-orchestration/lake-gateway/internal/entity"
)

type RemoveStagingJobUseCase struct {
	StagingJobRepository entity.StagingJobInterface
}

func NewRemoveStagingJobUseCase(
	repository entity.StagingJobInterface,
) *RemoveStagingJobUseCase {
	return &RemoveStagingJobUseCase{
		StagingJobRepository: repository,
	}
}

func (rsju *RemoveStagingJobUseCase) Execute(stagingJobId StagingJobInputIDDTO) (StagingJobOutputDTO, error) {
	err := rsju.StagingJobRepository.DeleteById(stagingJobId.ID)
	if err != nil {
		return StagingJobOutputDTO{}, err
	}

	dto := StagingJobOutputDTO{
		ID: string(stagingJobId.ID),
	}

	return dto, nil
}
