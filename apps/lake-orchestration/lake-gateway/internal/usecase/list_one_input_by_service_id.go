package usecase

import (
	"apps/lake-orchestation/lake-gateway/internal/entity"
)

type ListOneByIdAndServiceUseCase struct {
	InputRepository entity.InputInterface
}

func NewListOneByIdAndServiceUseCase(
	repository entity.InputInterface,
) *ListOneByIdAndServiceUseCase {
	return &ListOneByIdAndServiceUseCase{
		InputRepository: repository,
	}
}

func (lo *ListOneByIdAndServiceUseCase) Execute(service string, id string) (InputOutputDTO, error) {
     item, err := lo.InputRepository.FindOneByIdAndService(id, service)
     if err != nil {
          return InputOutputDTO{}, err
     }
     dto := InputOutputDTO{
          ID:   string(item.ID),
          Data: item.Data,
          Metadata: Metadata{
               ProcessingId:        item.Metadata.ProcessingId.String(),
               ProcessingTimestamp: item.Metadata.ProcessingTimestamp,
               Source:              item.Metadata.Source,
               Service:             item.Metadata.Service,
          },
          Status: Status{
               Code:   item.Status.Code,
               Detail: item.Status.Detail,
          },
     }
     return dto, nil
}
