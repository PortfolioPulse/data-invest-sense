package entity

type ConfigInterface interface {
	SaveConfig(config *Config) error
	FindAll() ([]*Config, error)
	FindAllByService(service string) ([]*Config, error)
	FindOneById(id string) (*Config, error)
	FindAllByDependentJod(service string, source string) ([]*Config, error)
     FindAllByServiceAndContext(service string, contextEnv string) ([]*Config, error)
}

type ProcessingJobDependenciesInterface interface {
     SaveProcessingJobDependencies(processingJobDependencies *ProcessingJobDependencies) error
     UpdateProcessingJobDependencies(jobDep *JobDependenciesWithProcessingData, id string) error
     DeleteProcessingJobDependencies(id string) error
     FindOneById(id string) (*ProcessingJobDependencies, error)
}
