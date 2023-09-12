package database

import (
	"apps/lake-orchestration/lake-gateway/internal/entity"
	"context"
	"libs/golang/goid/md5"
	"log"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StagingJobRepositorySuite struct {
	suite.Suite
	Client     *mongo.Client
	Database   string
	Source     string
	repo       *StagingJobRepository
}

func TestStagingJobRepositorySuite(t *testing.T) {
	suite.Run(t, new(StagingJobRepositorySuite))
}

func (suite *StagingJobRepositorySuite) SetupTest() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	mongoURI := "mongodb://localhost:27017"
	client, _ := mongo.Connect(ctx, options.Client().ApplyURI(mongoURI))
	suite.Client = client
	suite.Database = "test-database"
	suite.Source = "test-source"
	suite.repo = NewStagingJobRepository(suite.Client, suite.Database)
}

func (suite *StagingJobRepositorySuite) TearDownTest() {
	suite.Client.Database(suite.Database).Drop(context.Background())
	err := suite.Client.Disconnect(context.Background())
	suite.NoError(err)
}

func (suite *StagingJobRepositorySuite) TestSaveStagingJobWhenIsANewStagingJob() {
     service := "test-service-new"
	stagingJob, err := entity.NewStagingJob("test-input-id", map[string]interface{}{"test": "first-test"}, suite.Source, service, "test-processing-id")
	suite.NoError(err)
	err = suite.repo.SaveStagingJob(stagingJob)
	suite.NoError(err)
}

func (suite *StagingJobRepositorySuite) TestFindOneByIdWhenDocumentExists() {
	data := map[string]interface{}{"test": "test"}
	service := "test-service-new"
	documentId := md5.NewWithSourceAndServiceID(data, suite.Source, service)
	inputId := md5.NewWithSourceID(data, suite.Source)
	processingId := uuid.New()

	stagingJob, err := entity.NewStagingJob(string(inputId), data, suite.Source, service, processingId.String())
	suite.NoError(err)

	err = suite.repo.SaveStagingJob(stagingJob)
	suite.NoError(err)

	result, err := suite.repo.FindOneById(string(documentId))
	suite.NoError(err)
	suite.Equal(stagingJob.ID, result.ID)
	suite.Equal(stagingJob.Input, result.Input)
	suite.Equal(stagingJob.Source, result.Source)
	suite.Equal(stagingJob.Service, result.Service)
	suite.Equal(stagingJob.InputId, result.InputId)
	suite.Equal(stagingJob.ProcessingId, result.ProcessingId)
}

func (suite *StagingJobRepositorySuite) TestFindOneByIdWhenDocumentDoesNotExist() {
	result, err := suite.repo.FindOneById("test-id")
	suite.Error(err)
	suite.Nil(result)
}

func (suite *StagingJobRepositorySuite) TestFindAllByServiceAndSource() {
	data := map[string]interface{}{"test": "test"}
	service := "test-service-new"
	inputId := md5.NewWithSourceID(data, suite.Source)
	processingId := uuid.New()

	stagingJob, err := entity.NewStagingJob(string(inputId), data, suite.Source, service, processingId.String())
	suite.NoError(err)
	err = suite.repo.SaveStagingJob(stagingJob)
	suite.NoError(err)
	result, err := suite.repo.FindAllByServiceAndSource(service, suite.Source)
	log.Printf("result: %+v\n", result)
	suite.NoError(err)
	suite.Equal(stagingJob.ID, result[0].ID)
	suite.Equal(stagingJob.Input, result[0].Input)
	suite.Equal(stagingJob.Source, result[0].Source)
	suite.Equal(stagingJob.Service, result[0].Service)
	suite.Equal(stagingJob.InputId, result[0].InputId)
	suite.Equal(stagingJob.ProcessingId, result[0].ProcessingId)
}

func (suite *StagingJobRepositorySuite) TestFindAllByServiceAndSourceWhenNoDocumentsExist() {
	result, err := suite.repo.FindAllByServiceAndSource("test-service", "test-source")
	suite.NoError(err)
	suite.Empty(result)
}

