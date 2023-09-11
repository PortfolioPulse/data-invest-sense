package output

import (
	sharedDTO "libs/dtos/golang/dto-controller/shared"
)

type ConfigDTO struct {
	ID                string                      `json:"id"`
	Name              string                      `json:"name"`
	Active            bool                        `json:"active"`
	Frequency         string                      `json:"frequency"`
	Service           string                      `json:"service"`
	Source            string                      `json:"source"`
	Context           string                      `json:"context"`
	DependsOn         []sharedDTO.JobDependencies `json:"depends_on"`
	ServiceParameters map[string]interface{}      `json:"service_parameters"`
	JobParameters     map[string]interface{}      `json:"job_parameters"`
	CreatedAt         string                      `json:"created_at"`
	UpdatedAt         string                      `json:"updated_at"`
}

type ProcessingJobDependenciesDTO struct {
	ID              string                                `json:"id"`
	Service         string                                `json:"service"`
	Source          string                                `json:"source"`
	JobDependencies []sharedDTO.ProcessingJobDependencies `json:"job_dependencies"`
}
