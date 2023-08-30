package handlers

import (
	"apps/lake-orchestration/lake-gateway/internal/entity"
	"apps/lake-orchestration/lake-gateway/internal/usecase"
	"encoding/json"
	"libs/golang/events"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type WebInputStatusHandler struct {
	EventDispatcher         events.EventDispatcherInterface
	InputRepository   entity.InputInterface
	InputUpdatedEvent events.EventInterface
}

func NewWebInputStatusHandler(
	EventDispatcher events.EventDispatcherInterface,
	InputRepository entity.InputInterface,
	InputUpdatedEvent events.EventInterface,
) *WebInputStatusHandler {
	return &WebInputStatusHandler{
		EventDispatcher:         EventDispatcher,
		InputRepository:   InputRepository,
		InputUpdatedEvent: InputUpdatedEvent,
	}
}

func (h *WebInputStatusHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
     service := chi.URLParam(r, "service")
     // source := chi.URLParam(r, "source")
     id := chi.URLParam(r, "id")

     var dto usecase.InputStatusInputDTO
     err := json.NewDecoder(r.Body).Decode(&dto)
     if err != nil {
          http.Error(w, err.Error(), http.StatusBadRequest)
          return
     }
     dto.ID = id
     updateStatusInput := usecase.NewUpdateStatusInputUseCase(
          h.InputRepository,
          h.InputUpdatedEvent,
          h.EventDispatcher,
     )

     output, err := updateStatusInput.Execute(dto, service)
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

