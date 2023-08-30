package usecase

import (
	"apps/lake-orchestration/lake-repository/internal/entity"
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

func (la *ListAllSchemasByServiceUseCase) Execute(service string) ([]SchemaOutputDTO, error) {
	items, err := la.SchemaRepository.FindAllByService(service)
	if err != nil {
		return []SchemaOutputDTO{}, err
	}
	var result []SchemaOutputDTO
	for _, item := range items {
		dto := SchemaOutputDTO{
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
