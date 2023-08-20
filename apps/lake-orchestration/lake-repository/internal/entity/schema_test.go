package entity

import (
	"testing"

	"github.com/go-chi/jwtauth"
	"github.com/stretchr/testify/suite"
)

type SchemaSuite struct {
	suite.Suite
	tokenAuth *jwtauth.JWTAuth
}

func TestSchemaSuite(t *testing.T) {
	suite.Run(t, new(SchemaSuite))
}

func (suite *SchemaSuite) SetupTest() {
	// Initialize a test TokenAuth instance
	suite.tokenAuth = jwtauth.New("HS256", []byte("your-secret-key"), nil)
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

	schemaInstance, err := NewSchema(schemaType, service, source, jsonSchema, suite.tokenAuth)

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
          Service: "test",
          Source: "source",
		JsonSchema: jsonSchema,
	}

	err := schemaInstance.IsSchemaValid()

	suite.NoError(err)
}

func (suite *SchemaSuite) TestIsSchemaValid_EmptySchemaType() {
	schemaInstance := &Schema{
          Service: "test",
          Source: "source",
		JsonSchema: map[string]interface{}{},
	}

	err := schemaInstance.IsSchemaValid()

	suite.Error(err)
	suite.EqualError(err, "schemaType is empty")
}

func (suite *SchemaSuite) TestIsSchemaValid_EmptyJsonSchema() {
     schemaInstance := &Schema{
          SchemaType: "example",
          Service: "test",
          Source: "source",
	}

	err := schemaInstance.IsSchemaValid()

	suite.Error(err)
	suite.EqualError(err, "jsonSchema is empty")
}

func (suite *SchemaSuite) TestIsSchemaValid_InvalidJsonSchema() {
     schemaInstance := &Schema{
          SchemaType: "example",
          Service: "test",
          Source: "source",
          JsonSchema: nil,
     }

     err := schemaInstance.IsSchemaValid()

     suite.Error(err)
     suite.EqualError(err, "jsonSchema is empty")
}
