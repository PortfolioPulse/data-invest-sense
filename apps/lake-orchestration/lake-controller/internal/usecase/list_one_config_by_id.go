package usecase

import (
	"apps/lake-orchestration/lake-controller/internal/entity"
)

type ListOneConfigByIdUseCase struct {
	ConfigRepository entity.ConfigInterface
}

func NewListOneConfigByIdUseCase(
	repository entity.ConfigInterface,
) *ListOneConfigByIdUseCase {
	return &ListOneConfigByIdUseCase{
		ConfigRepository: repository,
	}
}

func (la *ListOneConfigByIdUseCase) Execute(id string) (ConfigOutputDTO, error) {
	item, err := la.ConfigRepository.FindOneById(id)
	if err != nil {
		return ConfigOutputDTO{}, err
	}

	dto := ConfigOutputDTO{
		ID:                string(item.ID),
		Name:              item.Name,
		Active:            item.Active,
		Service:           item.Service,
		Source:            item.Source,
		Context:           item.Context,
		DependsOn:         ConvertEntityToUsecaseDependencies(item.DependsOn),
		ServiceParameters: item.ServiceParamaters,
		JobParameters:     item.JobParameters,
		CreatedAt:         item.CreatedAt,
		UpdatedAt:         item.UpdatedAt,
	}

	return dto, nil

}
