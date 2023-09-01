package entity

import (
	"errors"
	"libs/golang/goid/md5"
)

var (
	ErrStagingJobIDEmpty           = errors.New("staging job id is empty")
	ErrStagingJobInputIDEmpty      = errors.New("staging job input id is empty")
	ErrStagingJobInputEmpty        = errors.New("staging job input is empty")
	ErrStagingJobSourceEmpty       = errors.New("staging job source is empty")
	ErrStagingJobServiceEmpty      = errors.New("staging job service is empty")
	ErrStagingJobProcessingIDEmpty = errors.New("staging job processing id is empty")
)

type StagingJob struct {
	ID           md5.ID                 `bson:"id"`
	InputId      string                 `bson:"input_id"`
	Input        map[string]interface{} `bson:"input"`
	Source       string                 `bson:"source"`
	Service      string                 `bson:"service"`
	ProcessingId string                 `bson:"processing_id"`
}

func NewStagingJob(
	inputId string,
	input map[string]interface{},
	source string,
	service string,
	processingId string,
) (*StagingJob, error) {
	stagingJob := &StagingJob{
		ID:           md5.NewWithSourceAndServiceID(input, source, service),
		InputId:      inputId,
		Input:        input,
		Source:       source,
		Service:      service,
		ProcessingId: processingId,
	}
	err := stagingJob.IsStagingJobValid()
	if err != nil {
		return nil, err
	}
	return stagingJob, nil
}

func (s *StagingJob) IsStagingJobValid() error {
	if s.ID == "" {
		return ErrStagingJobIDEmpty
	}
	if s.InputId == "" {
		return ErrStagingJobInputIDEmpty
	}
	if s.Input == nil {
		return ErrStagingJobInputEmpty
	}
	if s.Source == "" {
		return ErrStagingJobSourceEmpty
	}
	if s.Service == "" {
		return ErrStagingJobServiceEmpty
	}
	if s.ProcessingId == "" {
		return ErrStagingJobProcessingIDEmpty
	}
	return nil
}
