package md5

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type GoidMd5idSuite struct {
	suite.Suite
}

func TestGoidMd5idSuite(t *testing.T) {
	suite.Run(t, new(GoidMd5idSuite))
}

func (suite *GoidMd5idSuite) TestNewID_EmptyData() {
	data := map[string]interface{}{}
	expected := "d41d8cd98f00b204e9800998ecf8427e" // MD5 hash of an empty string
	assert.Equal(suite.T(), ID(expected), NewID(data), "Test case failed")
}

func (suite *GoidMd5idSuite) TestNewID_StringValues() {
	data := map[string]interface{}{
		"name":     "John Doe",
		"location": "New York",
	}
	expected := "faa1991af0e3f6034e7f4d7417571f25"
	assert.Equal(suite.T(), ID(expected), NewID(data), "Test case failed")
}

func (suite *GoidMd5idSuite) TestNewID_IntegerValues() {
	data := map[string]interface{}{
		"age": 30,
	}
	expected := "2aff98efe944ef37381ad3730b73eaa1"
	assert.Equal(suite.T(), ID(expected), NewID(data), "Test case failed")
}

func (suite *GoidMd5idSuite) TestNewID_MixedTypes() {
	data := map[string]interface{}{
		"name": "John Doe",
		"age":  30,
	}
	expected := "716af8daca73784a0907477ba684e5eb"
	assert.Equal(suite.T(), ID(expected), NewID(data), "Test case failed")
}

func (suite *GoidMd5idSuite) TestNewID_NestedMaps() {
	data := map[string]interface{}{
		"info": map[string]interface{}{
			"location": "New York",
			"name":     "John Doe",
		},
		"age": 30,
	}
	expected := "5ffa7346c3333e7182be9ca015cccc93"
	assert.Equal(suite.T(), ID(expected), NewID(data), "Test case failed")
}

func (suite *GoidMd5idSuite) TestNewWithSourceID_EmptyData() {
	data := map[string]interface{}{}
	source := "test"
	expected := "4aae176d98e48abf482d80d5151e6703" // MD5 hash of "source=test"
	assert.Equal(suite.T(), ID(expected), NewWithSourceID(data, source), "Test case failed")
}

func (suite *GoidMd5idSuite) TestNewWithSourceID_StringAndIntegerValues() {
	data := map[string]interface{}{
		"name":     "John Doe",
		"age":      30,
		"location": "New York",
	}
	source := "info"
	expected := "0582c16b19a20e2060ef9d7549f3e3ef" // MD5 hash of "age=30|location=New York|name=John Doe|source=info"
	assert.Equal(suite.T(), ID(expected), NewWithSourceID(data, source), "Test case failed")
}

func (suite *GoidMd5idSuite) TestNewWithSourceID_NestedMaps() {
	data := map[string]interface{}{
		"info": map[string]interface{}{
			"location": "Los Angeles",
			"name":     "Jane Smith",
		},
		"age": 25,
	}
	source := "nested"
	expected := "d64261ea2a42e61fbf07c03dc9c58082" // MD5 hash of "age=25|info.location=Los Angeles|info.name=Jane Smith|source=nested"
	assert.Equal(suite.T(), ID(expected), NewWithSourceID(data, source), "Test case failed")
}

func (suite *GoidMd5idSuite) TestNewWithSourceID_SpecialCharacters() {
	data := map[string]interface{}{
		"message":  "Hello, world!",
		"number":   123,
		"specials": "~!@#$%^&*()_+{}|:\"<>?-=[]\\;',./",
	}
	source := "special"
	expected := "d605ef42ab483ceec3995180aa838bef" // MD5 hash of "message=Hello, world!|number=123|specials=~!@#$%^&*()_+{}|:\"<>?-=[]\\;',./|source=special"
	assert.Equal(suite.T(), ID(expected), NewWithSourceID(data, source), "Test case failed")
}
