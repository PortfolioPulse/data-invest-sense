package usecase

import (
	apiClient "libs/api-client/golang/go-gateway"
	gatewayOutputDTO "libs/dtos/golang/dto-gateway/output"
)

type ListOneStagingJobUsingServiceSourcAndIDUseCase struct {
	GatewayAPIClient apiClient.Client
}

func NewListOneStagingJobUsingServiceSourceAndIDUseCase() *ListOneStagingJobUsingServiceSourcAndIDUseCase {
	return &ListOneStagingJobUsingServiceSourcAndIDUseCase{
		GatewayAPIClient: *apiClient.NewClient(),
	}
}

func (lo *ListOneStagingJobUsingServiceSourcAndIDUseCase) Execute(service string, source string, inputId string) (gatewayOutputDTO.StagingJobDTO, error) {
	stagingJob, err := lo.GatewayAPIClient.ListOneStagingJobUsingServiceSourceInputId(service, source, inputId)
	if err != nil {
		return gatewayOutputDTO.StagingJobDTO{}, err
	}
	return stagingJob, nil
}
