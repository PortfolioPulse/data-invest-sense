package entity

import (
	"errors"
	"libs/golang/goid/config"
	"time"
)

var (
	ErrConfigNameEmpty              = errors.New("config name is empty")
	ErrConfigFrequencyEmpty         = errors.New("config frequency is empty")
	ErrConfigServiceEmpty           = errors.New("config service is empty")
	ErrConfigSourceEmpty            = errors.New("config source is empty")
	ErrConfigContextEmpty           = errors.New("config context is empty")
	ErrConfigDependsOnEmpty         = errors.New("config dependsOn is empty")
	ErrConfigJobParametersEmpty     = errors.New("config jobParameters is empty")
	ErrConfigServiceParametersEmpty = errors.New("config serviceParameters is empty")
)

type JobDependencies struct {
	Service string `json:"service"`
	Source  string `json:"source"`
}

type Config struct {
	ID                config.ID              `bson:"id"`
	Name              string                 `bson:"name"`
	Active            bool                   `bson:"active"`
	Frequency         string                 `bson:"frequency"`
	Service           string                 `bson:"service"`
	Source            string                 `bson:"source"`
	Context           string                 `bson:"context"`
	DependsOn         []JobDependencies      `bson:"depends_on"`
	ServiceParameters map[string]interface{} `bson:"service_parameters"`
	JobParameters     map[string]interface{} `bson:"job_parameters"`
	CreatedAt         string                 `bson:"created_at"`
	UpdatedAt         string                 `bson:"updated_at"`
}

func NewConfig(
	name string,
	active bool,
	frequency string,
	service string,
	source string,
	context string,
	dependsOn []JobDependencies,
	jobParameters map[string]interface{},
	serviceParameters map[string]interface{},
) (*Config, error) {
	config := &Config{
		ID:                config.NewID(service, source),
		Name:              name,
		Active:            active,
		Frequency:         frequency,
		Service:           service,
		Source:            source,
		Context:           context,
		DependsOn:         dependsOn,
		ServiceParameters: serviceParameters,
		JobParameters:     jobParameters,
		CreatedAt:         time.Now().Format(time.RFC3339),
		UpdatedAt:         time.Now().Format(time.RFC3339),
	}
	err := config.IsConfigValid()
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (config *Config) IsConfigValid() error {
	if config.Name == "" {
		return ErrConfigNameEmpty
	}
	if config.Frequency == "" {
		return ErrConfigFrequencyEmpty
	}
	if config.Service == "" {
		return ErrConfigServiceEmpty
	}
	if config.Source == "" {
		return ErrConfigSourceEmpty
	}
	if config.Context == "" {
		return ErrConfigContextEmpty
	}
	if config.JobParameters == nil {
		return ErrConfigJobParametersEmpty
	}
	if config.ServiceParameters == nil {
		return ErrConfigServiceParametersEmpty
	}
	return nil
}
