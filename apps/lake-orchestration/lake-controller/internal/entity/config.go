package entity

import (
	"errors"
	"libs/golang/goid/config"
)

type JobDependencies struct {
	Service string `json:"service"`
	Source  string `json:"source"`
}

type Config struct {
	ID            config.ID              `json:"id"`
	Name          string                 `json:"name"`
	Active        bool                   `json:"active"`
	Service       string                 `json:"service"`
	Source        string                 `json:"source"`
	Context       string                 `json:"context"`
	DependsOn     []JobDependencies      `json:"dependsOn"`
	JobParameters map[string]interface{} `json:"jobParameters"`
}

func NewConfig(name string, active bool, service string, source string, context string, dependsOn []JobDependencies, jobParameters map[string]interface{}) (*Config, error) {
	config := &Config{
		ID:            config.NewID(service, source),
		Name:          name,
		Active:        active,
		Service:       service,
		Source:        source,
		Context:       context,
		DependsOn:     dependsOn,
		JobParameters: jobParameters,
	}
	err := config.IsConfigValid()
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (config *Config) IsConfigValid() error {
	if config.Name == "" {
		return errors.New("Name is empty")
	}
	if config.Service == "" {
		return errors.New("Service is empty")
	}
	if config.Source == "" {
		return errors.New("Source is empty")
	}
	if config.Context == "" {
		return errors.New("Context is empty")
	}
	if config.JobParameters == nil {
		return errors.New("JobParameters is empty")
	}
	return nil
}
