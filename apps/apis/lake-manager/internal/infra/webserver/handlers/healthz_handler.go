package handlers

import (
	"fmt"
	"net/http"
	"time"
)


type HealthzHandler struct {
     startedAt time.Time
}

func NewHealthzHandler() *HealthzHandler {
     return &HealthzHandler{
          startedAt: time.Now(),
     }
}

func (h *HealthzHandler) Healthz(w http.ResponseWriter, r *http.Request) {
     duration := time.Since(h.startedAt)
     if duration.Seconds() < 10 || duration.Seconds() > 10 * 60 {
          w.WriteHeader(http.StatusInternalServerError)
          w.Write([]byte(fmt.Sprintf("Healthz check failed after %v seconds", duration.Seconds())))
     } else {
          w.WriteHeader(http.StatusOK)
          w.Write([]byte("Healthz check passed"))
     }
}

