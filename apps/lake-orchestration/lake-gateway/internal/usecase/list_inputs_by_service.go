package usecase

import (
	"apps/lake-orchestration/lake-gateway/internal/entity"
)

type ListAllByServiceUseCase struct {
	InputRepository entity.InputInterface
}

func NewListAllByServiceUseCase(
	repository entity.InputInterface,
) *ListAllByServiceUseCase {
	return &ListAllByServiceUseCase{
		InputRepository: repository,
	}
}

func (la *ListAllByServiceUseCase) Execute(service string) ([]InputOutputDTO, error) {
     items, err := la.InputRepository.FindAllByService(service)
     if err != nil {
          return []InputOutputDTO{}, err
     }
     var result []InputOutputDTO
     for _, item := range items {
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
          result = append(result, dto)
     }
     return result, nil
}
