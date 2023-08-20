package jwt

import (
	"testing"

	"github.com/go-chi/jwtauth"
	"github.com/stretchr/testify/suite"
)

type GoJWTSuite struct {
	suite.Suite
	tokenAuth *jwtauth.JWTAuth
}

func TestGoJWTSuite(t *testing.T) {
	suite.Run(t, new(GoJWTSuite))
}

func (suite *GoJWTSuite) SetupTest() {
	// Initialize a test TokenAuth instance
	suite.tokenAuth = jwtauth.New("HS256", []byte("your-secret-key"), nil)
}

func (suite *GoJWTSuite) TestGenerateJWT() {
	jsonSchema := map[string]interface{}{
		"field1": "value1",
		"field2": 42,
	}

	token, err := GenerateSchemaJWT(suite.tokenAuth, jsonSchema)

	suite.NoError(err)
	suite.NotEmpty(token)
}

func (suite *GoJWTSuite) TestGenerateJWT_EmptyJsonSchema() {
	emptyJsonSchema := map[string]interface{}{}

	token, err := GenerateSchemaJWT(suite.tokenAuth, emptyJsonSchema)

	suite.Error(err)
	suite.Contains(err.Error(), "JsonSchema is empty")
	suite.Empty(token)
}
