package usecase

type JobDependencies struct {
	Service string `json:"service"`
	Source  string `json:"source"`
}

type ConfigInputDTO struct {
	Name              string                 `json:"name"`
	Active            bool                   `json:"active"`
	Frequency         string                 `json:"frequency"`
	Service           string                 `json:"service"`
	Source            string                 `json:"source"`
	Context           string                 `json:"context"`
	DependsOn         []JobDependencies      `json:"depends_on"`
	ServiceParameters map[string]interface{} `json:"service_parameters"`
	JobParameters     map[string]interface{} `json:"job_parameters"`
}

type ConfigOutputDTO struct {
	ID                string                 `json:"id"`
	Name              string                 `json:"name"`
	Active            bool                   `json:"active"`
	Frequency         string                 `json:"frequency"`
	Service           string                 `json:"service"`
	Source            string                 `json:"source"`
	Context           string                 `json:"context"`
	DependsOn         []JobDependencies      `json:"depends_on"`
	ServiceParameters map[string]interface{} `json:"service_parameters"`
	JobParameters     map[string]interface{} `json:"job_parameters"`
	CreatedAt         string                 `json:"created_at"`
	UpdatedAt         string                 `json:"updated_at"`
}
