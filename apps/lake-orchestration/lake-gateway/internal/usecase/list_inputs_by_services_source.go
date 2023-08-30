package usecase

import (
	"apps/lake-orchestration/lake-gateway/internal/entity"
)

type ListAllByServiceAndSourceUseCase struct {
	InputRepository entity.InputInterface
}

func NewListAllByServiceAndSourceUseCase(
	repository entity.InputInterface,
) *ListAllByServiceAndSourceUseCase {
	return &ListAllByServiceAndSourceUseCase{
		InputRepository: repository,
	}
}

func (la *ListAllByServiceAndSourceUseCase) Execute(service string, source string) ([]InputOutputDTO, error) {
     items, err := la.InputRepository.FindAllByServiceAndSource(service, source)
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
