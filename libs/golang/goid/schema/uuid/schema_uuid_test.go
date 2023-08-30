package encryptid

import (
     "testing"

     "github.com/stretchr/testify/assert"
     "github.com/stretchr/testify/suite"
)

type SchemaIDEncryptSuite struct {
	suite.Suite
}

func TestSchemaIDEncryptSuite(t *testing.T) {
     suite.Run(t, new(SchemaIDEncryptSuite))
}


func (suite *SchemaIDEncryptSuite) TestGenerateSchemaID() {
	properties := map[string]interface{}{
		"field1": map[string]interface{}{
			"type": "string",
		},
		"field2": map[string]interface{}{
			"type": "string",
		},
	}
	schemaType := "your-schema-type"

	schemaID, err := GenerateSchemaID(schemaType, properties)

	assert.NoError(suite.T(), err)
	assert.NotEmpty(suite.T(), schemaID)
     assert.Equal(suite.T(), "a68b228e-e042-5855-81a3-d0f803c86af9", schemaID)
}
