package uuid

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"github.com/google/uuid"
)

type GoidUuidSuite struct {
	suite.Suite
}

func TestGoidUuidSuite(t *testing.T) {
	suite.Run(t, new(GoidUuidSuite))
}

func (suite *GoidUuidSuite) TestNewID() {
	id := NewID()

	// Check if the ID is a valid UUID
	_, err := uuid.Parse(id.String())
	assert.NoError(suite.T(), err, "NewID() produced an invalid UUID")
}

func (suite *GoidUuidSuite) TestParseID() {
	// Valid UUID string
	validUUIDString := "f47ac10b-58cc-0372-8567-0e02b2c3d479"
	expectedID, _ := uuid.Parse(validUUIDString)

	// Test parsing a valid UUID string
	parsedID, err := ParseID(validUUIDString)
	assert.NoError(suite.T(), err, "ParseID(%s) returned an unexpected error", validUUIDString)
	assert.Equal(suite.T(), ID(expectedID), parsedID, "ParseID(%s) returned an incorrect ID", validUUIDString)
     assert.Equal(suite.T(), validUUIDString, parsedID.String(), "ParseID(%s) returned an incorrect ID", validUUIDString)

	// Invalid UUID string
	invalidUUIDString := "invalid-uuid"

	// Test parsing an invalid UUID string
	_, err = ParseID(invalidUUIDString)
	assert.Error(suite.T(), err, "ParseID(%s) should have returned an error, but it didn't", invalidUUIDString)
}