func (suite *StagingJobRepositorySuite) TestDeleteStagingJobByIdWhenDocumentExists() {
	data := map[string]interface{}{"test": "test"}
	service := "test-service-new"
	documentId := md5.NewWithSourceAndServiceID(data, suite.Source, service)
	inputId := md5.NewWithSourceID(data, suite.Source)
	processingId := uuid.New()

	stagingJob, err := entity.NewStagingJob(string(inputId), data, suite.Source, service, processingId.String())
	suite.NoError(err)
	err = suite.repo.SaveStagingJob(stagingJob)
	suite.NoError(err)

	err = suite.repo.DeleteById(string(documentId))
	suite.NoError(err)

	result, err := suite.repo.FindOneById(string(documentId))
	suite.Error(err)
	suite.Nil(result)
}

func (suite *StagingJobRepositorySuite) TestDeleteStagingJobByIdWhenDocumentDoesNotExist() {
	err := suite.repo.DeleteById("other-id")
	suite.Nil(err)
}

func (suite *StagingJobRepositorySuite) TestUpdateStagingJobWhenDocumentExists() {
     data := map[string]interface{}{"test": "test"}
     service := "test-service-new"
     documentId := md5.NewWithSourceAndServiceID(data, suite.Source, service)
     inputId := md5.NewWithSourceID(data, suite.Source)
     processingId := uuid.New()

     stagingJob, err := entity.NewStagingJob(string(inputId), data, suite.Source, service, processingId.String())
     suite.NoError(err)
     err = suite.repo.SaveStagingJob(stagingJob)
     suite.NoError(err)

     stagingJob.Input["test"] = "test-2"
     err = suite.repo.SaveStagingJob(stagingJob)
     suite.NoError(err)

     result, err := suite.repo.FindOneById(string(documentId))
     suite.NoError(err)
     suite.Equal(stagingJob.ID, result.ID)
     suite.Equal(stagingJob.Input, result.Input)
     suite.Equal(stagingJob.Source, result.Source)
     suite.Equal(stagingJob.Service, result.Service)
     suite.Equal(stagingJob.InputId, result.InputId)
     suite.Equal(stagingJob.ProcessingId, result.ProcessingId)
}

func (suite *StagingJobRepositorySuite) TestUpdateStagingJobWhenDocumentDoesNotExist() {
     //
     data := map[string]interface{}{"test": "test"}
     service := "test-service-new"
     documentId := md5.NewWithSourceAndServiceID(data, suite.Source, service)
     inputId := md5.NewWithSourceID(data, suite.Source)
     processingId := uuid.New()

     stagingJob, err := entity.NewStagingJob(string(inputId), data, suite.Source, service, processingId.String())
     suite.NoError(err)

     stagingJob.Input["test"] = "test-2"
     err = suite.repo.SaveStagingJob(stagingJob)
     suite.NoError(err)

     result, err := suite.repo.FindOneById(string(documentId))
     suite.NoError(err)
     suite.Equal(stagingJob.ID, result.ID)
     suite.Equal(stagingJob.Input, result.Input)
     suite.Equal(stagingJob.Source, result.Source)
     suite.Equal(stagingJob.Service, result.Service)
     suite.Equal(stagingJob.InputId, result.InputId)
     suite.Equal(stagingJob.ProcessingId, result.ProcessingId)
}

func (suite *StagingJobRepositorySuite) TestFindAllWhenNoDocumentsExist() {
     result, err := suite.repo.FindAll()
     suite.NoError(err)
     suite.Empty(result)
}

func (suite *StagingJobRepositorySuite) TestFindAllWhenDocumentsExist() {
     data := map[string]interface{}{"test": "test"}
     service := "test-service-new"
     inputId := md5.NewWithSourceID(data, suite.Source)
     processingId := uuid.New()

     stagingJob, err := entity.NewStagingJob(string(inputId), data, suite.Source, service, processingId.String())
     suite.NoError(err)
     err = suite.repo.SaveStagingJob(stagingJob)
     suite.NoError(err)

     result, err := suite.repo.FindAll()
     suite.NoError(err)
     suite.Equal(stagingJob.ID, result[0].ID)
     suite.Equal(stagingJob.Input, result[0].Input)
     suite.Equal(stagingJob.Source, result[0].Source)
     suite.Equal(stagingJob.Service, result[0].Service)
     suite.Equal(stagingJob.InputId, result[0].InputId)
     suite.Equal(stagingJob.ProcessingId, result[0].ProcessingId)
}

