package usecase

import (
	"apps/lake-orchestration/lake-controller/internal/entity"
	outputDTO "libs/dtos/golang/dto-controller/output"
)

type ListAllConfigsUseCase struct {
	ConfigRepository entity.ConfigInterface
}

func NewListAllConfigsUseCase(
	repository entity.ConfigInterface,
) *ListAllConfigsUseCase {
	return &ListAllConfigsUseCase{
		ConfigRepository: repository,
	}
}

func (la *ListAllConfigsUseCase) Execute() ([]outputDTO.ConfigDTO, error) {
	items, err := la.ConfigRepository.FindAll()
	if err != nil {
		return []outputDTO.ConfigDTO{}, err
	}
	var result []outputDTO.ConfigDTO
	for _, item := range items {
		dto := outputDTO.ConfigDTO{
			ID:                string(item.ID),
			Name:              item.Name,
			Active:            item.Active,
			Frequency:         item.Frequency,
			Service:           item.Service,
			Source:            item.Source,
			Context:           item.Context,
			DependsOn:         ConvertEntityToUsecaseDependencies(item.DependsOn),
			ServiceParameters: item.ServiceParameters,
			JobParameters:     item.JobParameters,
			CreatedAt:         item.CreatedAt,
			UpdatedAt:         item.UpdatedAt,
		}
		result = append(result, dto)
	}
	return result, nil
}
