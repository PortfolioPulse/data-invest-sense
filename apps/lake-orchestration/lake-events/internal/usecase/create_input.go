package usecase

import (
	apiClient "libs/api-client/golang/go-gateway"
	gatewayInputDTO "libs/dtos/golang/dto-gateway/input"
	gatewayOutputDTO "libs/dtos/golang/dto-gateway/output"
)

type CreateInputUseCase struct {
     GatewayAPIClient apiClient.Client
}

func NewCreateInputUseCase() *CreateInputUseCase {
     return &CreateInputUseCase{
          GatewayAPIClient: *apiClient.NewClient(),
     }
}

func (ciu *CreateInputUseCase) Execute(service string, source string, input gatewayInputDTO.InputDTO) (gatewayOutputDTO.InputDTO, error) {
     inputCreated, err := ciu.GatewayAPIClient.CreateInput(service, source, input)
     if err != nil {
          return gatewayOutputDTO.InputDTO{}, err
     }
     return inputCreated, nil
}
