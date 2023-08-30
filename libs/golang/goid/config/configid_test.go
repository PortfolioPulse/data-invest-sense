package config

import (
     "testing"
     "github.com/stretchr/testify/assert"
     "github.com/stretchr/testify/suite"
)

type GoidConfigidSuite struct {
     suite.Suite
}

func TestGoidConfigidSuite(t *testing.T) {
     suite.Run(t, new(GoidConfigidSuite))
}

func (suite *GoidConfigidSuite) TestNewID() {
     id := NewID("service", "source")

     assert.Equal(suite.T(), "service-source", id, "NewID() returned an incorrect ID")
}


