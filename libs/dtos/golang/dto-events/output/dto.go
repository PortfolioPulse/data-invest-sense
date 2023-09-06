package output

import (
	sharedDTO "libs/dtos/golang/dto-gateway/shared"
)

type ServiceFeedbackDTO struct {
	ID       string                 `json:"id"`
	Data     map[string]interface{} `json:"data"`
	Metadata sharedDTO.Metadata     `json:"metadata"`
	Status   sharedDTO.Status       `json:"status"`
}
