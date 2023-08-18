package schemas

import (
	entity "apps/lake-manager/internal/entity/schemas"
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type SchemaInputRepositoryTestSuite struct {
	suite.Suite
	Client *mongo.Client
}

func (suite *SchemaInputRepositoryTestSuite) SetupTest() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
     mongoURI := "mongodb://localhost:27017"
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
     suite.Client = client
}

func (suite *SchemaInputRepositoryTestSuite) TearDownTest() {
     suite.Client.Database("schemas").Collection("input").Drop(context.Background())
	err := suite.Client.Disconnect(context.Background())
	suite.NoError(err)
}

func TestSchemaInputRepositorySuite(t *testing.T) {
	suite.Run(t, new(SchemaInputRepositoryTestSuite))
}

func (suite *SchemaInputRepositoryTestSuite) TestSave() {
	input, err := entity.NewInput("new-test2", []string{"test"}, map[string]interface{}{"test": "test"})
	suite.NoError(err)
	repo := NewSchemaInputRepository(suite.Client)
	err = repo.Save(input)
	suite.NoError(err)
}
