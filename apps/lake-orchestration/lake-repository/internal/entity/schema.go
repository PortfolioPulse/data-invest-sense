package entity

import (
	"encoding/json"
	"errors"
	"strings"
	"time"

	schemaMD5ID "libs/golang/goid/schema/md5"
	uuid "libs/golang/goid/schema/uuid"

	"github.com/xeipuuv/gojsonschema"
)

const metaschemaURL = "http://json-schema.org/draft-07/schema#"

type Schema struct {
	ID         schemaMD5ID.ID         `json:"id"`
	SchemaType string                 `bson:"schema_type"`
	JsonSchema map[string]interface{} `bson:"json_schema"`
	SchemaID   uuid.ID                `bson:"schema_id"`
	Service    string                 `bson:"service"`
	Source     string                 `bson:"source"`
	CreatedAt  string
	UpdatedAt  string
}

func NewSchema(schemaType string, service string, source string, jsonSchema map[string]interface{}) (*Schema, error) {
	schemaId, err := uuid.GenerateSchemaID(schemaType, jsonSchema)
	if err != nil {
		return nil, err
	}

	schema := &Schema{
		ID:         schemaMD5ID.NewID(schemaType, service, source),
		SchemaType: schemaType,
		JsonSchema: jsonSchema,
		Service:    service,
		Source:     source,
		SchemaID:   schemaId,
		CreatedAt:  time.Now().Format(time.RFC3339),
		UpdatedAt:  time.Now().Format(time.RFC3339),
	}
	err = schema.IsSchemaValid()
	if err != nil {
		return nil, err
	}
	return schema, nil
}

func ValidateJSONSchema(jsonSchema map[string]interface{}) error {
	// Convert the JSON schema map to a JSON string
	jsonSchemaBytes, err := json.Marshal(jsonSchema)
	if err != nil {
		return err
	}

	// Create a JSONLoader for the JSON schema
	schemaLoader := gojsonschema.NewStringLoader(string(jsonSchemaBytes))

	// Validate the JSON Schema structure using the gojsonschema library
	metaschemaLoader := gojsonschema.NewReferenceLoader(metaschemaURL)
	compileResult, err := gojsonschema.Validate(metaschemaLoader, schemaLoader)
	if err != nil {
		return err
	}

	if !compileResult.Valid() {
		validationErrors := compileResult.Errors()
		errorMessages := make([]string, len(validationErrors))
		for i, err := range validationErrors {
			errorMessages[i] = err.String()
		}
		return errors.New("jsonSchema validation failed: " + strings.Join(errorMessages, ", "))
	}

	return nil
}

func (schema *Schema) IsSchemaValid() error {
	if schema.SchemaType == "" {
		return errors.New("schemaType is empty")
	}
	if schema.Service == "" {
		return errors.New("service is empty")
	}
	if schema.Source == "" {
		return errors.New("source is empty")
	}
	if schema.JsonSchema == nil {
		return errors.New("jsonSchema is empty")
	}
	err := ValidateJSONSchema(schema.JsonSchema)
	if err != nil {
		return err
	}
	return nil
}
