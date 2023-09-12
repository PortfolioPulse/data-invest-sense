package usecase

import (
	"apps/lake-orchestration/lake-controller/internal/entity"
	outputDTO "libs/dtos/golang/dto-controller/output"
)

type ListAllConfigsByServiceAndContextUseCase struct {
	ConfigRepository entity.ConfigInterface
}

func NewListAllConfigsByServiceAndContextUseCase(
	repository entity.ConfigInterface,
) *ListAllConfigsByServiceAndContextUseCase {
	return &ListAllConfigsByServiceAndContextUseCase{
		ConfigRepository: repository,
	}
}

func (la *ListAllConfigsByServiceAndContextUseCase) Execute(service string, contextEnv string) ([]outputDTO.ConfigDTO, error) {
	items, err := la.ConfigRepository.FindAllByServiceAndContext(service, contextEnv)
	if err != nil {
		return nil, err
	}

	var output []outputDTO.ConfigDTO
	for _, item := range items {
		output = append(output, outputDTO.ConfigDTO{
			ID:                item.ID,
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
		})
	}

	return output, nil
}
