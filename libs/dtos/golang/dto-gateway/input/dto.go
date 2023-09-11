package input

import (
	sharedDTO "libs/dtos/golang/dto-gateway/shared"
)

type InputDTO struct {
	Data map[string]interface{} `json:"data"`
}

type InputStatusDTO struct {
	ID     string           `json:"id"`
	Status sharedDTO.Status `json:"status"`
}

type StagingJobDTO struct {
	InputId      string                 `json:"input_id"`
	Input        map[string]interface{} `json:"input"`
	Source       string                 `json:"source"`
	Service      string                 `json:"service"`
	ProcessingId string                 `json:"processing_id"`
}

type StagingJobIDDTO struct {
	ID string `json:"id"`
}
