package usecase

import (
	"apps/lake-orchestration/lake-controller/internal/entity"
	outputDTO "libs/dtos/golang/dto-controller/output"
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

func (la *ListOneConfigByIdUseCase) Execute(id string) (outputDTO.ConfigDTO, error) {
	item, err := la.ConfigRepository.FindOneById(id)
	if err != nil {
		return outputDTO.ConfigDTO{}, err
	}

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

	return dto, nil

}
