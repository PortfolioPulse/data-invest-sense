package handlers

import (
	"apps/lake-orchestration/lake-controller/internal/entity"
	"apps/lake-orchestration/lake-controller/internal/usecase"
	"encoding/json"
	inputDTO "libs/dtos/golang/dto-controller/input"
	sharedDTO "libs/dtos/golang/dto-controller/shared"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type WebProcessingJobDependenciesHandler struct {
	ProcessingJobDependenciesRepository entity.ProcessingJobDependenciesInterface
}

func NewWebProcessingJobDependenciesHandler(
	repository entity.ProcessingJobDependenciesInterface,
) *WebProcessingJobDependenciesHandler {
	return &WebProcessingJobDependenciesHandler{
		ProcessingJobDependenciesRepository: repository,
	}
}

func (h *WebProcessingJobDependenciesHandler) CreateProcessingJobDependenciesHandler(w http.ResponseWriter, r *http.Request) {
	var inputDTO inputDTO.ProcessingJobDependenciesDTO
	err := json.NewDecoder(r.Body).Decode(&inputDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	useCase := usecase.NewCreateProcessingJobDependenciesUseCase(
		h.ProcessingJobDependenciesRepository,
	)

	output, err := useCase.Execute(inputDTO)
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

func (h *WebProcessingJobDependenciesHandler) ListOneProcessingJobDependenciesByIdHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	useCase := usecase.NewListOneProcessingJobDependenciesByIdUseCase(
		h.ProcessingJobDependenciesRepository,
	)

	output, err := useCase.Execute(id)
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

func (h *WebProcessingJobDependenciesHandler) RemoveProcessingJobDependenciesHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	useCase := usecase.NewRemoveProcessingJobDependenciesUseCase(
		h.ProcessingJobDependenciesRepository,
	)

	err := useCase.Execute(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (h *WebProcessingJobDependenciesHandler) UpdateProcessingJobDependenciesHandler(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	var jobDepsDTO sharedDTO.ProcessingJobDependencies
	err := json.NewDecoder(r.Body).Decode(&jobDepsDTO)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	log.Printf("UpdateProcessingJobDependenciesHandler: id=%s, jobDepsDTO=%v", id, jobDepsDTO)

	useCase := usecase.NewUpdateProcessingJobDependenciesUseCase(
		h.ProcessingJobDependenciesRepository,
	)

	err = useCase.Execute(jobDepsDTO, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	updatedDoc, err := h.ProcessingJobDependenciesRepository.FindOneById(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = json.NewEncoder(w).Encode(updatedDoc)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
