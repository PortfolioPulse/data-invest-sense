package usecase

import (
	"apps/lake-orchestration/lake-repository/internal/entity"
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

func (la *ListAllSchemasUseCase) Execute() ([]SchemaOutputDTO, error) {
	items, err := la.SchemaRepository.FindAll()
	if err != nil {
		return []SchemaOutputDTO{}, err
	}
	var result []SchemaOutputDTO
	for _, item := range items {
		dto := SchemaOutputDTO{
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
