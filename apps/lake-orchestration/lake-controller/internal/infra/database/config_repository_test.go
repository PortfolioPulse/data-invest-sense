package database

import (
	"apps/lake-orchestration/lake-controller/internal/entity"
	"apps/lake-orchestration/lake-controller/internal/infra/grpc/service"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConfigRepositorySuite struct {
     suite.Suite
     Client    *mongo.Client
     Database  string
     Collection string
     repo     *ConfigRepository
}

func TestConfigRepositorySuite(t *testing.T) {
     suite.Run(t, new(ConfigRepositorySuite))
}

func (suite *ConfigRepositorySuite) SetupTest() {
     ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
     defer cancel()
     mongoURI := "mongodb://localhost:27017"
     client, _ := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
     suite.Client = client
     suite.Database = "test-database"
     suite.Collection = "test-service"
     suite.repo = NewConfigRepository(suite.Client, suite.Database)
}

func (suite *ConfigRepositorySuite) TearDownTest() {
     suite.Client.Database(suite.Database).Drop(context.Background())
     err := suite.Client.Disconnect(context.Background())
     suite.NoError(err)
}

func (suite *ConfigRepositorySuite) TestSaveConfigWhenIsANewConfig() {
     jobParams := map[string]interface{}{
          "test": "test",
     }
     serviceParams := map[string]interface{}{
          "test": "test",
     }
     dependsOn := []entity.JobDependencies{
          {
               Service: "test",
               Source: "test",
          },
     }
     config, err := entity.NewConfig("test", true, "test", "test", "test", dependsOn, jobParams, serviceParams)
     suite.NoError(err)
     err = suite.repo.SaveConfig(config)
     suite.NoError(err)
}

func (suite *ConfigRepositorySuite) TestSaveConfigWhenIsAExistingConfig() {
     jobParams := map[string]interface{}{
          "test": "test",
     }
     serviceParams := map[string]interface{}{
          "test": "test",
     }
     dependsOn := []entity.JobDependencies{
          {
               Service: "test",
               Source: "test",
          },
     }
     config, err := entity.NewConfig("test", true, "test", "test", "test", dependsOn, jobParams, serviceParams)
     suite.NoError(err)
     err = suite.repo.SaveConfig(config)
     suite.NoError(err)
     err = suite.repo.SaveConfig(config)
     suite.NoError(err)
}

func (suite *ConfigRepositorySuite) TestFindOneById() {
     jobParams := map[string]interface{}{
          "test": "test",
     }
     serviceParams := map[string]interface{}{
          "test": "test",
     }
     dependsOn := []entity.JobDependencies{
          {
               Service: "test",
               Source: "test",
          },
     }
     config, err := entity.NewConfig("test", true, "test", "test", "test", dependsOn, jobParams, serviceParams)
     suite.NoError(err)
     err = suite.repo.SaveConfig(config)
     suite.NoError(err)
     result, err := suite.repo.FindOneById(string(config.ID))
     suite.NoError(err)
     suite.Equal(config, result)
}

func (suite *ConfigRepositorySuite) TestFindAllByService() {
     jobParams := map[string]interface{}{
          "test": "test",
     }
     serviceParams := map[string]interface{}{
          "test": "test",
     }
     dependsOn := []entity.JobDependencies{
          {
               Service: "test",
               Source: "test",
          },
     }
     config, err := entity.NewConfig("test", true, "test", "test", "test", dependsOn, jobParams, serviceParams)
     suite.NoError(err)
     err = suite.repo.SaveConfig(config)
     suite.NoError(err)
     results, err := suite.repo.FindAllByService("test")
     suite.NoError(err)
     suite.Len(results, 1)
}

func (suite *ConfigRepositorySuite) TestFindAll() {
     jobParams1 := map[string]interface{}{
          "url": "http://test1.com",
     }
     serviceParams1 := map[string]interface{}{
          "test": "test",
     }
     dependsOn1 := []entity.JobDependencies{
          {
               Service: "test",
               Source: "test3",
          },
     }
     config1, err := entity.NewConfig("test", true, "test", "test", "test", dependsOn1, jobParams1, serviceParams1)
     suite.NoError(err)

     jobParams2 := map[string]interface{}{
          "url": "http://test2.com",
     }
     serviceParams2 := map[string]interface{}{
          "test": "test",
     }
     dependsOn2 := []entity.JobDependencies{
          {
               Service: "test",
               Source: "test3",
          },
     }
     config2, err := entity.NewConfig("test", true, "test", "test", "test", dependsOn2, jobParams2, serviceParams2)
     suite.NoError(err)

     err = suite.repo.SaveConfig(config1)
     suite.NoError(err)
     err = suite.repo.SaveConfig(config2)
     suite.NoError(err)

     results, err := suite.repo.FindAll()
     suite.NoError(err)
     suite.Len(results, 2)
}
