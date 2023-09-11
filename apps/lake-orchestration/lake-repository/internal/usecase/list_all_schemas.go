package usecase

import (
	"apps/lake-orchestration/lake-repository/internal/entity"
     outputDTO "libs/dtos/golang/dto-repository/output"
)

type ListAllSchemasUseCase struct {
	SchemaRepository entity.SchemaInterface
}

func NewListAllSchemasUseCase(
	repository entity.SchemaInterface,
) *ListAllSchemasUseCase {
	return &ListAllSchemasUseCase{
		SchemaRepository: repository,
	}
}

func (la *ListAllSchemasUseCase) Execute() ([]outputDTO.SchemaDTO, error) {
	items, err := la.SchemaRepository.FindAll()
	if err != nil {
		return []outputDTO.SchemaDTO{}, err
	}
	var result []outputDTO.SchemaDTO
	for _, item := range items {
		dto := outputDTO.SchemaDTO{
			ID:         string(item.ID),
			SchemaType: item.SchemaType,
			JsonSchema: item.JsonSchema,
			Service:    item.Service,
			Source:     item.Source,
			SchemaID:   string(item.SchemaID),
			CreatedAt:  item.CreatedAt,
			UpdatedAt:  item.UpdatedAt,
		}
		result = append(result, dto)
	}
	return result, nil
}
