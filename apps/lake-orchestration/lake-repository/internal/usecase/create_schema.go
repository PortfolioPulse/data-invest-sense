package usecase

import (
	"apps/lake-orchestration/lake-repository/internal/entity"
	"fmt"

	"libs/golang/events"
)

type CreateSchemaUseCase struct {
	SchemaRepository entity.SchemaInterface
	SchemaCreated    events.EventInterface
	EventDispatcher  events.EventDispatcherInterface
}

func NewCreateSchemaUseCase(
	repository entity.SchemaInterface,
	SchemaCreated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *CreateSchemaUseCase {
	return &CreateSchemaUseCase{
		SchemaRepository: repository,
		SchemaCreated:    SchemaCreated,
		EventDispatcher:  EventDispatcher,
	}
}

func (csu *CreateSchemaUseCase) Execute(schema SchemaInputDTO) (SchemaOutputDTO, error) {
	schemaEntity, err := entity.NewSchema(schema.SchemaType, schema.Service, schema.Source, schema.JsonSchema)
	if err != nil {
		return SchemaOutputDTO{}, err
	}

	err = csu.SchemaRepository.SaveSchema(schemaEntity)
	if err != nil {
		return SchemaOutputDTO{}, err
	}

	dto := SchemaOutputDTO{
		ID:         string(schemaEntity.ID),
		SchemaType: schemaEntity.SchemaType,
		Service:    schemaEntity.Service,
		JsonSchema: schemaEntity.JsonSchema,
		SchemaID:   string(schemaEntity.SchemaID),
		CreatedAt:  schemaEntity.CreatedAt,
		UpdatedAt:  schemaEntity.UpdatedAt,
	}

	csu.SchemaCreated.SetPayload(dto)
	csu.EventDispatcher.Dispatch(csu.SchemaCreated, "schemas", fmt.Sprintf("%s.%s", dto.Service, dto.SchemaType))

	return dto, nil
}
