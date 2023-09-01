package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type StagingJobEntitySuite struct {
	suite.Suite
}

func TestStagingJobEntitySuite(t *testing.T) {
	suite.Run(t, new(StagingJobEntitySuite))
}

func (suite *StagingJobEntitySuite) TestGivenAnEmptySource_WhenCreateANewStagingJob_ThenShouldReceiveAnError() {
     stagingJob, err := NewStagingJob("test", map[string]interface{}{"test": "test"}, "", "file-downloader", "test")
     assert.Nil(suite.T(), stagingJob)
     assert.Error(suite.T(), err, "invalid source")
}

func (suite *StagingJobEntitySuite) TestGivenAnEmptyService_WhenCreateANewStagingJob_ThenShouldReceiveAnError() {
     stagingJob, err := NewStagingJob("test", map[string]interface{}{"test": "test"}, "test", "", "test")
     assert.Nil(suite.T(), stagingJob)
     assert.Error(suite.T(), err, "invalid service")
}

func (suite *StagingJobEntitySuite) TestGivenNoSourceAndService_WhenCreateANewStagingJob_ThenShouldReceiveAnError() {
     stagingJob, err := NewStagingJob("test", map[string]interface{}{"test": "test"}, "", "", "test")
     assert.Nil(suite.T(), stagingJob)
     assert.Error(suite.T(), err, "invalid source and service")
}

func (suite *StagingJobEntitySuite) TestGivenAnEmptyInputId_WhenCreateANewStagingJob_ThenShouldReceiveAnError() {
     stagingJob, err := NewStagingJob("", map[string]interface{}{"test": "test"}, "test", "file-downloader", "test")
     assert.Nil(suite.T(), stagingJob)
     assert.Error(suite.T(), err, "invalid input id")
}

func (suite *StagingJobEntitySuite) TestGivenAnEmptyInput_WhenCreateANewStagingJob_ThenShouldReceiveAnError() {
     stagingJob, err := NewStagingJob("test", nil, "test", "file-downloader", "test")
     assert.Nil(suite.T(), stagingJob)
     assert.Error(suite.T(), err, "invalid input")
}

func (suite *StagingJobEntitySuite) TestGivenAnEmptyProcessingId_WhenCreateANewStagingJob_ThenShouldReceiveAnError() {
     stagingJob, err := NewStagingJob("test", map[string]interface{}{"test": "test"}, "test", "file-downloader", "")
     assert.Nil(suite.T(), stagingJob)
     assert.Error(suite.T(), err, "invalid processing id")
}

func (suite *StagingJobEntitySuite) TestGivenAValidStagingJob_WhenCreateANewStagingJob_ThenShouldReceiveANewStagingJob() {
     stagingJob, err := NewStagingJob("test", map[string]interface{}{"test": "test"}, "test", "file-downloader", "test")
     assert.NotNil(suite.T(), stagingJob)
     assert.NoError(suite.T(), err)

     assert.Equal(suite.T(), "test", stagingJob.InputId)
     assert.Equal(suite.T(), map[string]interface{}{"test": "test"}, stagingJob.Input)
     assert.Equal(suite.T(), "test", stagingJob.Source)
     assert.Equal(suite.T(), "file-downloader", stagingJob.Service)
     assert.Equal(suite.T(), "test", stagingJob.ProcessingId)
}

func (suite *StagingJobEntitySuite) TestGivenAValidStagingJob_WhenValidateStagingJob_ThenShouldReceiveNoError() {
     stagingJob, _ := NewStagingJob("test", map[string]interface{}{"test": "test"}, "test", "file-downloader", "test")
     err := stagingJob.IsStagingJobValid()
     assert.NoError(suite.T(), err)
}
