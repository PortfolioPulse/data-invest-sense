package output

import (
	sharedDTO "libs/dtos/golang/dto-gateway/shared"
)

type InputDTO struct {
	ID       string                 `json:"id"`
	Data     map[string]interface{} `json:"data"`
	Metadata sharedDTO.Metadata     `json:"metadata"`
	Status   sharedDTO.Status       `json:"status"`
}

type ListInputDTO struct {
	Inputs []InputDTO `json:"inputs"`
}

type StagingJobDTO struct {
	ID           string                 `json:"id"`
	InputId      string                 `json:"input_id"`
	Input        map[string]interface{} `json:"input"`
	Source       string                 `json:"source"`
	Service      string                 `json:"service"`
	ProcessingId string                 `json:"processing_id"`
}
