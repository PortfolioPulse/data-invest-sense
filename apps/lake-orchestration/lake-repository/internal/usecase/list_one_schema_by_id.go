package usecase

import (
     "apps/lake-orchestration/lake-repository/internal/entity"
)

type ListOneSchemaByIdUseCase struct {
     SchemaRepository entity.SchemaInterface
}

func NewListOneSchemaByIdUseCase(
     repository entity.SchemaInterface,
) *ListOneSchemaByIdUseCase {
     return &ListOneSchemaByIdUseCase{
          SchemaRepository: repository,
     }
}

func (la *ListOneSchemaByIdUseCase) Execute(id string) (SchemaOutputDTO, error) {
     item, err := la.SchemaRepository.FindOneById(id)
     if err != nil {
          return SchemaOutputDTO{}, err
     }

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

     return dto, nil
}
