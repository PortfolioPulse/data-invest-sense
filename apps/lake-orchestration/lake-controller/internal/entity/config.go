package entity

import (
	"errors"
	"libs/golang/goid/config"
	"time"
)

type JobDependencies struct {
	Service string `json:"service"`
	Source  string `json:"source"`
}

type Config struct {
	ID                config.ID              `json:"id"`
	Name              string                 `json:"name"`
	Active            bool                   `json:"active"`
	Service           string                 `json:"service"`
	Source            string                 `json:"source"`
	Context           string                 `json:"context"`
	DependsOn         []JobDependencies      `json:"dependsOn"`
	ServiceParamaters map[string]interface{} `json:"serviceParamaters"`
	JobParameters     map[string]interface{} `json:"jobParameters"`
	CreatedAt         string                 `json:"createdAt"`
	UpdatedAt         string                 `json:"updatedAt"`
}

func NewConfig(name string, active bool, service string, source string, context string, dependsOn []JobDependencies, jobParameters map[string]interface{}, serviceParameters map[string]interface{}) (*Config, error) {
	config := &Config{
		ID:                config.NewID(service, source),
		Name:              name,
		Active:            active,
		Service:           service,
		Source:            source,
		Context:           context,
		DependsOn:         dependsOn,
		ServiceParamaters: serviceParameters,
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
		return errors.New("name is empty")
	}
	if config.Service == "" {
		return errors.New("service is empty")
	}
	if config.Source == "" {
		return errors.New("source is empty")
	}
	if config.Context == "" {
		return errors.New("context is empty")
	}
	if config.JobParameters == nil {
		return errors.New("jobParameters is empty")
	}
     if config.ServiceParamaters == nil {
          return errors.New("serviceParamaters is empty")
     }
	return nil
}
