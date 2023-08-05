package schemas

import (
	"apps/lake-manager/pkg/uid"
	"errors"
)

type SchemaInput struct {
	ID         string                 `json:"id"`
	Required   []string               `json:"required"`
	Properties map[string]interface{} `json:"properties"`
	SchemaId   uid.ID
}

func NewInput(id string, required []string, properties map[string]interface{}) (*SchemaInput, error) {
	schemaInput := &SchemaInput{
		ID:         id,
		Required:   required,
		Properties: properties,
		SchemaId:   uid.NewID(),
	}
	err := schemaInput.IsValid()
	if err != nil {
		return nil, err
	}
	return schemaInput, nil
}

func (i *SchemaInput) IsValid() error {
	if i.ID == "" {
		return errors.New("ID is required")
	}
	if len(i.Required) == 0 {
		return errors.New("required field is required")
	}
	if i.Properties == nil {
		return errors.New("properties field is required")
	}
	return nil
}
