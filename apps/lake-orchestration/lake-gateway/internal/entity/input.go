package entity

import (
	"errors"
	"libs/golang/goid/md5"
	"libs/golang/goid/uuid"
	"time"
)

type Metadata struct {
	ProcessingId        uuid.ID
	ProcessingTimestamp string
	Source              string
	Service             string
}

type Status struct {
	Code   int    `json:"code"`
	Detail string `json:"detail"`
}

type Input struct {
	ID       md5.ID                 `json:"id"`
	Data     map[string]interface{} `json:"data"`
	Metadata Metadata               `json:"metadata"`
	Status   Status                 `json:"status"`
}

type InputStatus struct {
     ID     md5.ID `json:"id"`
     Status Status `json:"status"`
}

func NewInputStatus(id string, status int, detail string) (*InputStatus, error) {
	inputStatus := &InputStatus{
          ID: md5.ID(id),
          Status: Status{
               Code: status,
               Detail: detail,
          },
     }
	err := inputStatus.IsStatusValid()
	if err != nil {
		return nil, err
	}
	return inputStatus, nil
}

func NewInput(data map[string]interface{}, source string, service string) (*Input, error) {
	input := &Input{
		ID:   md5.NewWithSourceID(data, source),
		Data: data,
		Metadata: Metadata{
			ProcessingId:        uuid.NewID(),
			ProcessingTimestamp: time.Now().Format(time.RFC3339),
			Source:              source,
			Service:             service,
		},
		Status: Status{
			Code:   0,
			Detail: "",
		},
	}
	err := input.IsValid()
	if err != nil {
		return nil, err
	}
	return input, nil
}

func (i *Input) IsValid() error {
	if i.Data == nil {
		return errors.New("data field is required")
	}
	if i.Metadata.Source == "" {
		return errors.New("source field is required")
	}
	if i.Metadata.Service == "" {
		return errors.New("service field is required")
	}
	return nil
}

func (is *InputStatus) IsStatusValid() error {
     if is.ID == "" {
          return errors.New("id is required")
     }
	if is.Status.Code == 0 && is.Status.Detail == "" {
		return errors.New("status code is required")
	}
	return nil
}


