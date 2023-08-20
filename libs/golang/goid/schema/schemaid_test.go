package schema

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GoidSchemaidSuite struct {
	suite.Suite
}

func TestGoidSchemaidSuite(t *testing.T) {
	suite.Run(t, new(GoidSchemaidSuite))
}

func (suite *GoidSchemaidSuite) TestNewID() {
	id := NewID("schemaType", "service", "source")

	assert.Equal(suite.T(), "schemaType-service-source", id, "NewID() returned an incorrect ID")
}
