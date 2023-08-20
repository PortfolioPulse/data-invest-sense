package usecase

import (
	"apps/lake-orchestration/lake-controller/internal/entity"
	"fmt"
	"libs/golang/events"
)

type CreateConfigUseCase struct {
	ConfigRepository entity.ConfigInterface
	ConfigCreated    events.EventInterface
	EventDispatcher  events.EventDispatcherInterface
}

func NewCreateConfigUseCase(
	repository entity.ConfigInterface,
	ConfigCreated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *CreateConfigUseCase {
	return &CreateConfigUseCase{
		ConfigRepository: repository,
		ConfigCreated:    ConfigCreated,
		EventDispatcher:  EventDispatcher,
	}
}

func ConvertEntityToUsecaseDependencies(entityDeps []entity.JobDependencies) []JobDependencies {
	usecaseDeps := make([]JobDependencies, len(entityDeps))
	for i, dep := range entityDeps {
		usecaseDeps[i] = JobDependencies{
			Service: dep.Service,
			Source:  dep.Source,
		}
	}
	return usecaseDeps
}

func (ccu *CreateConfigUseCase) Execute(config ConfigInputDTO) (ConfigOutputDTO, error) {
     entityJobDependencies := make([]entity.JobDependencies, len(config.DependsOn))
	for i, dep := range config.DependsOn {
		entityJobDependencies[i] = entity.JobDependencies{
			Service: dep.Service,
			Source:  dep.Source,
		}
	}

     configEntity, err := entity.NewConfig(
		config.Name,
		config.Active,
		config.Service,
		config.Source,
		config.Context,
		entityJobDependencies,
		config.JobParameters,
	)
	if err != nil {
		return ConfigOutputDTO{}, err
	}

	err = ccu.ConfigRepository.SaveConfig(configEntity)
	if err != nil {
		return ConfigOutputDTO{}, err
	}

     usecaseDeps := ConvertEntityToUsecaseDependencies(configEntity.DependsOn)

	dto := ConfigOutputDTO{
		ID:            string(configEntity.ID),
		Name:          configEntity.Name,
		Active:        configEntity.Active,
		Service:       configEntity.Service,
		Source:        configEntity.Source,
		Context:       configEntity.Context,
		DependsOn:     usecaseDeps,
		JobParameters: configEntity.JobParameters,
	}

	ccu.ConfigCreated.SetPayload(dto)
	ccu.EventDispatcher.Dispatch(ccu.ConfigCreated, "configs", fmt.Sprintf("%s.%s", dto.Service, dto.Source))

	return dto, nil
}
