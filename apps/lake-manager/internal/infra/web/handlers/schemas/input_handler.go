package schemas

import (
	entity "apps/lake-manager/internal/entity/schemas"
	usecase "apps/lake-manager/internal/usecase/schemas"
	"apps/lake-manager/pkg/events"
	"encoding/json"
	"net/http"
)

type WebSchemaInputHandler struct {
	EventDispatcher         events.EventDispatcherInterface
	SchemaInputRepository   entity.SchemaInputInterface
	SchemaInputCreatedEvent events.EventInterface
}

func NewWebSchemaInputHandler(
	EventDispatcher events.EventDispatcherInterface,
	SchemaInputRepository entity.SchemaInputInterface,
	SchemaInputCreatedEvent events.EventInterface,
) *WebSchemaInputHandler {
	return &WebSchemaInputHandler{
		EventDispatcher:         EventDispatcher,
		SchemaInputRepository:   SchemaInputRepository,
		SchemaInputCreatedEvent: SchemaInputCreatedEvent,
	}
}

func (h *WebSchemaInputHandler) Create(w http.ResponseWriter, r *http.Request) {
	var dto usecase.SchemaInputInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	createSchemaInput := usecase.NewCreateSchemaInputUseCase(
		h.SchemaInputRepository,
		h.SchemaInputCreatedEvent,
		h.EventDispatcher,
	)

	output, err := createSchemaInput.Execute(dto)
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
