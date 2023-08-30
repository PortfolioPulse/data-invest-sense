package handlers

import (
	"encoding/json"
	"net/http"

	"apps/lake-orchestration/lake-repository/internal/entity"
	"apps/lake-orchestration/lake-repository/internal/usecase"
	"libs/golang/events"

	"github.com/go-chi/chi/v5"
)

type WebSchemaHandler struct {
	EventDispatcher    events.EventDispatcherInterface
	SchemaRepository   entity.SchemaInterface
	SchemaCreatedEvent events.EventInterface
}

func NewWebSchemaHandler(
	EventDispatcher events.EventDispatcherInterface,
	SchemaRepository entity.SchemaInterface,
	SchemaCreatedEvent events.EventInterface,
) *WebSchemaHandler {
	return &WebSchemaHandler{
		EventDispatcher:    EventDispatcher,
		SchemaRepository:   SchemaRepository,
		SchemaCreatedEvent: SchemaCreatedEvent,
	}
}

func (h *WebSchemaHandler) CreateSchema(w http.ResponseWriter, r *http.Request) {
	var dto usecase.SchemaInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createSchema := usecase.NewCreateSchemaUseCase(
		h.SchemaRepository,
		h.SchemaCreatedEvent,
		h.EventDispatcher,
	)

	output, err := createSchema.Execute(dto)
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

func (h *WebSchemaHandler) ListOneSchemaById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	listOneSchemaById := usecase.NewListOneSchemaByIdUseCase(
		h.SchemaRepository,
	)

	output, err := listOneSchemaById.Execute(id)
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

func (h *WebSchemaHandler) ListAllSchemas(w http.ResponseWriter, r *http.Request) {
	listAllSchemas := usecase.NewListAllSchemasUseCase(
		h.SchemaRepository,
	)

	output, err := listAllSchemas.Execute()
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

func (h *WebSchemaHandler) ListAllSchemasByService(w http.ResponseWriter, r *http.Request) {
	service := chi.URLParam(r, "service")
	listAllSchemasByService := usecase.NewListAllSchemasByServiceUseCase(
		h.SchemaRepository,
	)

	output, err := listAllSchemasByService.Execute(service)
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
