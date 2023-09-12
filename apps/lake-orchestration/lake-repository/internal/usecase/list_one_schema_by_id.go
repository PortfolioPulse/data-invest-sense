package usecase

import (
     "apps/lake-orchestration/lake-repository/internal/entity"
     outputDTO "libs/dtos/golang/dto-repository/output"
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

func (la *ListOneSchemaByIdUseCase) Execute(id string) (outputDTO.SchemaDTO, error) {
     item, err := la.SchemaRepository.FindOneById(id)
     if err != nil {
          return outputDTO.SchemaDTO{}, err
     }

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

     return dto, nil
}
