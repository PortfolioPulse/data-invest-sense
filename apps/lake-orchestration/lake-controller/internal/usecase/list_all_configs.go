package usecase

import (
	"apps/lake-orchestration/lake-controller/internal/entity"
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

func (la *ListAllConfigsUseCase) Execute() ([]ConfigOutputDTO, error) {
	items, err := la.ConfigRepository.FindAll()
	if err != nil {
		return []ConfigOutputDTO{}, err
	}
	var result []ConfigOutputDTO
	for _, item := range items {
		dto := ConfigOutputDTO{
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
