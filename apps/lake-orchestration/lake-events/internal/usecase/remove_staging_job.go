package usecase

import (
	apiClient "libs/api-client/golang/go-gateway"
	gatewayInputDTO "libs/dtos/golang/dto-gateway/input"
	gatewayOutputDTO "libs/dtos/golang/dto-gateway/output"
)

type RemoveStagingJobUseCase struct {
	GatewayAPIClient apiClient.Client
}

func NewRemoveStagingJobUseCase() *RemoveStagingJobUseCase {
	return &RemoveStagingJobUseCase{
		GatewayAPIClient: *apiClient.NewClient(),
	}
}

func (rsju *RemoveStagingJobUseCase) Execute(stagingJobId gatewayInputDTO.StagingJobIDDTO) (gatewayOutputDTO.StagingJobDTO, error) {
	_, err := rsju.GatewayAPIClient.RemoveStagingJob(stagingJobId.ID)
	if err != nil {
		return gatewayOutputDTO.StagingJobDTO{}, err
	}
	return gatewayOutputDTO.StagingJobDTO{}, nil
}
