package entity

import (
	"libs/golang/goid/md5"
	"libs/golang/goid/uuid"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type InputEntitySuite struct {
	suite.Suite
}

func TestInputEntitySuite(t *testing.T) {
	suite.Run(t, new(InputEntitySuite))
}

func (suite *InputEntitySuite) TestGivenAnEmptySource_WhenCreateANewInput_ThenShouldReceiveAnError() {
	input := Input{
		Data: map[string]interface{}{"test": "test"},
		Metadata: Metadata{
			Source:  "",
			Service: "file-downloader",
		},
	}
	assert.Error(suite.T(), input.IsValid(), "invalid source")
}

func (suite *InputEntitySuite) TestGivenAnEmptyService_WhenCreateANewInput_ThenShouldReceiveAnError() {
	input := Input{
		Data: map[string]interface{}{"test": "test"},
		Metadata: Metadata{
			Source:  "test",
			Service: "",
		},
	}
	assert.Error(suite.T(), input.IsValid(), "invalid service")
}

func (suite *InputEntitySuite) TestGivenNoSourceAndService_WhenCreateANewInput_ThenShouldReceiveAnError() {
	input := Input{
		Data:     map[string]interface{}{"test": "test"},
		Metadata: Metadata{},
	}
	assert.Error(suite.T(), input.IsValid(), "invalid source and service")
}

func (suite *InputEntitySuite) TestGivenAnEmptyData_WhenCreateANewInput_ThenShouldReceiveAnError() {
	input := Input{}
	assert.Error(suite.T(), input.IsValid(), "invalid data")
}

func (suite *InputEntitySuite) TestGivenAValidParams_WhenCallNewInput_ThenShouldReceiveCreateInputWithAllParams() {
	input := Input{
		Data: map[string]interface{}{"test": "test"},
		Metadata: Metadata{
			Source:  "test",
			Service: "file-downloader",
		},
	}
	assert.Equal(suite.T(), map[string]interface{}{"test": "test"}, input.Data)
	assert.Nil(suite.T(), input.IsValid())
}

func (suite *InputEntitySuite) TestGivenAValidParams_WhenCallNewInputFunc_ThenShouldReceiveCreateInputWithAllParams() {
	input, err := NewInput(map[string]interface{}{"test": "test"}, "test", "file-downloader")
	assert.Nil(suite.T(), err)

	assert.Equal(suite.T(), map[string]interface{}{"test": "test"}, input.Data)
	assert.Equal(suite.T(), md5.ID("6653310993a16531c52a943792dd1767"), input.ID)
	assert.Equal(suite.T(), "test", input.Metadata.Source)
	assert.Equal(suite.T(), "file-downloader", input.Metadata.Service)
	assert.Equal(suite.T(), Status{Code: 0, Detail: ""}, input.Status)

	assert.IsType(suite.T(), Metadata{}, input.Metadata)
	assert.IsType(suite.T(), uuid.ID{}, input.Metadata.ProcessingId)
	assert.Equal(suite.T(), time.Now().Format(time.RFC3339), input.Metadata.ProcessingTimestamp)

	assert.Nil(suite.T(), input.IsValid())
}

func (suite *InputEntitySuite) TestGivenNonZeroCodeAndEmptyDetail_WhenCheckingStatusValid_ThenShouldReceiveASuccess() {
	inputStatus := InputStatus{
		ID: md5.ID("some-id"),
		Status: Status{
			Code:   1,
			Detail: "",
		},
	}
	assert.Nil(suite.T(), inputStatus.IsStatusValid())
}

func (suite *InputEntitySuite) TestGiveZeroCodeAndEmptyDetail_WhenCheckingStatusValid_ThenShouldReceiveAnError() {
	inputStatus := InputStatus{
		ID: md5.ID("some-id"),
		Status: Status{
			Code:   0,
			Detail: "",
		},
	}
	assert.Error(suite.T(), inputStatus.IsStatusValid())
}

func (suite *InputEntitySuite) TestGiveZeroCodeAndDetail_WhenCheckingStatusValid_ThenShouldReceiveASuccess() {
	inputStatus := InputStatus{
		ID: md5.ID("some-id"),
		Status: Status{
			Code:   0,
			Detail: "Reprocess input",
		},
	}
	assert.Nil(suite.T(), inputStatus.IsStatusValid())
}


func (suite *InputEntitySuite) TestValidInputStatusCreation() {
	inputStatus, err := NewInputStatus("some-id", 0, "No error")
	assert.NotNil(suite.T(), inputStatus)
	assert.Nil(suite.T(), err)
}

func (suite *InputEntitySuite) TestNonZeroStatusCodeWithEmptyDetail() {
	inputStatus, err := NewInputStatus("some-id", 1, "")
     assert.NotNil(suite.T(), inputStatus)
     assert.Nil(suite.T(), err)
}

func (suite *InputEntitySuite) TestValidInputStatusWithNonZeroCodeAndDetail() {
	inputStatus, err := NewInputStatus("some-id", 404, "Some error")
	assert.NotNil(suite.T(), inputStatus)
	assert.Nil(suite.T(), err)
}

func (suite *InputEntitySuite) TestEmptyID() {
     inputStatus, err := NewInputStatus("", 0, "No error")
     assert.Nil(suite.T(), inputStatus)
     assert.NotNil(suite.T(), err)
     assert.Equal(suite.T(), "id is required", err.Error())
}

func (suite *InputEntitySuite) TestNegativeCodeAndEmptyDetail() {
	inputStatus, err := NewInputStatus("some-id", -2, "")
	assert.Nil(suite.T(), err)
     assert.NotNil(suite.T(), inputStatus)
}
