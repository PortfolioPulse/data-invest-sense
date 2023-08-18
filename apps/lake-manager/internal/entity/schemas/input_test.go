package schemas

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"apps/lake-manager/pkg/uid"
)

type SchemaInputEntitySuite struct {
	suite.Suite
}

func (suite *SchemaInputEntitySuite) TestGivenAnEmptyID_WhenCreateANewSchemaInput_ThenShouldReceiveAnError() {
	input := SchemaInput{}
	assert.Error(suite.T(), input.IsValid(), "invalid id")
}

func (suite *SchemaInputEntitySuite) TestGivenAnEmptyRequired_WhenCreateANewSchemaInput_ThenShouldReceiveAnError() {
	input := SchemaInput{ID: "test"}
	assert.Error(suite.T(), input.IsValid(), "invalid required")
}

func (suite *SchemaInputEntitySuite) TestGivenAnEmptyProperties_WhenCreateANewSchemaInput_ThenShouldReceiveAnError() {
	input := SchemaInput{ID: "test", Required: []string{"test"}}
	assert.Error(suite.T(), input.IsValid(), "invalid properties")
}

func (suite *SchemaInputEntitySuite) TestGivenAValidParams_WhenCallNewSchemaInput_ThenShouldReceiveCreateSchemaInputWithAllParams() {
	input := SchemaInput{
		ID:         "test",
		Required:   []string{"test"},
		Properties: map[string]interface{}{"test": "test"},
	}
	assert.Equal(suite.T(), "test", input.ID)
	assert.Equal(suite.T(), []string{"test"}, input.Required)
	assert.Equal(suite.T(), map[string]interface{}{"test": "test"}, input.Properties)
	assert.Nil(suite.T(), input.IsValid())
}

func (suite *SchemaInputEntitySuite) TestGivenAValidParams_WhenCallNewSchemaInputFunc_ThenShouldReceiveCreateSchemaInputWithAllParams() {
	input, err := NewInput("test", []string{"test"}, map[string]interface{}{"test": "test"})
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "test", input.ID)
	assert.Equal(suite.T(), []string{"test"}, input.Required)
	assert.Equal(suite.T(), map[string]interface{}{"test": "test"}, input.Properties)
	assert.IsType(suite.T(), uid.ID{}, input.SchemaId)
	assert.Nil(suite.T(), input.IsValid())
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(SchemaInputEntitySuite))
}
