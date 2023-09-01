package usecase

import (
	"apps/lake-orchestration/lake-controller/internal/entity"
)

type ListAllConfigsByServiceUseCase struct {
	ConfigRepository entity.ConfigInterface
}

func NewListAllConfigsByServiceUseCase(
	repository entity.ConfigInterface,
) *ListAllConfigsByServiceUseCase {
	return &ListAllConfigsByServiceUseCase{
		ConfigRepository: repository,
	}
}

func (la *ListAllConfigsByServiceUseCase) Execute(service string) ([]ConfigOutputDTO, error) {
	items, err := la.ConfigRepository.FindAllByService(service)
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
