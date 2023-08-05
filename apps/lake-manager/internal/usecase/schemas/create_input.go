package schemas

import (
	entity "apps/lake-manager/internal/entity/schemas"
	"apps/lake-manager/pkg/events"
)

type SchemaInputInputDTO struct {
	ID         string                 `json:"id"`
	Required   []string               `json:"required"`
	Properties map[string]interface{} `json:"properties"`
}

type SchemaInputOutputDTO struct {
	ID         string                 `json:"id"`
	Required   []string               `json:"required"`
	Properties map[string]interface{} `json:"properties"`
	SchemaId   string                 `json:"schema_id"`
}

type CreateSchemaInputUseCase struct {
	SchemaInputRepository entity.SchemaInputInterface
	SchemaInputCreated    events.EventInterface
	EventDispatcher       events.EventDispatcherInterface
}

func NewCreateSchemaInputUseCase(
	repository entity.SchemaInputInterface,
	SchemaInputCreated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *CreateSchemaInputUseCase {
	return &CreateSchemaInputUseCase{
		SchemaInputRepository: repository,
		SchemaInputCreated:    SchemaInputCreated,
		EventDispatcher:       EventDispatcher,
	}
}

func (uc *CreateSchemaInputUseCase) Execute(input SchemaInputInputDTO) (SchemaInputOutputDTO, error) {
	schemaInput, err := entity.NewInput(input.ID, input.Required, input.Properties)
	if err != nil {
		return SchemaInputOutputDTO{}, err
	}

	err = uc.SchemaInputRepository.Save(schemaInput)
	if err != nil {
		return SchemaInputOutputDTO{}, err
	}

	dto := SchemaInputOutputDTO{
		ID:         schemaInput.ID,
		Required:   schemaInput.Required,
		Properties: schemaInput.Properties,
		SchemaId:   schemaInput.SchemaId.String(),
	}
	uc.SchemaInputCreated.SetPayload(dto)
	uc.EventDispatcher.Dispatch(uc.SchemaInputCreated)

	return dto, nil
}
