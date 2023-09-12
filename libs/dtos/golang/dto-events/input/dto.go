package input

import (
	sharedDTO "libs/dtos/golang/dto-events/shared"
)

type ServiceFeedbackDTO struct {
	Data     map[string]interface{} `json:"data"`
	Metadata sharedDTO.Metadata     `json:"metadata"`
	Status   sharedDTO.Status       `json:"status"`
}
