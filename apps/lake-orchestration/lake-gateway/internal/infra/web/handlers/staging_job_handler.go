package handlers

import (
	"apps/lake-orchestration/lake-gateway/internal/entity"
	"apps/lake-orchestration/lake-gateway/internal/usecase"
	"encoding/json"
	"libs/golang/events"
	"net/http"
)

type WebStagingJobHandler struct {
	EventDispatcher        events.EventDispatcherInterface
	StagingJobRepository   entity.StagingJobInterface
	StagingJobCreatedEvent events.EventInterface
}

func NewWebStagingJobHandler(
	EventDispatcher events.EventDispatcherInterface,
	StagingJobRepository entity.StagingJobInterface,
	StagingJobCreatedEvent events.EventInterface,
) *WebStagingJobHandler {
	return &WebStagingJobHandler{
		EventDispatcher:        EventDispatcher,
		StagingJobRepository:   StagingJobRepository,
		StagingJobCreatedEvent: StagingJobCreatedEvent,
	}
}

func (h *WebStagingJobHandler) CreateStagingJob(w http.ResponseWriter, r *http.Request) {
	var dto usecase.StagingJobInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createStagingJob := usecase.NewCreateStagingJobUseCase(
		h.StagingJobRepository,
		h.StagingJobCreatedEvent,
		h.EventDispatcher,
	)

	output, err := createStagingJob.Execute(dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
