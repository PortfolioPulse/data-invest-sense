package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type SchemaSuite struct {
	suite.Suite
}

func TestSchemaSuite(t *testing.T) {
	suite.Run(t, new(SchemaSuite))
}

func (suite *SchemaSuite) TestNewSchema() {
	jsonSchema := map[string]interface{}{
		"field1": map[string]interface{}{
			"type": "string",
		},
		"field2": map[string]interface{}{
			"type": "string",
		},
	}

	schemaType := "example"
	service := "test"
	source := "source"

	schemaInstance, err := NewSchema(schemaType, service, source, jsonSchema)

	suite.NoError(err)
	suite.NotNil(schemaInstance)
	suite.NotNil(schemaInstance.ID)
	suite.Equal(schemaType, schemaInstance.SchemaType)
	suite.Equal(jsonSchema, schemaInstance.JsonSchema)
	suite.NotNil(schemaInstance.SchemaID)
}

func (suite *SchemaSuite) TestIsSchemaValid_ValidSchema() {
	jsonSchema := map[string]interface{}{
		"field1": map[string]interface{}{
			"type": "string",
		},
		"field2": map[string]interface{}{
			"type": "string",
		},
	}

	schemaInstance := &Schema{
		SchemaType: "example",
		Service:    "test",
		Source:     "source",
		JsonSchema: jsonSchema,
	}

	err := schemaInstance.IsSchemaValid()

	suite.NoError(err)
}

func (suite *SchemaSuite) TestIsSchemaValid_EmptyService() {
	schemaInstance := &Schema{
		SchemaType: "example",
		Source:     "source",
		JsonSchema: map[string]interface{}{
			"field1": map[string]interface{}{
				"type": "string",
			},
		},
	}

	err := schemaInstance.IsSchemaValid()

	suite.Error(err)
	suite.EqualError(err, "service is empty")
}

func (suite *SchemaSuite) TestIsSchemaValid_EmptySource() {
	schemaInstance := &Schema{
		SchemaType: "example",
		Service:    "test",
		JsonSchema: map[string]interface{}{
			"field1": map[string]interface{}{
				"type": "string",
			},
		},
	}

	err := schemaInstance.IsSchemaValid()

	suite.Error(err)
	suite.EqualError(err, "source is empty")
}

func (suite *SchemaSuite) TestIsSchemaValid_EmptySchemaType() {
	schemaInstance := &Schema{
		Service:    "test",
		Source:     "source",
		JsonSchema: map[string]interface{}{},
	}

	err := schemaInstance.IsSchemaValid()

	suite.Error(err)
	suite.EqualError(err, "schemaType is empty")
}

func (suite *SchemaSuite) TestIsSchemaValid_EmptyJsonSchema() {
	schemaInstance := &Schema{
		SchemaType: "example",
		Service:    "test",
		Source:     "source",
		JsonSchema: nil,
	}

	err := schemaInstance.IsSchemaValid()

	suite.Error(err)
	suite.EqualError(err, "jsonSchema is empty")
}

func (suite *SchemaSuite) TestValidateJSONSchema_ValidSchema() {
	jsonSchema := map[string]interface{}{
		"type": "object",
		"properties": map[string]interface{}{
			"name": map[string]interface{}{
				"type": "string",
			},
		},
	}

	// schemaInstance := &Schema{
	// 	JsonSchema: jsonSchema,
	// }

	err := ValidateJSONSchema(jsonSchema)

	assert.NoError(suite.T(), err)
}

func (suite *SchemaSuite) TestValidateJSONSchema_InvalidSchema() {
	invalidJsonSchema := map[string]interface{}{
		"type": "invalid_type",
	}

	// schemaInstance := &Schema{
	// 	JsonSchema: invalidJsonSchema,
	// }

	err := ValidateJSONSchema(invalidJsonSchema)

	assert.Error(suite.T(), err)
	assert.Contains(suite.T(), err.Error(), "jsonSchema validation failed")
}

// func (suite *SchemaSuite) TestIsSchemaValid_InvalidJSONSchemaValidation() {
// 	// Create a JSON schema that is intentionally invalid
// 	invalidJsonSchema := map[string]interface{}{
// 		"invalidField": "invalidValue",
// 	}

// 	// Create a Schema instance with an invalid JSON schema
// 	schemaInstance := &Schema{
// 		SchemaType: "example",
// 		Service:    "test",
// 		Source:     "source",
// 		JsonSchema: invalidJsonSchema,
// 	}

// 	// Call IsSchemaValid which should return an error due to invalid JSON schema
// 	err := schemaInstance.IsSchemaValid()

// 	// Assert that an error is returned
// 	assert.Error(suite.T(), err)
// 	assert.Contains(suite.T(), err.Error(), "jsonSchema validation failed")
// }
