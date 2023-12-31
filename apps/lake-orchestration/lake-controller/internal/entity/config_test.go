package entity

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type ConfigSuite struct {
	suite.Suite
}

func TestConfigSuite(t *testing.T) {
	suite.Run(t, new(ConfigSuite))
}

func (suite *ConfigSuite) TestNewConfigWhenIsANewConfig() {
	jobParams := map[string]interface{}{
		"test": "test",
	}
	serviceParams := map[string]interface{}{
		"test": "test",
	}
	dependsOn := []JobDependencies{
		{
			Service: "test",
			Source:  "test",
		},
	}
	config, err := NewConfig("test", true, "test", "test", "test", dependsOn, jobParams, serviceParams)
	suite.NoError(err)
	suite.NotNil(config)
}

func (suite *ConfigSuite) TestNewConfigWhenNameIsEmpty() {
	jobParams := map[string]interface{}{
		"test": "test",
	}
     serviceParams := map[string]interface{}{
          "test": "test",
     }
	dependsOn := []JobDependencies{
		{
			Service: "test",
			Source:  "test",
		},
	}
	config, err := NewConfig("", true, "test", "test", "test", dependsOn, jobParams, serviceParams)
	suite.Error(err)
	suite.Nil(config)
}

func (suite *ConfigSuite) TestNewConfigWhenServiceIsEmpty() {
	jobParams := map[string]interface{}{
		"test": "test",
	}
     serviceParams := map[string]interface{}{
          "test": "test",
     }
	dependsOn := []JobDependencies{
		{
			Service: "test",
			Source:  "test",
		},
	}
	config, err := NewConfig("test", true, "", "test", "test", dependsOn, jobParams, serviceParams)
	suite.Error(err)
	suite.Nil(config)
}

func (suite *ConfigSuite) TestNewConfigWhenSourceIsEmpty() {
	jobParams := map[string]interface{}{
		"test": "test",
	}
     serviceParams := map[string]interface{}{
          "test": "test",
     }
	dependsOn := []JobDependencies{
		{
			Service: "test",
			Source:  "test",
		},
	}
	config, err := NewConfig("test", true, "test", "", "test", dependsOn, jobParams, serviceParams)
	suite.Error(err)
	suite.Nil(config)
}

func (suite *ConfigSuite) TestNewConfigWhenContextIsEmpty() {
	jobParams := map[string]interface{}{
		"test": "test",
	}
     serviceParams := map[string]interface{}{
          "test": "test",
     }
	dependsOn := []JobDependencies{
		{
			Service: "test",
			Source:  "test",
		},
	}
	config, err := NewConfig("test", true, "test", "test", "", dependsOn, jobParams, serviceParams)
	suite.Error(err)
	suite.Nil(config)
}

func (suite *ConfigSuite) TestNewConfigWhenJobParametersIsEmpty() {
     serviceParams := map[string]interface{}{
          "test": "test",
     }
	dependsOn := []JobDependencies{
		{
			Service: "test",
			Source:  "test",
		},
	}
	config, err := NewConfig("test", true, "test", "test", "test", dependsOn, nil, serviceParams)
	suite.Error(err)
	suite.Nil(config)
}

func (suite *ConfigSuite) TestIsConfigValid() {
	jobParams := map[string]interface{}{
		"test": "test",
	}
     serviceParams := map[string]interface{}{
          "test": "test",
     }
	dependsOn := []JobDependencies{
		{
			Service: "test",
			Source:  "test",
		},
	}
	config, err := NewConfig("test", true, "test", "test", "test", dependsOn, jobParams, serviceParams)
	suite.NoError(err)
	err = config.IsConfigValid()
	suite.NoError(err)
}
