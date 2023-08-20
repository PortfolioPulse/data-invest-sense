package database

import (
	"context"
	"testing"
	"time"

	"apps/lake-orchestration/lake-gateway/internal/entity"

	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type InputRepositorySuite struct {
	suite.Suite
	Client     *mongo.Client
	Database   string
	Collection string
	Source     string
	repo       *InputRepository
}

func TestInputRepositorySuite(t *testing.T) {
	suite.Run(t, new(InputRepositorySuite))
}

func (suite *InputRepositorySuite) SetupTest() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	mongoURI := "mongodb://localhost:27017"
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	suite.Client = client
	suite.Database = "test-database"
	suite.Collection = "test-service"
	suite.Source = "test-source"
	suite.repo = NewInputRepository(suite.Client, suite.Database)
}

func (suite *InputRepositorySuite) TearDownTest() {
	suite.Client.Database(suite.Database).Drop(context.Background())
	err := suite.Client.Disconnect(context.Background())
	suite.NoError(err)
}

func (suite *InputRepositorySuite) TestSaveInputWhenIsANewInput() {
	input, err := entity.NewInput(map[string]interface{}{"test": "test"}, suite.Source, suite.Collection)
	suite.NoError(err)
	err = suite.repo.SaveInput(input, suite.Collection)
	suite.NoError(err)
}

func (suite *InputRepositorySuite) TestFindAllByService() {
	data1 := map[string]interface{}{
		"test": "test1",
	}
	input1, err := entity.NewInput(data1, suite.Source, suite.Collection)
	suite.NoError(err)
	err = suite.repo.SaveInput(input1, suite.Collection)
	suite.NoError(err)

	data2 := map[string]interface{}{
		"test": "test2",
	}
	input2, err := entity.NewInput(data2, suite.Source, suite.Collection)
	suite.NoError(err)
	err = suite.repo.SaveInput(input2, suite.Collection)
	suite.NoError(err)

	results, err := suite.repo.FindAllByService(suite.Collection)
	suite.NoError(err)
	suite.Len(results, 2)
}

func (suite *InputRepositorySuite) TestFindAllByServiceAndSource() {
	data1 := map[string]interface{}{
		"test": "test1",
	}
	input1, err := entity.NewInput(data1, suite.Source, suite.Collection)
	suite.NoError(err)
	err = suite.repo.SaveInput(input1, suite.Collection)
	suite.NoError(err)

	data2 := map[string]interface{}{
		"test": "test2",
	}
	input2, err := entity.NewInput(data2, suite.Source, suite.Collection)
	suite.NoError(err)
	err = suite.repo.SaveInput(input2, suite.Collection)
	suite.NoError(err)

	results, err := suite.repo.FindAllByServiceAndSource(suite.Collection, suite.Source)
	suite.NoError(err)
	suite.Len(results, 2)
}

func (suite *InputRepositorySuite) TestFindAllByServiceAndSourceAndStatus() {
	data1 := map[string]interface{}{
		"test": "test1",
	}
	input1, err := entity.NewInput(data1, suite.Source, suite.Collection)
	suite.NoError(err)
	err = suite.repo.SaveInput(input1, suite.Collection)
	suite.NoError(err)

	data2 := map[string]interface{}{
		"test": "test2",
	}
	input2, err := entity.NewInput(data2, suite.Source, suite.Collection)
	suite.NoError(err)
	err = suite.repo.SaveInput(input2, suite.Collection)
	suite.NoError(err)

	results, err := suite.repo.FindAllByServiceAndSourceAndStatus(suite.Collection, suite.Source, 0)
	suite.NoError(err)
	suite.Len(results, 2)
}

func (suite *InputRepositorySuite) TestFindAllByServiceAndSourceAndStatus_NonExistentInput() {
	results, err := suite.repo.FindAllByServiceAndSourceAndStatus(suite.Collection, suite.Source, 0)
	suite.NoError(err)
	suite.Empty(results)
}

func (suite *InputRepositorySuite) TestFindAllByServiceAndSource_NonExistentSource() {
	results, err := suite.repo.FindAllByServiceAndSource(suite.Collection, "nonexistent-source")
	suite.NoError(err)
	suite.Empty(results)
}

func (suite *InputRepositorySuite) TestFindAllByService_NonExistentService() {
	results, err := suite.repo.FindAllByService("nonexistent-service")
	suite.NoError(err)
	suite.Empty(results)
}

func (suite *InputRepositorySuite) TestSaveInputAndUpdateInput_NewInput() {
	input, err := entity.NewInput(map[string]interface{}{"test": "test"}, suite.Source, suite.Collection)
	suite.NoError(err)

	err = suite.repo.SaveInput(input, suite.Collection)
	suite.NoError(err)

	// Update the existing input
	input.Status.Code = 1
	err = suite.repo.SaveInput(input, suite.Collection)
	suite.NoError(err)

	// Retrieve the updated input
	updatedInput, err := suite.repo.FindOneByIdAndService(string(input.ID), suite.Collection)
	suite.NoError(err)
	suite.Equal(1, updatedInput.Status.Code)
}

func (suite *InputRepositorySuite) TestSaveInputAndUpdateInput_ExistingInput() {
	input, err := entity.NewInput(map[string]interface{}{"test": "test"}, suite.Source, suite.Collection)
	suite.NoError(err)

	// Save the input for the first time
	err = suite.repo.SaveInput(input, suite.Collection)
	suite.NoError(err)

	// Update the existing input
	input.Status.Code = 1
	err = suite.repo.SaveInput(input, suite.Collection)
	suite.NoError(err)

	// Retrieve the updated input
	updatedInput, err := suite.repo.FindOneByIdAndService(string(input.ID), suite.Collection)
	suite.NoError(err)
	suite.Equal(1, updatedInput.Status.Code)
}

func (suite *InputRepositorySuite) TestFindById_ExistingInput() {
	input, err := entity.NewInput(map[string]interface{}{"test": "test"}, suite.Source, suite.Collection)
	suite.NoError(err)

	// Save the input
	err = suite.repo.SaveInput(input, suite.Collection)
	suite.NoError(err)

	// Retrieve the input by ID
	retrievedInput, err := suite.repo.FindOneByIdAndService(string(input.ID), suite.Collection)
	suite.NoError(err)
	suite.NotNil(retrievedInput)
}

func (suite *InputRepositorySuite) TestFindById_NonExistentInput() {
	nonexistentInputID := "nonexistent-id"
	_, err := suite.repo.FindOneByIdAndService(nonexistentInputID, suite.Collection)
	suite.Error(err)
}
