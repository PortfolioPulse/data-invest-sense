package entity

import (
	"errors"
	"libs/golang/goid/config"
)

var (
	ErrProcessingJobDependenciesInvalid      = errors.New("processing Job Dependencies is invalid")
	ErrProcessingJobDependenciesServiceEmpty = errors.New("processing Job Dependencies service is empty")
	ErrProcessingJobDependenciesSourceEmpty  = errors.New("processing Job Dependencies source is empty")
)

type JobDependenciesWithProcessingData struct {
	Service             string `bson:"service"`
	Source              string `bson:"source"`
	ProcessingId        string `bson:"processing_id"`
	ProcessingTimestamp string `bson:"processing_timestamp"`
	StatusCode          int    `bson:"status_code"`
}

type ProcessingJobDependencies struct {
	ID              config.ID                           `bson:"id"`
	Service         string                              `bson:"service"`
	Source          string                              `bson:"source"`
	JobDependencies []JobDependenciesWithProcessingData `bson:"job_dependencies"`
}

func NewProcessingJobDependencies(
	service string,
	source string,
	jobDependencies []JobDependenciesWithProcessingData,
) (*ProcessingJobDependencies, error) {
	processingJobDependencies := &ProcessingJobDependencies{
		ID:              config.NewID(service, source),
		Service:         service,
		Source:          source,
		JobDependencies: jobDependencies,
	}

	err := processingJobDependencies.IsProcessingJobDependenciesValid()
	if err != nil {
		return nil, err
	}
	return processingJobDependencies, nil
}

func (p *ProcessingJobDependencies) IsProcessingJobDependenciesValid() error {
	if p.Service == "" {
		return ErrProcessingJobDependenciesServiceEmpty
	}
	if p.Source == "" {
		return ErrProcessingJobDependenciesSourceEmpty
	}
	if len(p.JobDependencies) == 0 {
		return ErrProcessingJobDependenciesInvalid
	}
	return nil
}
