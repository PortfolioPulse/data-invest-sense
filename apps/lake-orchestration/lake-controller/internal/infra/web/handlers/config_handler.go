package handlers

import (
	"apps/lake-orchestration/lake-controller/internal/entity"
	"apps/lake-orchestration/lake-controller/internal/usecase"
	"encoding/json"
	"libs/golang/events"
	"net/http"
)

type WebConfigHandler struct {
	EventDispatcher    events.EventDispatcherInterface
	ConfigRepository   entity.ConfigInterface
	ConfigCreatedEvent events.EventInterface
}

func NewWebConfigHandler(
     EventDispatcher events.EventDispatcherInterface,
     ConfigRepository entity.ConfigInterface,
     ConfigCreatedEvent events.EventInterface,
) *WebConfigHandler {
     return &WebConfigHandler{
          EventDispatcher:    EventDispatcher,
          ConfigRepository:   ConfigRepository,
          ConfigCreatedEvent: ConfigCreatedEvent,
     }
}

func (h *WebConfigHandler) CreateConfig(w http.ResponseWriter, r *http.Request) {
     var dto usecase.ConfigInputDTO
     err := json.NewDecoder(r.Body).Decode(&dto)
     if err != nil {
          http.Error(w, err.Error(), http.StatusBadRequest)
          return
     }

     createConfig := usecase.NewCreateConfigUseCase(
          h.ConfigRepository,
          h.ConfigCreatedEvent,
          h.EventDispatcher,
     )

     output, err := createConfig.Execute(dto)
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
