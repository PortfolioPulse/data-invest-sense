package usecase

import (
	apiClient "libs/api-client/golang/go-gateway"
	gatewayInputDTO "libs/dtos/golang/dto-gateway/input"
	gatewayOutputDTO "libs/dtos/golang/dto-gateway/output"
	"log"
)

type UpdateStatusInputUseCase struct {
	GatewayAPIClient apiClient.Client
}

func NewUpdateStatusInputUseCase() *UpdateStatusInputUseCase {
	return &UpdateStatusInputUseCase{
		GatewayAPIClient: *apiClient.NewClient(),
	}
}

func (uiu *UpdateStatusInputUseCase) Execute(inputStatus gatewayInputDTO.InputStatusDTO, service string, source string) (gatewayOutputDTO.InputDTO, error) {
	// id := inputStatus.ID
     log.Printf("Updating input status id: %s", inputStatus.ID)
     log.Printf("Updating input status code: %v", inputStatus.Status.Code)
     input, err := uiu.GatewayAPIClient.UpdateInputStatus(inputStatus, service, source, inputStatus.ID)
	if err != nil {
		return gatewayOutputDTO.InputDTO{}, err
	}
	return input, nil
}
