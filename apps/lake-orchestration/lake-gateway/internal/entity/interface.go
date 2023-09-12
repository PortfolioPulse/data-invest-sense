package entity

type InputInterface interface {
     SaveInput(input *Input, service string) error
     FindOneByIdAndService(id string, service string) (*Input, error)
     FindAllByService(service string) ([]*Input, error)
     FindAllByServiceAndSource(service string, source string) ([]*Input, error)
     FindAllByServiceAndSourceAndStatus(service string, source string, status int) ([]*Input, error)
}

type StagingJobInterface interface {
     SaveStagingJob(stagingJob *StagingJob) error
     FindOneById(id string) (*StagingJob, error)
     DeleteById(id string) error
     FindAll() ([]*StagingJob, error)
     FindAllByService(service string) ([]*StagingJob, error)
     FindAllByServiceAndSource(service string, source string) ([]*StagingJob, error)
     FindAllByInputId(inputId string) ([]*StagingJob, error)
     FindAllByProcessingId(processingId string) ([]*StagingJob, error)
     FindOneStagingJobUsingServiceSourceAndId(service string, source string, inputId string) (*StagingJob, error)
}

