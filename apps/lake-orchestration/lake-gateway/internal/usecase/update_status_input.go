package usecase

import (
	"apps/lake-orchestration/lake-gateway/internal/entity"
	"fmt"
	"libs/golang/events"
	"time"
)

type UpdateStatusInputUseCase struct {
	InputRepository entity.InputInterface
	InputUpdated    events.EventInterface
	EventDispatcher events.EventDispatcherInterface
}

func NewUpdateStatusInputUseCase(
	repository entity.InputInterface,
	InputUpdated events.EventInterface,
	EventDispatcher events.EventDispatcherInterface,
) *UpdateStatusInputUseCase {
	return &UpdateStatusInputUseCase{
		InputRepository: repository,
		InputUpdated:    InputUpdated,
		EventDispatcher: EventDispatcher,
	}
}

func (uiu *UpdateStatusInputUseCase) Execute(inputStatus InputStatusInputDTO, service string) (InputOutputDTO, error) {
	inputStatusEntity, err := entity.NewInputStatus(inputStatus.ID, inputStatus.Status.Code, inputStatus.Status.Detail)
	if err != nil {
		return InputOutputDTO{}, err
	}

	input, err := uiu.InputRepository.FindOneByIdAndService(inputStatus.ID, service)
	if err != nil {
		return InputOutputDTO{}, err
	}

	input.Status = inputStatusEntity.Status
	input.Metadata.ProcessingTimestamp = time.Now().Format(time.RFC3339)

	err = uiu.InputRepository.SaveInput(input, service)
	if err != nil {
		return InputOutputDTO{}, err
	}

	dto := InputOutputDTO{
		ID:   string(input.ID),
		Data: input.Data,
		Metadata: Metadata{
			ProcessingId:        input.Metadata.ProcessingId.String(),
			ProcessingTimestamp: input.Metadata.ProcessingTimestamp,
			Source:              input.Metadata.Source,
			Service:             input.Metadata.Service,
		},
		Status: Status{
			Code:   input.Status.Code,
			Detail: input.Status.Detail,
		},
	}
	uiu.InputUpdated.SetPayload(dto)
	uiu.EventDispatcher.Dispatch(uiu.InputUpdated, "inputs", fmt.Sprintf("%s.%s", dto.Metadata.Service, dto.Metadata.Source))

	return dto, nil
}
