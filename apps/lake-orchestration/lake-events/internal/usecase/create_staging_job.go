package usecase

import (
	apiClient "libs/api-client/golang/go-gateway"
	gatewayInputDTO "libs/dtos/golang/dto-gateway/input"
	gatewayOutputDTO "libs/dtos/golang/dto-gateway/output"
)

type CreateStagingJobUseCase struct {
	GatewayAPIClient apiClient.Client
}

func NewCreateStagingJobUseCase() *CreateStagingJobUseCase {
	return &CreateStagingJobUseCase{
		GatewayAPIClient: *apiClient.NewClient(),
	}
}

func (csju *CreateStagingJobUseCase) Execute(stagingJob gatewayInputDTO.StagingJobDTO) (gatewayOutputDTO.StagingJobDTO, error) {
	stagingJobCreated, err := csju.GatewayAPIClient.CreateStagingJob(stagingJob)
	if err != nil {
		return gatewayOutputDTO.StagingJobDTO{}, err
	}
	return stagingJobCreated, nil
}
