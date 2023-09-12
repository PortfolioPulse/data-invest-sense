package handlers

import (
	"apps/lake-orchestration/lake-gateway/internal/entity"
	"apps/lake-orchestration/lake-gateway/internal/usecase"
	"encoding/json"
	"libs/golang/events"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type WebInputHandler struct {
	EventDispatcher   events.EventDispatcherInterface
	InputRepository   entity.InputInterface
	InputCreatedEvent events.EventInterface
}

func NewWebInputHandler(
	EventDispatcher events.EventDispatcherInterface,
	InputRepository entity.InputInterface,
	InputCreatedEvent events.EventInterface,
) *WebInputHandler {
	return &WebInputHandler{
		EventDispatcher:   EventDispatcher,
		InputRepository:   InputRepository,
		InputCreatedEvent: InputCreatedEvent,
	}
}

func (h *WebInputHandler) CreateInput(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	source := chi.URLParam(r, "source")
	var dto usecase.InputInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createInput := usecase.NewCreateInputUseCase(
		h.InputRepository,
		h.InputCreatedEvent,
		h.EventDispatcher,
	)

	output, err := createInput.Execute(dto, service, source)
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

func (h *WebInputHandler) ListAllByServiceAndSource(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	source := chi.URLParam(r, "source")
	listInputs := usecase.NewListAllByServiceAndSourceUseCase(
		h.InputRepository,
	)
	output, err := listInputs.Execute(service, source)
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

func (h *WebInputHandler) ListAllByService(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	listInputs := usecase.NewListAllByServiceUseCase(
		h.InputRepository,
	)
	output, err := listInputs.Execute(service)
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

func (h *WebInputHandler) ListOneByIdAndService(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	id := chi.URLParam(r, "id")
	listInputs := usecase.NewListOneByIdAndServiceUseCase(
		h.InputRepository,
	)
	output, err := listInputs.Execute(service, id)
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
