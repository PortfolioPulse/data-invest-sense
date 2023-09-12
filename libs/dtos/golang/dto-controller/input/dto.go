package input

import (
	sharedDTO "libs/dtos/golang/dto-controller/shared"
)

type ConfigDTO struct {
	Name              string                      `json:"name"`
	Active            bool                        `json:"active"`
	Frequency         string                      `json:"frequency"`
	Service           string                      `json:"service"`
	Source            string                      `json:"source"`
	Context           string                      `json:"context"`
	DependsOn         []sharedDTO.JobDependencies `json:"depends_on"`
	ServiceParameters map[string]interface{}      `json:"service_parameters"`
	JobParameters     map[string]interface{}      `json:"job_parameters"`
}

type ProcessingJobDependenciesDTO struct {
	Service         string                                `json:"service"`
	Source          string                                `json:"source"`
	JobDependencies []sharedDTO.ProcessingJobDependencies `json:"job_dependencies"`
}
