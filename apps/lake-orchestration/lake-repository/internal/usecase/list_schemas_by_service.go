package usecase

import (
	"apps/lake-orchestration/lake-repository/internal/entity"
	outputDTO "libs/dtos/golang/dto-repository/output"
)

type ListAllSchemasByServiceUseCase struct {
	SchemaRepository entity.SchemaInterface
}

func NewListAllSchemasByServiceUseCase(
	repository entity.SchemaInterface,
) *ListAllSchemasByServiceUseCase {
	return &ListAllSchemasByServiceUseCase{
		SchemaRepository: repository,
	}
}

func (la *ListAllSchemasByServiceUseCase) Execute(service string) ([]outputDTO.SchemaDTO, error) {
	items, err := la.SchemaRepository.FindAllByService(service)
	if err != nil {
		return []outputDTO.SchemaDTO{}, err
	}
	var result []outputDTO.SchemaDTO
	for _, item := range items {
		dto := outputDTO.SchemaDTO{
			ID:         string(item.ID),
			SchemaType: item.SchemaType,
			Service:    item.Service,
			Source:     item.Source,
			JsonSchema: item.JsonSchema,
			SchemaID:   string(item.SchemaID),
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
		}
		result = append(result, dto)
	}
	return result, nil
}
