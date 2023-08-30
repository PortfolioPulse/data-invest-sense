package usecase

type JobDependencies struct {
	Service string `json:"service"`
	Source  string `json:"source"`
}

type ConfigInputDTO struct {
	Name              string                 `json:"name"`
	Active            bool                   `json:"active"`
	Service           string                 `json:"service"`
	Source            string                 `json:"source"`
	Context           string                 `json:"context"`
	DependsOn         []JobDependencies      `json:"dependsOn"`
	ServiceParameters map[string]interface{} `json:"serviceParameters"`
	JobParameters     map[string]interface{} `json:"jobParameters"`
}

type ConfigOutputDTO struct {
	ID                string                 `json:"id"`
	Name              string                 `json:"name"`
	Active            bool                   `json:"active"`
	Service           string                 `json:"service"`
	Source            string                 `json:"source"`
	Context           string                 `json:"context"`
	DependsOn         []JobDependencies      `json:"dependsOn"`
	ServiceParameters map[string]interface{} `json:"serviceParameters"`
	JobParameters     map[string]interface{} `json:"jobParameters"`
	CreatedAt         string                 `json:"createdAt"`
	UpdatedAt         string                 `json:"updatedAt"`
}
