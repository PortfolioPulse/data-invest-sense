package usecase

// import (
//      "apps/lake-orchestration/lake-gateway/internal/entity"
//      "fmt"
//      "libs/golang/events"
//      "libs/golang/goid/uuid"
// )

// type CreateInputWithProcessingIdUseCase struct {
//      InputRepository entity.InputInterface
//      InputCreated    events.EventInterface
//      EventDispatcher events.EventDispatcherInterface
// }

// func NewCreateInputWithProcessingIdUseCase(
//      repository entity.InputInterface,
//      InputCreated events.EventInterface,
//      EventDispatcher events.EventDispatcherInterface,
// ) *CreateInputWithProcessingIdUseCase {
//      return &CreateInputWithProcessingIdUseCase{
//           InputRepository: repository,
//           InputCreated:    InputCreated,
//           EventDispatcher: EventDispatcher,
//      }
// }

// func (ciu *CreateInputWithProcessingIdUseCase) Execute(input InputInputDTO, service string, source string, processingId string) (InputOutputDTO, error) {
// 	inputEntity, err := entity.NewInput(input.Data, source, service)
// 	if err != nil {
// 		return InputOutputDTO{}, err
// 	}

//      inputEntity.Metadata.ProcessingId, _ = uuid.ParseID(processingId)

//      err = ciu.InputRepository.SaveInput(inputEntity, service)
//      if err != nil {
//           return InputOutputDTO{}, err
//      }

//      dto := InputOutputDTO{
//           ID:   string(inputEntity.ID),
//           Data: inputEntity.Data,
//           Metadata: Metadata{
//                ProcessingId:        inputEntity.Metadata.ProcessingId.String(),
//                ProcessingTimestamp: inputEntity.Metadata.ProcessingTimestamp,
//                Source:              inputEntity.Metadata.Source,
//                Service:             inputEntity.Metadata.Service,
//           },
//           Status: Status{
//                Code:   inputEntity.Status.Code,
//                Detail: inputEntity.Status.Detail,
//           },
//      }
//      ciu.InputCreated.SetPayload(dto)
//      ciu.EventDispatcher.Dispatch(ciu.InputCreated, "services", fmt.Sprintf("%s.inputs.%s", dto.Metadata.Service, dto.Metadata.Source))

//      return dto, nil
// }
