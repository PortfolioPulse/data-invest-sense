package handlers

import (
	"apps/lake-orchestration/lake-controller/internal/entity"
	"apps/lake-orchestration/lake-controller/internal/usecase"
	"encoding/json"
	inputDTO "libs/dtos/golang/dto-controller/input"
	"libs/golang/events"
	"net/http"

	"github.com/go-chi/chi/v5"
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
	var dto inputDTO.ConfigDTO
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

func (h *WebConfigHandler) ListAllConfigs(w http.ResponseWriter, r *http.Request) {
	listAllConfigs := usecase.NewListAllConfigsUseCase(
		h.ConfigRepository,
	)

	output, err := listAllConfigs.Execute()
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

func (h *WebConfigHandler) ListOneConfigById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	listOneConfigById := usecase.NewListOneConfigByIdUseCase(
		h.ConfigRepository,
	)

	output, err := listOneConfigById.Execute(id)
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

func (h *WebConfigHandler) ListAllConfigsByService(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	listAllConfigsByService := usecase.NewListAllConfigsByServiceUseCase(
		h.ConfigRepository,
	)

	output, err := listAllConfigsByService.Execute(service)
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

func (h *WebConfigHandler) ListAllConfigsByDependentJob(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	source := chi.URLParam(r, "source")

	listAllConfigsByDependentJob := usecase.NewListAllConfigsByDependentJobUseCase(
		h.ConfigRepository,
	)

	output, err := listAllConfigsByDependentJob.Execute(service, source)
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
