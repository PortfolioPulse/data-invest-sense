package handlers

import (
	"apps/lake-orchestration/lake-gateway/internal/entity"
	"apps/lake-orchestration/lake-gateway/internal/usecase"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type WebStagingJobHandler struct {
	StagingJobRepository entity.StagingJobInterface
}

func NewWebStagingJobHandler(
	StagingJobRepository entity.StagingJobInterface,
) *WebStagingJobHandler {
	return &WebStagingJobHandler{
		StagingJobRepository: StagingJobRepository,
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

func (h *WebStagingJobHandler) RemoveStagingJob(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	dto := usecase.StagingJobInputIDDTO{
		ID: id,
	}
	removeStagingJob := usecase.NewRemoveStagingJobUseCase(
		h.StagingJobRepository,
	)

	output, err := removeStagingJob.Execute(dto)
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

func (h *WebStagingJobHandler) ListOneStagingJobUsingServiceSourceId(w http.ResponseWriter, r *http.Request) {
     service := chi.URLParam(r, "service")
     source := chi.URLParam(r, "source")
     id := chi.URLParam(r, "id")

     listStagingJobs := usecase.NewListOneStagingJobUsingServiceSourceIdUseCase(
          h.StagingJobRepository,
     )

     output, err := listStagingJobs.Execute(service, source, id)
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
