package handlers

import (
	"encoding/json"
	// "fmt"
	// "io"
	"net/http"

	"apps/lake-orchestration/lake-repository/internal/entity"
	"apps/lake-orchestration/lake-repository/internal/usecase"
	"libs/golang/events"

	"github.com/go-chi/jwtauth"
)

type WebSchemaHandler struct {
	EventDispatcher    events.EventDispatcherInterface
	SchemaRepository   entity.SchemaInterface
	SchemaCreatedEvent events.EventInterface
     TokenAuth          *jwtauth.JWTAuth
}

func NewWebSchemaHandler(
	EventDispatcher events.EventDispatcherInterface,
	SchemaRepository entity.SchemaInterface,
	SchemaCreatedEvent events.EventInterface,
     TokenAuth *jwtauth.JWTAuth,
) *WebSchemaHandler {
	return &WebSchemaHandler{
		EventDispatcher:    EventDispatcher,
		SchemaRepository:   SchemaRepository,
		SchemaCreatedEvent: SchemaCreatedEvent,
          TokenAuth:          TokenAuth,
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
          h.TokenAuth,
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
