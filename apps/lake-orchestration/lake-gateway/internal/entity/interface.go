package entity

type InputInterface interface {
     SaveInput(input *Input, service string) error
     FindOneByIdAndService(id string, service string) (*Input, error)
     FindAllByService(service string) ([]*Input, error)
     FindAllByServiceAndSource(service string, source string) ([]*Input, error)
     FindAllByServiceAndSourceAndStatus(service string, source string, status int) ([]*Input, error)
}