func (suite *StagingJobRepositorySuite) TestFindAllByInputIdWhenNoDocumentsExist() {
     result, err := suite.repo.FindAllByInputId("test-input-id")
     suite.NoError(err)
     suite.Empty(result)
}

func (suite *StagingJobRepositorySuite) TestFindAllByInputIdWhenDocumentsExist() {
     data := map[string]interface{}{"test": "test"}
     service := "test-service-new"
     inputId := md5.NewWithSourceID(data, suite.Source)
     processingId := uuid.New()

     stagingJob, err := entity.NewStagingJob(string(inputId), data, suite.Source, service, processingId.String())
     suite.NoError(err)
     err = suite.repo.SaveStagingJob(stagingJob)
     suite.NoError(err)

     result, err := suite.repo.FindAllByInputId(string(inputId))
     suite.NoError(err)
     suite.Equal(stagingJob.ID, result[0].ID)
     suite.Equal(stagingJob.Input, result[0].Input)
     suite.Equal(stagingJob.Source, result[0].Source)
     suite.Equal(stagingJob.Service, result[0].Service)
     suite.Equal(stagingJob.InputId, result[0].InputId)
     suite.Equal(stagingJob.ProcessingId, result[0].ProcessingId)
}

func (suite *StagingJobRepositorySuite) TestFindAllByProcessingIdWhenNoDocumentsExist() {
     result, err := suite.repo.FindAllByProcessingId("test-processing-id")
     suite.NoError(err)
     suite.Empty(result)
}

func (suite *StagingJobRepositorySuite) TestFindAllByProcessingIdWhenDocumentsExist() {
     data := map[string]interface{}{"test": "test"}
     service := "test-service-new"
     inputId := md5.NewWithSourceID(data, suite.Source)
     processingId := uuid.New()

     stagingJob, err := entity.NewStagingJob(string(inputId), data, suite.Source, service, processingId.String())
     suite.NoError(err)
     err = suite.repo.SaveStagingJob(stagingJob)
     suite.NoError(err)

     result, err := suite.repo.FindAllByProcessingId(processingId.String())
     suite.NoError(err)
     suite.Equal(stagingJob.ID, result[0].ID)
     suite.Equal(stagingJob.Input, result[0].Input)
     suite.Equal(stagingJob.Source, result[0].Source)
     suite.Equal(stagingJob.Service, result[0].Service)
     suite.Equal(stagingJob.InputId, result[0].InputId)
     suite.Equal(stagingJob.ProcessingId, result[0].ProcessingId)
}

func (suite *StagingJobRepositorySuite) TestFindAllByServiceWhenNoDocumentsExist() {
     result, err := suite.repo.FindAllByService("test-service")
     suite.NoError(err)
     suite.Empty(result)
}

func (suite *StagingJobRepositorySuite) TestFindAllByServiceWhenDocumentsExist() {
     data := map[string]interface{}{"test": "test"}
     service := "test-service-new"
     inputId := md5.NewWithSourceID(data, suite.Source)
     processingId := uuid.New()

     stagingJob, err := entity.NewStagingJob(string(inputId), data, suite.Source, service, processingId.String())
     suite.NoError(err)
     err = suite.repo.SaveStagingJob(stagingJob)
     suite.NoError(err)

     result, err := suite.repo.FindAllByService(service)
     suite.NoError(err)
     suite.Equal(stagingJob.ID, result[0].ID)
     suite.Equal(stagingJob.Input, result[0].Input)
     suite.Equal(stagingJob.Source, result[0].Source)
     suite.Equal(stagingJob.Service, result[0].Service)
     suite.Equal(stagingJob.InputId, result[0].InputId)
     suite.Equal(stagingJob.ProcessingId, result[0].ProcessingId)
}
